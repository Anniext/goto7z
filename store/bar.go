package store

import (
	"fmt"
	"github.com/gosuri/uiprogress"
	"io/fs"
	"path/filepath"
	"strings"
)

var Bar *uiprogress.Bar

func InitBar(input string) {
	// 进度条
	count, _ := countFiles(input)
	Bar = uiprogress.AddBar(count)
	Bar.PrependFunc(func(b *uiprogress.Bar) string {
		return "Progress:"
	})
	Bar.Width = 50
	Bar.Total = count
	Bar.AppendCompleted()
	Bar.AppendFunc(func(b *uiprogress.Bar) string {
		return fmt.Sprintf("%v/%v", b.Current(), b.Total)
	})
	uiprogress.Start()
}

func countFiles(dirPath string) (int, error) {
	count := 0
	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		suffixStrLen := len(strings.Split(d.Name(), "."))
		if !d.IsDir() && strings.Split(d.Name(), ".")[suffixStrLen-2] != "" {
			count++
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}
