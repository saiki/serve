package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
)

var directory string
var address string
var port int

func init() {
	pwd, _ := os.Getwd()
	flag.StringVar(&directory, "d", pwd, "http serve directory.")
	flag.StringVar(&address, "a", "", "http listen url.")
	flag.IntVar(&port, "p", 80, "http listen port.")
}

func main() {
	flag.Parse()
	fileServer := http.FileServer(http.Dir(directory))
	http.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(address+":"+strconv.Itoa(port), nil))
}
