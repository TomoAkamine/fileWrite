package fileWrite

import "testing"

func TestFileWrite(t *testing.T) {
	result := Hello("TomoAkamine")
	want := "Hi, TomoAkamine. Welcome!"
	if result != want {
		t.Errorf("fileWrite.Hello() = %q want %q", result, want)
	}
}
