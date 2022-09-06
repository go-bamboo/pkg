package tools

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"

	"bls/pkg/ecode"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func PathCreate(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

func PathRemove(name string) (err error) {
	if err = os.RemoveAll(name); err != nil {
		return
	}
	return
}

func FileCreate(content bytes.Buffer, name string) (err error) {
	file, err := os.Create(name)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.WriteString(content.String())
	if err != nil {
		return
	}
	// for i := n; i < content.Cap(); i++ {
	// 	//写入byte的slice数据
	// 	file.Write(content)
	// 	//写入字符串
	// 	//
	// }
	return
}

func FileRemove(name string) (err error) {
	if err = os.Remove(name); err != nil {
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

// 获取文件大小
func FileGetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// 获取文件后缀
func FileGetExt(fileName string) string {
	return path.Ext(fileName)
}

// FileCheckExist 检查文件是否存在
func FileCheckExist(src string) (bool, error) {
	_, err := os.Stat(src)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else if os.IsExist(err) {
			return true, nil
		}
		return false, err
	}
	return true, nil
}

// FileCheckPermission 检查文件权限
func FileCheckPermission(src string) (bool, error) {
	_, err := os.Stat(src)
	if err != nil {
		if os.IsPermission(err) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

//IsNotExistMkDir 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	exist, err := FileCheckExist(src)
	if err != nil {
		return err
	}
	if !exist {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}

//MkDir 新建文件夹
func MkDir(src string) error {
	err := os.Mkdir(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// 打开文件
func Open(name string, flag int, perm os.FileMode) (f *os.File, err error) {
	f, err = os.OpenFile(name, flag, perm)
	if err != nil {
		err = ecode.WrapError(err)
		return
	}
	return
}

//GetCurrentPath 获取当前路径，比如：E:/abc/data/test
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
