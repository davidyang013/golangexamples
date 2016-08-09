package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func TarGz(src, dest string) {
	fw, err := os.Create(dest)
	checkErr(err)
	defer fw.Close()

	//Gzip
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	//Tar
	tw := tar.NewWriter(gw)
	defer tw.Close()

	//check file or directory
	f, err := os.Open(src)
	checkErr(err)

	fi, err := f.Stat()
	checkErr(err)

	if fi.IsDir() {
		fmt.Println("Create tar.gz from a directory...")
		tarGzDir(src, filepath.Base(src), tw)
	} else {
		fmt.Println("Create tar.gz from " + fi.Name() + "...")
		tarGzFile(src, fi.Name(), tw, fi)
	}
	fmt.Println("compress finished!")
}

//rec is the path inside of the tar.gz
func tarGzDir(src, rec string, tw *tar.Writer) {
	dir, err := os.Open(src)
	checkErr(err)
	defer dir.Close()

	//Get file info
	fis, err := dir.Readdir(0)
	checkErr(err)

	for _, fi := range fis {
		cur := src + "/" + fi.Name()

		if fi.IsDir() {
			fmt.Printf("Adding path ...%s\n", cur)
			tarGzDir(cur, rec+"/"+fi.Name(), tw)
		} else {
			fmt.Printf("Adding file .../%s\n", cur)
		}
		tarGzFile(cur, rec+"/"+fi.Name(), tw, fi)
	}
}

func tarGzFile(src, rec string, tw *tar.Writer, fi os.FileInfo) {
	if !fi.IsDir() {
		fr, err := os.Open(src)
		checkErr(err)
		defer fr.Close()

		//Create tar header
		hdr := new(tar.Header)
		hdr.Name = rec
		hdr.Size = fi.Size()
		hdr.Mode = int64(fi.Mode())
		hdr.ModTime = fi.ModTime()

		//Writer
		err = tw.WriteHeader(hdr)
		checkErr(err)

		//Write file data
		_, err = io.Copy(tw, fr)
		checkErr(err)
	}
}

func UnTarGz(src, dest string) {
	fmt.Println("Uncompress file from " + src)
	os.Mkdir(dest, os.ModePerm)

	fr, err := os.Open(src)
	checkErr(err)
	defer fr.Close()

	//Gzip
	gr, err := gzip.NewReader(fr)
	checkErr(err)
	defer gr.Close()

	tr := tar.NewReader(gr)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if hdr.Typeflag != tar.TypeDir {
			os.MkdirAll(dest+"/"+filepath.Dir(hdr.Name), os.ModePerm)
			//writer
			path := dest + "/" + hdr.Name
			fw, _ := os.Create(path)
			checkErr(err)
			_, err = io.Copy(fw, tr)
			fmt.Println("Uncompress file to " + path + "....")
			checkErr(err)
		}
	}
	fmt.Println("unpacking finshed!")
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	// TarGz(root, "c:\\Go\\test.tar.gz")
	UnTarGz("c:\\Go\\test.tar.gz", root)
}
