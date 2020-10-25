package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Data struct {
	Base64EncodedJPG string
}

// AJAX Request Handler
func ajaxHandler(w http.ResponseWriter, r *http.Request) {
	var d Data

	f, _ := os.Open("./front/image/tesla.jpg")

	// Read tesla.jpg with byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	d.Base64EncodedJPG = base64.StdEncoding.EncodeToString(content)

	// Create json response from struct
	a, err := json.Marshal(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(a)
	f.Close()
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./front")))
	http.HandleFunc("/ajax", ajaxHandler)
	http.ListenAndServe(":8080", nil)
}
