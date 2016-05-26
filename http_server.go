/*
*	Minimalist HTTP File Server with customisation
*	by Gau Bac Cuc 2016
*/

package main

import (
  "log"
  "net/http"
  "fmt"
)

var fs = http.FileServer(http.Dir("./"))

func foo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintf(w, "<http>\n<body>\n")
	fs.ServeHTTP(w,r)
	fmt.Fprintf(w, "</body>\n</html>")
}

func main() {
  http.HandleFunc("/", foo)

  log.Println("Listening...")
  http.ListenAndServe(":8080", nil)
}
