package zonerecord

import "regexp"

type RDataTXT struct {
	txt string
}

func (txt *RDataTXT) set(raw []byte) {
	valueRegex := regexp.MustCompile(`(?i)^.*IN\s+TXT\s+("[^"]*")`)
	valueRecord := valueRegex.FindSubmatch(raw)

	txt.txt = string(valueRecord[1])
}

func (txt *RDataTXT) string(zr ZoneRecord) string {
	return stringRDataSingleValue(zr)
}
