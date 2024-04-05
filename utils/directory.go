package utils

import (
	"github.com/Vladimir-Urik/cl/enums"
	"github.com/Vladimir-Urik/cl/structs"
	"github.com/dustin/go-humanize"
	"os"
	"sort"
)

func GetAllDirectories(path string) []structs.ListItem {
	var filesAndDirs []structs.ListItem

	data, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range data {
		info, err := file.Info()
		if err != nil {
			panic(err)
		}

		itemType := enums.File
		if file.IsDir() {
			itemType = enums.Directory
		}

		println(path)
		println(file.Name())

		filesAndDirs = append(filesAndDirs, structs.ListItem{
			Path:          path + "/" + file.Name(),
			Name:          file.Name(),
			Size:          info.Size(),
			FormattedSize: humanize.Bytes(uint64(info.Size())),
			Type:          itemType,
		})
	}

	sort.SliceStable(filesAndDirs, func(i, j int) bool {
		return filesAndDirs[i].Type < filesAndDirs[j].Type
	})

	return filesAndDirs
}

func GetSubdirectory(directory string) string {
	lastIndex := 0
	for i, char := range directory {
		if string(char) == "/" {
			lastIndex = i
		}
	}

	return directory[:lastIndex]
}

func DeleteFile(path string) interface{} {
	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
