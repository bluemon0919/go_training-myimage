package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"go_training-myimage/imageConverter"
)

var afterFormat string  // 変換後の画像形式
var beforeFormat string // 変換前の画像形式

// 画像変換する
func convert(path string, info os.FileInfo, err error) error {
	if filepath.Ext(path) != beforeFormat {
		return nil
	}

	c, err := imageConverter.New(path)
	if err != nil {
		return err
	}

	switch afterFormat {
	case ".png":
		err = c.ToPNG()
	case ".jpg":
		err = c.ToJPG()
	default:
		err = fmt.Errorf("after format is not supported")
	}
	if err != nil {
		return err
	}

	if err := c.RemoveOriginal(); err != nil {
		return err
	}
	return nil
}

// argsCheck checks that the arguments are correct
func argsCheck(args []string) error {
	// パラメータ数が一致することを確認する
	if 3 != len(args) {
		return fmt.Errorf("パラメータが一致しません")
	}

	// <search_dir>に指定されたディレクトリが存在することを確認する
	fInfo, err := os.Stat(args[0])
	if err != nil || !fInfo.IsDir() {
		return fmt.Errorf("%sはディレクトリではありません", args[0])
	}

	// <after_format>, <before_format>に指定された画像フォーマットが対応しているか確認する
	if !imageConverter.IsFormat(args[1]) {
		return fmt.Errorf("%sは使用できるフォーマットではありません", args[1])
	}
	if !imageConverter.IsFormat(args[2]) {
		return fmt.Errorf("%sは使用できるフォーマットではありません", args[2])
	}
	return nil
}

func main() {
	flag.Parse()
	args := flag.Args()

	if err := argsCheck(args); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	beforeFormat = args[1]
	afterFormat = args[2]

	if err := filepath.Walk(args[0], convert); err != nil {
		fmt.Println(err)
	}
}
