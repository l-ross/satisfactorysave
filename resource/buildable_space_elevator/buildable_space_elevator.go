package BuildableSpaceElevator

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableSpaceElevator struct {
	ClassName                               string
	MAddToSignificanceManager               bool
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCachedSkeletalMeshes                   string
	MCanChangePotential                     bool
	MDescription                            string
	MDisplayName                            string
	MEffectUpdateInterval                   float64
	MExcludeFromMaterialInstancing          string
	MFluidStackSizeDefault                  resource.StackSize
	MFluidStackSizeMultiplier               int
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsUseable                              bool
	MMapText                                string
	MMaterialNameToInstanceManager          string
	MMaxPotential                           float64
	MMaxPotentialIncreasePerCrystal         float64
	MMinPotential                           float64
	MMinimumProducingTime                   float64
	MMinimumStoppedTime                     float64
	MNumCyclesForProductivity               int
	MOnHasPowerChanged                      string
	MOnHasProductionChanged                 string
	MOnHasStandbyChanged                    string
	MOnProductionStatusChanged              string
	MPowerConsumption                       float64
	MPowerConsumptionExponent               float64
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSignificanceRange                      float64
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
	OnReplicationDetailActorCreatedEvent    string
}

var (
	SpaceElevator = FGBuildableSpaceElevator{
		ClassName:                               "Build_SpaceElevator_C",
		MAddToSignificanceManager:               false,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MDescription: `Requires deliveries of special Project Parts to complete Phases of Project Assembly.
Completing Phases in the Space Elevator will unlock new Tiers in the HUB Terminal.`,
		MDisplayName:                         `Space Elevator`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=-2010.000000,Y=0.000000,Z=240.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MMapText:                             `Space Elevator`,
		MMaterialNameToInstanceManager:       `()`,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MOnProductionStatusChanged:           `()`,
		MPowerConsumption:                    0.000000,
		MPowerConsumptionExponent:            1.600000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 true,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (*FGBuildableSpaceElevator, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableSpaceElevator with name %s", className)
}

var classNameToVar = map[string]*FGBuildableSpaceElevator{
	"Build_SpaceElevator_C": &SpaceElevator,
}
