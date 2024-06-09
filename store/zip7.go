package store

import (
	"github.com/essentialkaos/zip7"
)

func Decompression(fileInfo *zip7.Props) error {
	if fileInfo != nil {
		_, err := zip7.Extract(*fileInfo)
		if err != nil {
			return err
		}
	}
	return nil
}
