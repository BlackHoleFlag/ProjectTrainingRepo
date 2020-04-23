package main

import "fmt"

type everHydra struct {
	hydraHeadSet [3]hydraHead				//whole hydra have set oh heads
	body [5]string							//whole body status in text
	bodyHp float64							//how much hp
}

type hydraHead struct{
	name string								//each hydra head have own name
	headHp float64							//separate hp
	atackStruct							//special ability
	hydraGen [3]atackStruct					//every hydra able to use general abilities
	isCutOf bool							//status of head. true = is cut of, false=exist
}

type atackStruct struct {
	abilityName string						//each head using own abilities
	abilityDmg float64						//so have different damage
}

func main() {
	//creating general hydra abilities
	abilityStomp:= atackStruct {"stomp", 15}
	abilityBite := atackStruct {"bity-bite", 20}
	abilityTail := atackStruct {"tail swipe", 10}

	// creating set(array) of general abilities
	abilityAll  := [3]atackStruct {abilityStomp, abilityBite, abilityTail}

	//creating abilities for each head
	abilityFire := atackStruct {"fireblast", 50}				
	abilityCold := atackStruct {"frozen lance", 40}
	abilityAcid := atackStruct {"toxic spit", 30}

	//creating heads for hydra
	   headCeil := hydraHead {"Ceil", 200, abilityFire, abilityAll, false}
	headAquinox := hydraHead {"Aquinox", 200, abilityCold, abilityAll, false}
	  headBrokh := hydraHead {"Brokh", 200, abilityAcid, abilityAll, false}

	//creating set(array) of heads
	    headSet := [3]hydraHead {headCeil, headBrokh, headAquinox}

	//creating statuses for body
	 statusList := [5] string {"fine", "burned", "freezed", "poisoned", "dead"}

	//creating hydra from parts
	hydra1 := everHydra {headSet, statusList, 500}
	fmt.Println(hydra1)

	//thisIsMap := make map
}