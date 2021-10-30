package resource

type ItemDescriptor struct {
	Name                    string
	MAbbreviatedDisplayName string
	MCanBeDiscarded         bool
	MDescription            string
	MDisplayName            string
	MEnergyValue            float64
	MFluidColor             struct {
		B int
		G int
		R int
		A int
	}
	MForm     Form
	MGasColor struct {
		B int
		G int
		R int
		A int
	}
	MPersistentBigIcon  string
	MRadioactiveDecay   float64
	MRememberPickUp     bool
	MResourceSinkPoints int
	MSmallIcon          string
	MStackSize          StackSize
}

var (
	AluminaSolution = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Dissolved Alumina, extracted from Bauxite. Can be further refined into Aluminum Scrap for Aluminum Ingot production.`,
		MDisplayName:            `Alumina Solution`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 193,
			G: 193,
			R: 193,
			A: 255,
		},
		MForm: Liquid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Alumina/UI/LiquidAlumina_Pipe_512.LiquidAlumina_Pipe_512`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 20,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Alumina/UI/LiquidAlumina_Pipe_256.LiquidAlumina_Pipe_256`,
		MStackSize:          Fluid,
	}

	AluminumCasing = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `A versatile container cast from Aluminum Ingots.`,
		MDisplayName:            `Aluminum Casing`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/AluminumCasing/UI/IconDesc_AluminiumCasing_256.IconDesc_AluminiumCasing_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 393,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/AluminumCasing/UI/IconDesc_AluminiumCasing_64.IconDesc_AluminiumCasing_64`,
		MStackSize:          Big,
	}

	AluminumIngot = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Aluminum Ingots are made from Aluminum Scrap, which is refined from Alumina Solution.
Used to produce specialized aluminum-based parts.`,
		MDisplayName: `Aluminum Ingot`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/AluminumIngot/UI/IconDesc_AluminiumIngot_256.IconDesc_AluminiumIngot_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 131,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/AluminumIngot/UI/IconDesc_AluminiumIngot_64.IconDesc_AluminiumIngot_64`,
		MStackSize:          Medium,
	}

	AluminumPlate = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Thin, lightweight, and highly durable sheets that are mainly used for products that require high heat conduction or a high specific strength.`,
		MDisplayName:            `Alclad Aluminum Sheet`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/AluminumPlate/UI/IconDesc_AluminiumSheet_256.IconDesc_AluminiumSheet_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 266,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/AluminumPlate/UI/IconDesc_AluminiumSheet_64.IconDesc_AluminiumSheet_64`,
		MStackSize:          Big,
	}

	AluminumPlateReinforced = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Used to dissipate heat faster.`,
		MDisplayName:            `Heat Sink`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/HeatSink/UI/IconDesc_Heatsink_256.IconDesc_Heatsink_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 2804,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/HeatSink/UI/IconDesc_Heatsink_64.IconDesc_Heatsink_64`,
		MStackSize:          Medium,
	}

	AluminumScrap = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Aluminum Scrap is pure aluminum refined from Alumina. Can be smelted down to Aluminum Ingots for industrial usage.`,
		MDisplayName:            `Aluminum Scrap`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/AluminumScrap/UI/IconDesc_AluminiumScrap_256.IconDesc_AluminiumScrap_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 27,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/AluminumScrap/UI/IconDesc_AluminiumScrap_64.IconDesc_AluminiumScrap_64`,
		MStackSize:          Huge,
	}

	Battery = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Primarily used as fuel for Drones and Vehicles.`,
		MDisplayName:            `Battery`,
		MEnergyValue:            6000.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Battery/UI/IconDesc_Battery_256.IconDesc_Battery_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 465,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Battery/UI/IconDesc_Battery_64.IconDesc_Battery_64`,
		MStackSize:          Big,
	}

	Cable = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for crafting.
Primarily used to build power lines.`,
		MDisplayName: `Cable`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Cable/UI/IconDesc_Cables_256.IconDesc_Cables_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 24,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Cable/UI/IconDesc_Cables_64.IconDesc_Cables_64`,
		MStackSize:          Big,
	}

	CandyCane = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `A delicious Candy Cane to be enjoyed during the FICSMAS Holidays. 
*Disclaimer: Can't be consumed...`,
		MDisplayName: `Candy Cane`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_CanePart_256.IconDesc_CanePart_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_CanePart_64.IconDesc_CanePart_64`,
		MStackSize:          Huge,
	}

	CartridgeStandard = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Ammo for the Rifle.`,
		MDisplayName:            `Rifle Cartridge`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/CartridgeStandard/UI/Rifle_Magazine_256.Rifle_Magazine_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 664,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/CartridgeStandard/UI/Rifle_Magazine_64.Rifle_Magazine_64`,
		MStackSize:          Medium,
	}

	Cement = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for building.
Good for stable foundations.`,
		MDisplayName: `Concrete`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Cement/UI/IconDesc_Concrete_256.IconDesc_Concrete_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 12,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Cement/UI/IconDesc_Concrete_64.IconDesc_Concrete_64`,
		MStackSize:          Huge,
	}

	CircuitBoard = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Circuit Boards are advanced electronics that are used in a plethora of different ways.`,
		MDisplayName:            `Circuit Board`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/CircuitBoard/UI/IconDesc_CircuitBoard_256.IconDesc_CircuitBoard_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 696,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/CircuitBoard/UI/IconDesc_CircuitBoard_64.IconDesc_CircuitBoard_64`,
		MStackSize:          Big,
	}

	CircuitBoardHighSpeed = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `AI Limiters are super advanced electronics that are used to control AIs and keep them from evolving in malicious ways.`,
		MDisplayName:            `AI Limiter`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/AIlimiter/UI/IconDesc_AILimiter_256.IconDesc_AILimiter_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 920,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/AIlimiter/UI/IconDesc_AILimiter_64.IconDesc_AILimiter_64`,
		MStackSize:          Medium,
	}

	ColorCartridge = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Ammo for the Color Gun.
