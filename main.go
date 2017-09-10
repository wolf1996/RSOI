package main

import (
	"net/http"
	"fmt"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

// коммент чтоб было что коммитить
func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe("0.0.0.0:80", mux)
}
