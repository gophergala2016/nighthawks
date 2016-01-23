package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/eknkc/amber"
	// "github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Nighthawks server 1.0")

	r := mux.NewRouter()

	// front page
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl, _ := amber.CompileFile("amber/index.amber", amber.DefaultOptions)
		tpl.Execute(w, nil)
	})

	// static files
	r.HandleFunc("/static/{type}/{file}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		filepath := "static/" + params["type"] + "/" + params["file"]
		_, err := os.Stat(filepath)
		if err != nil {
			handleNotFound(w, r)
			return
		}

		http.ServeFile(w, r, "static/"+params["type"]+"/"+params["file"])
	})

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {

}
