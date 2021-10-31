package GolfCartDispenser

import (
	"fmt"
)

type FGGolfCartDispenser struct {
	ClassName              string
	MArmAnimation          string
	MAttachSocket          string
	MBackAnimation         string
	MCostToUse             string
	MEquipmentSlot         string
	MHasPersistentOwner    bool
	MPlaceDistanceMax      float64
	MUseDefaultPrimaryFire bool
}

var (
	GolfCartDispenser = FGGolfCartDispenser{
		ClassName:              "Equip_GolfCartDispenser_C",
		MArmAnimation:          `AE_Generic2Hand`,
		MAttachSocket:          `hand_rSocket`,
		MBackAnimation:         `BE_None`,
		MCostToUse:             ``,
		MEquipmentSlot:         `ES_ARMS`,
		MHasPersistentOwner:    false,
		MPlaceDistanceMax:      1000.000000,
		MUseDefaultPrimaryFire: false,
	}
)

func GetByClassName(className string) (*FGGolfCartDispenser, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGGolfCartDispenser with name %s", className)
}

var classNameToVar = map[string]*FGGolfCartDispenser{
	"Equip_GolfCartDispenser_C": &GolfCartDispenser,
}
