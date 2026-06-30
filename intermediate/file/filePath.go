package file

import (
	"fmt"
	"path/filepath"
)

func FilePath() {

	joinedPath := filepath.Join("folder1", "folder2", "file.txt")
	println("Joined Path:", joinedPath)

	cleanPath := filepath.Clean("/folder1/../folder2/file.txt")
	println("Cleaned Path:", cleanPath)

	dir, file := filepath.Split("/folder1/folder2/file.txt")
	println("Directory:", dir)
	println("File:", file)

	fmt.Print(filepath.Ext(file))
}
