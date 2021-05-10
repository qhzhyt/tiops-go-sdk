package utils

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
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

//压缩 使用gzip压缩成tar.gz
func tarGzipCompress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	gw := gzip.NewWriter(d)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	for _, file := range files {
		err := compress(file, "", tw)
		if err != nil {
			return err
		}
	}
	return nil
}

func TarGzipCompress(srcPath, dstFilePath string) error {
	file, err := os.Open(srcPath)
	if os.IsNotExist(err) {
		return err
	}
	stat, _ := file.Stat()
	if stat.IsDir() {
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		fileList := make([]*os.File, len(fileInfos))
		for i, info := range fileInfos {
			fp, err := os.Open(path.Join(srcPath, info.Name()))
			if err != nil {
				return err
			}
			fileList[i] = fp
		}
		return tarGzipCompress(fileList, dstFilePath)
	} else {
		return tarGzipCompress([]*os.File{file}, dstFilePath)
	}
}

func compress(file *os.File, prefix string, tw *tar.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, tw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := tar.FileInfoHeader(info, "")
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//解压 tar.gz
func DeCompress(tarFile, dest string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := dest + hdr.Name
		file, err := createFile(filename)
		if err != nil {
			return err
		}
		io.Copy(file, tr)
	}
	return nil
}

func createFile(name string) (*os.File, error) {
	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}