Necessary to change the color of factory buildings and vehicles.`,
		MDisplayName: `Color Cartridge`,
		MEnergyValue: 900.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/ColorCartridge/UI/IconDesc_ColorCartridge_256.IconDesc_ColorCartridge_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 10,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/ColorCartridge/UI/IconDesc_ColorCartridge_64.IconDesc_ColorCartridge_64`,
		MStackSize:          Big,
	}

	CompactedCoal = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `A much more efficient alternative for Coal. Used as fuel for vehicles & coal generators.`,
		MDisplayName:            `Compacted Coal`,
		MEnergyValue:            630.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/CompactedCoal/UI/CompactedCoal_256.CompactedCoal_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 28,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/CompactedCoal/UI/CompactedCoal_64.CompactedCoal_64`,
		MStackSize:          Medium,
	}

	Computer = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `A Computer is a complex logic machine that is used to control advanced behaviour in machines.`,
		MDisplayName:            `Computer`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Computer/UI/IconDesc_Computer_256.IconDesc_Computer_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 17260,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Computer/UI/IconDesc_Computer_64.IconDesc_Computer_64`,
		MStackSize:          Small,
	}

	ComputerSuper = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `The supercomputer is the next-gen version of the computer.`,
		MDisplayName:            `Supercomputer`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/ComputerSuper/UI/IconDesc_SuperComputer_256.IconDesc_SuperComputer_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 99576,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/ComputerSuper/UI/IconDesc_SuperComputer_64.IconDesc_SuperComputer_64`,
		MStackSize:          Small,
	}

	CoolingSystem = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Used to keep temperatures of advanced parts and buildings from exceeding to inefficient levels.`,
		MDisplayName:            `Cooling System`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/CoolingSystem/UI/IconDesc_CoolingSystem_256.IconDesc_CoolingSystem_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 12006,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/CoolingSystem/UI/IconDesc_CoolingSystem_64.IconDesc_CoolingSystem_64`,
		MStackSize:          Medium,
	}

	CopperDust = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Ground down Copper Ingots.
The high natural density of Copper, combined with the granularity of the powder, make this part fit for producing Nuclear Pasta in the Particle Accelerator.`,
		MDisplayName: `Copper Powder`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/CopperDust/UI/IconDesc_CopperDust_256.IconDesc_CopperDust_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 72,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/CopperDust/UI/IconDesc_CopperDust_64.IconDesc_CopperDust_64`,
		MStackSize:          Huge,
	}

	CopperIngot = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for crafting.
Crafted into the most basic parts.`,
		MDisplayName: `Copper Ingot`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/CopperIngot/UI/IconDesc_CopperIngot_256.IconDesc_CopperIngot_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 6,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/CopperIngot/UI/IconDesc_CopperIngot_64.IconDesc_CopperIngot_64`,
		MStackSize:          Medium,
	}

	CopperSheet = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for crafting.
Primarily used for pipelines due to its high corrosion resistance.`,
		MDisplayName: `Copper Sheet`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/CopperSheet/UI/IconDesc_CopperSheet_256.IconDesc_CopperSheet_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 24,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/CopperSheet/UI/IconDesc_CopperSheet_64.IconDesc_CopperSheet_64`,
		MStackSize:          Big,
	}

	Crystal = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         false,
		MDescription:            `A strange slug radiating a weak strange power.`,
		MDisplayName:            `Green Power Slug`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Environment/Crystal/UI/PowerSlugGreen_256.PowerSlugGreen_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     true,
		MResourceSinkPoints: 0,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Environment/Crystal/UI/PowerSlugGreen_64.PowerSlugGreen_64`,
		MStackSize:          Small,
	}

	CrystalOscillator = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `A crystal oscillator is an electronic oscillator circuit that uses the mechanical resonance of a vibrating crystal to create an electrical signal with a precise frequency.`,
		MDisplayName:            `Crystal Oscillator`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/CrystalOscillator/UI/IconDesc_CrystalOscillator_256.IconDesc_CrystalOscillator_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 3072,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/CrystalOscillator/UI/IconDesc_CrystalOscillator_64.IconDesc_CrystalOscillator_64`,
		MStackSize:          Medium,
	}

	CrystalShard = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Mucus from the power slugs compressed into a solid crystal-like shard. 
It radiates a strange power.`,
		MDisplayName: `Power Shard`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Environment/Crystal/UI/PowerShard_256.PowerShard_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 0,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Environment/Crystal/UI/PowerShard_64.PowerShard_64`,
		MStackSize:          Medium,
	}

	Crystalmk2 = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         false,
		MDescription:            `A strange slug radiating a strange power.`,
		MDisplayName:            `Yellow Power Slug`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Environment/Crystal/UI/PowerSlugYellow_256.PowerSlugYellow_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     true,
		MResourceSinkPoints: 0,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Environment/Crystal/UI/PowerSlugYellow_64.PowerSlugYellow_64`,
		MStackSize:          Small,
	}

	Crystalmk3 = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         false,
		MDescription:            `A strange slug radiating a powerful strange power.`,
		MDisplayName:            `Purple Power Slug`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Environment/Crystal/UI/PowerSlugPurple_256.PowerSlugPurple_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     true,
		MResourceSinkPoints: 0,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Environment/Crystal/UI/PowerSlugPurple_64.PowerSlugPurple_64`,
		MStackSize:          Small,
	}

	ElectromagneticControlRod = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Control Rods regulate power output via electromagnetism.`,
		MDisplayName:            `Electromagnetic Control Rod`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/ElectromagneticControlRod/UI/IconDesc_ElectromagneticControlRod_256.IconDesc_ElectromagneticControlRod_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 2560,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/ElectromagneticControlRod/UI/IconDesc_ElectromagneticControlRod_64.IconDesc_ElectromagneticControlRod_64`,
		MStackSize:          Medium,
	}

	Filter = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Used in gas masks to filter out pollutants in the air.`,
		MDisplayName:            `Gas Filter`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Filter/UI/IconDesc_GasMaskFilter_256.IconDesc_GasMaskFilter_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 830,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Filter/UI/IconDesc_GasMaskFilter_64.IconDesc_GasMaskFilter_64`,
		MStackSize:          Small,
	}

	FluidCanister = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Used to package fluids for transportation.`,
		MDisplayName:            `Empty Canister`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/FluidCanister/UI/IconDesc_EmptyCannister_256.IconDesc_EmptyCannister_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 60,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/FluidCanister/UI/IconDesc_EmptyCannister_64.IconDesc_EmptyCannister_64`,
		MStackSize:          Medium,
	}

	Fuel = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Fuel, packaged for alternative transport. Can be used as fuel for Vehicles or the Jetpack.`,
		MDisplayName:            `Packaged Fuel`,
		MEnergyValue:            750.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 21,
			G: 125,
			R: 235,
			A: 255,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Fuel/UI/IconDesc_Fuel_256.IconDesc_Fuel_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 270,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Fuel/UI/IconDesc_Fuel_64.IconDesc_Fuel_64`,
		MStackSize:          Medium,
	}

	GasTank = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Used to package gases and volatile liquids for transportation.`,
		MDisplayName:            `Empty Fluid Tank`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/GasTank/UI/IconDesc_PressureTank_256.IconDesc_PressureTank_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 225,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/GasTank/UI/IconDesc_PressureTank_64.IconDesc_PressureTank_64`,
		MStackSize:          Medium,
	}

	Gift = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Special FICSMAS buildings and parts can be obtained and produced from this FICSIT Holiday present.

*Watch the sky for deliveries from orbit!`,
		MDisplayName: `FICSMAS Gift`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Gift_256.IconDesc_Gift_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     true,
		MResourceSinkPoints: 1,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Gift_64.IconDesc_Gift_64`,
		MStackSize:          Huge,
	}

	GoldIngot = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Caterium Ingots are smelted from Caterium Ore. Caterium Ingots are mostly used for advanced electronics.`,
		MDisplayName:            `Caterium Ingot`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/GoldIngot/UI/IconDesc_CateriumIngot_256.IconDesc_CateriumIngot_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 42,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/GoldIngot/UI/IconDesc_CateriumIngot_64.IconDesc_CateriumIngot_64`,
		MStackSize:          Medium,
	}

	Gunpowder = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `An explosive powder that is commonly used in explosives and cartridges.`,
		MDisplayName:            `Black Powder`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/GunPowder/UI/IconDesc_Gunpowder_256.IconDesc_Gunpowder_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 50,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/GunPowder/UI/IconDesc_Gunpowder_64.IconDesc_Gunpowder_64`,
		MStackSize:          Medium,
	}

	HUBParts = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         false,
		MDescription:            `The parts required to build the basic structure of the HUB.`,
		MDisplayName:            `HUB Parts`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/HUBParts/UI/IconDesc_HubParts_256.IconDesc_HubParts_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 0,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/HUBParts/UI/IconDesc_HubParts_64.IconDesc_HubParts_64`,
		MStackSize:          One,
	}

	HazmatFilter = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Used in hazmat suits to filter out radioactive particles in the air.`,
		MDisplayName:            `Iodine Infused Filter`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/IodineInfusedFilter/UI/IconDesc_HazmatFilter_256.IconDesc_HazmatFilter_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 2718,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/IodineInfusedFilter/UI/IconDesc_HazmatFilter_64.IconDesc_HazmatFilter_64`,
		MStackSize:          Small,
	}

	HeavyOilResidue = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `A by-product of Plastic and Rubber production. Can be further refined into Fuel and Petroleum Coke.`,
		MDisplayName:            `Heavy Oil Residue`,
		MEnergyValue:            0.400000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 120,
			G: 45,
			R: 109,
			A: 255,
		},
		MForm: Liquid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/HeavyOilResidue/UI/IconDesc_LiquidHeavyOilResidue_Pipe_512.IconDesc_LiquidHeavyOilResidue_Pipe_512`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 30,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/HeavyOilResidue/UI/IconDesc_LiquidHeavyOilResidue_Pipe_256.IconDesc_LiquidHeavyOilResidue_Pipe_256`,
		MStackSize:          Fluid,
	}

	HighSpeedConnector = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `The high-speed connector connects several cables and wires in a very efficient way. Uses a standard pattern so its applications are many and varied.`,
		MDisplayName:            `High-Speed Connector`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/HighSpeedConnector/UI/IconDesc_HighSpeedConnector_256.IconDesc_HighSpeedConnector_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 3776,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/HighSpeedConnector/UI/IconDesc_HighSpeedConnector_64.IconDesc_HighSpeedConnector_64`,
		MStackSize:          Medium,
	}

	HighSpeedWire = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Caterium's high conductivity and resistance to corrosion makes it ideal for small, advanced electronics.`,
		MDisplayName:            `Quickwire`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/HighSpeedWire/UI/IconDesc_Quickwire_256.IconDesc_Quickwire_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 17,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/HighSpeedWire/UI/IconDesc_Quickwire_64.IconDesc_Quickwire_64`,
		MStackSize:          Huge,
	}

	IronIngot = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for crafting.
Crafted into the most basic parts.`,
		MDisplayName: `Iron Ingot`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/IronIngot/UI/IconDesc_IronIngot_256.IconDesc_IronIngot_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 2,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/IronIngot/UI/IconDesc_IronIngot_64.IconDesc_IronIngot_64`,
		MStackSize:          Medium,
	}

	IronPlate = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for crafting.
One of the most basic parts.`,
		MDisplayName: `Iron Plate`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/IronPlate/UI/IconDesc_IronPlates_256.IconDesc_IronPlates_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 6,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/IronPlate/UI/IconDesc_IronPlates_64.IconDesc_IronPlates_64`,
		MStackSize:          Big,
	}

	IronPlateReinforced = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for crafting.
A sturdier and more durable Iron Plate.`,
		MDisplayName: `Reinforced Iron Plate`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/IronPlateReinforced/UI/IconDesc_ReinforcedIronPlates_256.IconDesc_ReinforcedIronPlates_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 120,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/IronPlateReinforced/UI/IconDesc_ReinforcedIronPlates_64.IconDesc_ReinforcedIronPlates_64`,
		MStackSize:          Medium,
	}

	IronRod = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for crafting.
One of the most basic parts.`,
		MDisplayName: `Iron Rod`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/IronRod/UI/IconDesc_IronRods_256.IconDesc_IronRods_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 4,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/IronRod/UI/IconDesc_IronRods_64.IconDesc_IronRods_64`,
		MStackSize:          Big,
	}

	IronScrew = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for crafting.
One of the most basic parts.`,
		MDisplayName: `Screw`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/IronScrew/UI/IconDesc_IronScrews_256.IconDesc_IronScrews_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 2,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/IronScrew/UI/IconDesc_IronScrews_64.IconDesc_IronScrews_64`,
		MStackSize:          Huge,
	}

	LiquidFuel = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Fuel can be used to generate power or packaged to be used as fuel for Vehicles or the Jetpack.`,
		MDisplayName:            `Fuel`,
		MEnergyValue:            0.750000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 21,
			G: 125,
			R: 235,
			A: 255,
		},
		MForm: Liquid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Fuel/UI/IconDesc_LiquidFuel_Pipe_512.IconDesc_LiquidFuel_Pipe_512`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 75,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Fuel/UI/IconDesc_LiquidFuel_Pipe_256.IconDesc_LiquidFuel_Pipe_256`,
		MStackSize:          Fluid,
	}

	LiquidTurboFuel = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `A more efficient alternative to Fuel. Can be used to generate power or packaged to be used as fuel for Vehicles.`,
		MDisplayName:            `Turbofuel`,
		MEnergyValue:            2.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 46,
			G: 41,
			R: 212,
			A: 255,
		},
		MForm: Liquid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Turbofuel/UI/IconDesc_LiquidTurboFuel_Pipe_512.IconDesc_LiquidTurboFuel_Pipe_512`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 225,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Turbofuel/UI/IconDesc_LiquidTurboFuel_Pipe_256.IconDesc_LiquidTurboFuel_Pipe_256`,
		MStackSize:          Fluid,
	}

	ModularFrame = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for crafting.
