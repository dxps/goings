#!/bin/sh

## ----------------------------------------------------------------------------
## file: scripts/run-dev.sh
## desc: It runs the app in dev (reload on code changes) capable mode.
## meta: 2018-12-20 | The classic usage.
## ----------------------------------------------------------------------------

export GOINGS_APP_ROOT=/files/dev/go/src/github.com/vision8tech/goings
export GOINGS_APP_PORT=8080

kick -appPath=${GOINGS_APP_ROOT} -gopherjsAppPath=${GOINGS_APP_ROOT}/ui -mainSourceFile=goings-app.go

