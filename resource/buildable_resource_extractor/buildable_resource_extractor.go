package BuildableResourceExtractor

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableResourceExtractor struct {
	ClassName                               string
	CanPlayAfterStartUpStopped              bool
	MAddToSignificanceManager               bool
	MAllowCleranceSeparationEvenIfStackedOn bool
	MAllowedResourceForms                   string
	MAllowedResources                       string
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCachedSkeletalMeshes                   string
	MCanChangePotential                     bool
	MCanPlayAfterStartUpStopped             bool
	MDescription                            string
	MDisplayName                            string
	MEffectUpdateInterval                   float64
	MExcludeFromMaterialInstancing          string
	MExtractCycleTime                       float64
	MExtractStartupTime                     float64
	MExtractStartupTimer                    float64
	MExtractorTypeName                      string
	MFluidStackSizeDefault                  resource.StackSize
	MFluidStackSizeMultiplier               int
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MInternalMiningState_0                  string
	MIsUseable                              bool
	MItemsPerCycle                          int
	MMaterialNameToInstanceManager          string
	MMaxPotential                           float64
	MMaxPotentialIncreasePerCrystal         float64
	MMaximumDrillTime_0                     float64
	MMinPotential                           float64
	MMinimumDrillTime_0                     float64
	MMinimumProducingTime                   float64
	MMinimumStoppedTime                     float64
	MNumCyclesForProductivity               int
	MOnHasPowerChanged                      string
	MOnHasProductionChanged                 string
	MOnHasStandbyChanged                    string
	MOnProductionStatusChanged              string
	MOnlyAllowCertainResources              bool
	MParticleMap                            string
	MPipeOutputConnections                  string
	MPowerConsumption                       float64
	MPowerConsumptionExponent               float64
	MReplicatedFlowRate                     float64
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSignificanceRange                      float64
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MToggleMiningStateHandle_0              string
	MaxRenderDistance                       float64
	OnReplicationDetailActorCreatedEvent    string
}

