package tools

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type ReplaceHelper struct {
	Root    string //路径
	OldText string //需要替换的文本
	NewText string //新的文本
}

func (h *ReplaceHelper) DoWalk() error {
	return filepath.Walk(h.Root, h.walkCallback)
}

func (h *ReplaceHelper) walkCallback(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if f == nil {
		return nil
	}
	if f.IsDir() {
		// log.Println("DIR:", path)
		return nil
	}

	//文件类型需要进行过滤
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(buf)

	//替换
	newContent := strings.Replace(content, h.OldText, h.NewText, -1)

	//重新写入
	ioutil.WriteFile(path, []byte(newContent), 0)

	return err
}
