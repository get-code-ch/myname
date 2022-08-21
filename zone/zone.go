package zone

import (
	"errors"
	"fmt"
	"github.com/get-code-ch/myname/zonerecord"
	"os"
	"regexp"
)

type Zone struct {
	name     string
	file     string
	ttl      string
	origin   string
	include  string
	generate []string
	records  []*zonerecord.ZoneRecord
}

// ParseZoneFile function parse DNS zone file and fill records in Zone
// Return an error if zone file could not be parsed
func (z *Zone) ParseZoneFile(fileName string) error {
	// zoneRecordRegex is used to validate en extract zone record in file
	// [1] - Type
	zoneRecordRegex := regexp.MustCompile(`(?mi)^(?:(?:[^$|^;]|\S*)\s+(?:\d*\s+)?IN\s+(` + zonerecord.RegexHandledRecordType() + `)\s+[^(|\r|\n]*)(?:\((?:\D*\d*.*){5}\))?`)
	//zoneDirectiveRegex := regexp.MustCompile(`(?mi)^\$INCLUDE|^\$ORIGIN|^\$TTL`)

	var zoneRecords []*zonerecord.ZoneRecord

	// Reading zone file
	z.file = fileName
	data, err := os.ReadFile(z.file)
	if err != nil {
		return errors.New(fmt.Sprintf("error reading zone file --> %v", err))
	}
	rawRecords := zoneRecordRegex.FindAllSubmatch(data, -1)
	if rawRecords == nil {
		return errors.New(fmt.Sprintf("error no DNS zone record found in zone file --> %v", err))
	}

	for _, rr := range rawRecords {
		zr := zonerecord.NewZoneRecord(rr[1], rr[0])
		if zr == nil {
			return errors.New(fmt.Sprintf("error something wrong with zone record format for --> %s", rr[0]))
		}
		zoneRecords = append(zoneRecords, zr)
	}
	z.records = zoneRecords
	return nil
}

func (z *Zone) String() string {
	zrStr := ""
	for _, zr := range z.records {
		zrStr += fmt.Sprintf("%s", zr) + "\n"
	}
	return zrStr
}
