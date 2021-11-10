# covid-meter
This project has an api development that retrives total covid cases in a state of India.

## Setup

* Install Go (see go.mod file for version). 
* Install docker.
* Get Api key for locationiq.com api access.
* Clone project.
* cd into project root directory.
* Run these command:
```
$ docker run -p 27017:27017 mongo
$ export ACCESS_KEY_IQ=PASTE_YOUR_LOCATIONIQ_API_KEY
$ go mod tidy
$ go run .
```
* open localhost:1323/swagger/
