package mine

import (
	"encoding/json"
	"time"
)

type Return struct {
	Status   bool `json:"status"`
	Code     any  `json:"code"`
	Msg      any  `json:"msg"`
	Data     any  `json:"data"`
	Datetime any  `json:"datetime,omitempty"`
}

func Succ(a ...any) (r Return) {
	r.Status = true
	if len(a) > 0 {
		r.Data = a[0]
	}
	if len(a) > 1 {
		r.Code = a[1]
	}
	if len(a) > 2 {
		r.Msg = a[2]
	} else {
		r.Msg = "successful"
	}
	if len(a) > 3 {
		r.Datetime = a[3]
	} else {
		r.Datetime = time.Now().In(time.FixedZone(TimeZoneName, TimeZoneValue)).Format(TimeFormart)
	}
	return
}

func Err(a ...any) (r Return) {
	if len(a) > 0 {
		r.Code = a[0]
	}
	if len(a) > 1 {
		r.Msg = a[1]
	} else {
		r.Msg = "failed"
	}
	if len(a) > 2 {
		r.Data = a[2]
	}
	if len(a) > 3 {
		r.Datetime = a[3]
	} else {
		r.Datetime = time.Now().In(time.FixedZone(TimeZoneName, TimeZoneValue)).Format(TimeFormart)
	}
	return
}

func (r Return) JsonEnCode() string {
	marshal, _ := json.Marshal(r)
	return string(marshal)
}
