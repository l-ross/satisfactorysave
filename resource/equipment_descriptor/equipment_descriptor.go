package EquipmentDescriptor

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGEquipmentDescriptor struct {
	ClassName               string
	MAbbreviatedDisplayName string
	MCanBeDiscarded         bool
	MDescription            string
	MDisplayName            string
	MEnergyValue            float64
	MFluidColor             string
	MForm                   resource.Form
	MGasColor               string
	MPersistentBigIcon      string
	MRadioactiveDecay       float64
	MRememberPickUp         bool
	MResourceSinkPoints     int
	MSmallIcon              string
	MStackSize              resource.StackSize
}

var (
	Chainsaw = FGEquipmentDescriptor{
		ClassName:               "Desc_Chainsaw_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Slot: Hands
Fuel: Biofuel

Used to clear an area of flora that is too difficult to remove by hand.`,
		MDisplayName:        `Chainsaw`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Equipment/Chainsaw/UI/Chainsaw_256.Chainsaw_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 2760,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Equipment/Chainsaw/UI/Chainsaw_64.Chainsaw_64`,
		MStackSize:          resource.One,
	}

	EqDescZipLine = FGEquipmentDescriptor{
		ClassName:               "BP_EqDescZipLine_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Slot: Hand

Provides faster traversal in factories by allowing Pioneers to zip along Power lines.
Activate the Zipline and aim at a nearby Power Line to connect to it.`,
		MDisplayName:        `Zipline`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Equipment/Zipline/UI/IconDesc_Zipline_256.IconDesc_Zipline_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 5284,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Equipment/Zipline/UI/IconDesc_Zipline_64.IconDesc_Zipline_64`,
		MStackSize:          resource.One,
	}

	EquipmentDescriptorCandyCane = FGEquipmentDescriptor{
		ClassName:               "BP_EquipmentDescriptorCandyCane_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Slot: Hands

Heavy delicious self defense weapon for melee range.`,
		MDisplayName:        `Candy Cane Basher`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Equipment/StunSpear/UI/IconDesc_CaneEquipment_256.IconDesc_CaneEquipment_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Equipment/StunSpear/UI/IconDesc_CaneEquipment_64.IconDesc_CaneEquipment_64`,
		MStackSize:          resource.One,
	}

	EquipmentDescriptorColorGun = FGEquipmentDescriptor{
		ClassName:               "BP_EquipmentDescriptorColorGun_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Slot: Hands
Ammo: Color Cartridge

Paints factory buildings and vehicles. The color can be adjusted prior to painting.`,
		MDisplayName:        `Color Gun`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Equipment/ColorGun/UI/ColorGun_256.ColorGun_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 860,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Equipment/ColorGun/UI/ColorGun_64.ColorGun_64`,
		MStackSize:          resource.One,
	}

	EquipmentDescriptorGasmask = FGEquipmentDescriptor{
		ClassName:               "BP_EquipmentDescriptorGasmask_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Slot: Body

Allows you to breathe normally in poison gas. Consumes Gas Filters from your inventory when you are in poison gas.`,
		MDisplayName:        `Gas Mask`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Equipment/GasMask/UI/GasMask_256.GasMask_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 55000,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Equipment/GasMask/UI/GasMask_64.GasMask_64`,
		MStackSize:          resource.One,
	}

	EquipmentDescriptorHazmatSuit = FGEquipmentDescriptor{
		ClassName:               "BP_EquipmentDescriptorHazmatSuit_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Slot: Body

Shields you from the adverse effects of radiation. 
Consumes Iodine Infused Filters from your inventory when you are in radioactive areas.`,
		MDisplayName:        `Hazmat Suit`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/HazmatSuit/UI/IconDesc_HazmatSuit_256.IconDesc_HazmatSuit_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 54100,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/HazmatSuit/UI/IconDesc_HazmatSuit_64.IconDesc_HazmatSuit_64`,
		MStackSize:          resource.One,
	}

	EquipmentDescriptorHoverPack = FGEquipmentDescriptor{
		ClassName:               "BP_EquipmentDescriptorHoverPack_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Slot: Body
Power consumption: 100 MW

Allows for vertical movement and hovering in mid-air to improve building efficiency and factory traversal. Wirelessly connects to nearby power connections, such as Power Poles and Buildings, for power consumption.

Slow-fall: Hold ascend input after losing connection mid-air.
Disable Hoverpack: Double tap descend input while hovering.`,
		MDisplayName:        `Hover Pack`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Equipment/HoverPack/UI/IconDesc_Hoverpack_256.IconDesc_Hoverpack_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 413920,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Equipment/HoverPack/UI/IconDesc_Hoverpack_64.IconDesc_Hoverpack_64`,
		MStackSize:          resource.One,
	}

	EquipmentDescriptorJetPack = FGEquipmentDescriptor{
		ClassName:               "BP_EquipmentDescriptorJetPack_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Slot: Body

Allows you to move more freely in the air. Consumes Fuel when used and refills with Fuel from your inventory when you're on the ground.`,
		MDisplayName:        `Jetpack`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Equipment/JetPack/UI/Jetpack_256.Jetpack_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 49580,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Equipment/JetPack/UI/Jetpack_64.Jetpack_64`,
		MStackSize:          resource.One,
	}

	EquipmentDescriptorJumpingStilts = FGEquipmentDescriptor{
		ClassName:               "BP_EquipmentDescriptorJumpingStilts_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Slot: Body

An exoskeleton for your lower legs that assists movement, allowing you to sprint faster and jump higher.
Also dampens the impact of landing.`,
		MDisplayName:        `Blade Runners`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Equipment/JumpingStilts/UI/IconDesc_SprintingStilts_256.IconDesc_SprintingStilts_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 4988,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Equipment/JumpingStilts/UI/IconDesc_SprintingStilts_64.IconDesc_SprintingStilts_64`,
		MStackSize:          resource.One,
	}

	EquipmentDescriptorNobeliskDetonator = FGEquipmentDescriptor{
		ClassName:               "BP_EquipmentDescriptorNobeliskDetonator_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Slot: Hands
Ammo: Nobelisk

Used to blow up cracked boulders, rocks and invasive vegetation.`,
		MDisplayName:        `Nobelisk Detonator`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Equipment/NobeliskDetonator/UI/Detonator_256.Detonator_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 39520,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Equipment/NobeliskDetonator/UI/Detonator_64.Detonator_64`,
		MStackSize:          resource.One,
	}

	EquipmentDescriptorObjectScanner = FGEquipmentDescriptor{
		ClassName:               "BP_EquipmentDescriptorObjectScanner_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Slot: Hands

Scans the area for a set item. Beeps at a rate proportional to proximity and direction.`,
		MDisplayName:        `Object Scanner`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Equipment/GemstoneScanner/UI/ObjectScanner_256.ObjectScanner_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 3080,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Equipment/GemstoneScanner/UI/ObjectScanner_64.ObjectScanner_64`,
		MStackSize:          resource.One,
	}

	EquipmentDescriptorRifle = FGEquipmentDescriptor{
		ClassName:               "BP_EquipmentDescriptorRifle_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Slot: Hands
Ammo: Rifle Cartridges

Rapid fire weapon for self-defense. Has a mag size of 10.`,
		MDisplayName:        `Rifle`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Equipment/Rifle/UI/RifleMk1_256.RifleMk1_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 99160,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Equipment/Rifle/UI/RifleMK1_64.RifleMK1_64`,
		MStackSize:          resource.One,
	}

	EquipmentDescriptorShockShank = FGEquipmentDescriptor{
		ClassName:               "BP_EquipmentDescriptorShockShank_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Slot: Hands

Standard issue electroshock self defense weapon for melee range.`,
		MDisplayName:        `Xeno-Zapper`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Equipment/ShockShank/UI/XenoZapper_256.XenoZapper_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1880,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Equipment/ShockShank/UI/XenoZapper_64.XenoZapper_64`,
		MStackSize:          resource.One,
	}

	EquipmentDescriptorSnowballMittens = FGEquipmentDescriptor{
		ClassName:               "BP_EquipmentDescriptorSnowballMittens_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Slot: Hands
Ammo: Snowball

*Disclaimer: Snowballs in pile not included in ammo reserve.`,
		MDisplayName:        `Snowball Pile`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_SnowballMittens_256.IconDesc_SnowballMittens_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_SnowballMittens_64.IconDesc_SnowballMittens_64`,
		MStackSize:          resource.One,
	}

	EquipmentDescriptorStunSpear = FGEquipmentDescriptor{
		ClassName:               "BP_EquipmentDescriptorStunSpear_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Slot: Hands

Heavy electroshock self defense weapon for melee range.`,
		MDisplayName:        `Xeno-Basher`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Equipment/StunSpear/UI/ShockBaton_256.ShockBaton_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 18800,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Equipment/StunSpear/UI/ShockBaton_64.ShockBaton_64`,
		MStackSize:          resource.One,
	}

	GolfCart = FGEquipmentDescriptor{
		ClassName:               "Desc_GolfCart_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `The one and only FICSIT Factory Cart™
Now with special - FICSIT Foundation only - Grip Wheels, for an even smoother and faster factory floor experience!`,
		MDisplayName:        `Factory Cart™`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Equipment/GolfCart/UI/IconDesc_GolfCart_256.IconDesc_GolfCart_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1552,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Equipment/GolfCart/UI/IconDesc_GolfCart_64.IconDesc_GolfCart_64`,
		MStackSize:          resource.One,
	}

	ItemDescriptorPortableMiner = FGEquipmentDescriptor{
		ClassName:               "BP_ItemDescriptorPortableMiner_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Slot: Hands

Can be set up on a resource node to automatically extract the resource. Very limited storage space.`,
		MDisplayName:        `Portable Miner`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/PortableMiner_256.PortableMiner_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 56,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Equipment/PortableMiner/UI/PortableMiner_64.PortableMiner_64`,
		MStackSize:          resource.One,
	}

	RebarGunProjectile = FGEquipmentDescriptor{
		ClassName:               "Desc_RebarGunProjectile_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Slot: Hands
Ammo: Spiked Rebar

Improvised ranged weapon for self defense. Has to be reloaded after each use.`,
		MDisplayName:        `Rebar Gun`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Equipment/NailGun/UI/RebarGun_256.RebarGun_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1968,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Equipment/NailGun/UI/RebarGun_64.RebarGun_64`,
		MStackSize:          resource.One,
	}
)

