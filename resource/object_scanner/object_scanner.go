package ObjectScanner

import (
	"fmt"
)

type FGObjectScanner struct {
	ClassName                    string
	MArmAnimation                string
	MAttachSocket                string
	MBackAnimation               string
	MBeepDelayMax                float64
	MBeepDelayMin                float64
	MCostToUse                   string
	MDetectionRange              float64
	MEquipmentSlot               string
	MHasPersistentOwner          bool
	MNormalizedCloesnessToObject float64
	MObjectDetails               string
	MObjectIsWithinRange         bool
	MPlayingSound                bool
	MScanlineLerpT               float64
	MScreenUpdateTime            float64
	MScreenUpdateTimer           string
	MShouldBeepEvenIfNoObject    bool
	MUpdateClosestObjectTime     float64
	MUseDefaultPrimaryFire       bool
}

var (
	ObjectScanner = FGObjectScanner{
		ClassName:                    "Equip_ObjectScanner_C",
		MArmAnimation:                `AE_Generic1Hand`,
		MAttachSocket:                `hand_rSocket`,
		MBackAnimation:               `BE_None`,
		MBeepDelayMax:                3.000000,
		MBeepDelayMin:                0.100000,
		MCostToUse:                   ``,
		MDetectionRange:              25000.000000,
		MEquipmentSlot:               `ES_ARMS`,
		MHasPersistentOwner:          false,
		MNormalizedCloesnessToObject: 0.000000,
		MObjectDetails:               `((ScannableClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Environment/Crystal/BP_Crystal.BP_Crystal_C"',DisplayText=NSLOCTEXT("[1EC05A034D581C9DF4E2CABBDFC08F96]", "86EC608140665F49DDDC7CAB3C24F13F", "Power Slugs"),ScannerLightColor=(B=177,G=146,R=0,A=0),RequiredSchematic=/Game/FactoryGame/Schematics/Research/PowerSlugs_RS/Research_PowerSlugs_3.Research_PowerSlugs_3_C),(ScannableClass=Class'"/Script/FactoryGame.FGEnemy"',DisplayText=NSLOCTEXT("[1EC05A034D581C9DF4E2CABBDFC08F96]", "5B74A8BE4A1C9CD855BEFC811B4B3B25", "Enemies"),ScannerLightColor=(B=0,G=101,R=255,A=0),PreCacheAllOfType=False,NewDetectionRange=10000.000000,RequiredSchematic=/Game/FactoryGame/Schematics/Research/AlienOrganisms_RS/Research_AOrganisms_2.Research_AOrganisms_2_C),(ScannableClass=BlueprintGeneratedClass'"/Game/FactoryGame/World/Benefit/DropPod/BP_DropPod.BP_DropPod_C"',DisplayText=NSLOCTEXT("[1EC05A034D581C9DF4E2CABBDFC08F96]", "445D760D4B31F089FAB651BB8F5A483A", "Crash Sites"),ScannerLightColor=(B=0,G=255,R=25,A=0),ShouldOverrideDetectionRange=True,NewDetectionRange=10000.000000,RequiredSchematic=/Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_4_1.Research_Quartz_4_1_C),(ScannableClass=BlueprintGeneratedClass'"/Game/FactoryGame/Prototype/WAT/BP_WAT1.BP_WAT1_C"',DisplayText=NSLOCTEXT("[1EC05A034D581C9DF4E2CABBDFC08F96]", "F4D2653A4A5880AF210B01A5603FB563", "Somersloop"),ScannerLightColor=(B=39,G=0,R=218,A=0),RequiredSchematic=/Game/FactoryGame/Schematics/Story/Schematic_ObjectScanner_WAT1.Schematic_ObjectScanner_WAT1_C),(ScannableClass=BlueprintGeneratedClass'"/Game/FactoryGame/Prototype/WAT/BP_WAT2.BP_WAT2_C"',DisplayText=NSLOCTEXT("[1EC05A034D581C9DF4E2CABBDFC08F96]", "B7981DAF4A4E916397EAE49AB80D441C", "Mercer Sphere"),ScannerLightColor=(B=255,G=0,R=170,A=0),RequiredSchematic=/Game/FactoryGame/Schematics/Story/Schematic_ObjectScanner_WAT2.Schematic_ObjectScanner_WAT2_C),(ScannableClass=BlueprintGeneratedClass'"/Game/FactoryGame/World/Benefit/BerryBush/BP_BerryBush.BP_BerryBush_C"',DisplayText=NSLOCTEXT("[1EC05A034D581C9DF4E2CABBDFC08F96]", "ADD8E7C04B264BAFB77408BEF1094886", "Berry"),ScannerLightColor=(B=141,G=0,R=255,A=0),RequiredSchematic=/Game/FactoryGame/Schematics/Research/Nutrients_RS/Research_Nutrients_0.Research_Nutrients_0_C),(ScannableClass=BlueprintGeneratedClass'"/Game/FactoryGame/World/Benefit/NutBush/BP_NutBush.BP_NutBush_C"',DisplayText=NSLOCTEXT("[1EC05A034D581C9DF4E2CABBDFC08F96]", "5427CF4542BC62E4F49D0BAD9B8BEFB0", "Nut"),ScannerLightColor=(B=0,G=108,R=186,A=0),RequiredSchematic=/Game/FactoryGame/Schematics/Research/Nutrients_RS/Research_Nutrients_1.Research_Nutrients_1_C),(ScannableClass=BlueprintGeneratedClass'"/Game/FactoryGame/World/Benefit/Mushroom/BP_Shroom_01.BP_Shroom_01_C"',DisplayText=NSLOCTEXT("[1EC05A034D581C9DF4E2CABBDFC08F96]", "8FB555B946C751A2EA7A278D771C8161", "Mushroom"),ScannerLightColor=(B=208,G=195,R=255,A=0),RequiredSchematic=/Game/FactoryGame/Schematics/Research/Nutrients_RS/Research_Nutrients_2.Research_Nutrients_2_C))`,
		MObjectIsWithinRange:         false,
		MPlayingSound:                false,
		MScanlineLerpT:               1.000000,
		MScreenUpdateTime:            0.050000,
		MScreenUpdateTimer:           `()`,
		MShouldBeepEvenIfNoObject:    false,
		MUpdateClosestObjectTime:     2.000000,
		MUseDefaultPrimaryFire:       false,
	}
)

func GetByClassName(className string) (*FGObjectScanner, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGObjectScanner with name %s", className)
}

var classNameToVar = map[string]*FGObjectScanner{
	"Equip_ObjectScanner_C": &ObjectScanner,
}
