package test

import (
	"github.com/get-code-ch/myname/zone"
	"testing"
)

func TestZoneFiles(t *testing.T) {
	want := make(map[string]interface{})

	want["ZoneFilesCount"] = 3
	want["ZoneFilesCountFiltered"] = 1

	if zoneFiles, err := zone.NewZoneFiles("./resources/zones/"); err != nil {
		t.Fatalf("Error getting zone files --> %v", err)
	} else {
		if len(zoneFiles) != want["ZoneFilesCount"] {
			t.Errorf("Failed, zoneFiles lenght %d want %d", len(zoneFiles), want["ZoneFilesCount"])
		}
	}

	if zoneFiles, err := zone.NewZoneFiles("./resources/zones/", `.*example\.com.*`); err != nil {
		t.Fatalf("Error getting zone files --> %v", err)
	} else {
		if len(zoneFiles) != want["ZoneFilesCountFiltered"] {
			t.Errorf("Failed, zoneFiles lenght %d want %d", len(zoneFiles), want["ZoneFilesCountFiltered"])
		}
	}

}
