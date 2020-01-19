package myimage

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// myFile is file information.
type myFile struct {
	path string // ファイル名を除くパス
	name string // 拡張子を除くファイル名
	ext  string // 拡張子
}

// MyImage is image object.
type MyImage struct {
	data image.Image // デコードした画像データ
	file myFile      // ファイル名情報
}

// NewImage creates a new image object.
func NewImage(path string) (*MyImage, error) {
	dir, name, ext := exceptExt(path)
	r, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var m image.Image
	switch filepath.Ext(path) {
	case ".jpg":
		m, err = jpeg.Decode(r)
	case ".png":
		m, err = png.Decode(r)
	}
	if err != nil {
		return nil, err
	}

	return &MyImage{
		data: m,
		file: myFile{
			path: dir,
			name: name,
			ext:  ext,
		},
	}, nil
}

// originalFile gets original file path.
func (i *MyImage) originalFile() string {
	return filepath.Join(i.file.path, i.file.name+i.file.ext)
}

// GetExt gets extension
func (i *MyImage) GetExt() string {
	return i.file.ext
}

// ConvertToPNG converts original image to PNG format.
// The original image is not deleted.
func (i *MyImage) ConvertToPNG() error {
	// 変換後の画像ファイルを作る
	path := filepath.Join(i.file.path, i.file.name+".png")
	w, err := os.Create(path)
	if err != nil {
		return err
	}
	defer w.Close()

	if err := png.Encode(w, i.data); err != nil {
		return err
	}
	return nil
}

// ConvertToJPG converts original image to JPG format.
// The original image is not deleted.
func (i *MyImage) ConvertToJPG() error {
	// 変換後の画像ファイルを作る
	path := filepath.Join(i.file.path, i.file.name+".jpg")
	w, err := os.Create(path)
	if err != nil {
		return err
	}
	defer w.Close()

	opts := &jpeg.Options{Quality: 100}
	if err := jpeg.Encode(w, i.data, opts); err != nil {
		return err
	}
	return nil
}

// Remove remove origial file.
func (i *MyImage) Remove() error {
	err := os.Remove(i.originalFile())
	if err != nil {
		return err
	}
	return nil
}

func exceptExt(filename string) (string, string, string) {
	dir := filepath.Dir(filename)
	tmp := filepath.Base(filename)
	ext := filepath.Ext(filename)
	name := strings.Split(tmp, ext)
	return dir, name[0], ext
}

// IsFormat returns whether the format is supported
func IsFormat(f string) bool {
	switch f {
	case ".jpg", ".png":
		return true
	default:
		return false
	}
}
