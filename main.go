package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World1")
}

// коммент чтоб было что коммитить
func main() {
	port := os.Getenv("PORT")
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorldHandler)
	log.Fatal(http.ListenAndServe(addr, mux))
}
