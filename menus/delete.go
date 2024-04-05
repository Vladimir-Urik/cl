package menus

import (
	"github.com/Vladimir-Urik/cl/structs"
	"github.com/fatih/color"
)

func PrintFileDeletionConfirmation(file structs.ListItem) {
	print("\033[H\033[2J")
	color.Red("Are you sure you want to delete the following file?")
	println("")
	println(file.FormatDeleteConfirmation())
	println("")
	println("Press [y] to confirm, [n] to cancel")
}
