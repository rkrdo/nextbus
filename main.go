package main

import(
  "log"
  "net/http"
  "os"
  "encoding/json"
  "fmt"
)

type Response struct {
  RouteDesc string
  StopDesc string
  List []struct {
    Sched string
    Est string
  }
}

func printResults(r Response) {
  fmt.Println("Route:", r.RouteDesc)
  fmt.Println("Stop:", r.StopDesc)
  if len(r.List) > 0 {
    for i, list := range r.List {
      fmt.Printf("Trip #%d\n", i+1)
      fmt.Printf("\tScheduled Time: %s\n", list.Sched)
      fmt.Printf("\tEstimated Time: %s\n", list.Est)
    }
  } else {
    fmt.Println("No available trips found")
  }
}

func checkArgs(args []string) {
  if len(args) < 2 {
    log.Fatal("Usage: cm-nextbus StopID")
  }
}

func main() {

  checkArgs(os.Args)

  stopID := os.Args[1]

  url := "http://www.capmetro.org/planner/s_nextbus2.asp?stopid=" + stopID + "&opt=2"

  resp, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }

  if resp.StatusCode != http.StatusOK {
    log.Fatal(resp.Status)
  }

  r := new(Response)
  err = json.NewDecoder(resp.Body).Decode(r)

  if err != nil {
    log.Fatal(err)
  }

  printResults(*r)
}
