// Code generated by ../../gen/docs_json. DO NOT EDIT.

package VehicleDescriptor

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGVehicleDescriptor struct {
	ClassName               string
	MAbbreviatedDisplayName string
	MBuildMenuPriority      float64
	MCanBeDiscarded         bool
	MDescription            string
	MDisplayName            string
	MEnergyValue            float64
	MFluidColor             string
	MForm                   resource.Form
	MFuelConsumption        float64
	MGasColor               string
	MInventorySize          int
	MPersistentBigIcon      string
	MPowerConsumption       string
	MRadioactiveDecay       float64
	MRememberPickUp         bool
	MSmallIcon              string
	MStackSize              resource.StackSize
	MSubCategories          string
}

var (
	CyberWagon = FGVehicleDescriptor{
		ClassName:               "Desc_CyberWagon_C",
		MAbbreviatedDisplayName: ``,
		MBuildMenuPriority:      4.000000,
		MCanBeDiscarded:         true,
		MDescription:            ``,
		MDisplayName:            ``,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Invalid,
		MFuelConsumption:        150.000000,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MInventorySize:          1,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Buildable/Vehicle/Cyberwagon/UI/Cyberwagon_512.Cyberwagon_512`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Buildable/Vehicle/Cyberwagon/UI/Cyberwagon_256.Cyberwagon_256`,
		MStackSize:              resource.One,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Interface/UI/InGame/BuildMenu/BuildCategories/Sub_Transport/SC_Vehicles.SC_Vehicles_C"')`,
	}

	DroneTransport = FGVehicleDescriptor{
		ClassName:               "Desc_DroneTransport_C",
		MAbbreviatedDisplayName: ``,
		MBuildMenuPriority:      6.000000,
		MCanBeDiscarded:         true,
		MDescription:            ``,
		MDisplayName:            ``,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Invalid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MInventorySize:          9,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Buildable/Factory/DroneStation/UI/IconDesc_Drone_512.IconDesc_Drone_512`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Buildable/Factory/DroneStation/UI/IconDesc_Drone_256.IconDesc_Drone_256`,
		MStackSize:              resource.Medium,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Interface/UI/InGame/BuildMenu/BuildCategories/Sub_Transport/SC_Vehicles.SC_Vehicles_C"')`,
	}

	Explorer = FGVehicleDescriptor{
		ClassName:               "Desc_Explorer_C",
		MAbbreviatedDisplayName: ``,
		MBuildMenuPriority:      3.000000,
		MCanBeDiscarded:         true,
		MDescription:            ``,
		MDisplayName:            ``,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Invalid,
		MFuelConsumption:        90.000000,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MInventorySize:          24,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Buildable/Vehicle/Explorer/UI/Explorer_512.Explorer_512`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Buildable/Vehicle/Explorer/UI/Explorer_256.Explorer_256`,
		MStackSize:              resource.One,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Interface/UI/InGame/BuildMenu/BuildCategories/Sub_Transport/SC_Vehicles.SC_Vehicles_C"')`,
	}

	FreightWagon = FGVehicleDescriptor{
		ClassName:               "Desc_FreightWagon_C",
		MAbbreviatedDisplayName: ``,
		MBuildMenuPriority:      50.000000,
		MCanBeDiscarded:         true,
		MDescription:            ``,
		MDisplayName:            ``,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Invalid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MInventorySize:          32,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Buildable/Vehicle/Train/Wagon/UI/FreightCar_512.FreightCar_512`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Buildable/Vehicle/Train/Wagon/UI/FreightCar_256.FreightCar_256`,
		MStackSize:              resource.Medium,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Interface/UI/InGame/BuildMenu/BuildCategories/Sub_Transport/SC_Trains.SC_Trains_C"')`,
	}

	Locomotive = FGVehicleDescriptor{
		ClassName:               "Desc_Locomotive_C",
		MAbbreviatedDisplayName: ``,
		MBuildMenuPriority:      49.000000,
		MCanBeDiscarded:         true,
		MDescription:            ``,
		MDisplayName:            ``,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Invalid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Buildable/Vehicle/Train/Locomotive/UI/Locomotive_512.Locomotive_512`,
		MPowerConsumption:       `(Min=25.000000,Max=110.000000)`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Buildable/Vehicle/Train/Locomotive/UI/Locomotive_256.Locomotive_256`,
		MStackSize:              resource.One,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Interface/UI/InGame/BuildMenu/BuildCategories/Sub_Transport/SC_Trains.SC_Trains_C"')`,
	}

	Tractor = FGVehicleDescriptor{
		ClassName:               "Desc_Tractor_C",
		MAbbreviatedDisplayName: ``,
		MBuildMenuPriority:      1.000000,
		MCanBeDiscarded:         true,
		MDescription:            ``,
		MDisplayName:            ``,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Invalid,
		MFuelConsumption:        55.000000,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MInventorySize:          25,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Buildable/Vehicle/Tractor/UI/Tractor_512.Tractor_512`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Buildable/Vehicle/Tractor/UI/Tractor_256.Tractor_256`,
		MStackSize:              resource.One,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Interface/UI/InGame/BuildMenu/BuildCategories/Sub_Transport/SC_Vehicles.SC_Vehicles_C"')`,
	}

	Truck = FGVehicleDescriptor{
		ClassName:               "Desc_Truck_C",
		MAbbreviatedDisplayName: ``,
		MBuildMenuPriority:      2.000000,
		MCanBeDiscarded:         true,
		MDescription:            ``,
		MDisplayName:            ``,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Invalid,
		MFuelConsumption:        75.000000,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MInventorySize:          48,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Buildable/Vehicle/Truck/UI/Truck_512.Truck_512`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Buildable/Vehicle/Truck/UI/Truck_256.Truck_256`,
		MStackSize:              resource.One,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Interface/UI/InGame/BuildMenu/BuildCategories/Sub_Transport/SC_Vehicles.SC_Vehicles_C"')`,
	}
)

func GetByClassName(className string) (*FGVehicleDescriptor, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGVehicleDescriptor with name %s", className)
}

var classNameToVar = map[string]*FGVehicleDescriptor{
	"Desc_CyberWagon_C":     &CyberWagon,
	"Desc_DroneTransport_C": &DroneTransport,
	"Desc_Explorer_C":       &Explorer,
	"Desc_FreightWagon_C":   &FreightWagon,
	"Desc_Locomotive_C":     &Locomotive,
	"Desc_Tractor_C":        &Tractor,
	"Desc_Truck_C":          &Truck,
}
