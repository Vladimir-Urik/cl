package utils

import "os"

func GetDirectory() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return path
}
