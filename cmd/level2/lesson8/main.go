package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var path string
	var listen string
	flag.StringVar(&path, "path", "/rand", "/rand")
	flag.StringVar(&listen, "address", ":80", "127.0.0.1:80")
	flag.Parse()

	http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Request ", request.URL, "Method:", request.Method,"From:", request.RemoteAddr)
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r, err := strconv.Atoi(request.URL.Query().Get("a")); err != nil {
			writer.WriteHeader(http.StatusInternalServerError);
			return
		} else {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(strconv.Itoa(r+1)))
		}
	})

	http.HandleFunc("/admin", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Query().Get("path")
		info, err := os.Stat(path)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return;
		}
		if info.IsDir() {
			files, err  := os.ReadDir(path)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			for _, file :=  range files {
				writer.Write([]byte(fmt.Sprintf("%s %v\n", file.Name(), file.Type())))
			}
		} else {
			f, err := os.Open(path)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer f.Close()
			io.Copy(writer, f)
		}

	})
	http.ListenAndServe(listen, nil)

}
