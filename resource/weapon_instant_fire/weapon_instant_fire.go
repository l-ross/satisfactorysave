package WeaponInstantFire

import (
	"fmt"
)

type FGWeaponInstantFire struct {
	ClassName              string
	Fire                   string
	MArmAnimation          string
	MAttachSocket          string
	MBackAnimation         string
	MBlockSprintWhenFiring bool
	MCostToUse             string
	MCurrentAmmo           int
	MEquipmentSlot         string
	MFireRate              float64
	MHasPersistentOwner    bool
	MHasReloadedOnce       bool
	MInstantHitDamage      float64
	MLockAngle             float64
	MMagSize               int
	MReloadTime            float64
	MUseDefaultPrimaryFire bool
	MWeaponTraceLength     float64
}

var (
	Rifle = FGWeaponInstantFire{
		ClassName:              "Equip_Rifle_C",
		Fire:                   `()`,
		MArmAnimation:          `AE_Rifle`,
		MAttachSocket:          `hand_rSocket`,
		MBackAnimation:         `BE_None`,
		MBlockSprintWhenFiring: true,
		MCostToUse:             ``,
		MCurrentAmmo:           0,
		MEquipmentSlot:         `ES_ARMS`,
		MFireRate:              0.200000,
		MHasPersistentOwner:    false,
		MHasReloadedOnce:       false,
		MInstantHitDamage:      6.000000,
		MLockAngle:             0.000000,
		MMagSize:               10,
		MReloadTime:            3.200000,
		MUseDefaultPrimaryFire: false,
		MWeaponTraceLength:     15000.000000,
	}
)

func GetByClassName(className string) (*FGWeaponInstantFire, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGWeaponInstantFire with name %s", className)
}

var classNameToVar = map[string]*FGWeaponInstantFire{
	"Equip_Rifle_C": &Rifle,
}
