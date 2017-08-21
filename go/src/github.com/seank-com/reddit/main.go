package main

import (
  "io"
  "log"
  "net/http"
  "os"
)

func main() {
  client := &http.Client{}

  req, err := http.NewRequest("GET", "http://reddit.com/r/golang.json", nil)
  if err != nil {
    log.Fatal(err)
  }

  req.Header.Set("User-Agent", "seank-com")

  resp, err := client.Do(req)
  if err != nil {
    log.Fatalln(err)
  }

  if resp.StatusCode != http.StatusOK {
    log.Fatal(resp.Status)
  }

  _, err = io.Copy(os.Stdout, resp.Body)
  if err != nil {
    log.Fatal(err)
  }
}
