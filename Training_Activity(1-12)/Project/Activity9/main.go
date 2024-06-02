package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("Activity9/index", indexHandler)

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 not found.", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(writer, "Invalid method.", http.StatusNotFound)
		return
	}
	fmt.Fprint(writer, "Hello World!")
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(writer, "POST request successful\n")
	name := request.FormValue("name")
	address := request.FormValue("address")

	fmt.Fprintf(writer, "name=%s\n", name)
	fmt.Fprintf(writer, "address=%s\n", address)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "Activity9/index/index.html" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method.", http.StatusMethodNotAllowed)
		return
	}

	file, err := os.Open("index.html")
	if err != nil {
		http.Error(w, "File not found", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		http.Error(w, "File not found", http.StatusInternalServerError)
		return
	}

	http.ServeContent(w, r, file.Name(), fi.ModTime(), file)
}
