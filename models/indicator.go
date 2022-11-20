package models

//pakcage contains data models for threat intel entities
//type Indicator implements STIX Indicator object
//https://docs.oasis-open.org/cti/stix/v2.1/os/stix-v2.1-os.html#_muftrcpnf89v

type Indicator struct {
	typeString  string
	name        string
	pattern     AttackPattern
	patternType string
	validFrom   string
	validUntil  string
}
