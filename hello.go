package main
import (
	"io"
	"log"
	"net/http"
	)
func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!\n")
	io.WriteString(w, "Hello, Dream!")
}
func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}