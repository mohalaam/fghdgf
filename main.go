package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", HelloHandler)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, """{
  "Description": "Fichier de vérification de la propriété du domaine pour Microsoft 365 : placer dans la racine du site web",
  "Domain": "inappropriate-mahalia-csdfzs-6fb1f4b8.koyeb.app",
  "Id": "f58a1e05-6512-4491-9510-7b16670249fe"" /n")
}"""
