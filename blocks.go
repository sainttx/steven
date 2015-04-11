// Copyright 2015 Matthew Collins
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

// Valid blocks.
var (
	BlockAir                        = initSimpleConfig("air", simpleConfig{NotCullAgainst: true, NoCollision: true})
	BlockStone                      = initStone("stone")
	BlockGrass                      = initGrass()
	BlockDirt                       = initSimple("dirt")
	BlockCobblestone                = initSimple("cobblestone")
	BlockPlanks                     = initPlanks("planks")
	BlockSapling                    = initSapling("sapling")
	BlockBedrock                    = initSimple("bedrock")
	BlockFlowingWater               = initLiquid("flowing_water", false)
	BlockWater                      = initLiquid("water", false)
	BlockFlowingLava                = initLiquid("flowing_lava", true)
	BlockLava                       = initLiquid("lava", true)
	BlockSand                       = initSimple("sand")
	BlockGravel                     = initSimple("gravel")
	BlockGoldOre                    = initSimple("gold_ore")
	BlockIronOre                    = initSimple("iron_ore")
	BlockCoalOre                    = initSimple("coal_ore")
	BlockLog                        = initLog("log", false)
	BlockLeaves                     = initLeaves("leaves", false)
	BlockSponge                     = initSponge("sponge")
	BlockGlass                      = initSimpleConfig("glass", simpleConfig{NotCullAgainst: true})
	BlockLapisOre                   = initSimple("lapis_ore")
	BlockLapisBlock                 = initSimple("lapis_block")
	BlockDispenser                  = initDispenser("dispenser")
	BlockSandstone                  = initSimple("sandstone")
	BlockNote                       = initSimple("noteblock")
	BlockBed                        = initBed("bed")
	BlockGoldenRail                 = initPoweredRail("golden_rail")
	BlockDetectorRail               = initPoweredRail("detector_rail")
	BlockStickyPiston               = initSimple("sticky_piston")
	BlockWeb                        = initSimpleConfig("web", simpleConfig{NotCullAgainst: true, NoCollision: true})
	BlockTallGrass                  = initTallGrass()
	BlockDeadBush                   = initDeadBush("deadbush")
	BlockPiston                     = initSimple("piston")
	BlockPistonHead                 = initSimple("piston_head")
	BlockWool                       = initSimple("wool")
	BlockPistonExtension            = initSimple("piston_extension")
	BlockYellowFlower               = initSimple("yellow_flower")
	BlockRedFlower                  = initSimple("red_flower")
	BlockBrownMushroom              = initSimpleConfig("brown_mushroom", simpleConfig{NotCullAgainst: true, NoCollision: true})
	BlockRedMushrrom                = initSimpleConfig("red_mushroom", simpleConfig{NotCullAgainst: true, NoCollision: true})
	BlockGoldBlock                  = initSimple("gold_block")
	BlockIronBlock                  = initSimple("iron_block")
	BlockDoubleStoneSlab            = initSimple("double_stone_slab")
	BlockStoneSlab                  = initSimple("stone_slab")
	BlockBrickBlock                 = initSimple("brick_block")
	BlockTNT                        = initSimple("tnt")
	BlockBookShelf                  = initSimple("blockshelf")
	BlockMossyCobblestone           = initSimple("mossy_cobblestone")
	BlockObsidian                   = initSimple("obsidian")
	BlockTorch                      = initSimple("torch")
	BlockFire                       = initSimple("fire")
	BlockMobSpawner                 = initSimple("mob_spawner")
	BlockOakStairs                  = initStairs("oak_stairs")
	BlockChest                      = initSimple("chest")
	BlockRedstoneWire               = initSimple("redstone_wire")
	BlockDiamondOre                 = initSimple("diamond_ore")
	BlockDiamondBlock               = initSimple("diamond_block")
	BlockCraftingTable              = initSimple("crafting_table")
	BlockWheat                      = initSimple("wheat")
	BlockFarmland                   = initSimple("farmland")
	BlockFurnace                    = initSimple("furnace")
	BlockFurnaceLit                 = initSimple("lit_furnace")
	BlockStandingSign               = initSimple("standing_sign")
	BlockWoodenDoor                 = initDoor("wooden_door")
	BlockLadder                     = initSimple("ladder")
	BlockRail                       = initRail("rail")
	BlockStoneStairs                = initStairs("stone_stairs")
	BlockWallSign                   = initSimple("wall_sign")
	BlockLever                      = initSimple("lever")
	BlockStonePressurePlate         = initSimple("stone_pressure_plate")
	BlockIronDoor                   = initDoor("iron_door")
	BlockWoodenPressurePlate        = initSimple("wooden_pressure_plate")
	BlockRedstoneOre                = initSimple("redstone_ore")
	BlockRedstoneOreLit             = initSimple("lit_redstone_ore")
	BlockRedstoneTorchUnlit         = initSimple("unlit_redstone_torch")
	BlockRedstoneTorch              = initSimple("redstone_torch")
	BlockStoneButton                = initSimple("stone_button")
	BlockSnowLayer                  = initSimple("snow_layer")
	BlockIce                        = initSimple("ice")
	BlockSnow                       = initSimple("snow")
	BlockCactus                     = initSimpleConfig("cactus", simpleConfig{NotCullAgainst: true})
	BlockClay                       = initSimple("clay")
	BlockReeds                      = initSimpleConfig("reeds", simpleConfig{NotCullAgainst: true, NoCollision: true})
	BlockJukebox                    = initSimple("jukebox")
	BlockFence                      = initFence("fence", true)
	BlockPumpkin                    = initSimple("pumpkin")
	BlockNetherrack                 = initSimple("netherrack")
	BlockSoulSand                   = initSimple("soul_sand")
	BlockGlowstone                  = initSimple("glowstone")
	BlockPortal                     = initSimple("portal")
	BlockPumpkinLit                 = initSimple("lit_pumpkin")
	BlockCake                       = initSimple("cake")
	BlockRepeaterUnpowered          = initSimple("unpowered_repeater")
	BlockRepeaterPowered            = initSimple("powered_repeater")
	BlockStainedGlass               = initStainedGlass("stained_glass")
	BlockTrapDoor                   = initSimple("trapdoor")
	BlockMonsterEgg                 = initSimple("monster_egg")
	BlockStoneBrick                 = initSimple("stonebrick")
	BlockBrownMushroomBlock         = initSimple("brown_mushroom_block")
	BlockRedMushroomBlock           = initSimple("red_mushroom_block")
	BlockIronBars                   = initConnectable("iron_bars")
	BlockGlassPane                  = initConnectable("glass_pane")
	BlockMelonBlock                 = initSimple("melon_block")
	BlockPumpkinStem                = initSimple("pumpkin_stem")
	BlockMelonStem                  = initSimple("melon_stem")
	BlockVine                       = initVines("vine")
	BlockFenceGate                  = initSimple("fence_gate")
	BlockBrickStairs                = initStairs("brick_stairs")
	BlockStoneBrickStairs           = initStairs("stone_brick_stairs")
	BlockMycelium                   = initSimple("mycelium")
	BlockWaterlily                  = initSimple("waterlily")
	BlockNetherBrick                = initSimple("nether_brick")
	BlockNetherBrickFence           = initFence("nether_brick_fence", false)
	BlockNetherBrickStairs          = initStairs("nether_brick_stairs")
	BlockNetherWart                 = initSimple("nether_wart")
	BlockEnchantingTable            = initSimple("enchanting_table")
	BlockBrewingStand               = initSimple("brewing_stand")
	BlockCauldron                   = initSimple("cauldron")
	BlockEndPortal                  = initSimple("end_portal")
	BlockEndPortalFrame             = initSimple("end_portal_frame")
	BlockEndStone                   = initSimple("end_stone")
	BlockDragonEgg                  = initSimple("dragon_egg")
	BlockRedstoneLamp               = initSimple("redstone_lamp")
	BlockRedstoneLampLit            = initSimple("lit_redstone_lamp")
	BlockDoubleWoodenSlab           = initSimple("double_wooden_slab")
	BlockWoodenSlab                 = initSimple("wooden_slab")
	BlockCocoa                      = initSimple("cocoa")
	BlockSandstoneStairs            = initStairs("sandstone_stairs")
	BlockEmeraldOre                 = initSimple("emerald_ore")
	BlockEnderChest                 = initSimple("ender_Chest")
	BlockTripwireHook               = initSimple("tripwire_hook")
	BlockTripwire                   = initSimple("tripwire")
	BlockEmeraldBlock               = initSimple("emerald_block")
	BlockSpruceStairs               = initStairs("spruce_stairs")
	BlockBirchStairs                = initStairs("birch_stairs")
	BlockJungleStairs               = initStairs("jungle_stairs")
	BlockCommandBlock               = initSimple("command_block")
	BlockBeacon                     = initSimple("beacon")
	BlockCobblestoneWall            = initSimple("cobblestone_wall")
	BlockFlowerPot                  = initSimple("flower_pot")
	BlockCarrots                    = initSimple("carrots")
	BlockPotatoes                   = initSimple("potatoes")
	BlockWoodenButton               = initSimple("wooden_button")
	BlockSkull                      = initSimple("skull")
	BlockAnvil                      = initSimple("anvil")
	BlockTrappedChest               = initSimple("trapped_chest")
	BlockLightWeightedPressurePlate = initSimple("light_weighted_pressure_plate")
	BlockHeavyWeightedPressurePlate = initSimple("heavy_weighted_pressure_plate")
	BlockComparatorUnpowered        = initSimple("unpowered_comparator")
	BlockComparatorPowered          = initSimple("powered_comparator")
	BlockDaylightDetector           = initSimple("daylight_detector")
	BlockRedstoneBlock              = initSimple("redstone_block")
	BlockQuartzOre                  = initSimple("quartz_ore")
	BlockHopper                     = initSimple("hopper")
	BlockQuartzBlock                = initSimple("quartz_block")
	BlockQuartzStairs               = initStairs("quartz_stairs")
	BlockActivatorRail              = initPoweredRail("activator_rail")
	BlockDropper                    = initDispenser("dropper")
	BlockStainedHardenedClay        = initStainedClay("stained_hardened_clay")
	BlockStainedGlassPane           = initStainedGlassPane("stained_glass_pane")
	BlockLeaves2                    = initLeaves("leaves2", true)
	BlockLog2                       = initLog("log2", true)
	BlockAcaciaStairs               = initStairs("acacia_stairs")
	BlockDarkOakStairs              = initStairs("dark_oak_stairs")
	BlockSlime                      = initSimple("slime")
	BlockBarrier                    = initSimple("barrier")
	BlockIronTrapDoor               = initSimple("iron_trapdoor")
	BlockPrismarine                 = initSimple("prismarine")
	BlockSeaLantern                 = initSimple("sea_lantern")
	BlockHayBlock                   = initSimple("hay_block")
	BlockCarpet                     = initSimple("carpet")
	BlockHardenedClay               = initSimple("hardened_clay")
	BlockCoalBlock                  = initSimple("coal_block")
	BlockPackedIce                  = initSimple("packed_ice")
	BlockDoublePlant                = initSimple("double_plant")
	BlockStandingBanner             = initSimple("standing_banner")
	BlockWallBanner                 = initSimple("wall_banner")
	BlockDaylightDetectorInverted   = initSimple("daylight_detector_inverted")
	BlockRedSandstone               = initSimple("red_sandstone")
	BlockRedSandstoneStairs         = initStairs("red_sandstone_stairs")
	BlockDoubleStoneSlab2           = initSimple("double_stone_slab2")
	BlockStoneSlab2                 = initSimple("stone_slab2")
	BlockSpruceFenceGate            = initSimple("spruce_fence_gate")
	BlockBirchFenceGate             = initSimple("birch_fence_gate")
	BlockJungleFenceGate            = initSimple("jungle_fence_gate")
	BlockDarkOakFenceGate           = initSimple("dark_oak_fence_gate")
	BlockAcaciaFenceGate            = initSimple("acacia_fence_gate")
	BlockSpruceFence                = initFence("spruce_fence", true)
	BlockBirchFence                 = initFence("birch_fence", true)
	BlockJungleFence                = initFence("jungle_fence", true)
	BlockDarkOakFence               = initFence("dark_oak_fence", true)
	BlockAcaciaFence                = initFence("acacia_fence", true)
	BlockSpruceDoor                 = initDoor("spruce_door")
	BlockBirchDoor                  = initDoor("birch_door")
	BlockJungleDoor                 = initDoor("jungle_door")
	BlockAcaciaDoor                 = initDoor("acacia_door")
	BlockDarkOakDoor                = initDoor("dark_oak_door")
)
