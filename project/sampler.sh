#!/bin/bash

CONTAINER=$1
FORMAT='{{.CPUPerc}}|{{.MemPerc}}|{{.MemUsage}}|{{.Name}}'

docker stats --format $FORMAT $CONTAINER | sed -u 's/\x1b\[[0-9;]*[a-zA-Z]//g' | tee $2