/*
Copyright © 2024 Vladimír Urík <info@gggedr.lol>
*/
package main

import (
	"github.com/Vladimir-Urik/cl/enums"
	"github.com/Vladimir-Urik/cl/menus"
	"github.com/Vladimir-Urik/cl/structs"
	"github.com/Vladimir-Urik/cl/utils"
	"github.com/eiannone/keyboard"
	"golang.org/x/crypto/ssh/terminal"
)

var isInMainMenu = true
var fileToDelete = structs.ListItem{}

func main() {
	activeIndex := 0
	directory := utils.GetDirectory()
	filesAndDirs := utils.GetAllDirectories(utils.GetDirectory())

	lastWidth := 0
	lastHeight := 0

	go func() {
		for {
			width, height, err := terminal.GetSize(0)
			if err != nil {
				panic(err)
			}

			if width != lastWidth || height != lastHeight {
				lastWidth = width
				lastHeight = height
				if isInMainMenu {
					menus.PrintExplorer(directory, activeIndex, filesAndDirs)
				}
			}
		}
	}()

	for {
		if isInMainMenu {
			menus.PrintExplorer(directory, activeIndex, filesAndDirs)
		}

		if !isInMainMenu && fileToDelete != (structs.ListItem{}) {
			menus.PrintFileDeletionConfirmation(fileToDelete)
		}

		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}

		if isInMainMenu {
			if char == 119 {
				if activeIndex == 0 {
					activeIndex = len(filesAndDirs) - 1
				} else {
					activeIndex--
				}
			}

			if char == 115 {
				if activeIndex == len(filesAndDirs)-1 {
					activeIndex = 0
				} else {
					activeIndex++
				}
			}

			if char == 100 {
				if filesAndDirs[activeIndex].Type == enums.File {
					continue
				}

				dir := filesAndDirs[activeIndex]
				directory = filesAndDirs[activeIndex].Path
				filesAndDirs = utils.GetAllDirectories(dir.Path)
				activeIndex = 0
			}

			if char == 97 {
				if directory == utils.GetDirectory() {
					continue
				}

				directory = utils.GetSubdirectory(directory)
				filesAndDirs = utils.GetAllDirectories(directory)
				activeIndex = 0
			}

			if char == 113 {
				print("\033[H\033[2J")
				break
			}

			if char == 114 {
				if filesAndDirs[activeIndex].Type == enums.Directory {
					continue
				}

				isInMainMenu = false
				fileToDelete = filesAndDirs[activeIndex]
			}
		}

		if !isInMainMenu && fileToDelete != (structs.ListItem{}) {
			if char == 121 {
				err := utils.DeleteFile(fileToDelete.Path)
				if err != nil {
					panic(err)
				}

				isInMainMenu = true
				filesAndDirs = utils.GetAllDirectories(utils.GetDirectory())
				activeIndex = 0
			}

			if char == 110 {
				isInMainMenu = true
				activeIndex = 0
			}
		}
	}
}
