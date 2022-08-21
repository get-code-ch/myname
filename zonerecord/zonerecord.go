package zonerecord

import (
	"reflect"
	"regexp"
	"strings"
)

type RecordData interface {
	set([]byte)
	string(ZoneRecord) string
}

type RDataSingleValueConstraint interface {
	RDataNS | RDataCNAME | RDataA | RDataAAAA
}

type ZoneRecord struct {
	name    string
	recType string
	class   string
	ttl     string
	rdata   RecordData
}

// HandledRecordType function return a map[string] of Type from handled DNS record types
func HandledRecordType() map[string]reflect.Type {
	var RDataTypes = make(map[string]reflect.Type)

	// Defining record data type
	RDataTypes["SOA"] = reflect.TypeOf(RDataSOA{})
	RDataTypes["NS"] = reflect.TypeOf(RDataNS{})
	RDataTypes["CNAME"] = reflect.TypeOf(RDataCNAME{})
	RDataTypes["A"] = reflect.TypeOf(RDataA{})
	RDataTypes["AAAA"] = reflect.TypeOf(RDataAAAA{})
	RDataTypes["TXT"] = reflect.TypeOf(RDataTXT{})
	RDataTypes["CAA"] = reflect.TypeOf(RDataCAA{})
	RDataTypes["MX"] = reflect.TypeOf(RDataMX{})
	RDataTypes["SRV"] = reflect.TypeOf(RDataSRV{})

	return RDataTypes
}

func RegexHandledRecordType() string {
	regexStr := ""

	for zt := range HandledRecordType() {
		if regexStr != "" {
			regexStr += "|"
		}
		regexStr += zt
	}
	return regexStr
}

func RegexRDataSingleValueRecordType() string {
	return "NS|CNAME|A|AAAA"
}

// IsZoneRecordTypeHandled function return true if zone record type is handled
func IsZoneRecordTypeHandled(recordType string) bool {
	if _, ok := HandledRecordType()[recordType]; ok {
		return true
	}
	return false
}

// NewZoneRecord function create a new zone record
func NewZoneRecord(zoneType []byte, raw []byte) *ZoneRecord {

	// [1] - Name / Service
	// [2] - TTL
	// [3] - Class
	zr := new(ZoneRecord)
	zr.recType = strings.ToUpper(string(zoneType))

	zoneRecordRegex := regexp.MustCompile(`(?mi)^([^$]|[^;]|\S.*?)\s+(\d*)(?:\s+)?(IN)\s+(?:` + zr.recType + `)`)
	if zrInfo := zoneRecordRegex.FindSubmatch(raw); zrInfo != nil {
		zr.name = string(zrInfo[1])
		zr.ttl = string(zrInfo[2])
		zr.class = string(zrInfo[3])
		zr.rdata = NewRecordData(zr.recType)
		zr.rdata.set(raw)
	} else {
		zr = nil
	}

	return zr
}

// NewRecordData function create a new RData... object depending on DNS record type.
func NewRecordData(recordType string) RecordData {
	return reflect.New(HandledRecordType()[recordType]).Interface().(RecordData)
}

// setRDataSingleValue is a generic function who set record Data with unique value for DNS record of type NS, A, AAAA, etc.
func setRDataSingleValue[R RDataSingleValueConstraint](rData *R, raw []byte) {
	valueRegex := regexp.MustCompile(`(?i)^.*IN\s+(?:` + RegexRDataSingleValueRecordType() + `)\s+([^;\s]*)`)
	valueRecord := valueRegex.FindSubmatch(raw)

	value := string(valueRecord[1])
	reflect.ValueOf(rData).Elem().Field(0).Set(reflect.ValueOf(value))
}

func stringRDataSingleValue(zr ZoneRecord) string {
	value := reflect.ValueOf(zr.rdata).Elem().Field(0).String()
	recordStr := zr.name + " \t" + zr.ttl + " \t" + zr.class + " \t" + zr.recType + " \t" + value
	return recordStr
}

// String function return zone record as a formatted string
func (zr *ZoneRecord) String() string {
	return zr.rdata.string(*zr)
}