Multi-purpose building block.`,
		MDisplayName: `Modular Frame`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/ModularFrame/UI/IconDesc_ModularFrame_256.IconDesc_ModularFrame_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 408,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/ModularFrame/UI/IconDesc_ModularFrame_64.IconDesc_ModularFrame_64`,
		MStackSize:          Small,
	}

	ModularFrameFused = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `A corrosion resistant, nitride hardened, highly robust, yet lightweight modular frame.`,
		MDisplayName:            `Fused Modular Frame`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/ModularFrameFused/UI/IconDesc_FusedModularFrame_256.IconDesc_FusedModularFrame_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 62840,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/ModularFrameFused/UI/IconDesc_FusedModularFrame_64.IconDesc_FusedModularFrame_64`,
		MStackSize:          Small,
	}

	ModularFrameHeavy = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `A more robust multi-purpose frame.`,
		MDisplayName:            `Heavy Modular Frame`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/ModularFrameHeavy/UI/IconDesc_ModularFrameHeavy_256.IconDesc_ModularFrameHeavy_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 11520,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/ModularFrameHeavy/UI/IconDesc_ModularFrameHeavy_64.IconDesc_ModularFrameHeavy_64`,
		MStackSize:          Small,
	}

	ModularFrameLightweight = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Enhances and directs radio signals.`,
		MDisplayName:            `Radio Control Unit`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/RadioControlUnit/UI/IconDesc_RadioControlUnit_256.IconDesc_RadioControlUnit_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 32908,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/RadioControlUnit/UI/IconDesc_RadioControlUnit_64.IconDesc_RadioControlUnit_64`,
		MStackSize:          Small,
	}

	Motor = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `The Motor creates a mechanical force that is used to move things from machines to vehicles.`,
		MDisplayName:            `Motor`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Motor/UI/IconDesc_Engine_256.IconDesc_Engine_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1520,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Motor/UI/IconDesc_Engine_64.IconDesc_Engine_64`,
		MStackSize:          Small,
	}

	MotorLightweight = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `The Turbo Motor is a more complex, and more powerful, version of the regular Motor.`,
		MDisplayName:            `Turbo Motor`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/TurboMotor/UI/IconDesc_TurboMotor_256.IconDesc_TurboMotor_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 242720,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/TurboMotor/UI/IconDesc_TurboMotor_64.IconDesc_TurboMotor_64`,
		MStackSize:          Small,
	}

	NitricAcid = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Produced by reaction of Nitrogen Gas with Water. Its high corrosiveness and oxidizing properties make it an excellent choice for refinement and fuel production processes.`,
		MDisplayName:            `Nitric Acid`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 162,
			G: 217,
			R: 217,
			A: 255,
		},
		MForm: Liquid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/NitricAcid/UI/IconDesc_NitricAcid_256.IconDesc_NitricAcid_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 94,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/NitricAcid/UI/IconDesc_NitricAcid_64.IconDesc_NitricAcid_64`,
		MStackSize:          Fluid,
	}

	NobeliskExplosive = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Can be used with the Nobelisk Detonator to blow up cracked boulders, vegetation or other vulnerable targets.`,
		MDisplayName:            `Nobelisk`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/NobeliskExplosive/UI/IconDesc_Explosive_256.IconDesc_Explosive_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 980,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/NobeliskExplosive/UI/IconDesc_Explosive_64.IconDesc_Explosive_64`,
		MStackSize:          Small,
	}

	NonFissibleUranium = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         false,
		MDescription: `The isotope Uranium-238 is non-fissile, meaning it cannot be used for nuclear fission. It can, however, be conversed into fissile Plutonium in the Particle Accelerator.

Caution: Mildly Radioactive.`,
		MDisplayName: `Non-fissile Uranium`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Non-FissibleUranium/UI/IconDesc_NonFissileUranium_256.IconDesc_NonFissileUranium_256`,
		MRadioactiveDecay:   0.750000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 0,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Non-FissibleUranium/UI/IconDesc_NonFissileUranium_64.IconDesc_NonFissileUranium_64`,
		MStackSize:          Huge,
	}

	NuclearWaste = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         false,
		MDescription: `The by-product of consuming Uranium Fuel Rods in the Nuclear Power Plant.
Non-fissible Uranium can be extracted. Handle with caution.

Caution: HIGHLY Radioactive.`,
		MDisplayName: `Uranium Waste`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/NuclearWaste/UI/IconDesc_NuclearWaste_256.IconDesc_NuclearWaste_256`,
		MRadioactiveDecay:   10.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 0,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/NuclearWaste/UI/IconDesc_NuclearWaste_64.IconDesc_NuclearWaste_64`,
		MStackSize:          Huge,
	}

	PackagedOil = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Crude Oil, packaged for alternative transport. Can be used as fuel for Vehicles.`,
		MDisplayName:            `Packaged Oil`,
		MEnergyValue:            320.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/RawResources/CrudeOil/UI/Oil_256.Oil_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 180,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/RawResources/CrudeOil/UI/Oil_64.Oil_64`,
		MStackSize:          Medium,
	}

	PackagedOilResidue = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Heavy Oil Residue, packaged for alternative transport. Can be used as fuel for Vehicles.`,
		MDisplayName:            `Packaged Heavy Oil Residue`,
		MEnergyValue:            400.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 120,
			G: 45,
			R: 109,
			A: 255,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/RawResources/CrudeOil/UI/OilResidue_256.OilResidue_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 180,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/RawResources/CrudeOil/UI/OilResidue_64.OilResidue_64`,
		MStackSize:          Medium,
	}

	PackagedWater = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Water, packaged for alternative transport.`,
		MDisplayName:            `Packaged Water`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 212,
			G: 176,
			R: 122,
			A: 255,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/PackagedWater/UI/IconDesc_PackagedWater_256.IconDesc_PackagedWater_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 130,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/PackagedWater/UI/IconDesc_PackagedWater_64.IconDesc_PackagedWater_64`,
		MStackSize:          Medium,
	}

	PetroleumCoke = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for crafting.
