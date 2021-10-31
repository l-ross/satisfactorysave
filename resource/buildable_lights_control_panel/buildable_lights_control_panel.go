package BuildableLightsControlPanel

import (
	"fmt"
)

type FGBuildableLightsControlPanel struct {
	ClassName                               string
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MConnections                            string
	MControlledBuildables                   string
	MDescription                            string
	MDisplayName                            string
	MExcludeFromMaterialInstancing          string
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsBridgeConnected                      bool
	MIsEnabled                              bool
	MIsUseable                              bool
	MLightControlData                       string
	MMaterialNameToInstanceManager          string
	MOnControlledBuildablesChanged          string
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
}

var (
	LightsControlPanel = FGBuildableLightsControlPanel{
		ClassName:                               "Build_LightsControlPanel_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MConnections:                            ``,
		MControlledBuildables:                   ``,
		MDescription: `Useful for sectioning and modifying many lights at once.

Controls all Lights connected to the Power Grid attached to the 'Light Power Connector'.
(Other Control Panels and Power Switches interrupt the connection.)`,
		MDisplayName:                    `Lights Control Panel`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsBridgeConnected:              false,
		MIsEnabled:                      false,
		MIsUseable:                      true,
		MLightControlData:               `(Intensity=1.000000)`,
		MMaterialNameToInstanceManager:  `()`,
		MOnControlledBuildablesChanged:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}
)

func GetByClassName(className string) (*FGBuildableLightsControlPanel, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableLightsControlPanel with name %s", className)
}

var classNameToVar = map[string]*FGBuildableLightsControlPanel{
	"Build_LightsControlPanel_C": &LightsControlPanel,
}
