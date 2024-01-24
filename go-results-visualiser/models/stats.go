package models

import "fmt"

type ScenarioStats struct {
	NrOfConn    uint64  `csv:"Connections" json:"nr-of-connections"`
	NrOfReq     uint64  `csv:"NrOfReq" json:"nr-of-req"`
	TimeTaken   float64 `csv:"TimeTaken" json:"time-taken"`
	LatencyMean float64 `csv:"Latency.Mean" json:"latency-mean"`
	LatencyMax  float64 `csv:"Latency.Max" json:"latency-max"`
	ReqMean     float64 `csv:"Req.Mean" json:"req-mean"`
	ReqMax      float64 `csv:"Req.Max" json:"req-max"`
	Req1xx      uint64  `csv:"Req2XX" json:"req2xx"`
	Others      uint64  `csv:"Other" json:"other"`
}

func (scenarioStats *ScenarioStats) String() string {
	return fmt.Sprintf("%d %d %s %s %s %f %f %d %d",
		scenarioStats.NrOfConn,
		scenarioStats.NrOfReq,
		formatTimeUs(scenarioStats.TimeTaken),
		formatTimeUs(scenarioStats.LatencyMean),
		formatTimeUs(scenarioStats.LatencyMax),
		scenarioStats.ReqMean,
		scenarioStats.ReqMax,
		scenarioStats.Req1xx,
		scenarioStats.Others)
}

type ServerStats struct {
	ServerName string
	Stats      *ScenarioStats
}

func (serverStat ServerStats) String() string {
	return fmt.Sprintf("ServerName: %s, stats: %v",
		serverStat.ServerName,
		serverStat.Stats)
}

type Scenario struct {
	Name  string
	Stats []ServerStats
}

func (scenario Scenario) String() string {
	return fmt.Sprintf("ScenarioName: %s, Stats: %v", scenario.Name, scenario.Stats)
}