A carbon-rich material distilled from Heavy Oil Residue. 
Used as a less efficient coal replacement or for aluminum refinement.`,
		MDisplayName: `Petroleum Coke`,
		MEnergyValue: 180.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/PetroleumCoke/UI/IconDesc_PetroleumCoke_256.IconDesc_PetroleumCoke_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 20,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/PetroleumCoke/UI/IconDesc_PetroleumCoke_64.IconDesc_PetroleumCoke_64`,
		MStackSize:          Big,
	}

	Plastic = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `A versatile and easy to manufacture material that can be used for a lot of things.`,
		MDisplayName:            `Plastic`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Plastic/UI/IconDesc_Plastic_256.IconDesc_Plastic_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 75,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Plastic/UI/IconDesc_Plastic_64.IconDesc_Plastic_64`,
		MStackSize:          Big,
	}

	PlutoniumCell = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         false,
		MDescription: `Plutonium Cells are concrete encased Plutonium Pellets.
Used to produce Plutonium Fuel Rods for Nuclear Power production.

Caution: Moderately Radioactive.`,
		MDisplayName: `Encased Plutonium Cell`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/PlutoniumCell/UI/IconDesc_EncasedPlutoniumCell_256.IconDesc_EncasedPlutoniumCell_256`,
		MRadioactiveDecay:   120.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 0,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/PlutoniumCell/UI/IconDesc_EncasedPlutoniumCell_64.IconDesc_EncasedPlutoniumCell_64`,
		MStackSize:          Big,
	}

	PlutoniumPellet = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         false,
		MDescription: `Produced in the Particle Accelerator through conversion of Non-fissile Uranium.
