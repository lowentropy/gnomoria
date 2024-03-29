package gnomoria

import (
	"math/rand"
)

var goods []string = []string{
	"ammo_pouch",
	"anvil",
	"armor_plate",
	"axle",
	"bag",
	"ballpeen_hammer",
	"bandage",
	"bar",
	"barrel",
	"battle_axe",
	"battle_axe_head",
	"bed",
	"bed_frame",
	"beer",
	"bellows",
	"blade_trap",
	"block",
	"blunderbuss",
	"blunderbuss_barrel",
	"blunderbuss_stock",
	"bolt",
	"bolt_of_cloth",
	"bone",
	"bone_needle",
	"bone_shirt",
	"bone_statuette",
	"boot",
	"bread",
	"breastplate",
	"brick_block",
	"chair",
	"chisel",
	"clay",
	"claymore",
	"claymore_blade",
	"coal",
	"commemorative_coin",
	"crate",
	"crossbow",
	"crossbow_bow",
	"crossbow_stock",
	"cut_gem",
	"cutting_wheel",
	"door",
	"felling_axe",
	"felling_axe_head",
	"fiber",
	"file",
	"flintlock_pistol",
	"four_poster_bed",
	"four_poster_bed_frame",
	"furnace",
	"gauntlet",
	"gear",
	"gearbox",
	"gemmed_necklace",
	"gemmed_ring",
	"grain",
	"greave",
	"haft",
	"hammer",
	"hammer_head",
	"hand_axe",
	"hand_axe_head",
	"handcrank",
	"hatch",
	"hearth",
	"helmet",
	"hide",
	"hilt",
	"knife",
	"knife_blade",
	"leather_boot",
	"leather_bracer",
	"leather_cuirass",
	"leather_glove",
	"leather_greave",
	"leather_helm",
	"leather_panel",
	"lever",
	"log",
	"loom",
	"mattress",
	"meat",
	"mechanical_wall",
	"mechanism_base",
	"mold",
	"musket_round",
	"necklace",
	"ore",
	"padding",
	"pauldron",
	"pet_rock",
	"pickaxe",
	"pickaxe_head",
	"pillar",
	"pistol_barrel",
	"pistol_stock",
	"plank",
	"pressure_plate",
	"puzzle_box",
	"quiver",
	"raw_gem",
	"raw_stone",
	"ring",
	"rod",
	"sandwich",
	"sausage",
	"sawblade",
	"shield",
	"shield_backing",
	"shield_boss",
	"skull",
	"skull_helmet",
	"spike",
	"spike_trap",
	"spring",
	"statue",
	"statuette",
	"stick",
	"strap",
	"straw",
	"strawberry",
	"string",
	"sword",
	"sword_blade",
	"table",
	"torch",
	"training_dummy",
	"trap_base",
	"warhammer",
	"warhammer_head",
	"windmill_blade",
	"wine",
	"workbench",
	"wrench",
	"yak",
}

var workshops []string = []string{
	"armorer",
	"blacksmith",
	"bonecarver",
	"butcher_shop",
	"carpenter",
	"distillery",
	"engineer_shop",
	"forge",
	"furnace",
	"gemcutter",
	"jeweler",
	"kiln",
	"kitchen",
	"leatherworker",
	"loom",
	"machine_shop",
	"metalworker",
	"sawmill",
	"smelter",
	"stonecarver",
	"stonecutter",
	"stonemason",
	"tailor",
	"tinker_bench",
	"training_grounds",
	"weaponsmith",
	"well",
	"woodcarver",
}

