package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"./handler"
)

const (
	port = "7777"
)

//webroot = flag.String("root", os.Getenv("PWD"), "web root directory")
var (
	webroot = flag.String("root", "./", "web root directory")
	user    = []byte("go")
	passwd  = []byte("go")
)

func main() {
	p := port

	if len(os.Args) == 2 {
		p = os.Args[1]
	}

	http.HandleFunc("/video", handler.BasicAuth(handler.VideoHandler, user, passwd))
	http.HandleFunc("/encode", handler.BasicAuth(handler.EncoMp3Handler, user, passwd))
	http.HandleFunc("/download/", handler.BasicAuth(handler.FileServerHandle(*webroot, "/download/"), user, passwd))

	err := http.ListenAndServe(":"+p, nil)
	if err != nil {
		log.Panicln("ListenAndServe:", err)
	}
}
