package imageConverter

import "testing"

func TestExceptExt(t *testing.T) {
	in := "/testroot/test/image.png"
	dir, name, ext := exceptExt(in)

	expectDir := "/testroot/test"
	expectName := "image"
	expectExt := ".png"

	if dir != expectDir {
		t.Error("\nActual： ", dir, "\nExpect： ", expectDir)
	}
	if name != expectName {
		t.Error("\nActual： ", name, "\nExpect： ", expectName)
	}
	if ext != expectExt {
		t.Error("\nActual: ", ext, "\nExpect: ", expectExt)
	}

	t.Log("TestExceptExt finish")
}
