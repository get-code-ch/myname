package zonerecord

import (
	"regexp"
	"strconv"
)

type RDataMX struct {
	preference int16
	exchange   string
}

func (mx *RDataMX) set(raw []byte) {
	var err error
	var num int64

	mxRegex := regexp.MustCompile(`(?i)^.*IN\s+MX\s+(\d*)\s+([^;\s]*)`)
	mxRecord := mxRegex.FindSubmatch(raw)

	if num, err = strconv.ParseInt(string(mxRecord[1]), 10, 16); err == nil {
		mx.preference = int16(num)
	}

	mx.exchange = string(mxRecord[2])
}

func (mx *RDataMX) string(zr ZoneRecord) string {
	recordStr := zr.name + " \t" + zr.ttl + " \t" + zr.class + " \t" + zr.recType + " \t"
	recordStr += strconv.FormatInt(int64(mx.preference), 10) + " \t"
	recordStr += mx.exchange
	return recordStr
}
