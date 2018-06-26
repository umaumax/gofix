package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
)

func loadDict(filetype string) (list []string, err error) {
	configDirPath, err := homedir.Expand("~/.config/gofix")
	if err != nil {
		return
	}
	filenames, err := filepath.Glob(filepath.Join(configDirPath, fmt.Sprintf("%s.gofix", filetype)))
	if err != nil {
		return
	}
	{
		var commonFilenames []string
		commonFilenames, err = filepath.Glob(filepath.Join(configDirPath, "common.gofix"))
		if err != nil {
			return
		}
		filenames = append(filenames, commonFilenames...)
	}
	for _, filename := range filenames {
		var f *os.File
		f, err = os.Open(filename)
		if err != nil {
			return
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "#") {
				continue
			}
			list = append(list, line)
		}
		f.Close()
		if err = scanner.Err(); err != nil {
			return
		}
	}
	return
}
