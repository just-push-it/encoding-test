package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "" {
		fmt.Println("No input file specified")
		return
	}
	input := os.Args[1]
	ext := filepath.Ext(input)
	if !isSupportedExtension(ext) {
		fmt.Println("Unsupported file type")
		return
	}
	filename := strings.Split(filepath.Base(input), ".")[0]
	outdir := filepath.Dir(input)
	outfile := filepath.Join(outdir, filename+"_encoded.mp4")

	fmt.Println("input", input)
	fmt.Println("outfile", outfile)

	ffmpeg.Input(input, nil).Output(outfile, ffmpeg.KwArgs{"c:v": "hevc_nvenc"}).
		OverWriteOutput().ErrorToStdOut().Run()
}

func isSupportedExtension(ext string) bool {
	supportedExtensions := []string{".mp4"}
	for _, v := range supportedExtensions {
		if strings.ToLower(ext) == v {
			return true
		}
	}
	return false
}
