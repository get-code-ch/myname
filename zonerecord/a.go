package zonerecord

type RDataA struct {
	Ipv4 string
}

func (a *RDataA) set(raw []byte) {
	setRDataSingleValue(a, raw)
}

func (a *RDataA) string(zr ZoneRecord) string {
	return stringRDataSingleValue(zr)
}
