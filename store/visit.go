package store

import (
	"anniext.natapp4.cc/xt/goto7z/profile"
	"fmt"
	"github.com/essentialkaos/zip7"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var (
	fileCount = 1
)

func Visit(path string, info os.FileInfo, err error) error {
	var fileNeme string
	config, _ := profile.GetProfile()
	Bar.Set(fileCount)

	if err != nil {
		fmt.Printf("Error accessing path %q: %v\n", path, err)
		return err
	}

	if info.IsDir() {
		return nil
	}

	suffixStrLen := len(strings.Split(info.Name(), "."))
	if strings.Split(info.Name(), ".")[suffixStrLen-2] == "" {
		return nil
	}

	suffix := strings.Split(info.Name(), ".")[suffixStrLen-1]

	if suffix != config.Mode {
		fileNeme = strings.Split(info.Name(), ".")[suffixStrLen-2] + "." + config.Mode
		err := os.Rename(path, filepath.Dir(path)+"/"+fileNeme)
		if err != nil {
			fmt.Println(err)
			return err
		}
	} else {
		fileNeme = info.Name()
	}

	fileInfo := &zip7.Props{
		File:      filepath.Dir(path) + "/" + fileNeme,
		Password:  config.Passwd,
		OutputDir: config.Output,
		Threads:   8,
	}

	err = Decompression(fileInfo)
	if err != nil {
		fmt.Println(err)
	}

	os.Remove(filepath.Dir(path) + "/" + fileNeme)
	fname, err := checkFilesExist(config.Output)
	if err != nil {
		return err
	}

	if fname != "" {
		fnamelist := strings.Split(fname, ".")
		mstuffix := strings.Split(fname, ".")[len(fnamelist)-1]
		if mstuffix != config.Mode {
			err := os.Rename(config.Output+"/"+fname, filepath.Dir(config.Output)+"/"+fnamelist[len(fnamelist)-2]+"."+config.Mode)
			if err != nil {
				fmt.Println(err)
				return err
			}
			fname = fnamelist[len(fnamelist)-2] + "." + config.Mode
		}
		fi := &zip7.Props{
			File:      filepath.Dir(config.Output) + "/" + fname,
			Password:  config.Passwd,
			OutputDir: config.Output,
			Threads:   8,
		}
		err = Decompression(fi)
		os.Remove(filepath.Dir(config.Output) + "/" + fname)
	}

	fileCount += 1

	return nil
}

func checkFilesExist(dirPath string) (string, error) {
	var fileName string
	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		suffixStrLen := len(strings.Split(d.Name(), "."))
		if !d.IsDir() && strings.Split(d.Name(), ".")[suffixStrLen-2] != "" {
			fileName = d.Name()
			return filepath.SkipAll
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return fileName, nil
}
