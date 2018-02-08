package worker

import (
	"bytes"
	"testing"

	rt "github.com/appscode/g2/pkg/runtime"
)

var (
	outpackcases = map[rt.PT]map[string]string{
		rt.PT_CanDo: {
			"src":  "\x00REQ\x00\x00\x00\x01\x00\x00\x00\x01a",
			"data": "a",
		},
		rt.PT_CanDoTimeout: {
			"src":  "\x00REQ\x00\x00\x00\x17\x00\x00\x00\x06a\x00\x00\x00\x00\x01",
			"data": "a\x00\x00\x00\x00\x01",
		},
		rt.PT_CantDo: {
			"src":  "\x00REQ\x00\x00\x00\x02\x00\x00\x00\x01a",
			"data": "a",
		},
		rt.PT_ResetAbilities: {
			"src": "\x00REQ\x00\x00\x00\x03\x00\x00\x00\x00",
		},
		rt.PT_PreSleep: {
			"src": "\x00REQ\x00\x00\x00\x04\x00\x00\x00\x00",
		},
		rt.PT_GrabJob: {
			"src": "\x00REQ\x00\x00\x00\x09\x00\x00\x00\x00",
		},
		rt.PT_GrabJobUniq: {
			"src": "\x00REQ\x00\x00\x00\x1E\x00\x00\x00\x00",
		},
		rt.PT_WorkData: {
			"src":  "\x00REQ\x00\x00\x00\x1C\x00\x00\x00\x03a\x00b",
			"data": "a\x00b",
		},
		rt.PT_WorkWarning: {
			"src":  "\x00REQ\x00\x00\x00\x1D\x00\x00\x00\x03a\x00b",
			"data": "a\x00b",
		},
		rt.PT_WorkStatus: {
			"src":  "\x00REQ\x00\x00\x00\x0C\x00\x00\x00\x08a\x0050\x00100",
			"data": "a\x0050\x00100",
		},
		rt.PT_WorkComplete: {
			"src":  "\x00REQ\x00\x00\x00\x0D\x00\x00\x00\x03a\x00b",
			"data": "a\x00b",
		},
		rt.PT_WorkFail: {
			"src":    "\x00REQ\x00\x00\x00\x0E\x00\x00\x00\x01a",
			"handle": "a",
		},
		rt.PT_WorkException: {
			"src":  "\x00REQ\x00\x00\x00\x19\x00\x00\x00\x03a\x00b",
			"data": "a\x00b",
		},
		rt.PT_SetClientId: {
			"src":  "\x00REQ\x00\x00\x00\x16\x00\x00\x00\x01a",
			"data": "a",
		},
		rt.PT_AllYours: {
			"src": "\x00REQ\x00\x00\x00\x18\x00\x00\x00\x00",
		},
	}
)

func TestOutPack(t *testing.T) {
	for k, v := range outpackcases {
		outpack := getOutPack()
		outpack.dataType = k
		if handle, ok := v["handle"]; ok {
			outpack.handle = handle
		}
		if data, ok := v["data"]; ok {
			outpack.data = []byte(data)
		}
		data := outpack.Encode()
		if bytes.Compare([]byte(v["src"]), data) != 0 {
			t.Errorf("%d: %X expected, %X got.", k, v["src"], data)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, v := range outpackcases {
			outpack := getOutPack()
			outpack.dataType = k
			if handle, ok := v["handle"]; ok {
				outpack.handle = handle
			}
			if data, ok := v["data"]; ok {
				outpack.data = []byte(data)
			}
			outpack.Encode()
		}
	}
}
