package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	// ProjectName 模板项目名
	ProjectName = "console-template"
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		os.Exit(2)
		return
	}

	if args[0] == "init" {
		newProjectName := args[1]
		currentPath, _ := os.Getwd()
		filepath.Walk(currentPath, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				if info.Name() == ".git" || info.Name() == "node_modules" || info.Name() == "vendor" || info.Name() == "cmd" {
					return filepath.SkipDir
				}

				return nil
			}

			content := ReadFile(path)

			if strings.Contains(content, ProjectName) {
				content = strings.Replace(content, ProjectName, newProjectName, -1)
				WriteToFile(path, content)
				fmt.Printf("\t%s%supdate file%s\t %s%s\n", "\x1b[36m", "\x1b[1m", "\x1b[21m", path, "\x1b[0m")
			}
			return nil
		})
		return
	}

}

// MustCheck panics when the error is not nil
func MustCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// WriteToFile creates a file and writes content to it
func WriteToFile(filename, content string) {
	f, err := os.Create(filename)
	MustCheck(err)
	defer CloseFile(f)
	_, err = f.WriteString(content)
	MustCheck(err)
}

// ReadFile 读取文件
func ReadFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	MustCheck(err)
	return string(b)
}

// CloseFile attempts to close the passed file
// or panics with the actual error
func CloseFile(f *os.File) {
	err := f.Close()
	MustCheck(err)
}
