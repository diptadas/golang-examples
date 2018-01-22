package abe

type ProtobufDuration struct {
	Seconds string `json:"seconds,omitempty"`
	Nanos   int32  `json:"nanos,omitempty"`
}
