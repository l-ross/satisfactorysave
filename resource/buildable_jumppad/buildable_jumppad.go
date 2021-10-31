package BuildableJumppad

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableJumppad struct {
	ClassName                               string
	CharactersToLaunch                      string
	ComponentsToLaunch                      string
	MAddToSignificanceManager               bool
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCachedSkeletalMeshes                   string
	MCanChangePotential                     bool
	MChargeRateMultiplier                   float64
	MCurrentPowerLevel                      float64
	MDescription                            string
	MDestinationMeshHeightOffset            float64
	MDisplayName                            string
	MEffectUpdateInterval                   float64
	MExcludeFromMaterialInstancing          string
	MFluidStackSizeDefault                  resource.StackSize
	MFluidStackSizeMultiplier               int
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHasPowerForLaunch                      bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsUseable                              bool
	MKillTimer                              string
	MLaunchAngle                            float64
	MLaunchPowerCost                        float64
	MLaunchVelocity                         float64
	MMaterialNameToInstanceManager          string
	MMaxPotential                           float64
	MMaxPotentialIncreasePerCrystal         float64
	MMinPotential                           float64
	MMinimumProducingTime                   float64
	MMinimumStoppedTime                     float64
	MNumArrows                              int
	MNumCyclesForProductivity               int
	MOnHasPowerChanged                      string
	MOnHasProductionChanged                 string
	MOnHasStandbyChanged                    string
	MOnProductionStatusChanged              string
	MPlayerChainJumpResetTime               float64
	MPowerBankCapacity                      float64
	MPowerConsumption                       float64
	MPowerConsumptionExponent               float64
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MShowTrajectoryCounter                  int
	MSignificanceRange                      float64
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MTrajectoryData                         string
	MTrajectoryGravityMultiplier            float64
	MTrajectoryMeshRotation                 string
	MTrajectoryMeshScale                    string
	MaxRenderDistance                       float64
	OnReplicationDetailActorCreatedEvent    string
	VehiclesToLaunch                        string
}

var (
	JumpPadAdjustable = FGBuildableJumppad{
		ClassName:                               "Build_JumpPadAdjustable_C",
		CharactersToLaunch:                      ``,
		ComponentsToLaunch:                      ``,
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MChargeRateMultiplier:                   1.000000,
		MCurrentPowerLevel:                      0.000000,
		MDescription: `Used for quick, vertical traversal.
The launch angle can be adjusted while building.
Caution: Be sure to land safely!`,
		MDestinationMeshHeightOffset:         430.000000,
		MDisplayName:                         `Jump Pad`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHasPowerForLaunch:                   false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           false,
		MKillTimer:                           `()`,
		MLaunchAngle:                         -1.000000,
		MLaunchPowerCost:                     20.000000,
		MLaunchVelocity:                      2500.000000,
		MMaterialNameToInstanceManager:       `()`,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                0.000000,
		MMinimumStoppedTime:                  0.000000,
		MNumArrows:                           10,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MOnProductionStatusChanged:           `()`,
		MPlayerChainJumpResetTime:            8.000000,
		MPowerBankCapacity:                   100.000000,
		MPowerConsumption:                    5.000000,
		MPowerConsumptionExponent:            1.600000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MShowTrajectoryCounter:               0,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MTrajectoryData:                      `()`,
		MTrajectoryGravityMultiplier:         1.000000,
		MTrajectoryMeshRotation:              `(Pitch=-90.000000,Yaw=0.000000,Roll=-90.000000)`,
		MTrajectoryMeshScale:                 `(X=0.200000,Y=1.000000,Z=1.000000)`,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
		VehiclesToLaunch:                     ``,
	}
)

func GetByClassName(className string) (*FGBuildableJumppad, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableJumppad with name %s", className)
}

var classNameToVar = map[string]*FGBuildableJumppad{
	"Build_JumpPadAdjustable_C": &JumpPadAdjustable,
}
