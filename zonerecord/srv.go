package zonerecord

import (
	"regexp"
	"strconv"
)

type RDataSRV struct {
	service  string
	proto    string
	priority uint16
	weight   uint16
	port     uint16
	target   string
}

func (srv *RDataSRV) set(raw []byte) {
	var err error
	var num int64

	// [1] - service [2] - proto [3] - priority [4] - weight
	// [5] - port [6] - target
	srvRegex := regexp.MustCompile(`(?i)^(_[^.]*)\.(_[^.]*).*IN\s+SRV\s+(\d+)\s+(\d+)\s+(\d+)\s+([^;\s]*)`)
	srvRecord := srvRegex.FindSubmatch(raw)

	srv.service = string(srvRecord[1])
	srv.proto = string(srvRecord[2])

	if num, err = strconv.ParseInt(string(srvRecord[3]), 10, 16); err == nil {
		srv.priority = uint16(num)
	}
	if num, err = strconv.ParseInt(string(srvRecord[4]), 10, 16); err == nil {
		srv.priority = uint16(num)
	}
	if num, err = strconv.ParseInt(string(srvRecord[5]), 10, 16); err == nil {
		srv.weight = uint16(num)
	}
	if num, err = strconv.ParseInt(string(srvRecord[6]), 10, 16); err == nil {
		srv.port = uint16(num)
	}

	srv.target = string(srvRecord[6])
}

func (srv *RDataSRV) string(zr ZoneRecord) string {
	recordStr := zr.name + " \t" + zr.ttl + " \t" + zr.class + " \t" + zr.recType + " \t"
	return recordStr
}
