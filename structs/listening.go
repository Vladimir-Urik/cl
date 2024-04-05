package structs

import (
	"github.com/Vladimir-Urik/cl/enums"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

type ListItem struct {
	Path          string
	Name          string
	Size          int64
	FormattedSize string
	Type          enums.ItemType
}

func (li ListItem) GetTitle() string {
	directoryIcon := "🗀"
	fileIcon := "🖹"

	text := ""
	if li.Type == enums.Directory {
		text = directoryIcon
		return " " + text + " " + li.Name + "  "
	} else {
		text = fileIcon
		return " " + text + " " + li.Name + " (" + li.FormattedSize + ")  "
	}
}

func (li ListItem) FormatDisplay() string {
	c := color.New(color.FgBlue)
	if li.Type == enums.File {
		c = color.New(color.FgGreen)
	}

	return c.Sprint(li.GetTitle())
}

func (li ListItem) FormatDeleteConfirmation() string {
	c := color.New(color.FgGreen)
	return c.Sprint(li.Path)
}

func (li ListItem) FormatAsSelected() string {
	endText := " [d] to enter directory"

	c := color.New(color.BgBlue).Add(color.FgBlack)
	if li.Type == enums.File {
		c = color.New(color.BgGreen).Add(color.FgBlack)
		endText = " [r] to delete file"
	}

	value := c.Sprint(li.GetTitle())

	width, _, err := terminal.GetSize(0)
	if err == nil {
		for i := len(li.GetTitle()); i < (width - len(endText)); i++ {
			value += " "
		}
	}

	return value + endText
}
