package menus

import (
	"github.com/Vladimir-Urik/cl/structs"
	"github.com/fatih/color"
)

func PrintExplorer(directory string, activeIndex int, filesAndDirs []structs.ListItem) {
	print("\033[H\033[2J")
	color.Blue("🗀 Directory: ")
	println(color.CyanString("  - " + directory))
	println("")

	indexFrom := 0
	indexTo := 9
	if activeIndex > 5 {
		indexFrom = activeIndex - 5
		indexTo = activeIndex + 4
	}

	if activeIndex > len(filesAndDirs)-5 {
		indexFrom = len(filesAndDirs) - 10
		indexTo = len(filesAndDirs) - 1
	}

	if indexFrom > 9 {
		println("...")
	}

	for index, fileOrDir := range filesAndDirs {
		if index < indexFrom || index > indexTo {
			continue
		}

		if index == activeIndex {
			println(fileOrDir.FormatAsSelected())
			continue
		}

		println(fileOrDir.FormatDisplay())
	}

	if indexTo < (len(filesAndDirs) - 1) {
		println("...")
	}

	println("")
	println("w - up, s - down, d - enter, a - back, q - quit, r - delete")
}

func printFileDeletionConfirmation(file structs.ListItem) {
	print("\033[H\033[2J")
	color.Red("Are you sure you want to delete the following file?")
	println("")
	println(file.FormatDeleteConfirmation())
	println("")
	println("Press [y] to confirm, [n] to cancel")
}
