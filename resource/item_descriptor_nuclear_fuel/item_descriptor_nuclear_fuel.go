package ItemDescriptorNuclearFuel

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGItemDescriptorNuclearFuel struct {
	ClassName               string
	MAbbreviatedDisplayName string
	MAmountOfWaste          int
	MCanBeDiscarded         bool
	MDescription            string
	MDisplayName            string
	MEnergyValue            float64
	MFluidColor             string
	MForm                   resource.Form
	MGasColor               string
	MPersistentBigIcon      string
	MRadioactiveDecay       float64
	MRememberPickUp         bool
	MResourceSinkPoints     int
	MSmallIcon              string
	MSpentFuelClass         string
	MStackSize              resource.StackSize
}

var (
	NuclearFuelRod = FGItemDescriptorNuclearFuel{
		ClassName:               "Desc_NuclearFuelRod_C",
		MAbbreviatedDisplayName: ``,
		MAmountOfWaste:          50,
		MCanBeDiscarded:         true,
		MDescription: `Used as fuel for the Nuclear Power Plant.

Caution: Produces radioactive Uranium Waste when consumed.
Caution: Moderately Radioactive.`,
		MDisplayName:        `Uranium Fuel Rod`,
		MEnergyValue:        750000.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/NuclearFuelRod/UI/IconDesc_NuclearFuelRod_256.IconDesc_NuclearFuelRod_256`,
		MRadioactiveDecay:   50.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 44092,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/NuclearFuelRod/UI/IconDesc_NuclearFuelRod_64.IconDesc_NuclearFuelRod_64`,
		MSpentFuelClass:     `BlueprintGeneratedClass'/Game/FactoryGame/Resource/Parts/NuclearWaste/Desc_NuclearWaste.Desc_NuclearWaste_C'`,
		MStackSize:          resource.Small,
	}

	PlutoniumFuelRod = FGItemDescriptorNuclearFuel{
		ClassName:               "Desc_PlutoniumFuelRod_C",
		MAbbreviatedDisplayName: ``,
		MAmountOfWaste:          10,
		MCanBeDiscarded:         true,
		MDescription: `Used as fuel for the Nuclear Power Plant.

Caution: Produces radioactive Plutonium Waste when consumed.
Caution: HIGHLY Radioactive.`,
		MDisplayName:        `Plutonium Fuel Rod`,
		MEnergyValue:        1500000.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/PlutoniumFuelRods/UI/IconDesc_PlutoniumFuelRod_256.IconDesc_PlutoniumFuelRod_256`,
		MRadioactiveDecay:   250.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 153184,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/PlutoniumFuelRods/UI/IconDesc_PlutoniumFuelRod_64.IconDesc_PlutoniumFuelRod_64`,
		MSpentFuelClass:     `BlueprintGeneratedClass'/Game/FactoryGame/Resource/Parts/NuclearWaste/Desc_PlutoniumWaste.Desc_PlutoniumWaste_C'`,
		MStackSize:          resource.Small,
	}
)

func GetByClassName(className string) (*FGItemDescriptorNuclearFuel, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGItemDescriptorNuclearFuel with name %s", className)
}

var classNameToVar = map[string]*FGItemDescriptorNuclearFuel{
	"Desc_NuclearFuelRod_C":   &NuclearFuelRod,
	"Desc_PlutoniumFuelRod_C": &PlutoniumFuelRod,
}
