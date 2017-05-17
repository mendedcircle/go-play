# ffmpeg_convert_video
Convert {.mkv, .mp4, .webm} video files to .mp3

```golang
func convertVideoToMp3(videoFilename, mp3Filename) error {
     fmpegCmd := "/usr/local/bin/ffmpeg"
     argsList := []string{"-i", videoFilename, "-vn", "-c:a", "libmp3lame", "-y", mp3Filename}

     if err := exec.Command(ffmpegCmd, argsList...).Run(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return err
        os.Exit(1)
     }
     return nil
}
```

