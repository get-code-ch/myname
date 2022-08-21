package zonerecord

type RDataAAAA struct {
	Ipv6 string
}

func (aaaa *RDataAAAA) set(raw []byte) {
	setRDataSingleValue(aaaa, raw)
}

func (aaaa *RDataAAAA) string(zr ZoneRecord) string {
	return stringRDataSingleValue(zr)
}