Used to produce Encased Plutonium Cells for Plutonium Fuel Rods.

Power Usage: 250-750 MW (500 MW average).
Caution: Moderately Radioactive.`,
		MDisplayName: `Plutonium Pellet`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/PlutoniumPellet/UI/IconDesc_PlutoniumPellet_256.IconDesc_PlutoniumPellet_256`,
		MRadioactiveDecay:   20.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 0,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/PlutoniumPellet/UI/IconDesc_PlutoniumPellet_64.IconDesc_PlutoniumPellet_64`,
		MStackSize:          Medium,
	}

	PolymerResin = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for crafting.
A by-product of oil refinement into fuel. Commonly used to manufacture plastics.`,
		MDisplayName: `Polymer Resin`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/PolymerResin/UI/IconDesc_PolymerResin_256.IconDesc_PolymerResin_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 12,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/PolymerResin/UI/IconDesc_PolymerResin_64.IconDesc_PolymerResin_64`,
		MStackSize:          Big,
	}

	PressureConversionCube = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Converts outgoing force into internal pressure. Required to contain unstable, high-energy matter.`,
		MDisplayName:            `Pressure Conversion Cube`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/PressureConversionCube/UI/IconDesc_ConversionCube_256.IconDesc_ConversionCube_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 257312,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/PressureConversionCube/UI/IconDesc_ConversionCube_64.IconDesc_ConversionCube_64`,
		MStackSize:          Small,
	}

	QuartzCrystal = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Used for radar and quantum technology.`,
		MDisplayName:            `Quartz Crystal`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/QuartzCrystal/UI/IconDesc_QuartzResource_256.IconDesc_QuartzResource_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 50,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/QuartzCrystal/UI/IconDesc_QuartzResource_64.IconDesc_QuartzResource_64`,
		MStackSize:          Medium,
	}

	Rotor = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for crafting.
The moving parts of a motor.`,
		MDisplayName: `Rotor`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Rotor/UI/IconDesc_Rotor_256.IconDesc_Rotor_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 140,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Rotor/UI/IconDesc_Rotor_64.IconDesc_Rotor_64`,
		MStackSize:          Medium,
	}

	Rubber = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Rubber is a material that is very flexible and has a lot of friction.`,
		MDisplayName:            `Rubber`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Rubber/UI/IconDesc_Rubber_256.IconDesc_Rubber_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 60,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Rubber/UI/IconDesc_Rubber_64.IconDesc_Rubber_64`,
		MStackSize:          Big,
	}

	Silica = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Derived from quartz, commonly used in structural materials and microelectronics.`,
		MDisplayName:            `Silica`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Silica/UI/IconDesc_Silica_256.IconDesc_Silica_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 20,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Silica/UI/IconDesc_Silica_64.IconDesc_Silica_64`,
		MStackSize:          Medium,
	}

	Snow = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `It's snow. Not the nice, thick, crunchy kind though... more the disgustingly wet, slushy kind... Guess we can make stuff from it.`,
		MDisplayName:            `Actual Snow`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Snow_256.IconDesc_Snow_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Snow_64.IconDesc_Snow_64`,
		MStackSize:          Huge,
	}

	SnowballProjectile = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Compressed dihydrogen monoxide crystals.`,
		MDisplayName:            `Snowball`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_SnowballProjectile_256.IconDesc_SnowballProjectile_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_SnowballProjectile_64.IconDesc_SnowballProjectile_64`,
		MStackSize:          Huge,
	}

	SpaceElevatorPart1 = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Project Part #1. Ship with the Space Elevator to complete phases of Project Assembly.`,
		MDisplayName:            `Smart Plating`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/SmartPlate/UI/IconDesc_SpelevatorPart_1_256.IconDesc_SpelevatorPart_1_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 520,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/SmartPlate/UI/IconDesc_SpelevatorPart_1_64.IconDesc_SpelevatorPart_1_64`,
		MStackSize:          Small,
	}

	SpaceElevatorPart2 = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Project Part #2. Ship with the Space Elevator to complete phases of Project Assembly.`,
		MDisplayName:            `Versatile Framework`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Skeletalframework/UI/IconDesc_SpelevatorPart_2_256.IconDesc_SpelevatorPart_2_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1176,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Skeletalframework/UI/IconDesc_SpelevatorPart_2_64.IconDesc_SpelevatorPart_2_64`,
		MStackSize:          Small,
	}

	SpaceElevatorPart3 = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Project Part #3. Ship with the Space Elevator to complete phases of Project Assembly.`,
		MDisplayName:            `Automated Wiring`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/AutomatedWiring/UI/SpelevatorPart_3_256.SpelevatorPart_3_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1440,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/AutomatedWiring/UI/SpelevatorPart_3_64.SpelevatorPart_3_64`,
		MStackSize:          Small,
	}

	SpaceElevatorPart4 = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Project Part #4. Ship with the Space Elevator to complete phases of Project Assembly.`,
		MDisplayName:            `Modular Engine`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/ModularEngine/UI/IconDesc_SpelevatorPart_4_256.IconDesc_SpelevatorPart_4_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 9960,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/ModularEngine/UI/IconDesc_SpelevatorPart_4_64.IconDesc_SpelevatorPart_4_64`,
		MStackSize:          Small,
	}

	SpaceElevatorPart5 = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Project Part #5. Ship with the Space Elevator to complete phases of Project Assembly.`,
		MDisplayName:            `Adaptive Control Unit`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/ControlSystem/UI/IconDesc_SpelevatorPart_5_256.IconDesc_SpelevatorPart_5_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 86120,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/ControlSystem/UI/IconDesc_SpelevatorPart_5_64.IconDesc_SpelevatorPart_5_64`,
		MStackSize:          Small,
	}

	SpaceElevatorPart6 = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Project Part #7. Ship with the Space Elevator to complete phases of Project Assembly.

These modular generators use superconducting magnets and vast amounts of electricity to produce an easily expandible and powerful magnet field.`,
		MDisplayName: `Magnetic Field Generator`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/MagneticFieldGenerator/UI/IconDesc_MagneticFieldGenerator_256.IconDesc_MagneticFieldGenerator_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 15650,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/MagneticFieldGenerator/UI/IconDesc_MagneticFieldGenerator_64.IconDesc_MagneticFieldGenerator_64`,
		MStackSize:          Small,
	}

	SpaceElevatorPart7 = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Project Part #6. Ship with the Space Elevator to complete phases of Project Assembly.

This extremely fast and precise computing system is specifically designed to direct the Project Assembly: Assembly Phase.`,
		MDisplayName: `Assembly Director System`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/AssemblyControlUnit/UI/IconDesc_AssemblyDirectorSystem_256.IconDesc_AssemblyDirectorSystem_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 543632,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/AssemblyControlUnit/UI/IconDesc_AssemblyDirectorSystem_64.IconDesc_AssemblyDirectorSystem_64`,
		MStackSize:          Small,
	}

	SpaceElevatorPart8 = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Project Part #8. Ship with the Space Elevator to complete phases of Project Assembly.

Uses extreme heat to produce the high pressure plasma required to get Project Assembly into motion.`,
		MDisplayName: `Thermal Propulsion Rocket`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/ThermalAntimatterPropulsionRocket/UI/IconDesc_ThermalPropulsionRocket_256.IconDesc_ThermalPropulsionRocket_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 732956,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/ThermalAntimatterPropulsionRocket/UI/IconDesc_ThermalPropulsionRocket_64.IconDesc_ThermalPropulsionRocket_64`,
		MStackSize:          Small,
	}

	SpaceElevatorPart9 = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Project Part #9. Ship with the Space Elevator to complete phases of Project Assembly.
Power Usage: 500-1500 MW (1000 MW average).

Nuclear Pasta is extremely dense degenerate matter, formed when extreme pressure forces protons and electrons together into neutrons. It is theorized to exist naturally within the crust of neutron stars.`,
		MDisplayName: `Nuclear Pasta`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/NuclearPasta/UI/IconDesc_NuclearPasta_256.IconDesc_NuclearPasta_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 543424,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/NuclearPasta/UI/IconDesc_NuclearPasta_64.IconDesc_NuclearPasta_64`,
		MStackSize:          Small,
	}

	SpikedRebar = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Ammo for the Rebar Gun.`,
		MDisplayName:            `Spiked Rebar`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/SpikedRebar/UI/IconDesc_SpikedRebar_256.IconDesc_SpikedRebar_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 8,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/SpikedRebar/UI/IconDesc_SpikedRebar_64.IconDesc_SpikedRebar_64`,
		MStackSize:          Small,
	}

	Stator = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for crafting.
The static parts of a motor.`,
		MDisplayName: `Stator`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Stator/UI/IconDesc_Stator_256.IconDesc_Stator_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 240,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Stator/UI/IconDesc_Stator_64.IconDesc_Stator_64`,
		MStackSize:          Medium,
	}

	SteelIngot = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Steel Ingots are made from Iron Ore that's been smelted with Coal. They are made into several parts used in building construction.`,
		MDisplayName:            `Steel Ingot`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/SteelIngot/UI/IconDesc_SteelIngot_256.IconDesc_SteelIngot_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 8,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/SteelIngot/UI/IconDesc_SteelIngot_64.IconDesc_SteelIngot_64`,
		MStackSize:          Medium,
	}

	SteelPipe = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Steel Pipes are used most often when constructing a little more advanced buildings.`,
		MDisplayName:            `Steel Pipe`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/SteelPipe/UI/IconDesc_SteelPipe_256.IconDesc_SteelPipe_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 24,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/SteelPipe/UI/IconDesc_SteelPipe_64.IconDesc_SteelPipe_64`,
		MStackSize:          Big,
	}

	SteelPlate = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Steel Beams are used most often when constructing a little more advanced buildings.`,
		MDisplayName:            `Steel Beam`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/SteelPlate/UI/IconDesc_SteelBeam_256.IconDesc_SteelBeam_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 64,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/SteelPlate/UI/IconDesc_SteelBeam_64.IconDesc_SteelBeam_64`,
		MStackSize:          Big,
	}

	SteelPlateReinforced = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Encased Industrial Beams utilizes the compressive strength of concrete and tensile strength of steel simultaneously.
Mostly used as a stable basis for constructing buildings.`,
		MDisplayName: `Encased Industrial Beam`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/SteelPlateReinforced/UI/IconDesc_EncasedSteelBeam_256.IconDesc_EncasedSteelBeam_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 632,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/SteelPlateReinforced/UI/IconDesc_EncasedSteelBeam_64.IconDesc_EncasedSteelBeam_64`,
		MStackSize:          Medium,
	}

	SulfuricAcid = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `A mineral acid produced by combining Sulfur and Water in a complex reaction. Primarily used in refinement processes and Battery production.`,
		MDisplayName:            `Sulfuric Acid`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 255,
			R: 255,
			A: 255,
		},
		MForm: Liquid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/SulfuricAcid/UI/IconDesc_LiquidSulfuricAcid_Pipe_512.IconDesc_LiquidSulfuricAcid_Pipe_512`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 16,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/SulfuricAcid/UI/IconDesc_LiquidSulfuricAcid_Pipe_256.IconDesc_LiquidSulfuricAcid_Pipe_256`,
		MStackSize:          Fluid,
	}

	TurboFuel = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Turbofuel, packaged for alternative transport. Can be used as fuel for Vehicles.`,
		MDisplayName:            `Packaged Turbofuel`,
		MEnergyValue:            2000.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 46,
			G: 41,
			R: 212,
			A: 255,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Turbofuel/UI/IconDesc_TurboFuel_256.IconDesc_TurboFuel_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 570,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Turbofuel/UI/IconDesc_TurboFuel_64.IconDesc_TurboFuel_64`,
		MStackSize:          Medium,
	}

	UraniumCell = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Uranium Cells are produced from Uranium Ore. 
Used to produce Uranium Fuel Rods for Nuclear Power production.

Caution: Mildly Radioactive.`,
		MDisplayName: `Encased Uranium Cell`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/UraniumCell/UI/IconDesc_NuclearCell_256.IconDesc_NuclearCell_256`,
		MRadioactiveDecay:   0.500000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 147,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/UraniumCell/UI/IconDesc_NuclearCell_64.IconDesc_NuclearCell_64`,
		MStackSize:          Big,
	}

	Wire = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for crafting.
