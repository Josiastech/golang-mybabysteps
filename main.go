package main

import (
  "fmt" // implements formated I/O with functions analogus to C's printf and scanf
  "net/http" // Package http provides HTTP client and server implementation
)

func main(){
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, "Hello, Go web development")
  })

  fmt.Println(http.ListenAndServe(":8080", nil))
}
