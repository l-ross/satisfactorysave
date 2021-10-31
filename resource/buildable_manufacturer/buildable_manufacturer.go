package BuildableManufacturer

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableManufacturer struct {
	ClassName                               string
	CurrentPackagingMode                    string
	IsPowered                               bool
	MAddToSignificanceManager               bool
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCachedSkeletalMeshes                   string
	MCanChangePotential                     bool
	MColor                                  string
	MCurrentColorVFX                        string
	MCurrentColor_VFX                       string
	MCurrentRecipeChanged                   string
	MCurrentRecipeCheck                     string
	MDescription                            string
	MDisplayName                            string
	MEffectUpdateInterval                   float64
	MExcludeFromMaterialInstancing          string
	MFactoryInputConnections                string
	MFactoryOutputConnections               string
	MFluidStackSizeDefault                  resource.StackSize
	MFluidStackSizeMultiplier               int
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsRadioActive                          bool
	MIsUseable                              bool
	MManufacturingSpeed                     float64
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
	MPipeInputConnections                   string
	MPipeOutputConnections                  string
	MPowerConsumption                       float64
	MPowerConsumptionExponent               float64
	MPreviousRecipeCheck                    string
	MProductionEffectsRunning               bool
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSignificanceRange                      float64
	MSkipBuildEffect                        bool
	MSocketStoppedAkComponents              string
	MStoppedAkComponents                    string
	MStoppedProducingAnimationSounds        bool
	MToggleDormancyOnInteraction            bool
	M_NotifyNameREferences                  string
	MaxRenderDistance                       float64
	OnReplicationDetailActorCreatedEvent    string
}

