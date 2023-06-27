package task

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Tasker interface {
	Process() error
}

type dirCtx struct {
	SrcDir string
	DstDir string
	files  []string
}

func buildFileList(srcDir string) []string {
	files := []string{}
	fmt.Println("Generation du fichier...")
	filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix((path), ".jpg") {
			return nil
		}

		files = append(files, path)
		return nil
	})

	return files
}
