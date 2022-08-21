package zonerecord

type RDataCNAME struct {
	Cname string
}

func (cname *RDataCNAME) set(raw []byte) {
	setRDataSingleValue(cname, raw)
}

func (cname *RDataCNAME) string(zr ZoneRecord) string {
	return stringRDataSingleValue(zr)
}
