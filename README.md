# Goings

Goings is my journey into the _isomorphic Go_ space.


## Run

Prereqs:
- Go 1.x
- resolve the dependencies by running `go get` in the project location

Run the app:
- in dev mode using `go run goings-app.go`
    - as currently it is using two environment variables for the initial setup, use as an example:
      `GOINGS_APP_ROOT=/files/dev/go/src/github.com/vision8tech/goings GOINGS_APP_PORT=8080 go run goings-app.go`

## Usage

Get the list of projects using `curl -X GET http://localhost:8080/api/projects`

