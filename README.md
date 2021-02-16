# JumpCloud Programming Assignment - Go

## Overview
This implements a module called actionTracker which enables an application to initialize a struct that has 2 key methods:

1. AddAction - accepts a single key-value pair of action and time, formatted as JSON
* For example: {"action":"cartwheel", "time":5}
* The addAction method maintains a running count of instance of each action, and a time total so averages can be easily calculated.

2. GetStatistics
* Calculates the average time for each action submitted via AddAction, and returns JSON
* For example: [{"Action":"jump","Avg":150},{"Action":"run","Avg":75}]

## Steps to get started:
1. Make sure you have Go installed.
* For details, [click here](https://golang.org/doc/install)

2. git clone git@github.com:barton-w/jumpcloud_assignment_go.git

3. cd jumpcloud_assignment_go/actionTracker/

4. go test -v
* Feel free to check out the unit-test file actionTracker_test.go

5. Alternatively, you can also run a simple calling program which imports the actionTracker module:
* cd ../caller/
* go run caller.go

6. Enjoy! And please feel free to reach out to me with any questions.
