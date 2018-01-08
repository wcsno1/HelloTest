package main
import (
	"io"
	"log"
	"net/http"
	)
func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, My Dear Sister\n")
	io.WriteString(w, "Good luck and have a nice day,from WuChao")
}
func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}