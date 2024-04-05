package utils

import "fmt"

func ReadInput() string {
	var input string

	_, err := fmt.Scanln(&input)
	if err != nil {
		if err.Error() == "unexpected newline" {
			return ""
		}

		panic(err)
	}

	return input
}
