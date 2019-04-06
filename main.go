package main

import (
	"fmt"
	// "html/template"
	"io"
	"os"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, "<html><head></head><body><h1>Welcome to GaiaStack!</h1><img height=\"100\" src=\"/static/logo.png\"></body></html>")
		return
	}
}

func main() {
	logfile,err := os.OpenFile("/data/log/mylog",os.O_RDWR|os.O_CREATE,0666)
	if err != nil {
		fmt.Printf("%s\r\n",err.Error())
		os.Exit(-1)
	}
	defer logfile.Close()
	logger:=log.New(logfile,"INFO:",log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("hello^^")
	logger.Println("server begin...")

	http.HandleFunc("/", hello)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	err1 := http.ListenAndServe(":80", nil)
	if err1 != nil {
		fmt.Println("Listen And Server", err.Error())
	}
}
