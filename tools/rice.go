package tools

import (
	"bytes"
	"github.com/GeertJohan/go.rice/embedded"
	"io"
	"os"
)

func RiceWirteToLocal(key string) {

	ef, err := os.Open(key)
	if os.IsExist(err) || ef != nil {
		return
	}
	err = os.MkdirAll(key, os.ModePerm)
	if err != nil {
		return
	}

	data, ok := embedded.EmbeddedBoxes[key]
	if !ok {
		return
	}
	for _, v := range data.Dirs {
		riceWriteDir([]*embedded.EmbeddedDir{v}, key)
	}
	for _, v := range data.Files {
		riceWriteFile([]*embedded.EmbeddedFile{v}, key)
	}
}
func riceWriteDir(dir []*embedded.EmbeddedDir, prefix string) {
	for _, v := range dir {
		os.MkdirAll(prefix+"/"+v.Filename, os.ModePerm)
	}
}
func riceWriteFile(file []*embedded.EmbeddedFile, prefix string) {
	for _, v := range file {
		f, err := os.Create(prefix + "/" + v.Filename)
		if err == nil {
			io.Copy(f, bytes.NewBuffer([]byte(v.Content)))
		}
		f.Close()
	}
}
