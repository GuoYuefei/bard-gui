#!/bin/sh

go build cmd/main.go
nohup ./main > ./nohup.out 2>&1 & echo $! > ./run.pid