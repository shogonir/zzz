package main

import (
  "fmt"
  "time"
  "flag"
  "net/http"
  "strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
  ms := r.URL.Query().Get("ms")
  msInt, _ := strconv.Atoi(ms)
  if (msInt > 20000) {
    fmt.Fprint(w, "Bad Request")
  } else {
    time.Sleep(time.Duration(msInt) * time.Millisecond)
    fmt.Fprint(w, "Good morning.")
  }
}

func main() {

  port := flag.Int("port", 7999, "listen port number")
  flag.Parse()

  http.HandleFunc("/zzz/", handler)

  err := http.ListenAndServe(":" + strconv.Itoa(*port), nil)
  if err != nil {
    fmt.Println(err)
  }
}
