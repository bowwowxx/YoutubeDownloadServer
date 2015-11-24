package handler

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)

var (
	singlepath = "https://www.youtube.com/watch?v="
	listpath   = "https://www.youtube.com/playlist?list="
	BUF_LEN    = 1024
)

type ViewFunc func(http.ResponseWriter, *http.Request)

func DefHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "oops", 404)
}

func BasicAuth(f ViewFunc, user, passwd []byte) ViewFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		basicAuthPrefix := "Basic "
		auth := r.Header.Get("Authorization")

		if strings.HasPrefix(auth, basicAuthPrefix) {

			payload, err := base64.StdEncoding.DecodeString(
				auth[len(basicAuthPrefix):],
			)
			if err == nil {
				pair := bytes.SplitN(payload, []byte(":"), 2)
				if len(pair) == 2 && bytes.Equal(pair[0], user) &&
					bytes.Equal(pair[1], passwd) {

					f(w, r)
					return
				}
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="Login Check"`)
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func VideoHandler(w http.ResponseWriter, r *http.Request) {
	rq := r.URL.RawQuery
	qs, _ := url.ParseQuery(rq)
	ur := singlepath
	kind := qs.Get("list")
	id := qs.Get("v")

	if len(id) > 0 {
		ur = singlepath + id
	}

	if len(kind) > 0 {
		ur = listpath + kind
	}

	mp3_link := exec.Command("youtube-dl", "-i", "-c", ur)
	go pippCmd(w, mp3_link)
	fmt.Fprintf(w, "Successfully....\n")
}

func FileServerHandle(dir, prefix string) ViewFunc {
	fs := http.FileServer(http.Dir(dir))
	realHandler := http.StripPrefix(prefix, fs).ServeHTTP
	return func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.URL)
		realHandler(w, req)
	}
}

func EncoMp3Handler(w http.ResponseWriter, r *http.Request) {
	rq := r.URL.RawQuery
	qs, _ := url.ParseQuery(rq)
	ur := singlepath
	kind := qs.Get("list")
	id := qs.Get("v")
	// fmt.Println("youtube id:", id, len(id))
	if len(id) > 0 {
		ur = singlepath + id
	}

	if len(kind) > 0 {
		ur = listpath + kind
	}

	// } else {
	// 	http.Error(w, "YouTubeId is null", 500)
	// 	return
	// }

	mp3_link := exec.Command("youtube-dl", "-f", "140", "-x", "--audio-format", "mp3", "-i", "-c", ur)
	go pippCmd(w, mp3_link)
	fmt.Fprintf(w, "Successfully....\n")
}

func pippCmd(wr http.ResponseWriter, mp3_link *exec.Cmd) {
	pipeReader, pipeWriter := io.Pipe()
	mp3_link.Stdout = pipeWriter
	mp3_link.Stderr = pipeWriter
	go writeCmdOutput(wr, pipeReader)
	mp3_link.Run()
	pipeWriter.Close()
}

func writeCmdOutput(res http.ResponseWriter, pipeReader *io.PipeReader) {
	buffer := make([]byte, BUF_LEN)
	for {
		n, err := pipeReader.Read(buffer)
		if err != nil {
			pipeReader.Close()
			break
		}

		data := buffer[0:n]
		//res.Write(data)
		fmt.Println(string(data))
		if f, ok := res.(http.Flusher); ok {
			f.Flush()
		}
		//reset buffer
		for i := 0; i < n; i++ {
			buffer[i] = 0
		}
	}
}
