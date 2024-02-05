#!/bin/bash

CONTAINER=$1
FORMAT='{{.CPUPerc}}|{{.MemPerc}}|{{.MemUsage}}|{{.Name}}'

touch ./$2
docker stats --format $FORMAT $CONTAINER | sed -u 's/\x1b\[[0-9;]*[a-zA-Z]//g' | sed -u 's/[//]/|/g' | sed -u 's/ //g' | tee ./$2