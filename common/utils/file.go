package utils

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

//func CopyFile(srcName, dstName string) (written int64, err error) {
//	src, err := os.Open(srcName)
//	if err != nil {
//		return
//	}
//	defer src.Close()
//	dst, err := os.Create(dstName)
//	if err != nil {
//		return
//	}
//	defer dst.Close()
//	return io.Copy(dst, src) //
//}

func WritePID(filePath string) bool {
	pidFile, err := os.Create(filePath)
	if err != nil {
		return false
	}
	defer pidFile.Close()
	pidFile.WriteString(fmt.Sprintln(os.Getpid()))
	return true
}

func GetPID(filePath string) int {
	pidFile, err := os.Open(filePath)
	if err != nil {
		return 0
	}
	defer pidFile.Close()
	var pid int
	fmt.Fscanf(pidFile, "%d", &pid)
	return pid
}

// File copies a single file from src to dst
func CopyFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

// CopyDir copies a whole directory recursively
func CopyDir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	srcinfo, err = os.Stat(src)
	if os.IsNotExist(err) {
		return errors.New("src dir not exist")
	}
	_, err = os.Stat(dst)

	if os.IsNotExist(err) {
		if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
			return err
		}
	}
	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = CopyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

func ShouldReadFile(filePath string) []byte {
	fp, err := os.Open(filePath)
	var result []byte
	if err != nil {
		return result
	}
	defer fp.Close()
	result, err = ioutil.ReadAll(fp)
	return result
}
