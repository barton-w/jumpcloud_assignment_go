package actionTracker

import (
  "testing"
  "fmt"
  "encoding/json"
)

//Struct used for test-case JSON parsing
type resultObj struct {
  Action string
  Avg float32
}

//TestCase1 - matching the assignment criteria
func TestCase1(t *testing.T) {
  //Instantiate an actionTracker and pass data to its AddAction method
  at := New()
  action := `{"action":"jump","time":100}`
  at.AddAction(action)
  action = `{"action":"run","time":75}`
  at.AddAction(action)
  action = `{"action":"jump","time":200}`
  at.AddAction(action)

  //Define the expected end result as JSON
  expected := `[{"Action":"jump","Avg":150},{"Action":"run","Avg":75}]`

  //Call GetStatistics
  result, error := at.GetStatistics()

  //Log expected and result
  fmt.Println("Expected: ", expected)
  fmt.Println("Result: ", result)

  //Parse JSON from both expected and result
  var resultSlc []resultObj
  var expectedSlc []resultObj
  _ = json.Unmarshal([]byte(result), &resultSlc)
  _ = json.Unmarshal([]byte(expected), &expectedSlc)

  //Using the compare function to check test case accuracy.
  //The order of objects within the JSON arrays could be different between expected and result, but contain the same data
  if error != nil || !compare(resultSlc, expectedSlc) {
    t.Fatalf(`GetStatistics() fail. %#q, %v, should match: %#q, nil`, result, error, expected)
  }
}

//TestCase2 - testing a bunch of actions
func TestCase2(t *testing.T) {
  //Instantiate an actionTracker
  at := New()

  //Create a slice with a bunch of actions and call AddAction 20x for each, with time 1-20
  actions := []string {"climb","sprint","flip","drum","fish"}
  for _, action := range actions {
    for i := 1; i <=20; i++ {
      //Set a JSON string as the argument via concatenation
      actionArg := `{"action":"` + action + `","time":` + fmt.Sprint(i) + `}`
      at.AddAction(actionArg)
    }
  }

  //Define the expected end result as JSON - average of 1-20 would be 10.5 for each action
  expected := `[{"Action":"drum","Avg":10.5},{"Action":"fish","Avg":10.5},{"Action":"climb","Avg":10.5},{"Action":"sprint","Avg":10.5},{"Action":"flip","Avg":10.5}]`

  //Call GetStatistics
  result, error := at.GetStatistics()

  //Log expected and result
  fmt.Println("Expected: ", expected)
  fmt.Println("Result: ", result)

  //Parse JSON from both expected and result
  var resultSlc []resultObj
  var expectedSlc []resultObj
  _ = json.Unmarshal([]byte(result), &resultSlc)
  _ = json.Unmarshal([]byte(expected), &expectedSlc)

  //Using the compare function to check test case accuracy.
  //The order of objects within the JSON arrays could be different between expected and result, but contain the same data
  if error != nil || !compare(resultSlc, expectedSlc) {
    t.Fatalf(`GetStatistics() fail. %#q, %v, should match: %#q, nil`, result, error, expected)
  }
}

//Testing bad data
func TestErrorCase(t *testing.T) {
  at := New()
  action := `This is not JSON`
  err := at.AddAction(action)
  if err == nil {
    t.Fatal("AddAction did not result in a validation error")
  }
}


func compare(rSlc []resultObj, eSlc []resultObj) bool {
  matchCount := 0
  for _, eitem := range eSlc {
    for _, ritem := range rSlc {
      if eitem == ritem {
        matchCount ++
      }
    }
  }
  if matchCount == len(eSlc) {
    return true
  }
  return false
}
