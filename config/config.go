package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type OtxConfig struct {
	otxApiKey     string
	otxPulses     []string
	otxIndicators []string
}

type DBConfig struct {
	dbType     string
	dbAddress  string
	dbPort     int
	dbLogin    string
	dbPassword string
}

func ReadConfigFile(filename string) bool {
	config, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var otx OtxConfig
	parseOtxConfig(string(config), &otx)
	fmt.Println(otx.otxApiKey)
	fmt.Println(otx.otxPulses[1])
	fmt.Println(otx.otxIndicators[1])
	return true
}

func parseOtxConfig(otxConfigString string, otx *OtxConfig) {
	var otxGlobalFlag bool = false
	var otxPulsesFlag bool = false
	var otxIndicatorsFlag bool = false

	reader := bufio.NewScanner(strings.NewReader(otxConfigString))

	for reader.Scan() {
		if reader.Text() == "[otx]" {
			otxGlobalFlag = true
			otxIndicatorsFlag = false
			otxPulsesFlag = false
		}
		if reader.Text() == "[otx.pulses]" {
			otxGlobalFlag = false
			otxIndicatorsFlag = false
			otxPulsesFlag = true
		}
		if reader.Text() == "[otx.indicators]" {
			otxGlobalFlag = false
			otxIndicatorsFlag = true
			otxPulsesFlag = false
		}
		if otxGlobalFlag {
			if strings.Contains(reader.Text(), "api_key") {
				apikey := strings.Split(reader.Text(), "=")[1]
				otx.otxApiKey = apikey
			}
		}
		if otxPulsesFlag {
			if strings.Contains(reader.Text(), "pulses") {
				otx.otxPulses = strings.Split(reader.Text(), "=")
			}
		}
		if otxIndicatorsFlag {
			if strings.Contains(reader.Text(), "indicators") {
				otx.otxIndicators = strings.Split(reader.Text(), "=")
			}
		}
	}
}
