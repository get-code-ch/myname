package test

import (
	"github.com/get-code-ch/myname/zone"
	"log"
	"testing"
)

const ZonesResourcesDirectory = "./resources/zones/"

func TestZoneLoad(t *testing.T) {
	var zoneFiles zone.Files
	var z *zone.Zone
	var err error

	want := make(map[string]interface{})

	want["ZoneFile"] = "example.com.txt"

	// Getting zoneFiles
	if zoneFiles, err = zone.NewZoneFiles(ZonesResourcesDirectory, `.*example\.com.*`); err != nil {
		t.Fatalf("Error getting zone files --> %v", err)
	}

	// Checking if z file exist
	fileIdx := -1
	for idx, file := range zoneFiles {
		if file.Name() == want["ZoneFile"] {
			fileIdx = idx
			break
		}
	}

	if fileIdx == -1 {
		t.Errorf("Failed, zone file not found want %s", want["ZoneFile"])
	}

	// Creating Zone for each existing zone file
	file := zoneFiles[0]
	t.Logf("Testing zone file %s", file.Name())
	z = new(zone.Zone)
	if err := z.ParseZoneFile(ZonesResourcesDirectory + file.Name()); err != nil {
		t.Logf("Error parsing file -> %v", err)
	}

	log.Printf("%s", z)

}
