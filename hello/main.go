package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	log.Print(hostname)
	fmt.Fprintln(w, "こんぺこのせかいぺこ～")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
