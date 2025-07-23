package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hola desde el servidor Go!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor Go escuchando en el puerto 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
} 