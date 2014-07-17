package main

import(
  "io"
  "log"
  "net/http"
  "os"
)

func main() {
  if len(os.Args) < 2 {
    log.Fatal("Usage: cm-nextbus StopID")
  }

  stopID := os.Args[1]

  url := "http://www.capmetro.org/planner/s_nextbus2.asp?stopid=" + stopID + "&opt=2"

  resp, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }

  if resp.StatusCode != http.StatusOK {
    log.Fatal(resp.Status)
  }

  _, err = io.Copy(os.Stdout, resp.Body)

  if err != nil {
    log.Fatal(err)
  }

}
