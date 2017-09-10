package main

import (
	"net/http"
	"fmt"
	"log"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

// коммент чтоб было что коммитить
func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorldHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:80", mux))
}
