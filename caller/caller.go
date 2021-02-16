package main

import (
  "fmt"
  "actionTracker"
)

func main() {
  //Generate an instance of actionTracker
  at := actionTracker.New()

  //Call AddAction with data as supplied in the assignment spec
  action := "{\"action\":\"jump\",\"time\":100}"
  at.AddAction(action)
  action = "{\"action\":\"run\",\"time\":75}"
  at.AddAction(action)
  action = "{\"action\":\"jump\",\"time\":200}"
  at.AddAction(action)

  //Call GetStatistics to return and print averages
  response, error := at.GetStatistics()
  if error != nil {
    fmt.Println(error)
  }
  fmt.Println(response)
}