var (
	MinerMk1 = FGBuildableResourceExtractor{
		ClassName:                               "Build_MinerMk1_C",
		CanPlayAfterStartUpStopped:              false,
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MAllowedResourceForms:                   `(RF_SOLID)`,
		MAllowedResources:                       ``,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MDescription: `Extracts solid resources from the resource node it is built on. 
The normal extraction rate is 60 resources per minute. 
The extraction rate is modified depending on resource node purity. Outputs all extracted resources onto connected conveyor belts.`,
		MDisplayName:                         `Miner Mk.1`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MExtractCycleTime:                    1.000000,
		MExtractStartupTime:                  15.000000,
		MExtractStartupTimer:                 10.000000,
		MExtractorTypeName:                   `Miner`,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  `((Translation=(X=0.000000,Y=907.217163,Z=177.884918)))`,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MItemsPerCycle:                       1,
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
		MOnlyAllowCertainResources:           false,
		MParticleMap:                         `((ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/Coal/Desc_Coal.Desc_Coal_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Coal_01.P_Mining_Coal_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/Stone/Desc_Stone.Desc_Stone_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Iron_01.P_Mining_Iron_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/OreIron/Desc_OreIron.Desc_OreIron_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Iron_01.P_Mining_Iron_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/OreBauxite/Desc_OreBauxite.Desc_OreBauxite_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Bauxite_01.P_Mining_Bauxite_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/OreCopper/Desc_OreCopper.Desc_OreCopper_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Copper_01.P_Mining_Copper_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/CrudeOil/Desc_LiquidOil.Desc_LiquidOil_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/Buildable/Factory/MinerMk2/Particle/Mining.Mining"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/OreGold/Desc_OreGold.Desc_OreGold_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Gold_01.P_Mining_Gold_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/RawQuartz/Desc_RawQuartz.Desc_RawQuartz_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Quartz_01.P_Mining_Quartz_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/Sulfur/Desc_Sulfur.Desc_Sulfur_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Sulfur_01.P_Mining_Sulfur_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/OreUranium/Desc_OreUranium.Desc_OreUranium_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Uranium_01.P_Mining_Uranium_01"'))`,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    5.000000,
		MPowerConsumptionExponent:            1.600000,
		MReplicatedFlowRate:                  0.000000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	MinerMk2 = FGBuildableResourceExtractor{
		ClassName:                               "Build_MinerMk2_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MAllowedResourceForms:                   `(RF_SOLID)`,
		MAllowedResources:                       ``,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MCanPlayAfterStartUpStopped:             false,
		MDescription: `Extracts solid resources from the resource node it is built on. 
The normal extraction rate is 120 resources per minute. 
The extraction rate is modified depending on resource node purity. Outputs all extracted resources onto connected conveyor belts.`,
		MDisplayName:                         `Miner Mk.2`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MExtractCycleTime:                    0.500000,
		MExtractStartupTime:                  15.000000,
		MExtractStartupTimer:                 10.000000,
		MExtractorTypeName:                   `Miner`,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  `((Translation=(X=-0.000900,Y=914.780884,Z=175.000000)))`,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MItemsPerCycle:                       1,
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
		MOnlyAllowCertainResources:           false,
		MParticleMap:                         `((ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/Coal/Desc_Coal.Desc_Coal_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Coal_01.P_Mining_Coal_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/Stone/Desc_Stone.Desc_Stone_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Iron_01.P_Mining_Iron_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/OreIron/Desc_OreIron.Desc_OreIron_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Iron_01.P_Mining_Iron_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/OreBauxite/Desc_OreBauxite.Desc_OreBauxite_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Bauxite_01.P_Mining_Bauxite_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/OreCopper/Desc_OreCopper.Desc_OreCopper_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Copper_01.P_Mining_Copper_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/CrudeOil/Desc_LiquidOil.Desc_LiquidOil_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/Buildable/Factory/MinerMk2/Particle/Mining.Mining"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/OreGold/Desc_OreGold.Desc_OreGold_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Gold_01.P_Mining_Gold_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/RawQuartz/Desc_RawQuartz.Desc_RawQuartz_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Quartz_01.P_Mining_Quartz_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/Sulfur/Desc_Sulfur.Desc_Sulfur_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Sulfur_01.P_Mining_Sulfur_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/OreUranium/Desc_OreUranium.Desc_OreUranium_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Uranium_01.P_Mining_Uranium_01"'))`,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    12.000000,
		MPowerConsumptionExponent:            1.600000,
		MReplicatedFlowRate:                  0.000000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	MinerMk3 = FGBuildableResourceExtractor{
		ClassName:                               "Build_MinerMk3_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MAllowedResourceForms:                   `(RF_SOLID)`,
		MAllowedResources:                       ``,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MCanPlayAfterStartUpStopped:             false,
		MDescription: `Extracts solid resources from the resource node it is built on. 
The normal extraction rate is 240 resources per minute. 
The extraction rate is modified depending on resource node purity. Outputs all extracted resources onto connected conveyor belts.`,
		MDisplayName:                         `Miner Mk.3`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MExtractCycleTime:                    0.250000,
		MExtractStartupTime:                  15.000000,
		MExtractStartupTimer:                 10.000000,
		MExtractorTypeName:                   `Miner`,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  `((Translation=(X=-0.000900,Y=914.780884,Z=175.000000)))`,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MInternalMiningState_0:               `NewEnumerator0`,
		MIsUseable:                           true,
		MItemsPerCycle:                       1,
		MMaterialNameToInstanceManager:       `()`,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMaximumDrillTime_0:                  0.000000,
		MMinPotential:                        0.010000,
		MMinimumDrillTime_0:                  0.000000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MOnProductionStatusChanged:           `()`,
		MOnlyAllowCertainResources:           false,
		MParticleMap:                         `((ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/Coal/Desc_Coal.Desc_Coal_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Coal_01.P_Mining_Coal_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/Stone/Desc_Stone.Desc_Stone_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Iron_01.P_Mining_Iron_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/OreIron/Desc_OreIron.Desc_OreIron_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Iron_01.P_Mining_Iron_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/OreBauxite/Desc_OreBauxite.Desc_OreBauxite_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Bauxite_01.P_Mining_Bauxite_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/OreCopper/Desc_OreCopper.Desc_OreCopper_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Copper_01.P_Mining_Copper_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/CrudeOil/Desc_LiquidOil.Desc_LiquidOil_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/Buildable/Factory/MinerMk2/Particle/Mining.Mining"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/OreGold/Desc_OreGold.Desc_OreGold_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Gold_01.P_Mining_Gold_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/RawQuartz/Desc_RawQuartz.Desc_RawQuartz_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Quartz_01.P_Mining_Quartz_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/Sulfur/Desc_Sulfur.Desc_Sulfur_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Sulfur_01.P_Mining_Sulfur_01"'),(ResourceNode_16_2100B5C34EE8DF7958D78A974512F3C3=/Game/FactoryGame/Resource/RawResources/OreUranium/Desc_OreUranium.Desc_OreUranium_C,ParticleSystem1_9_F0CF81514E1E1C5007AC99B0C59C63CD=ParticleSystem'"/Game/FactoryGame/VFX/Factory/Miner/P_Mining_Uranium_01.P_Mining_Uranium_01"'))`,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    30.000000,
		MPowerConsumptionExponent:            1.600000,
		MReplicatedFlowRate:                  0.000000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MToggleMiningStateHandle_0:           `()`,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	OilPump = FGBuildableResourceExtractor{
		ClassName:                               "Build_OilPump_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MAllowedResourceForms:                   `(RF_LIQUID)`,
		MAllowedResources:                       `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/CrudeOil/Desc_LiquidOil.Desc_LiquidOil_C"')`,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MDescription: `Normal extraction rate: 120m³ oil per minute.
Head Lift: 10 meters.
(Allows fluids to be transported 10 meters upwards.)

Can be built on an Oil Node to extract Crude Oil. Extraction rates depend on node purity.`,
		MDisplayName:                         `Oil Extractor`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MExtractCycleTime:                    1.000000,
		MExtractStartupTime:                  -10.000000,
		MExtractStartupTimer:                 10.000000,
		MExtractorTypeName:                   `None`,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MItemsPerCycle:                       2000,
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
		MOnlyAllowCertainResources:           true,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    40.000000,
		MPowerConsumptionExponent:            1.600000,
		MReplicatedFlowRate:                  0.000000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (*FGBuildableResourceExtractor, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableResourceExtractor with name %s", className)
}

var classNameToVar = map[string]*FGBuildableResourceExtractor{
	"Build_MinerMk1_C": &MinerMk1,
	"Build_MinerMk2_C": &MinerMk2,
	"Build_MinerMk3_C": &MinerMk3,
	"Build_OilPump_C":  &OilPump,
}