func GetByClassName(className string) (*FGEquipmentDescriptor, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGEquipmentDescriptor with name %s", className)
}

var classNameToVar = map[string]*FGEquipmentDescriptor{
	"Desc_Chainsaw_C":                           &Chainsaw,
	"BP_EqDescZipLine_C":                        &EqDescZipLine,
	"BP_EquipmentDescriptorCandyCane_C":         &EquipmentDescriptorCandyCane,
	"BP_EquipmentDescriptorColorGun_C":          &EquipmentDescriptorColorGun,
	"BP_EquipmentDescriptorGasmask_C":           &EquipmentDescriptorGasmask,
	"BP_EquipmentDescriptorHazmatSuit_C":        &EquipmentDescriptorHazmatSuit,
	"BP_EquipmentDescriptorHoverPack_C":         &EquipmentDescriptorHoverPack,
	"BP_EquipmentDescriptorJetPack_C":           &EquipmentDescriptorJetPack,
	"BP_EquipmentDescriptorJumpingStilts_C":     &EquipmentDescriptorJumpingStilts,
	"BP_EquipmentDescriptorNobeliskDetonator_C": &EquipmentDescriptorNobeliskDetonator,
	"BP_EquipmentDescriptorObjectScanner_C":     &EquipmentDescriptorObjectScanner,
	"BP_EquipmentDescriptorRifle_C":             &EquipmentDescriptorRifle,
	"BP_EquipmentDescriptorShockShank_C":        &EquipmentDescriptorShockShank,
	"BP_EquipmentDescriptorSnowballMittens_C":   &EquipmentDescriptorSnowballMittens,
	"BP_EquipmentDescriptorStunSpear_C":         &EquipmentDescriptorStunSpear,
	"Desc_GolfCart_C":                           &GolfCart,
	"BP_ItemDescriptorPortableMiner_C":          &ItemDescriptorPortableMiner,
	"Desc_RebarGunProjectile_C":                 &RebarGunProjectile,
}
