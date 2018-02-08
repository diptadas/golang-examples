package server

import (
	"bytes"
	"encoding/json"
	"net"

	. "github.com/appscode/g2/pkg/runtime"
)

const (
	wsRunning         = 1
	wsSleep           = 2
	wsPrepareForSleep = 3
)

func status2str(status int) string {
	switch status {
	case wsRunning:
		return "running"
	case wsSleep:
		return "sleep"
	case wsPrepareForSleep:
		return "prepareForSleep"
	}

	return "unknown"
}

type Worker struct {
	net.Conn
	Session

	workerId    string
	status      int
	runningJobs map[string]*Job
	canDo       map[string]int32
}

func (w *Worker) MarshalJSON() ([]byte, error) {
	b := &bytes.Buffer{}
	enc := json.NewEncoder(b)
	m := make(map[string]interface{})
	m["sessionId"] = w.SessionId
	m["Id"] = w.workerId
	m["status"] = status2str(w.status)

	type FuncWithDuration struct {
		FunctionName string `json:"function_name,omitempty"`
		Duration     int32  `json:"timeout_duration,omitempty"`
	}
	canDoSlice := make([]FuncWithDuration, 0, len(w.canDo))
	for k, d := range w.canDo {
		canDoSlice = append(canDoSlice, FuncWithDuration{k, d})
	}
	m["canDo"] = canDoSlice
	jobSlice := make([]string, 0, len(w.canDo))
	for k := range w.runningJobs {
		jobSlice = append(jobSlice, k)
	}
	m["runningJobs"] = jobSlice

	if err := enc.Encode(m); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