var workshopGoods map[string][]string = map[string][]string{
	"armorer": []string{
		"armor_plate",
		"bar",
		"boot",
		"breastplate",
		"gauntlet",
		"greave",
		"helmet",
		"padding",
		"pauldron",
		"strap",
	},
	"blacksmith": []string{
		"ballpeen_hammer",
		"bar",
		"cutting_wheel",
		"felling_axe",
		"felling_axe_head",
		"file",
		"haft",
		"pickaxe",
		"pickaxe_head",
	},
	"bonecarver": []string{
		"bone",
		"bone_needle",
		"bone_shirt",
		"bone_statuette",
		"skull",
		"skull_helmet",
		"strap",
	},
	"butcher_shop": []string{
		"bone",
		"hide",
		"meat",
		"skull",
		"yak",
	},
	"carpenter": []string{
		"barrel",
		"bed",
		"bed_frame",
		"bellows",
		"blunderbuss_stock",
		"chair",
		"coal",
		"crate",
		"crossbow_stock",
		"door",
		"four_poster_bed",
		"four_poster_bed_frame",
		"haft",
		"hilt",
		"loom",
		"mattress",
		"pistol_stock",
		"plank",
		"stick",
		"table",
		"torch",
		"training_dummy",
		"workbench",
	},
	"distillery": []string{
		"beer",
		"grain",
		"strawberry",
		"wine",
	},
	"engineer_shop": []string{
		"axle",
		"bar",
		"blade_trap",
		"blunderbuss",
		"blunderbuss_barrel",
		"blunderbuss_stock",
		"bolt_of_cloth",
		"crossbow",
		"crossbow_bow",
		"crossbow_stock",
		"flintlock_pistol",
		"gear",
		"gearbox",
		"handcrank",
		"hatch",
		"lever",
		"mechanical_wall",
		"mechanism_base",
		"pistol_barrel",
		"pistol_stock",
		"plank",
		"pressure_plate",
		"rod",
		"spike",
		"spike_trap",
		"spring",
		"sword_blade",
		"trap_base",
		"windmill_blade",
	},
	"forge": []string{
		"anvil",
		"bar",
		"coal",
		"ore",
	},
	"furnace": []string{
		"coal",
		"log",
	},
	"gemcutter": []string{
		"cut_gem",
		"raw_gem",
	},
	"jeweler": []string{
		"bar",
		"cut_gem",
		"gemmed_necklace",
		"gemmed_ring",
		"necklace",
		"ring",
	},
	"kiln": []string{
		"brick_block",
		"clay",
		"coal",
		"statuette",
	},
	"kitchen": []string{
		"bread",
		"grain",
		"meat",
		"sandwich",
		"sausage",
	},
	"leatherworker": []string{
		"hide",
		"leather_boot",
		"leather_bracer",
		"leather_cuirass",
		"leather_glove",
		"leather_greave",
		"leather_helm",
		"leather_panel",
		"padding",
		"quiver",
		"strap",
		"string",
	},
	"loom": []string{
		"bolt_of_cloth",
		"fiber",
	},
	"machine_shop": []string{
		"bar",
		"blunderbuss_barrel",
		"bolt",
		"crossbow_bow",
		"gear",
		"musket_round",
		"pistol_barrel",
		"rod",
		"spike",
		"spring",
		"wrench",
	},
	"metalworker": []string{
		"bar",
		"commemorative_coin",
		"statue",
		"statuette",
	},
	"sawmill": []string{
		"log",
		"plank",
	},
	"smelter": []string{},
	"stonecarver": []string{
		"block",
		"pet_rock",
		"pillar",
		"statue",
		"statuette",
	},
	"stonecutter": []string{
		"block",
		"raw_stone",
	},
	"stonemason": []string{
		"block",
		"bolt",
		"chair",
		"chisel",
		"door",
		"furnace",
		"hearth",
		"hilt",
		"knife",
		"knife_blade",
		"mold",
		"musket_round",
		"sawblade",
		"table",
	},
	"tailor": []string{
		"ammo_pouch",
		"bag",
		"bandage",
		"bolt_of_cloth",
		"fiber",
		"mattress",
		"padding",
		"straw",
		"string",
	},
	"tinker_bench":     []string{},
	"training_grounds": []string{},
	"weaponsmith": []string{
		"bar",
		"battle_axe",
		"battle_axe_head",
		"claymore",
		"claymore_blade",
		"haft",
		"hammer",
		"hammer_head",
		"hand_axe",
		"hand_axe_head",
		"hilt",
		"shield",
		"shield_backing",
		"shield_boss",
		"sword",
		"sword_blade",
		"warhammer",
		"warhammer_head",
	},
	"well": []string{},
	"woodcarver": []string{
		"plank",
		"puzzle_box",
		"statue",
		"statuette",
	},
}

func RandomGood() string {
	return goods[rand.Intn(len(goods))]
}

func RandomWorkshop() string {
	return workshops[rand.Intn(len(workshops))]
}

func WorkshopGoods(workshop string) []string {
	return workshopGoods[workshop]
}

func NumWorkshops() int {
	return len(workshops)
}

func NumGoods() int {
	return len(goods)
}
