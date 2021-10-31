package NobeliskDetonator

import (
	"fmt"
)

type FGNobeliskDetonator struct {
	ClassName               string
	MArmAnimation           string
	MAttachSocket           string
	MBackAnimation          string
	MBlockSprintWhenFiring  bool
	MCostToUse              string
	MCurrentAmmo            int
	MDelayBetweenExplosions float64
	MDispensedExplosives    string
	MEquipmentSlot          string
	MExplosiveData          string
	MFireRate               float64
	MHasPersistentOwner     bool
	MIsPendingExecuteFire   bool
	MMagSize                int
	MMaxChargeTime          float64
	MMaxThrowForce          int
	MMinThrowForce          int
	MReloadTime             float64
	MUseDefaultPrimaryFire  bool
}

var (
	NobeliskDetonator = FGNobeliskDetonator{
		ClassName:               "Equip_NobeliskDetonator_C",
		MArmAnimation:           `AE_Nobelisk`,
		MAttachSocket:           `hand_lSocket`,
		MBackAnimation:          `BE_None`,
		MBlockSprintWhenFiring:  true,
		MCostToUse:              ``,
		MCurrentAmmo:            0,
		MDelayBetweenExplosions: 0.150000,
		MDispensedExplosives:    ``,
		MEquipmentSlot:          `ES_ARMS`,
		MExplosiveData:          `(ProjectileClass=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/NobeliskDetonator/BP_NobeliskExplosive.BP_NobeliskExplosive_C"',ProjectileLifeSpan=-1.000000,ProjectileStickSpan=-1.000000,ExplosionDamage=50,ExplosionRadius=750.000000,ImpactDamage=1,ExplodeAtEndOfLife=True,DamageType=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/NobeliskDetonator/DamageType_NobeliskExplosiveImpact.DamageType_NobeliskExplosiveImpact_C"',DamageTypeExplode=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/NobeliskDetonator/DamageType_NobeliskExplosive.DamageType_NobeliskExplosive_C"')`,
		MFireRate:               0.500000,
		MHasPersistentOwner:     false,
		MIsPendingExecuteFire:   false,
		MMagSize:                1,
		MMaxChargeTime:          1.000000,
		MMaxThrowForce:          3000,
		MMinThrowForce:          200,
		MReloadTime:             1.000000,
		MUseDefaultPrimaryFire:  false,
	}
)

func GetByClassName(className string) (*FGNobeliskDetonator, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGNobeliskDetonator with name %s", className)
}

var classNameToVar = map[string]*FGNobeliskDetonator{
	"Equip_NobeliskDetonator_C": &NobeliskDetonator,
}
