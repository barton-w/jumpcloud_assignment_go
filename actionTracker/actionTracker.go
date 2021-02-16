package actionTracker

import (
  "encoding/json"
  "strings"
  "errors"
  "sync"
)

//Defining a struct with instances and timeTotals maps where action data will be set and accessed
type actionTracker struct {
  instances map[string]float32
  timeTotals map[string]float32
  mtx sync.RWMutex
}

//Exported construtor function
func New() actionTracker {
  var at actionTracker
  at.instances = make(map[string]float32)
  at.timeTotals = make(map[string]float32)
  return at
}

//Exported AddAction method
func (at actionTracker) AddAction(action string) error {
  //Define a struct to deserialize JSON into
  type actionObj struct {
    Action string
    Time float32
  }
  var act actionObj
  //Deserialize JSON argument
  byteArray := []byte(action)
  error := json.Unmarshal(byteArray, &act)

  //Throw a custom error in the event we're unable to deserialize
  if error != nil {
    return errors.New("unprocessable JSON argument")
  }

  at.mtx.Lock()
  //Maintain a running count of instances of each action
  at.instances[act.Action] += 1
  //Maintain a cumulative total time for each action
  at.timeTotals[act.Action] += act.Time
  at.mtx.Unlock()

  return error
}

//Exported GetStatistics method
func (at actionTracker) GetStatistics() (string, error) {
  //Define a struct for JSON serialization
  type statsObj struct {
    Action string
    Avg float32
  }
  //Establish a slice to collect JSON strings
  var averages []string

  //Iterate over timeTotals, compute averages, and create JSON strings
  at.mtx.RLock()
  for action, time := range at.timeTotals {
    var avg statsObj
    avg.Action = action
    avg.Avg = time / at.instances[action]
    byte, error := json.Marshal(avg)

    if error != nil {
      at.mtx.RUnlock()
      return "", error
    }
    //Add JSON string to averages slice
    averages = append(averages, string(byte))
  }
  at.mtx.RUnlock()

  //Return formatted JSON
  returnJson := "[" + strings.Join(averages, ",") + "]"
  return returnJson, nil
}
