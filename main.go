package main

import (
	"bytes"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/goodbye", handleGoodbye)
	mux.HandleFunc("/hello/", handleHelloParameterized)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("HomePage!\n"))
	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}

}

func handleGoodbye(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Goodbye!\n"))
	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}

}

func handleHelloParameterized(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	userList := params["user"]

	userName := "User"
	if len(userList) > 0 {
		userName = userList[0]
	}

	var output bytes.Buffer
	output.WriteString("Hello, ")
	output.WriteString(userName)
	output.WriteString("!\n")

	_, err := w.Write(output.Bytes())
	if err != nil {
		slog.Error("error writing response body", "err", err)
		return
	}
}
