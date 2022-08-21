package zonerecord

import (
	"regexp"
	"strconv"
)

type RDataSOA struct {
	mname   string
	rname   string
	serial  uint32
	refresh int32
	retry   int32
	expire  int32
	minimum uint32
}

func (soa *RDataSOA) set(raw []byte) {
	var err error
	var num int64

	// [1] - mname [2] - rname [3] - Serial [4] - Refresh
	// [5] - Retry [6] - Expire[7] - Minimum
	SOARegex := regexp.MustCompile(`(?mi)^.*IN\s+SOA\s*(\S*)\s+(\S*)\s+\(\s+(\d*)\D+(\d*)\D+(\d*)\D+(\d*)\D+(\d*)\D+$`)
	SOARecord := SOARegex.FindSubmatch(raw)

	soa.mname = string(SOARecord[1])
	soa.rname = string(SOARecord[2])

	// Serial
	if num, err = strconv.ParseInt(string(SOARecord[3]), 10, 64); err == nil {
		soa.serial = uint32(num)
	}

	// Refresh
	if num, err = strconv.ParseInt(string(SOARecord[4]), 10, 64); err == nil {
		soa.refresh = int32(num)
	}

	// Retry
	if num, err = strconv.ParseInt(string(SOARecord[5]), 10, 64); err == nil {
		soa.retry = int32(num)
	}

	// Expire
	if num, err = strconv.ParseInt(string(SOARecord[6]), 10, 64); err == nil {
		soa.expire = int32(num)
	}

	// Minimum
	if num, err = strconv.ParseInt(string(SOARecord[7]), 10, 64); err == nil {
		soa.minimum = uint32(num)
	}
}

func (soa *RDataSOA) string(zr ZoneRecord) string {
	recordStr := zr.name + " \t" + zr.ttl + " \t" + zr.class + " \t" + zr.recType + " " + soa.mname + " " + soa.rname
	recordStr += " \t(\n"
	recordStr += " \t\t\t\t" + strconv.FormatUint(uint64(soa.serial), 10) + " \t\t; serial\n"
	recordStr += " \t\t\t\t" + strconv.FormatInt(int64(soa.refresh), 10) + " \t\t; refresh\n"
	recordStr += " \t\t\t\t" + strconv.FormatInt(int64(soa.retry), 10) + " \t\t; retry\n"
	recordStr += " \t\t\t\t" + strconv.FormatInt(int64(soa.expire), 10) + " \t\t; expire\n"
	recordStr += " \t\t\t\t" + strconv.FormatUint(uint64(soa.minimum), 10) + " )\t\t ; minimum\n"
	return recordStr
}
