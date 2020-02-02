package imageConverter

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// original is original file information.
type original struct {
	path string // ファイル名を除くパス
	name string // 拡張子を除くファイル名
	ext  string // 拡張子
}

// ImageConverter is image object.
type ImageConverter struct {
	data image.Image // デコードした画像データ
	org  original    // 元ファイル情報
}

// New creates a new image object.
func New(path string) (*ImageConverter, error) {
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

	return &ImageConverter{
		data: m,
		org: original{
			path: dir,
			name: name,
			ext:  ext,
		},
	}, nil
}

// originalFile gets original file path.
func (i *ImageConverter) originalFile() string {
	return filepath.Join(i.org.path, i.org.name+i.org.ext)
}

// ToPNG converts original image to PNG format.
// The original image is not deleted.
func (i *ImageConverter) ToPNG() error {
	// 変換後の画像ファイルを作る
	path := filepath.Join(i.org.path, i.org.name+".png")
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

// ToJPG converts original image to JPG format.
// The original image is not deleted.
func (i *ImageConverter) ToJPG() error {
	// 変換後の画像ファイルを作る
	path := filepath.Join(i.org.path, i.org.name+".jpg")
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

// OriginalRemove remove origial file.
func (i *ImageConverter) OriginalRemove() error {
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
