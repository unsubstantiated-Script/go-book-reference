package webDemo

import (
	"log"
	"net/http"
)

func WebAppDemo() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/arabesque", arabieeHandler)
	http.HandleFunc("/espagnole", espagnoleHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Howdy, chumps!")
}

func arabieeHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "انظر يا خبب")
}

func espagnoleHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Adios, cholos!")
}
