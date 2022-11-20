package models

//pakcage contains data models for threat intel entities
//type AttackPattern implements STIX Attack Pattern object
//https://docs.oasis-open.org/cti/stix/v2.1/os/stix-v2.1-os.html#_axjijf603msy

import "fmt"

type AttackPattern struct {
	typeString string
	name       string
}

func Test() {
	fmt.Println("Import works!")
}

func CreateAttackPattern(patternName string) AttackPattern {
	return AttackPattern{typeString: "attack-pattern", name: patternName}
}
