package zonerecord

type RDataNS struct {
	Nsdname string
}

func (ns *RDataNS) set(raw []byte) {
	setRDataSingleValue(ns, raw)
}

func (ns *RDataNS) string(zr ZoneRecord) string {
	return stringRDataSingleValue(zr)
}
