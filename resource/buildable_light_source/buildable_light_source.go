// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildableLightSource

import (
	"fmt"
)

type FGBuildableLightSource struct {
	ClassName                               string
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MDescription                            string
	MDisplayName                            string
	MExcludeFromMaterialInstancing          string
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHasPower                               bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsDay                                  bool
	MIsEnabled                              bool
	MIsUseable                              bool
	MLightControlData                       string
	MMaterialNameToInstanceManager          string
	MPowerConsumption                       float64
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
}

var (
	CeilingLight = FGBuildableLightSource{
		ClassName:                               "Build_CeilingLight_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Can be placed on ceilings to light up indoor factory spaces.

Light color and intensity can be modified.
Allows up to two Power Line connections.`,
		MDisplayName:                    `Ceiling Light`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHasPower:                       false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsDay:                          false,
		MIsEnabled:                      true,
		MIsUseable:                      false,
		MLightControlData:               `(Intensity=1.000000)`,
		MMaterialNameToInstanceManager:  `()`,
		MPowerConsumption:               2.000000,
		MShouldModifyWorldGrid:          false,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	StreetLight = FGBuildableLightSource{
		ClassName:                               "Build_StreetLight_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `A small Street Light, perfect for lighting up factory pathways and roads.

Light color and intensity can be modified.
Allows up to two Power Line connections.`,
		MDisplayName:                    `Street Light`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHasPower:                       false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsDay:                          false,
		MIsEnabled:                      true,
		MIsUseable:                      true,
		MLightControlData:               `(Intensity=1.000000)`,
		MMaterialNameToInstanceManager:  `()`,
		MPowerConsumption:               1.000000,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}
)

func GetByClassName(className string) (*FGBuildableLightSource, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableLightSource with name %s", className)
}

var classNameToVar = map[string]*FGBuildableLightSource{
	"Build_CeilingLight_C": &CeilingLight,
	"Build_StreetLight_C":  &StreetLight,
}
