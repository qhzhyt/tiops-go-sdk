package utils

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
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

func TarGzipCompressFuture(srcPath, dstFilePath string) *Future {
	return RunFuture(func(future *Future) {

		dirSize, err := GetDirSize(srcPath)

		if err != nil {
			future.SetError(err)
			return
		}

		unitSize := dirSize / 1000

		processSize := int64(0)
		lastProgressSize := int64(0)

		compress := func(file *os.File, prefix string, tw *tar.Writer) error {
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
				header.Format = tar.FormatGNU
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
				processSize += info.Size()
				if processSize-unitSize > lastProgressSize {
					lastProgressSize = processSize
					future.SetProgress(float64(processSize) / float64(dirSize))
				}
			}
			return nil
		}

		tarGzipCompress := func(files []*os.File, dest string) error {
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

		file, err := os.Open(srcPath)
		if os.IsNotExist(err) {
			future.SetError(err)
			return
		}
		stat, _ := file.Stat()
		if stat.IsDir() {
			fileInfos, err := file.Readdir(-1)
			if err != nil {
				future.SetError(err)
				return
			}
			fileList := make([]*os.File, len(fileInfos))
			for i, info := range fileInfos {
				fp, err := os.Open(path.Join(srcPath, info.Name()))
				if err != nil {
					future.SetError(err)
					return
				}
				fileList[i] = fp
			}
			if err := tarGzipCompress(fileList, dstFilePath); err != nil {
				future.SetError(err)
				return
			}
		} else {
			if err := tarGzipCompress([]*os.File{file}, dstFilePath); err != nil {
				future.SetError(err)
				return
			}
		}

		future.SetResult(dirSize)
	})
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

type TempWriter struct {
	data     []byte
	callback func(string)
}

func (t *TempWriter) Write(p []byte) (n int, err error) {
	t.data = append(t.data, p...)
	if t.callback != nil {
		t.callback(string(p))
	}
	return len(p), nil
}

func (t *TempWriter) String() string {
	return string(t.data)
}

func TarZCF(srcPath, dstFilePath string) {

	dirSize, _ := GetDirSize(srcPath)

	block := dirSize/512/100 + 1

	//--blocking-factor=$block_size --checkpoint=1 --checkpoint-action=dot

	tarCMD := exec.Command("tar", fmt.Sprintf("--blocking-factor=%d", block), "--checkpoint=1", "--checkpoint-action=dot", "-zcf", dstFilePath, "-C", srcPath, ".")

	//fmt.Println(tarCMD.String())
	tarCMD.Stdout = &TempWriter{
		callback: func(s string) {
			fmt.Print(s)
		},
	}

	tarCMD.Run()
}

func TarJCFFuture(srcPath, dstFilePath string) *Future {

	return RunFuture(func(future *Future) {

		if err := os.MkdirAll(path.Dir(dstFilePath), 0777); err != nil {
			future.SetError(err)
			return
		}

		dirSize, err := GetDirSize(srcPath)

		if err != nil {
			future.SetError(err)
			return
		}

		block := dirSize/512/100 + 1

		os.Setenv("XZ_OPT", "-T0")

		//--blocking-factor=$block_size --checkpoint=1 --checkpoint-action=dot

		tarCMD := exec.Command("tar", fmt.Sprintf("--blocking-factor=%d", block), "--checkpoint=1", "--checkpoint-action=dot", "-Jcf", dstFilePath, "-C", srcPath, ".")

		//tarCMD.Dir = srcPath

		//fmt.Println(tarCMD.String())
		tarCMD.Stdout = &TempWriter{
			callback: func(s string) {
				//fmt.Print(s)
				for _, r := range s {
					if r == '.' {
						future.AddProgress(float64(1) / float64(100))
					}
				}
			},
		}

		tarCMD.Stderr = &TempWriter{
			callback: func(s string) {
				//fmt.Print(s)
				fmt.Print(s)
			},
		}

		err = tarCMD.Run()
		if err != nil {
			future.SetError(err)
			return
		}
		future.SetResult(dirSize)
	})
}

func TarJXFFuture(srcPath, dstFilePath string) *Future {

	return RunFuture(func(future *Future) {
		//dirSize := GetFileSize(srcPath)

		//xz -l /home/hyt/thesis.tar.xz

		os.Setenv("XZ_OPT", "-T0")

		xzCMD := exec.Command("xz", "--robot", "-l", srcPath)

		tmpWriter := &TempWriter{}
		xzCMD.Stdout = tmpWriter
		err := xzCMD.Run()
		if err != nil {
			future.SetError(err)
			return
		}

		temp := 1
		fileSize := 0

		_, _ = fmt.Sscanf(strings.TrimSpace(strings.Split(tmpWriter.String(), "totals")[1]), "%d %d %d %d", &temp, &temp, &temp, &fileSize)

		fmt.Println(fileSize)

		block := fileSize/512/100 + 1

		//--blocking-factor=$block_size --checkpoint=1 --checkpoint-action=dot

		tarCMD := exec.Command("tar", fmt.Sprintf("--blocking-factor=%d", block), "--checkpoint=1", "--checkpoint-action=dot", "-xf", srcPath, "-C", dstFilePath)

		//fmt.Println(tarCMD.String())
		tarCMD.Stdout = &TempWriter{
			callback: func(s string) {
				//fmt.Print(s)
				for _, r := range s {
					if r == '.' {
						future.AddProgress(float64(1) / float64(100))
					}
				}
			},
		}

		err = tarCMD.Run()
		if err != nil {
			future.SetError(err)
			return
		}
		future.SetResult(fileSize)
	})
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
		header.Format = tar.FormatGNU
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

//DeCompress 解压 tar.gz
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

//GetDirSize get dir size by path
func GetDirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		//fmt.Println(err)
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

//GetFileSize get file size by path
func GetFileSize(path string) int64 {
	if !exists(path) {
		return 0
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

//exists Whether the path exists
func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
