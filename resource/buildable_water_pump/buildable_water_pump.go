package BuildableWaterPump

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableWaterPump struct {
	ClassName                                                                  string
	HasLostSignificance                                                        bool
	MAddToSignificanceManager                                                  bool
	MAllowCleranceSeparationEvenIfStackedOn                                    bool
	MAllowedResourceForms                                                      string
	MAllowedResources                                                          string
	MAudioTimelineCounter                                                      float64
	MAudioTimerCounter                                                         float64
	MAudioTimerMS                                                              float64
	MAudioTimerReference                                                       string
	MBlockSharingMaterialInstanceMapping                                       bool
	MBuildEffectSpeed                                                          float64
	MCachedSkeletalMeshes                                                      string
	MCanChangePotential                                                        bool
	MDepthTraceOriginOffset                                                    string
	MDescription                                                               string
	MDisplayName                                                               string
	MEffectUpdateInterval                                                      float64
	MExcludeFromMaterialInstancing                                             string
	MExtractCycleTime                                                          float64
	MExtractStartupTime                                                        float64
	MExtractStartupTimer                                                       float64
	MExtractorTypeName                                                         string
	MFluidStackSizeDefault                                                     resource.StackSize
	MFluidStackSizeMultiplier                                                  int
	MFogPlaneTransforms                                                        string
	MForceNetUpdateOnRegisterPlayer                                            bool
	MHideOnBuildEffectStart                                                    bool
	MHighlightVector                                                           string
	MInteractingPlayers                                                        string
	MIsUseable                                                                 bool
	MItemsPerCycle                                                             int
	MMaterialNameToInstanceManager                                             string
	MMaxPotential                                                              float64
	MMaxPotentialIncreasePerCrystal                                            float64
	MMinPotential                                                              float64
	MMinimumDepthForPlacement                                                  float64
	MMinimumProducingTime                                                      float64
	MMinimumStoppedTime                                                        float64
	MNumCyclesForProductivity                                                  int
	MOnHasPowerChanged                                                         string
	MOnHasProductionChanged                                                    string
	MOnHasStandbyChanged                                                       string
	MOnProductionStatusChanged                                                 string
	MOnlyAllowCertainResources                                                 bool
	MPipeOutputConnections                                                     string
	MPowerConsumption                                                          float64
	MPowerConsumptionExponent                                                  float64
	MReplicatedFlowRate                                                        float64
	MShouldModifyWorldGrid                                                     bool
	MShouldShowHighlight                                                       bool
	MSignificanceRange                                                         float64
	MSkipBuildEffect                                                           bool
	MToggleDormancyOnInteraction                                               bool
	MWaterpumpTimeline_RTPC_B8FA6F944E717E3B7A286E84901F620E                   float64
	MWaterpumpTimeline__Direction_B8FA6F944E717E3B7A286E84901F620E             string
	MaxRenderDistance                                                          float64
	OnReplicationDetailActorCreatedEvent                                       string
	PlayPitchAndVolumeRTPCTimeline_RTPC_2B435F41466C37D2AD809A88AA21BA89       float64
	PlayPitchAndVolumeRTPCTimeline__Direction_2B435F41466C37D2AD809A88AA21BA89 string
}

var (
	WaterPump = FGBuildableWaterPump{
		ClassName:                               "Build_WaterPump_C",
		HasLostSignificance:                     false,
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MAllowedResourceForms:                   `(RF_LIQUID)`,
		MAllowedResources:                       `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/Water/Desc_Water.Desc_Water_C"')`,
		MAudioTimelineCounter:                   0.000000,
		MAudioTimerCounter:                      0.000000,
		MAudioTimerMS:                           0.100000,
		MAudioTimerReference:                    `()`,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MDepthTraceOriginOffset:                 `(X=0.000000,Y=300.000000,Z=200.000000)`,
		MDescription: `Default extraction rate: 120m³ water per minute.
Head Lift: 10 meters.
(Allows fluids to be transported 10 meters upwards.)

Extracts water from the body of water it is built on.
Note that the water needs to be deep enough and that rivers do not commonly suffice.`,
		MDisplayName:                    `Water Extractor`,
		MEffectUpdateInterval:           0.000000,
		MExcludeFromMaterialInstancing:  ``,
		MExtractCycleTime:               1.000000,
		MExtractStartupTime:             10.000000,
		MExtractStartupTimer:            10.000000,
		MExtractorTypeName:              `None`,
		MFluidStackSizeDefault:          resource.Fluid,
		MFluidStackSizeMultiplier:       4,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      true,
		MItemsPerCycle:                  2000,
		MMaterialNameToInstanceManager:  `()`,
		MMaxPotential:                   1.000000,
		MMaxPotentialIncreasePerCrystal: 0.500000,
		MMinPotential:                   0.010000,
		MMinimumDepthForPlacement:       50.000000,
		MMinimumProducingTime:           2.000000,
		MMinimumStoppedTime:             5.000000,
		MNumCyclesForProductivity:       20,
		MOnHasPowerChanged:              `()`,
		MOnHasProductionChanged:         `()`,
		MOnHasStandbyChanged:            `()`,
		MOnProductionStatusChanged:      `()`,
		MOnlyAllowCertainResources:      true,
		MPipeOutputConnections:          ``,
		MPowerConsumption:               20.000000,
		MPowerConsumptionExponent:       1.600000,
		MReplicatedFlowRate:             0.000000,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSignificanceRange:              18000.000000,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWaterpumpTimeline_RTPC_B8FA6F944E717E3B7A286E84901F620E:       0.000000,
		MWaterpumpTimeline__Direction_B8FA6F944E717E3B7A286E84901F620E: `Forward`,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
		PlayPitchAndVolumeRTPCTimeline_RTPC_2B435F41466C37D2AD809A88AA21BA89:       0.000000,
		PlayPitchAndVolumeRTPCTimeline__Direction_2B435F41466C37D2AD809A88AA21BA89: `Forward`,
	}
)

func GetByClassName(className string) (*FGBuildableWaterPump, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableWaterPump with name %s", className)
}

var classNameToVar = map[string]*FGBuildableWaterPump{
	"Build_WaterPump_C": &WaterPump,
}
