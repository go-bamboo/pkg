package filex

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	} else {
		return true
	}
}

// OpenFile 打开文件
func OpenFile(fpath string, flag int, perm os.FileMode) (f *os.File, err error) {
	xpath, _ := path.Split(fpath)
	if !IsExist(xpath) {
		if err = os.MkdirAll(xpath, os.ModePerm); err != nil {
			return nil, err
		}
	}
	f, err = os.OpenFile(fpath, flag, perm)
	if err != nil {
		return
	}
	return
}

func CreateFile(fpath string) (err error) {
	xpath, _ := path.Split(fpath)
	if !IsExist(xpath) {
		if err = os.MkdirAll(xpath, os.ModePerm); err != nil {
			return err
		}
	}
	_, err = os.Create(fpath)
	if err != nil {
		return
	}
	return
}

func Remove(name string) (err error) {
	if err = os.RemoveAll(name); err != nil {
		return
	}
	return
}

// FileZip
// 文件ZIP压缩
func FileZip(dst, src string, notContPath string) (err error) {
	//创建准备写入的文件
	fw, err := os.Create(dst)
	defer fw.Close()
	if err != nil {
		return err
	}

	// 通过 fw 来创建 zip.Write
	zw := zip.NewWriter(fw)
	defer func() {
		// 检测一下是否成功关闭
		if err := zw.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	return filepath.Walk(src, func(path string, fi os.FileInfo, errBack error) (err error) {
		if errBack != nil {
			return errBack
		}

		fh, err := zip.FileInfoHeader(fi)
		if err != nil {
			return
		}

		fh.Name = strings.TrimPrefix(path, string(filepath.Separator))

		if fi.IsDir() {
			fh.Name += "/"
		}
		fh.Name = strings.Replace(fh.Name, notContPath, "", -1)

		w, err := zw.CreateHeader(fh)
		if err != nil {
			return
		}

		if !fh.Mode().IsRegular() {
			return nil
		}

		fr, err := os.Open(path)
		defer fr.Close()
		if err != nil {
			return
		}

		n, err := io.Copy(w, fr)
		if err != nil {
			return
		}

		fmt.Printf("成功压缩文件： %s, 共写入了 %d 个字符的数据\n", path, n)

		return nil
	})
}

// Size 获取文件大小
func Size(r io.Reader) int {
	bufr := bufio.NewReader(r)
	return bufr.Size()
}

// Ext 获取文件后缀
func Ext(fileName string) string {
	return path.Ext(fileName)
}

// CheckPermission 检查文件权限
func CheckPermission(src string, mode os.FileMode) (bool, error) {
	info, err := os.Stat(src)
	if err != nil {
		if os.IsPermission(err) {
			return true, nil
		}
		return false, err
	}
	if info.Mode() == mode {
		return true, nil
	}
	return false, nil
}

// IsNotExistMkDir 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	exist := IsExist(src)
	if !exist {
		if err := Mkdir(src); err != nil {
			return err
		}
	}
	return nil
}

// Mkdir 新建文件夹
func Mkdir(src string) error {
	err := os.Mkdir(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// GetCurrentPath 获取当前路径，比如：E:/abc/data/test
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}