One of the most basic parts.`,
		MDisplayName: `Wire`,
		MEnergyValue: 0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/Wire/UI/IconDesc_Wire_256.IconDesc_Wire_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 6,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/Wire/UI/IconDesc_Wire_64.IconDesc_Wire_64`,
		MStackSize:          Huge,
	}

	XmasBall1 = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Used for making FICSMAS Decorations.`,
		MDisplayName:            `Red FICSMAS Ornament`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_XmasBall_Red_256.IconDesc_XmasBall_Red_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_XmasBall_Red_64.IconDesc_XmasBall_Red_64`,
		MStackSize:          Huge,
	}

	XmasBall2 = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Again, used for making FICSMAS Decorations.`,
		MDisplayName:            `Blue FICSMAS Ornament`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_XmasBall_Blue_256.IconDesc_XmasBall_Blue_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_XmasBall_Blue_64.IconDesc_XmasBall_Blue_64`,
		MStackSize:          Huge,
	}

	XmasBall3 = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Still used for making FICSMAS Decorations.`,
		MDisplayName:            `Copper FICSMAS Ornament`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_XmasBall_Yellow_256.IconDesc_XmasBall_Yellow_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_XmasBall_Yellow_64.IconDesc_XmasBall_Yellow_64`,
		MStackSize:          Big,
	}

	XmasBall4 = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `This super special... nope... still just used for making FICSMAS Decorations.`,
		MDisplayName:            `Iron FICSMAS Ornament`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_XmasBall_Silver_256.IconDesc_XmasBall_Silver_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_XmasBall_Silver_64.IconDesc_XmasBall_Silver_64`,
		MStackSize:          Big,
	}

	XmasBallCluster = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `All the FICSMAS Ornaments smashed together to make even more FICSMAS Decorations!`,
		MDisplayName:            `FICSMAS Ornament Bundle`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Balls_256.IconDesc_Balls_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Balls_64.IconDesc_Balls_64`,
		MStackSize:          Medium,
	}

	XmasBow = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `A fancy Bow, maybe someone can wear this? You certainly can't! Probably, some parts and decorations can be made from this.`,
		MDisplayName:            `FICSMAS Bow`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Bow_256.IconDesc_Bow_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Bow_64.IconDesc_Bow_64`,
		MStackSize:          Huge,
	}

	XmasBranch = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `A special Tree Branch, used to produce parts and buildings during the FICSMAS Event.`,
		MDisplayName:            `FICSMAS Tree Branch`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Branch_256.IconDesc_Branch_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Branch_64.IconDesc_Branch_64`,
		MStackSize:          Huge,
	}

	XmasStar = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `This special FICSMAS Star signifies the productivity of FICSIT all across the universe. It also signifies the fact that you have nearly completed the Holiday Event, so it's time to get back to work.`,
		MDisplayName:            `FICSMAS Wonder Star`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Star_256.IconDesc_Star_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Star_64.IconDesc_Star_64`,
		MStackSize:          Small,
	}

	XmasWreath = ItemDescriptor{
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `A decoration used to make decorations. Its use cases are questionable.`,
		MDisplayName:            `FICSMAS Decoration`,
		MEnergyValue:            0.000000,
		MFluidColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MForm: Solid,
		MGasColor: struct {
			B int
			G int
			R int
			A int
		}{
			B: 0,
			G: 0,
			R: 0,
			A: 0,
		},
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Wreath_256.IconDesc_Wreath_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 1,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Wreath_64.IconDesc_Wreath_64`,
		MStackSize:          Medium,
	}
)
