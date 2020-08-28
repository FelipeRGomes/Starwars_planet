#!/bin/bash

# go get github.com/BurntSushi/toml gopkg.in/mgo.v2 github.com/gorilla/mux gopkg.in/snowplow/snowplow-golang-tracker.v1/tracker

go mod init statwars_planets

go run *.go

