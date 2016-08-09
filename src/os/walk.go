package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	IsDirectory = iota
	IsRegular
	IsSymlink
)

type sysFile struct {
	fType  int
	fName  string
	fLink  string
	fSize  int64
	fMtime time.Time
	fPerm  os.FileMode
}

type F struct {
	files []*sysFile
}

func (self *F) Visit(path string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	}
	var tp int
	if f.IsDir() {
		tp = IsDirectory
	} else if f.Mode()&os.ModeSymlink > 0 {
		tp = IsSymlink
	} else {
		tp = IsRegular
	}

	inoFile := &sysFile{
		fName:  path,
		fType:  tp,
		fPerm:  f.Mode(),
		fMtime: f.ModTime(),
		fSize:  f.Size(),
	}
	self.files = append(self.files, inoFile)
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	self := F{
		files: make([]*sysFile, 0),
	}

	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		return self.Visit(path, f, err)
	})

	if err != nil {
		fmt.Printf("filepath.Walk() return %v\n", err)
	}

	for _, v := range self.files {
		fmt.Println(v.fName, v.fSize)
	}
}
