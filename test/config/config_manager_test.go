package config_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/rainbowsthill/copper_backend/config"
)

func TestAddConfigFile(t *testing.T) {
	testCases := [2]string{}
	var err error

	testCases[0], err = filepath.Abs("./config.yaml")
	if err != nil {
		t.Fatalf("Failed to load config.yaml: %v", err)
	}
	testCases[1], err = filepath.Abs("./general.yaml")
	if err != nil {
		t.Fatalf("Failed to load general.yaml: %v", err)
	}

	for _, testCase := range testCases {
		if err = config.AddConfigFile(testCase); err != nil {
			t.Fatalf("Failed to register config file %s: %v", testCase, err)
		}
	}
}

func TestGetConfiguration(t *testing.T) {
	TestAddConfigFile(t)

	p1, _ := filepath.Abs("./config.yaml")
	p2, _ := filepath.Abs("./general.yaml")
	testCases := map[string]map[string]any{
		p1: {
			"service.ip":                       "192.168.1.10",
			"service.port":                     1789,
			"service.service_info.id":          1,
			"service.service_info.data_center": 1,
		},
		p2: {
			"registry_service_master.ip":                        "127.0.0.1",
			"registry_service_master.port":                      1789,
			"registry_service_replica01.load_balancer.ip":       "127.0.0.1",
			"registry_service_replica01.load_balancer.username": nil,
		},
	}

	for p, records := range testCases {
		for recordPath, record := range records {
			if r := config.GetConfiguration(p, strings.Split(recordPath, ".")); r != record {
				t.Fatalf("Result of GetConfiguation(%s, %s):\n%v\nExpected:\n%v", p, recordPath, r, record)
			}
		}
	}
}
