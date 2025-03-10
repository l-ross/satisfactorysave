// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildablePipeReservoir

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildablePipeReservoir struct {
	Name                                 string
	ClassName                            string
	MAddToSignificanceManager            bool
	MAllowColoring                       bool
	MAttachmentPoints                    string
	MBuildEffectSpeed                    float64
	MCachedSkeletalMeshes                string
	MCanChangePotential                  bool
	MCreateClearanceMeshRepresentation   bool
	MDescription                         string
	MDisplayName                         string
	MDoesHaveShutdownAnimation           bool
	MEffectUpdateInterval                float64
	MFluidBox                            string
	MFluidStackSizeDefault               resource.StackSize
	MFluidStackSizeMultiplier            int
	MForceNetUpdateOnRegisterPlayer      bool
	MHideOnBuildEffectStart              bool
	MHighlightVector                     string
	MIndicatorData                       string
	MInteractingPlayers                  string
	MIsUseable                           bool
	MMaxPotential                        float64
	MMaxPotentialIncreasePerCrystal      float64
	MMinPotential                        float64
	MMinimumProducingTime                float64
	MMinimumStoppedTime                  float64
	MNumCyclesForProductivity            int
	MOnHasPowerChanged                   string
	MOnHasProductionChanged              string
	MOnHasStandbyChanged                 string
	MPipeConnections                     string
	MPowerConsumption                    float64
	MPowerConsumptionExponent            float64
	MShouldModifyWorldGrid               bool
	MShouldShowAttachmentPointVisuals    bool
	MShouldShowHighlight                 bool
	MSignificanceRange                   float64
	MSkipBuildEffect                     bool
	MStackingHeight                      float64
	MStorageCapacity                     float64
	MToggleDormancyOnInteraction         bool
	MaxRenderDistance                    float64
	OnReplicationDetailActorCreatedEvent string
}

var (
	IndustrialTank = FGBuildablePipeReservoir{
		Name:                               "IndustrialTank",
		ClassName:                          "Build_IndustrialTank_C",
		MAddToSignificanceManager:          true,
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCachedSkeletalMeshes:              ``,
		MCanChangePotential:                false,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Can contain up to 2400m³ of fluid.
Has an input and output for pipes.`,
		MDisplayName:                         `Industrial Fluid Buffer`,
		MDoesHaveShutdownAnimation:           false,
		MEffectUpdateInterval:                0.000000,
		MFluidBox:                            `()`,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MIndicatorData:                       `()`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MPipeConnections:                     ``,
		MPowerConsumption:                    0.000000,
		MPowerConsumptionExponent:            1.600000,
		MShouldModifyWorldGrid:               true,
		MShouldShowAttachmentPointVisuals:    false,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MStackingHeight:                      1200.000000,
		MStorageCapacity:                     2400.000000,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	PipeStorageTank = FGBuildablePipeReservoir{
		Name:                               "PipeStorageTank",
		ClassName:                          "Build_PipeStorageTank_C",
		MAddToSignificanceManager:          true,
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCachedSkeletalMeshes:              ``,
		MCanChangePotential:                false,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Can contain up to 400m³ of fluid.
Has an input and output for pipes.`,
		MDisplayName:                         `Fluid Buffer`,
		MDoesHaveShutdownAnimation:           false,
		MEffectUpdateInterval:                0.000000,
		MFluidBox:                            `()`,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MIndicatorData:                       `()`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MPipeConnections:                     ``,
		MPowerConsumption:                    0.000000,
		MPowerConsumptionExponent:            1.600000,
		MShouldModifyWorldGrid:               true,
		MShouldShowAttachmentPointVisuals:    false,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MStackingHeight:                      800.000000,
		MStorageCapacity:                     400.000000,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (FGBuildablePipeReservoir, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGBuildablePipeReservoir{}, fmt.Errorf("failed to find FGBuildablePipeReservoir with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGBuildablePipeReservoir{
	"Build_IndustrialTank_C":  IndustrialTank,
	"Build_PipeStorageTank_C": PipeStorageTank,
}
