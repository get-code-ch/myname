package test

import (
	"github.com/get-code-ch/myname/config"
	"testing"
)

func TestConfig(t *testing.T) {
	want := make(map[string]interface{})

	want["ZoneFilesPath"] = "./resources/zones/"

	cfg, err := config.NewConfig("./resources/config.json")
	if err != nil {
		t.Fatalf("Error loading configuration file --> %v", err)
	}
	if cfg == nil {
		t.Fatalf("Error parsing configuration file")
	}

	if cfg.ZonesFilesPath == "" {
		t.Fatalf("Error ZonesFilesPath undefined")
	}

	if cfg.ZonesFilesPath != want["ZoneFilesPath"] {
		t.Logf("got \"%s\" want \"%s\"", cfg.ZonesFilesPath, want["ZoneFilesPath"])
	}

}
