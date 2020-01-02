package main

import (
	"flag"
	"os"
	"path/filepath"
)

var root = flag.String("dir", "", "ディレクトリ") // ディレクトリの指定
var af = flag.String("a", ".png", "変換後の形式") // 変換後の画像形式を指定
var bf = flag.String("b", ".jpg", "変換前の形式") // 変換前の画像形式を指定

// 画像変換する
var reconv = func(path string, info os.FileInfo, err error) error {
	if filepath.Ext(path) == *bf {
		i := NewImage(path)

		switch i.GetExt() {
		case ".jpg":
			i.ConvertToPNG()
		case ".png":
			i.ConvertToJPG()
		}
		i.Remove()
	}
	return nil
}

func main() {
	flag.Parse()

	err := filepath.Walk(*root, reconv)
	if err != nil {
		return
	}
}
