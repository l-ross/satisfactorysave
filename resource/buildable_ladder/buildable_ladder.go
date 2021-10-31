package BuildableLadder

import (
	"fmt"
)

type FGBuildableLadder struct {
	ClassName                               string
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MDescription                            string
	MDisplayName                            string
	MExcludeFromMaterialInstancing          string
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsUseable                              bool
	MLadderMeshes                           string
	MMaterialNameToInstanceManager          string
	MMaxSegmentCount                        int
	MMeshHeight                             float64
	MNumSegments                            int
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MWidth                                  float64
	MaxRenderDistance                       float64
}

var (
	Ladder = FGBuildableLadder{
		ClassName:                               "Build_Ladder_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription:                            `A ladder with a default height of 2 meters, which can be extended while building. Snaps to walls and foundations.`,
		MDisplayName:                            `Ladder`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MLadderMeshes:                           ``,
		MMaterialNameToInstanceManager:          `()`,
		MMaxSegmentCount:                        10,
		MMeshHeight:                             200.000000,
		MNumSegments:                            0,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MToggleDormancyOnInteraction:            false,
		MWidth:                                  80.000000,
		MaxRenderDistance:                       -1.000000,
	}
)

func GetByClassName(className string) (*FGBuildableLadder, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableLadder with name %s", className)
}

var classNameToVar = map[string]*FGBuildableLadder{
	"Build_Ladder_C": &Ladder,
}
