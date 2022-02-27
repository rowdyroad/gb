package main

import (
	bytes2 "bytes"
	"fmt"
	"io"
	"os"
	"path"
)

func main() {
	f, err := os.Open("cmd/lesson9/text.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var b [100]byte
	for {
		if n, err := f.Read(b[:]); err == nil {
			fmt.Println(string(b[:n]))
			if n < len(b) {
				break
			}
		} else {
			break
		}
	}

	f.Seek(10, io.SeekStart)
	f.Seek(-2, io.SeekCurrent)
	n, err := f.Read(b[:])
	fmt.Println(n, err, string(b[:]))


	fw,err := os.OpenFile("cmd/lesson9/text_write1.txt", os.O_CREATE | os.O_APPEND  | os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer fw.Close()
	fmt.Println(fw.Write([]byte{'h','e','l','l','o','\n'}))

	fw.WriteString("hello\n")

	fs, _ := fw.Stat()

	fmt.Println("mode:", fs.Mode(), "name:", fs.Name(), "time:", fs.ModTime(), "dir:", fs.IsDir(), "size:", fs.Size(), "sys:", fs.Sys())


	bytes := bytes2.NewBufferString("bye")
	fw.ReadFrom(bytes)
	fw.Truncate(10)

	dirs, err := os.ReadDir("cmd")
	if err != nil {
		panic(err)
	}
	for _, dir := range dirs {
		fmt.Println(dir.Name())
	}

	fmt.Println(os.WriteFile("cmd/lesson9/text_write3.txt", []byte("byebye"), 0755))

	data, _ := os.ReadFile("cmd/lesson9/text.txt")
	fmt.Println(string(data))

	os.Remove("cmd/lesson9/text_write2.txt")

	os.Rename("cmd/lesson9/text_write.txt", "cmd/lesson9/text_renamed.txt")

	os.Mkdir("cmd/lesson9/dir", 0755)
	fmt.Println(os.MkdirAll(path.Join("cmd/lesson9", "a","c","b"),0755))

	fmt.Println("base:", path.Base("cmd/lesson9/text_write.txt"))
	fmt.Println("dir:",path.Dir("cmd/lesson9/text_write.txt"))
	fmt.Println("ext:",path.Ext("cmd/lesson9/text_write.txt"))

	fmt.Println(path.Split("cmd/lesson9/text_write.txt"))

	fmt.Println(path.Clean("cmd/lesson9/text_write.txt"))

	fmt.Println(path.IsAbs("/cmd/lesson9/text_write.txt"))

	fmt.Println(path.Match("*.tx1t", "text_write.txt"))



}
