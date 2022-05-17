package mine

import "time"

func Now(format ...string) string {
	if len(format) > 0 {
		return time.Now().In(time.FixedZone(TimeZoneName, TimeZoneValue)).Format(format[0])
	}
	return time.Now().In(time.FixedZone(TimeZoneName, TimeZoneValue)).Format(TimeFormart)
}

func Time() int64 {
	return time.Now().Unix()
}
