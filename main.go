package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

// коммент чтоб было что коммитить
func main(){
	port := os.Getenv("PORT")
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorldHandler)
	log.Fatal(http.ListenAndServe(addr, mux))
}
