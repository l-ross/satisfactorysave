package BuildableGeneratorGeoThermal

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableGeneratorGeoThermal struct {
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
	MLoadPercentage                         float64
	MMaterialNameToInstanceManager          string
	MMaxPotential                           float64
	MMaxPotentialIncreasePerCrystal         float64
	MMaxPowerProduction                     float64
	MMinPotential                           float64
	MMinPowerProduction                     float64
	MMinimumProducingTime                   float64
	MMinimumStoppedTime                     float64
	MNumCyclesForProductivity               int
	MOnHasPowerChanged                      string
	MOnHasProductionChanged                 string
	MOnHasStandbyChanged                    string
	MOnProductionStatusChanged              string
	MPowerConsumption                       float64
	MPowerConsumptionExponent               float64
	MPowerProduction                        float64
	MPowerProductionExponent                float64
	MProductionEffectsRunning               bool
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSignificanceRange                      float64
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MVariablePowerProductionConstant        float64
	MVariablePowerProductionCycleLength     float64
	MVariablePowerProductionCycleOffset     float64
	MVariablePowerProductionFactor          float64
	MaxRenderDistance                       float64
	OnReplicationDetailActorCreatedEvent    string
}

var (
	GeneratorGeoThermal = FGBuildableGeneratorGeoThermal{
		ClassName:                               "Build_GeneratorGeoThermal_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MDescription: `Has to be built on a Geyser. Generates power.

Caution: Power production fluctuates.

Power Production:
Impure 50-150 MW (100 MW average)
Normal 100-300 MW (200 MW average)
Pure 200-600 MW (400 MW average)`,
		MDisplayName:                         `Geothermal Generator`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MLoadPercentage:                      0.000000,
		MMaterialNameToInstanceManager:       `()`,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMaxPowerProduction:                  0.000000,
		MMinPotential:                        0.010000,
		MMinPowerProduction:                  0.000000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MOnProductionStatusChanged:           `()`,
		MPowerConsumption:                    0.000000,
		MPowerConsumptionExponent:            1.600000,
		MPowerProduction:                     0.000000,
		MPowerProductionExponent:             1.300000,
		MProductionEffectsRunning:            false,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MVariablePowerProductionConstant:     0.000000,
		MVariablePowerProductionCycleLength:  60.000000,
		MVariablePowerProductionCycleOffset:  0.000000,
		MVariablePowerProductionFactor:       200.000000,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (*FGBuildableGeneratorGeoThermal, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableGeneratorGeoThermal with name %s", className)
}

var classNameToVar = map[string]*FGBuildableGeneratorGeoThermal{
	"Build_GeneratorGeoThermal_C": &GeneratorGeoThermal,
}
