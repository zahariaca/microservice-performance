### Goal 
Simple comparisson of Spring Boot Rest vs Webflux, also benefits of virtual threads on non reactive spring.
Go thrown in out of curiosity

### Conditions 

Test Machine: AMD Ryzen 7 3700x 8-Core Processor, 32GB Ram

All tests run with 2 docker containers: server + postgresql database

Tests executed using: https://github.com/codesenberg/bombardier

### Results:
Results can be seen at: https://zahariaca.github.io/microservice-performance/

### Test Scenario 1: Simeple GET, server return static string. 6 cpu cores, 1Gb RAM

#### 25 connections, 500k requests


##### GET tests:
E.g command used
```sh
bombardier.exe --fasthttp -c 25 -n 500000 -l http://docker:8080/hello -t 10s --format=j -p r | jq -r '["Connections","NrOfReq","TimeTaken", "Latency.Mean", "Latency.Max", "Req.Mean", "Req.Max", "Req2XX", "Other"], (. | [.spec.numberOfConnections, .spec.numberOfRequests, .result.timeTakenSeconds, .result.latency.mean, .result.latency.max, .result.rps.mean, .result.rps.max, .result.req2xx, .result.others]) | @csv' > go.csv
```


```sh
bombardier.exe --fasthttp -m POST -H "Content-Type: application/json" -f .\body.json -c 100 -n 100 -l http://docker:8080/add -t 10s --format=j -p r | jq -r '[\"TimeTaken\", \"Latency.Mean\", \"Latency.Max\", \"Req.Mean\", \"Req.Max\", \"Req2XX\", \"Other\"], (. | [.result.timeTakenSeconds, .result.latency.mean, .result.latency.max, .result.rps.mean, .result.rps.max, .result.req2xx, .result.others]) | @csv'
```