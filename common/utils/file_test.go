package utils

import (
	"fmt"
	"testing"
)

func TestGetDirSize(t *testing.T) {
	t.Error(GetDirSize("."))
}

func TestTarGzipCompressFuture(t *testing.T) {
	value, err := TarGzipCompressFuture("E:/Music", "E:/music.tar.gz").OnProgress(func(p float64) {
		fmt.Println(p * 100)
	}).Get()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(value)
	}
}
