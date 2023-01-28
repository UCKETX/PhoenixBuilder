package blockNBT_depends

// 此表用于记录现阶段支持了的方块实体
var index = map[string]string{
	"command_block":           "CommandBlock",
	"chain_command_block":     "CommandBlock",
	"repeating_command_block": "CommandBlock",
	// 命令方块
	"blast_furnace":      "Container",
	"lit_blast_furnace":  "Container",
	"smoker":             "Container",
	"lit_smoker":         "Container",
	"furnace":            "Container",
	"lit_furnace":        "Container",
	"chest":              "Container",
	"barrel":             "Container",
	"trapped_chest":      "Container",
	"hopper":             "Container",
	"dispenser":          "Container",
	"dropper":            "Container",
	"cauldron":           "Container",
	"lava_cauldron":      "Container",
	"jukebox":            "Container",
	"brewing_stand":      "Container",
	"undyed_shulker_box": "Container",
	"shulker_box":        "Container",
	// 容器
	"standing_sign":          "Sign",
	"spruce_standing_sign":   "Sign",
	"birch_standing_sign":    "Sign",
	"jungle_standing_sign":   "Sign",
	"acacia_standing_sign":   "Sign",
	"darkoak_standing_sign":  "Sign",
	"mangrove_standing_sign": "Sign",
	"bamboo_standing_sign":   "Sign",
	"crimson_standing_sign":  "Sign",
	"warped_standing_sign":   "Sign",
	"wall_sign":              "Sign",
	"spruce_wall_sign":       "Sign",
	"birch_wall_sign":        "Sign",
	"jungle_wall_sign":       "Sign",
	"acacia_wall_sign":       "Sign",
	"darkoak_wall_sign":      "Sign",
	"mangrove_wall_sign":     "Sign",
	"bamboo_wall_sign":       "Sign",
	"crimson_wall_sign":      "Sign",
	"warped_wall_sign":       "Sign",
	"sign":                   "Sign",
	"spruce_sign":            "Sign",
	"birch_sign":             "Sign",
	"jungle_sign":            "Sign",
	"acacia_sign":            "Sign",
	"darkoak_sign":           "Sign",
	"mangrove_sign":          "Sign",
	"bamboo_sign":            "Sign",
	"crimson_sign":           "Sign",
	"warped_sign":            "Sign",
	"oak_hanging_sign":       "Sign",
	"spruce_hanging_sign":    "Sign",
	"birch_hanging_sign":     "Sign",
	"jungle_hanging_sign":    "Sign",
	"acacia_hanging_sign":    "Sign",
	"dark_oak_hanging_sign":  "Sign",
	"mangrove_hanging_sign":  "Sign",
	"bamboo_hanging_sign":    "Sign",
	"crimson_hanging_sign":   "Sign",
	"warped_hanging_sign":    "Sign",
	// 告示牌
	"glow_frame": "Frame",
	"frame":      "Frame",
	// 物品展示框
	"lectern": "lectern",
	// 讲台
}

// 检查这个方块实体是否已被支持
func CheckIfIsEffectiveNBTBlock(blockName string) string {
	value, ok := index[blockName]
	if ok {
		return value
	}
	return ""
}