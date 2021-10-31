package BuildableTradingPost

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableTradingPost struct {
	ClassName                               string
	ABClass                                 string
	MAddToSignificanceManager               bool
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCachedSkeletalMeshes                   string
	MCanChangePotential                     bool
	MDefaultResources                       string
	MDescription                            string
	MDisplayName                            string
	MEffectUpdateInterval                   float64
	MExcludeFromMaterialInstancing          string
	MFluidStackSizeDefault                  resource.StackSize
	MFluidStackSizeMultiplier               int
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MGenerators                             string
	MGroundSearchZDistance                  float64
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsUseable                              bool
	MMamFreeText                            string
	MMamOccupiedText                        string
	MMapText                                string
	MMaterialNameToInstanceManager          string
	MMaxPotential                           float64
	MMaxPotentialIncreasePerCrystal         float64
	MMeshes                                 string
	MMinPotential                           float64
	MMinimumProducingTime                   float64
	MMinimumStoppedTime                     float64
	MNeedPlayingBuildEffectNotification     bool
	MNumCyclesForProductivity               int
	MOnHasPowerChanged                      string
	MOnHasProductionChanged                 string
	MOnHasStandbyChanged                    string
	MOnProductionStatusChanged              string
	MPowerConsumption                       float64
	MPowerConsumptionExponent               float64
	MShipUpgradeLevel                       int
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSignificanceRange                      float64
	MSkeletalMeshSoftPtr                    string
	MSkipBuildEffect                        bool
	MSpawningGroundZOffset                  float64
	MStorageInventorySize                   int
	MStorageText                            string
	MStorageVisibilityLevel                 int
	MToggleDormancyOnInteraction            bool
	MWorkBenchFree                          string
	MWorkBenchOccupied                      string
	MaxRenderDistance                       float64
	Meshes                                  string
	OnReplicationDetailActorCreatedEvent    string
}

var (
	TradingPost = FGBuildableTradingPost{
		ClassName:                               "Build_TradingPost_C",
		ABClass:                                 `/Game/FactoryGame/Buildable/Factory/TradingPost/BPA_Tradingpost.BPA_Tradingpost_C`,
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MDefaultResources:                       ``,
		MDescription:                            `The heart of your factory. This is where you complete FICSIT milestones to unlock additional blueprints of buildings, vehicles, parts, equipment etc.`,
		MDisplayName:                            `The HUB`,
		MEffectUpdateInterval:                   0.000000,
		MExcludeFromMaterialInstancing:          ``,
		MFluidStackSizeDefault:                  resource.Fluid,
		MFluidStackSizeMultiplier:               1,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MGenerators:                             ``,
		MGroundSearchZDistance:                  250.000000,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MMamFreeText:                            `Use MAM`,
		MMamOccupiedText:                        `MAM is occupied`,
		MMapText:                                `The HUB`,
		MMaterialNameToInstanceManager:          `()`,
		MMaxPotential:                           1.000000,
		MMaxPotentialIncreasePerCrystal:         0.500000,
		MMeshes:                                 `(/Game/FactoryGame/Buildable/Factory/TradingPost/Mesh/Tradingpost_Stage_1.Tradingpost_Stage_1,/Game/FactoryGame/Buildable/Factory/TradingPost/Mesh/Tradingpost_Stage_2.Tradingpost_Stage_2,/Game/FactoryGame/Buildable/Factory/TradingPost/Mesh/Tradingpost_Stage_3.Tradingpost_Stage_3,/Game/FactoryGame/Buildable/Factory/TradingPost/Mesh/Tradingpost_Stage_4.Tradingpost_Stage_4,/Game/FactoryGame/Buildable/Factory/TradingPost/Mesh/Tradingpost_Stage_5.Tradingpost_Stage_5,/Game/FactoryGame/Buildable/Factory/TradingPost/Mesh/Tradingpost_Stage_6.Tradingpost_Stage_6)`,
		MMinPotential:                           0.010000,
		MMinimumProducingTime:                   2.000000,
		MMinimumStoppedTime:                     5.000000,
		MNeedPlayingBuildEffectNotification:     false,
		MNumCyclesForProductivity:               20,
		MOnHasPowerChanged:                      `()`,
		MOnHasProductionChanged:                 `()`,
		MOnHasStandbyChanged:                    `()`,
		MOnProductionStatusChanged:              `()`,
		MPowerConsumption:                       0.000000,
		MPowerConsumptionExponent:               1.600000,
		MShipUpgradeLevel:                       6,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSignificanceRange:                      18000.000000,
		MSkeletalMeshSoftPtr:                    `/Game/FactoryGame/Buildable/Factory/TradingPost/Mesh/SK_Tradingpost.SK_Tradingpost`,
		MSkipBuildEffect:                        false,
		MSpawningGroundZOffset:                  5.000000,
		MStorageInventorySize:                   10,
		MStorageText:                            `Open Storage`,
		MStorageVisibilityLevel:                 1,
		MToggleDormancyOnInteraction:            false,
		MWorkBenchFree:                          `Use Craft Bench`,
		MWorkBenchOccupied:                      `Craft Bench occupied`,
		MaxRenderDistance:                       -1.000000,
		Meshes:                                  `(StaticMesh'"/Game/FactoryGame/Buildable/Factory/TradingPost/Mesh/Tradingpost_Stage_1.Tradingpost_Stage_1"',StaticMesh'"/Game/FactoryGame/Buildable/Factory/TradingPost/Mesh/Tradingpost_Stage_2.Tradingpost_Stage_2"',StaticMesh'"/Game/FactoryGame/Buildable/Factory/TradingPost/Mesh/Tradingpost_Stage_3.Tradingpost_Stage_3"',StaticMesh'"/Game/FactoryGame/Buildable/Factory/TradingPost/Mesh/Tradingpost_Stage_4.Tradingpost_Stage_4"',StaticMesh'"/Game/FactoryGame/Buildable/Factory/TradingPost/Mesh/Tradingpost_Stage_5.Tradingpost_Stage_5"',StaticMesh'"/Game/FactoryGame/Buildable/Factory/TradingPost/Mesh/Tradingpost_Stage_6.Tradingpost_Stage_6"')`,
		OnReplicationDetailActorCreatedEvent:    `()`,
	}
)

func GetByClassName(className string) (*FGBuildableTradingPost, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableTradingPost with name %s", className)
}

var classNameToVar = map[string]*FGBuildableTradingPost{
	"Build_TradingPost_C": &TradingPost,
}
