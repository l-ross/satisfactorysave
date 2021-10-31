package BuildableFloodlight

import (
	"fmt"
)

type FGBuildableFloodlight struct {
	ClassName                               string
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MDescription                            string
	MDisplayName                            string
	MExcludeFromMaterialInstancing          string
	MFixtureAngle                           int
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
	FloodlightPole = FGBuildableFloodlight{
		ClassName:                               "Build_FloodlightPole_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `A tall light tower, specifically designed for illuminating large or outdoor spaces.

Light color and intensity can be modified.
Allows up to two Power Line connections.`,
		MDisplayName:                    `Flood Light Tower`,
		MExcludeFromMaterialInstancing:  ``,
		MFixtureAngle:                   50,
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
		MPowerConsumption:               6.000000,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	FloodlightWall = FGBuildableFloodlight{
		ClassName:                               "Build_FloodlightWall_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Can be attached to Walls and Foundations to illuminate large spaces.

Light color and intensity can be modified.
Allows up to two Power Line connections.`,
		MDisplayName:                    `Wall Mounted Flood Light`,
		MExcludeFromMaterialInstancing:  ``,
		MFixtureAngle:                   30,
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
		MPowerConsumption:               6.000000,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}
)

func GetByClassName(className string) (*FGBuildableFloodlight, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableFloodlight with name %s", className)
}

var classNameToVar = map[string]*FGBuildableFloodlight{
	"Build_FloodlightPole_C": &FloodlightPole,
	"Build_FloodlightWall_C": &FloodlightWall,
}
