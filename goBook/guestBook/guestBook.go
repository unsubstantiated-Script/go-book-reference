package guestBook

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func GuestBook() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new-sig", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("goBook/guestBook/signatures.txt")
	fmt.Printf("%#v\n", signatures)
	html, err := template.ParseFiles("goBook/guestBook/view.html")
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(writer, guestbook)
	check(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("goBook/guestBook/signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}
func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("goBook/guestBook/new-sig.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}
