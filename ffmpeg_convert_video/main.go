package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	extensions  = []string{".mp4", ".mkv", ".webm"}
	srcFolder   = "." + string(filepath.Separator)
	destination = "finish"
)

func main() {
	watchFolderForChanges(srcFolder, extensions)
}

// watchForChanges - watch a particular folder for changes.
func watchFolderForChanges(targetFolder string, extensions []string) {
	for { // while
		// get a hold of the handle to the folder we are watching
		watchedDir, err := os.Open(targetFolder)
		check(err)

		// enumerate the contents of the watched directory
		files, err := watchedDir.Readdir(-1) // -1, returns as many files as it find
		check(err)
		// already have the contents, let close the file handle
		watchedDir.Close()

		// process files
		processFileByExt(files, extensions)
	}
}

// Utils
func processFileByExt(files []os.FileInfo, extenions []string) {
	for _, ext := range extensions {
		for _, videoFile := range files {
			if filepath.Ext(videoFile.Name()) == ext {
				// process videoFile
				fmt.Println("Converting ", videoFile.Name(), "...")
				videoFileName := videoFile.Name()
				mp3file := getMp3Ext(videoFile.Name())
				converVideoToMp3(videoFileName, mp3file)
				fmt.Printf("Moving %s to finish folder\n", mp3file)
				moveToFinishFolder(mp3file)
				fmt.Printf("Removing %s video file\n", videoFileName)
				os.Remove(videoFileName)
				fmt.Printf("Successfully converted video: %s to audio: %s\n", videoFileName, mp3file)
			}
		}
	}
	//fmt.Println("Successfully converted video file")
}

// getMp3Ext - remove video extenstion set it as .mp3
func getMp3Ext(fileName string) string {
	extension := filepath.Ext(fileName)
	mp3FileName := fileName[0 : len(fileName)-len(extension)]
	return fmt.Sprintf("%s.mp3", mp3FileName)
}

// moveToFinishFolder
func moveToFinishFolder(fileName string) {
	moveCmd := "/bin/mv"
	// create a directory if it does not exist, otherwise do noting.
	if _, err := os.Stat(destination); os.IsNotExist(err) {
		err = os.MkdirAll(destination, 0755)
		if err != nil {
			panic(err)
		}
	}
	args := []string{fileName, destination}
	if err := exec.Command(moveCmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// convervideoToMp3 - convert video to mp3 with ffmpeg
func converVideoToMp3(videoFileName, mp3FileName string) error {
	ffmpegCmd := "/usr/local/bin/ffmpeg"
	args := []string{"-i", videoFileName, "-vn", "-c:a", "libmp3lame", "-y", mp3FileName}

	if err := exec.Command(ffmpegCmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
		os.Exit(1)
	}
	return nil
}

// check error not nil
func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
