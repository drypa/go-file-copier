package main

import (
	"fmt"
	"github.com/mkideal/cli"
	"io"
	"log"
	"os"
	"path/filepath"
)

type args struct {
	FilePath string `cli:"file" usage:"file"`
	Count    int    `cli:"count" usage:"copies count"`
}

func main() {
	cli.Run(new(args), func(ctx *cli.Context) error {

		argv := ctx.Argv().(*args)
		file, err := os.Open(argv.FilePath)
		if err != nil {
			return err
		}
		defer file.Close()
		dir := filepath.Dir(argv.FilePath)
		ext := filepath.Ext(argv.FilePath)
		copyDir := filepath.Join(dir, "copies")
		os.Mkdir(copyDir, os.ModeDir)

		for i := 0; i < argv.Count; i++ {
			createFile(file, copyDir, i, ext)
		}
		return nil
	})
}

func createFile(file *os.File, copyDir string, num int, ext string) {
	file.Seek(0, io.SeekStart)
	copyName := filepath.Join(copyDir, fmt.Sprintf("%d%s", num, ext))
	copy, err := os.Create(copyName)
	if err != nil {
		log.Printf("Error creating copy file %v", err)
	}
	defer copy.Close()
	_, err = io.Copy(copy, file)
	if err != nil {
		log.Printf("Error while copy file %v", err)
	}
	err = copy.Sync()
	if err != nil {
		log.Printf("Error while flush copy file %v", err)
	}
	log.Println(copyName)
}
