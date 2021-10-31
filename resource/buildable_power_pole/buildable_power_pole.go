package BuildablePowerPole

import (
	"fmt"
)

type FGBuildablePowerPole struct {
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
	MIsUseable                              bool
	MMaterialNameToInstanceManager          string
	MPowerConnections                       string
	MPowerPoleType                          string
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
}

var (
	PowerPoleMk1 = FGBuildablePowerPole{
		ClassName:                               "Build_PowerPoleMk1_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Can handle up to 4 Power Line connections.

Connect Power Poles, Power Generators and factory buildings together with Power Lines to create a power grid. The power grid supplies the connected buildings with power.`,
		MDisplayName:                    `Power Pole Mk.1`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHasPower:                       false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      true,
		MMaterialNameToInstanceManager:  `()`,
		MPowerConnections:               ``,
		MPowerPoleType:                  `PPT_POLE`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	PowerPoleMk2 = FGBuildablePowerPole{
		ClassName:                               "Build_PowerPoleMk2_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Can handle up to 7 Power Line connections.

Connect Power Poles, Power Generators and factory buildings together with Power Lines to create a power grid. The power grid supplies the connected buildings with power.`,
		MDisplayName:                    `Power Pole Mk.2`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHasPower:                       false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      true,
		MMaterialNameToInstanceManager:  `()`,
		MPowerConnections:               ``,
		MPowerPoleType:                  `PPT_POLE`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	PowerPoleMk3 = FGBuildablePowerPole{
		ClassName:                               "Build_PowerPoleMk3_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Can handle up to 10 Power Line connections.

Connect Power Poles, Power Generators and factory buildings together with Power Lines to create a power grid. The power grid supplies the connected buildings with power.`,
		MDisplayName:                    `Power Pole Mk.3`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHasPower:                       false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      true,
		MMaterialNameToInstanceManager:  `()`,
		MPowerConnections:               ``,
		MPowerPoleType:                  `PPT_POLE`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	PowerPoleWall = FGBuildablePowerPole{
		ClassName:                               "Build_PowerPoleWall_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Power Pole that attaches to a wall.

Can handle up to 4 Power Line connections.

Connect Power Poles, Power Generators and factory buildings together with Power Lines to create a power grid. The power grid supplies the connected buildings with power.`,
		MDisplayName:                    `Wall Outlet Mk.1`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHasPower:                       false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MPowerConnections:               ``,
		MPowerPoleType:                  `PPT_WALL`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	PowerPoleWallDouble = FGBuildablePowerPole{
		ClassName:                               "Build_PowerPoleWallDouble_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Power Pole that attaches to a wall. Has one connector on each side of the wall.

Can handle up to 4 Power Line connections.

Connect Power Poles, Power Generators and factory buildings together with Power Lines to create a power grid. The power grid supplies the connected buildings with power.`,
		MDisplayName:                    `Double Wall Outlet Mk.1`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHasPower:                       false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MPowerConnections:               ``,
		MPowerPoleType:                  `PPT_WALL_DOUBLE`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	PowerPoleWallDoubleMk2 = FGBuildablePowerPole{
		ClassName:                               "Build_PowerPoleWallDouble_Mk2_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Power Pole that attaches to a wall. Has one connector on each side of the wall.

Can handle up to 7 Power Line connections.

Connect Power Poles, Power Generators and factory buildings together with Power Lines to create a power grid. The power grid supplies the connected buildings with power.`,
		MDisplayName:                    `Double Wall Outlet Mk.2`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHasPower:                       false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MPowerConnections:               ``,
		MPowerPoleType:                  `PPT_WALL_DOUBLE`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	PowerPoleWallDoubleMk3 = FGBuildablePowerPole{
		ClassName:                               "Build_PowerPoleWallDouble_Mk3_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Power Pole that attaches to a wall. Has one connector on each side of the wall.

Can handle up to 10 Power Line connections.

Connect Power Poles, Power Generators and factory buildings together with Power Lines to create a power grid. The power grid supplies the connected buildings with power.`,
		MDisplayName:                    `Double Wall Outlet Mk.3`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHasPower:                       false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MPowerConnections:               ``,
		MPowerPoleType:                  `PPT_WALL_DOUBLE`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	PowerPoleWallMk2 = FGBuildablePowerPole{
		ClassName:                               "Build_PowerPoleWall_Mk2_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Power Pole that attaches to a wall.

Can handle up to 7 Power Line connections.

Connect Power Poles, Power Generators and factory buildings together with Power Lines to create a power grid. The power grid supplies the connected buildings with power.`,
		MDisplayName:                    `Wall Outlet Mk.2`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHasPower:                       false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MPowerConnections:               ``,
		MPowerPoleType:                  `PPT_WALL`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	PowerPoleWallMk3 = FGBuildablePowerPole{
		ClassName:                               "Build_PowerPoleWall_Mk3_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Power Pole that attaches to a wall.

Can handle up to 10 Power Line connections.

Connect Power Poles, Power Generators and factory buildings together with Power Lines to create a power grid. The power grid supplies the connected buildings with power.`,
		MDisplayName:                    `Wall Outlet Mk.3`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHasPower:                       false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MPowerConnections:               ``,
		MPowerPoleType:                  `PPT_WALL`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}
)

func GetByClassName(className string) (*FGBuildablePowerPole, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildablePowerPole with name %s", className)
}

var classNameToVar = map[string]*FGBuildablePowerPole{
	"Build_PowerPoleMk1_C":            &PowerPoleMk1,
	"Build_PowerPoleMk2_C":            &PowerPoleMk2,
	"Build_PowerPoleMk3_C":            &PowerPoleMk3,
	"Build_PowerPoleWall_C":           &PowerPoleWall,
	"Build_PowerPoleWallDouble_C":     &PowerPoleWallDouble,
	"Build_PowerPoleWallDouble_Mk2_C": &PowerPoleWallDoubleMk2,
	"Build_PowerPoleWallDouble_Mk3_C": &PowerPoleWallDoubleMk3,
	"Build_PowerPoleWall_Mk2_C":       &PowerPoleWallMk2,
	"Build_PowerPoleWall_Mk3_C":       &PowerPoleWallMk3,
}