var (
	AssemblerMk1 = FGBuildableManufacturer{
		ClassName:                               "Build_AssemblerMk1_C",
		IsPowered:                               false,
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MCurrentRecipeChanged:                   `()`,
		MDescription: `Crafts two parts into another part.

Can be automated by feeding parts into it with a conveyor belt connected to the input. The produced parts can be automatically extracted by connecting a conveyor belt to the output.`,
		MDisplayName:                         `Assembler`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFactoryInputConnections:             ``,
		MFactoryOutputConnections:            ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  `((Translation=(X=0.000230,Y=669.259399,Z=180.000000)),(Rotation=(X=-0.000000,Y=0.000000,Z=-1.000000,W=0.000004),Translation=(X=200.001419,Y=-646.223206,Z=180.000000)),(Rotation=(X=-0.000000,Y=0.000000,Z=-1.000000,W=0.000007),Translation=(X=-199.998581,Y=-646.223206,Z=180.000000)))`,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MManufacturingSpeed:                  1.000000,
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
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    15.000000,
		MPowerConsumptionExponent:            1.600000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	Blender = FGBuildableManufacturer{
		ClassName:                               "Build_Blender_C",
		IsPowered:                               false,
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MColor:                                  ``,
		MCurrentColorVFX:                        `(R=0.000000,G=0.000000,B=0.000000,A=0.000000)`,
		MCurrentRecipeChanged:                   `()`,
		MDescription: `The Blender is capable of blending fluids and combining them with solid parts in various processes.
Head Lift: 10 meters.
(Allows fluids to be transported 10 meters upwards).

Contains both Conveyor Belt and Pipe inputs and outputs.`,
		MDisplayName:                         `Blender`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFactoryInputConnections:             ``,
		MFactoryOutputConnections:            ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsRadioActive:                       false,
		MIsUseable:                           true,
		MManufacturingSpeed:                  1.000000,
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
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    75.000000,
		MPowerConsumptionExponent:            1.600000,
		MProductionEffectsRunning:            false,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		M_NotifyNameREferences:               `("Arm_04_ClawBase","Arm_02_SFXSocket")`,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	ConstructorMk1 = FGBuildableManufacturer{
		ClassName:                               "Build_ConstructorMk1_C",
		IsPowered:                               false,
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MCurrentRecipeChanged:                   `()`,
		MCurrentRecipeCheck:                     ``,
		MDescription: `Crafts one part into another part.

Can be automated by feeding parts into it with a conveyor belt connected to the input. The produced parts can be automatically extracted by connecting a conveyor belt to the output.`,
		MDisplayName:                         `Constructor`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFactoryInputConnections:             ``,
		MFactoryOutputConnections:            ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MManufacturingSpeed:                  1.000000,
		MMaterialNameToInstanceManager:       `()`,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                0.000000,
		MMinimumStoppedTime:                  0.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MOnProductionStatusChanged:           `()`,
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    4.000000,
		MPowerConsumptionExponent:            1.600000,
		MPreviousRecipeCheck:                 ``,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	FoundryMk1 = FGBuildableManufacturer{
		ClassName:                               "Build_FoundryMk1_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MCurrentRecipeChanged:                   `()`,
		MDescription: `Smelts two resources into alloy ingots.

Can be automated by feeding ore into it with a conveyor belt connected to the inputs. The produced ingots can be automatically extracted by connecting a conveyor belt to the output.`,
		MDisplayName:                         `Foundry`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFactoryInputConnections:             ``,
		MFactoryOutputConnections:            ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  `((Rotation=(X=0.000000,Y=0.000000,Z=-1.000000,W=-0.000001),Translation=(X=199.744049,Y=-462.004303,Z=173.798508)),(Rotation=(X=0.000000,Y=0.000000,Z=-1.000000,W=-0.000001),Translation=(X=-200.372986,Y=-462.004547,Z=173.798508)),(Rotation=(X=-0.000000,Y=0.000000,Z=-0.000001,W=1.000000),Translation=(X=-198.404800,Y=274.042969,Z=173.798508)))`,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MManufacturingSpeed:                  1.000000,
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
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    16.000000,
		MPowerConsumptionExponent:            1.600000,
		MProductionEffectsRunning:            false,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	ManufacturerMk1 = FGBuildableManufacturer{
		ClassName:                               "Build_ManufacturerMk1_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MCurrentRecipeChanged:                   `()`,
		MDescription: `Crafts three or four parts into another part.

Can be automated by feeding parts into it with a conveyor belt connected to the input. The produced parts can be automatically extracted by connecting a conveyor belt to the output.`,
		MDisplayName:                         `Manufacturer`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFactoryInputConnections:             ``,
		MFactoryOutputConnections:            ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  `((Rotation=(X=-0.000000,Y=0.000000,Z=-1.000000,W=0.000005),Translation=(X=-200.000000,Y=-900.000000,Z=140.000000)),(Rotation=(X=-0.000000,Y=0.000000,Z=-1.000000,W=0.000005),Translation=(X=600.000000,Y=-900.000000,Z=140.000000)),(Rotation=(X=-0.000000,Y=0.000000,Z=-1.000000,W=0.000005),Translation=(X=200.000000,Y=-900.000000,Z=140.000000)),(Rotation=(X=-0.000000,Y=0.000000,Z=-1.000000,W=0.000005),Translation=(X=-600.000000,Y=-900.000000,Z=140.000000)),(Rotation=(X=0.000000,Y=-0.000000,Z=0.000010,W=1.000000),Translation=(X=0.000007,Y=840.000000,Z=140.000000)))`,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MManufacturingSpeed:                  1.000000,
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
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    55.000000,
		MPowerConsumptionExponent:            1.600000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MSocketStoppedAkComponents:           ``,
		MStoppedAkComponents:                 ``,
		MStoppedProducingAnimationSounds:     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	OilRefinery = FGBuildableManufacturer{
		ClassName:                               "Build_OilRefinery_C",
		IsPowered:                               false,
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MCurrentRecipeChanged:                   `()`,
		MDescription: `Refines fluid and/or solid parts into other parts.
Head Lift: 10 meters.
(Allows fluids to be transported 10 meters upwards.)

Contains both a Conveyor Belt and Pipe input and output, to allow for the automation of various recipes.`,
		MDisplayName:                         `Refinery`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFactoryInputConnections:             ``,
		MFactoryOutputConnections:            ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MManufacturingSpeed:                  1.000000,
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
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    30.000000,
		MPowerConsumptionExponent:            1.600000,
		MProductionEffectsRunning:            false,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   17000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	Packager = FGBuildableManufacturer{
		ClassName:                               "Build_Packager_C",
		CurrentPackagingMode:                    ``,
		IsPowered:                               false,
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MCurrentColor_VFX:                       `(R=1.000000,G=1.000000,B=1.000000,A=0.000000)`,
		MCurrentRecipeChanged:                   `()`,
		MDescription: `Used for the packaging and unpacking of fluids.
Head Lift: 10 meters.
(Allows fluids to be transported 10 meters upwards.)

Contains both a Conveyor Belt and Pipe input and output, to allow for the automation of various recipes.`,
		MDisplayName:                         `Packager`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFactoryInputConnections:             ``,
		MFactoryOutputConnections:            ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MManufacturingSpeed:                  1.000000,
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
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    10.000000,
		MPowerConsumptionExponent:            1.600000,
		MProductionEffectsRunning:            false,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	SmelterMk1 = FGBuildableManufacturer{
		ClassName:                               "Build_SmelterMk1_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MCurrentRecipeChanged:                   `()`,
		MDescription: `Smelts ore into ingots.

Can be automated by feeding ore into it with a conveyor belt connected to the input. The produced ingots can be automatically extracted by connecting a conveyor belt to the output.`,
		MDisplayName:                         `Smelter`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFactoryInputConnections:             ``,
		MFactoryOutputConnections:            ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MManufacturingSpeed:                  1.000000,
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
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    4.000000,
		MPowerConsumptionExponent:            1.600000,
		MProductionEffectsRunning:            false,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (*FGBuildableManufacturer, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableManufacturer with name %s", className)
}

var classNameToVar = map[string]*FGBuildableManufacturer{
	"Build_AssemblerMk1_C":    &AssemblerMk1,
	"Build_Blender_C":         &Blender,
	"Build_ConstructorMk1_C":  &ConstructorMk1,
	"Build_FoundryMk1_C":      &FoundryMk1,
	"Build_ManufacturerMk1_C": &ManufacturerMk1,
	"Build_OilRefinery_C":     &OilRefinery,
	"Build_Packager_C":        &Packager,
	"Build_SmelterMk1_C":      &SmelterMk1,
}
