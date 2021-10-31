package PipeHyperStart

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGPipeHyperStart struct {
	ClassName                                                                      string
	AudioCounterTimer                                                              string
	InterpolateEngineSound_InterpolateEngineAlpha_064FA8194B7224F6F187999413D1C8A6 float64
	InterpolateEngineSound__Direction_064FA8194B7224F6F187999413D1C8A6             string
	IsEnginePlaying                                                                bool
	MAddToSignificanceManager                                                      bool
	MAllowCleranceSeparationEvenIfStackedOn                                        bool
	MAudioTimerCounter                                                             float64
	MBlockSharingMaterialInstanceMapping                                           bool
	MBuildEffectSpeed                                                              float64
	MCachedSkeletalMeshes                                                          string
	MCanChangePotential                                                            bool
	MCanStack                                                                      bool
	MDescription                                                                   string
	MDisplayName                                                                   string
	MEffectUpdateInterval                                                          float64
	MExcludeFromMaterialInstancing                                                 string
	MFluidStackSizeDefault                                                         resource.StackSize
	MFluidStackSizeMultiplier                                                      int
	MFogPlaneTransforms                                                            string
	MForceNetUpdateOnRegisterPlayer                                                bool
	MHideOnBuildEffectStart                                                        bool
	MHighlightVector                                                               string
	MInitialMinSpeedFactor                                                         float64
	MInteractingPlayers                                                            string
	MIsUseable                                                                     bool
	MIsWindSoundPlaying                                                            bool
	MLength                                                                        float64
	MMaterialNameToInstanceManager                                                 string
	MMaxPotential                                                                  float64
	MMaxPotentialIncreasePerCrystal                                                float64
	MMinPotential                                                                  float64
	MMinimumProducingTime                                                          float64
	MMinimumStoppedTime                                                            float64
	MNumCyclesForProductivity                                                      int
	MOnHasPowerChanged                                                             string
	MOnHasProductionChanged                                                        string
	MOnHasStandbyChanged                                                           string
	MOnProductionStatusChanged                                                     string
	MOpeningOffset                                                                 float64
	MPowerConsumption                                                              float64
	MPowerConsumptionExponent                                                      float64
	MShouldModifyWorldGrid                                                         bool
	MShouldShowHighlight                                                           bool
	MSignificanceRange                                                             float64
	MSkipBuildEffect                                                               bool
	MStackHeight                                                                   float64
	MToggleDormancyOnInteraction                                                   bool
	MUseStaticHeight                                                               bool
	MWindDirectionFromTurbine                                                      string
	MaxRenderDistance                                                              float64
	OnReplicationDetailActorCreatedEvent                                           string
}

var (
	PipeHyperStart = FGPipeHyperStart{
		ClassName:         "Build_PipeHyperStart_C",
		AudioCounterTimer: `()`,
		InterpolateEngineSound_InterpolateEngineAlpha_064FA8194B7224F6F187999413D1C8A6: 0.000000,
		InterpolateEngineSound__Direction_064FA8194B7224F6F187999413D1C8A6:             `Forward`,
		IsEnginePlaying:                         false,
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MAudioTimerCounter:                      0.000000,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       300.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MCanStack:                               false,
		MDescription:                            `Used to enter and power a Hyper Tube.`,
		MDisplayName:                            `Hyper Tube Entrance`,
		MEffectUpdateInterval:                   0.000000,
		MExcludeFromMaterialInstancing:          ``,
		MFluidStackSizeDefault:                  resource.Fluid,
		MFluidStackSizeMultiplier:               1,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInitialMinSpeedFactor:                  1.200000,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MIsWindSoundPlaying:                     false,
		MLength:                                 200.000000,
		MMaterialNameToInstanceManager:          `()`,
		MMaxPotential:                           1.000000,
		MMaxPotentialIncreasePerCrystal:         0.500000,
		MMinPotential:                           0.010000,
		MMinimumProducingTime:                   2.000000,
		MMinimumStoppedTime:                     5.000000,
		MNumCyclesForProductivity:               20,
		MOnHasPowerChanged:                      `()`,
		MOnHasProductionChanged:                 `()`,
		MOnHasStandbyChanged:                    `()`,
		MOnProductionStatusChanged:              `()`,
		MOpeningOffset:                          350.000000,
		MPowerConsumption:                       10.000000,
		MPowerConsumptionExponent:               1.600000,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSignificanceRange:                      18000.000000,
		MSkipBuildEffect:                        false,
		MStackHeight:                            0.000000,
		MToggleDormancyOnInteraction:            false,
		MUseStaticHeight:                        false,
		MWindDirectionFromTurbine:               `()`,
		MaxRenderDistance:                       -1.000000,
		OnReplicationDetailActorCreatedEvent:    `()`,
	}
)

func GetByClassName(className string) (*FGPipeHyperStart, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGPipeHyperStart with name %s", className)
}

var classNameToVar = map[string]*FGPipeHyperStart{
	"Build_PipeHyperStart_C": &PipeHyperStart,
}
