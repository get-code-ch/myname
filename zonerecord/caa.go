package zonerecord

import (
	"regexp"
	"strconv"
)

type RDataCAA struct {
	flag  uint8
	tag   string
	value string
}

func (caa *RDataCAA) set(raw []byte) {
	var err error
	var num int64

	caaRegex := regexp.MustCompile(`(?i)^.*IN\s+CAA\s+(\d*)\s+(\S*)\s+("[^"]*")`)
	caaRecord := caaRegex.FindSubmatch(raw)

	if num, err = strconv.ParseInt(string(caaRecord[1]), 10, 8); err == nil {
		caa.flag = uint8(num)
	}

	caa.tag = string(caaRecord[2])
	caa.value = string(caaRecord[3])
}

func (caa *RDataCAA) string(zr ZoneRecord) string {
	recordStr := zr.name + " \t" + zr.ttl + " \t" + zr.class + " \t" + zr.recType + " \t"
	recordStr += strconv.FormatUint(uint64(caa.flag), 10) + " \t"
	recordStr += caa.tag + " \t"
	recordStr += caa.value
	return recordStr
}
