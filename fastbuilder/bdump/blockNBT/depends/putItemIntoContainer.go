package blockNBT_depends

import (
	"fmt"
	"phoenixbuilder/fastbuilder/environment"
	"phoenixbuilder/fastbuilder/types"
	"phoenixbuilder/io/commands"
	"phoenixbuilder/minecraft"
	"phoenixbuilder/minecraft/protocol"
	"phoenixbuilder/minecraft/protocol/packet"
	"time"
)

var ContainerOpenData interface{}
var Container_Hotbar_0_StackNetworkID int32

type EnchItem struct {
	WantPutItem   types.ChestSlot
	ContainerInfo *types.Module
}

type EnchItemList []EnchItem

func containerBlockStatesToBlockData(name string, states string) (uint16, error) {
	got, err := ParseStringBlockStates(&states)
	if err != nil {
		return 0, fmt.Errorf("containerBlockStatesToBlockData: %v", err)
	}
	blockStates := *got
	// prepare
	if name == "lava_cauldron" || name == "cauldron" {
		_, ok := blockStates["fill_level"]
		if !ok {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not find blockStates[\"fill_level\"]; blockStates = %#v", name, blockStates)
		}
		fill_level, normal := blockStates["fill_level"].(int32)
		if !normal {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not parse blockStates[\"fill_level\"]; blockStates = %#v", name, blockStates)
		}
		_, ok = blockStates["cauldron_liquid"]
		if !ok {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not find blockStates[\"cauldron_liquid\"]; blockStates = %#v", name, blockStates)
		}
		cauldron_liquid, normal := blockStates["cauldron_liquid"].(string)
		if !normal {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not parse blockStates[\"cauldron_liquid\"]; blockStates = %#v", name, blockStates)
		}
		if cauldron_liquid == "water" {
			return uint16(fill_level), nil
		}
		if cauldron_liquid == "lava" {
			return uint16(fill_level) + 8, nil
		}
		if cauldron_liquid == "powder_snow" {
			return uint16(fill_level) + 16, nil
		}
	}
	// 炼药锅（岩浆炼药锅）
	if name == "smoker" || name == "lit_smoker" || name == "trapped_chest" || name == "chest" || name == "hopper" || name == "lit_blast_furnace" || name == "blast_furnace" || name == "furnace" || name == "lit_furnace" {
		_, ok := blockStates["facing_direction"]
		if !ok {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not find blockStates[\"facing_direction\"]; blockStates = %#v", name, blockStates)
		}
		facing_direction, normal := blockStates["facing_direction"].(int32)
		if !normal {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not parse blockStates[\"facing_direction\"]; blockStates = %#v", name, blockStates)
		}
		return uint16(facing_direction), nil
	}
	// 烟熏炉（发光的烟熏炉）, 陷阱箱, 箱子, 漏斗, 熔炉（发光的熔炉）
	if name == "undyed_shulker_box" || name == "jukebox" {
		return 0, nil
	}
	// 未染色的潜影盒, 唱片机
	if name == "barrel" {
		_, ok := blockStates["facing_direction"]
		if !ok {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not find blockStates[\"facing_direction\"]; blockStates = %#v", name, blockStates)
		}
		facing_direction, normal := blockStates["facing_direction"].(int32)
		if !normal {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not parse blockStates[\"facing_direction\"]; blockStates = %#v", name, blockStates)
		}
		_, ok = blockStates["open_bit"]
		if !ok {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not find blockStates[\"open_bit\"]; blockStates = %#v", name, blockStates)
		}
		got, normal := blockStates["open_bit"].(string)
		if !normal {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not parse blockStates[\"open_bit\"]; blockStates = %#v", name, blockStates)
		}
		var open_bit uint8 = 0
		if got == "true" {
			open_bit = 1
		} else if got != "false" {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could unexpected blockStates[\"open_bit\"]; blockStates = %#v", name, blockStates)
		}
		if open_bit == 0 {
			return uint16(facing_direction), nil
		} else {
			return uint16(facing_direction) + 8, nil
		}
	}
	// 木桶
	if name == "shulker_box" {
		_, ok := blockStates["color"]
		if !ok {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not find blockStates[\"color\"]; blockStates = %#v", name, blockStates)
		}
		color, normal := blockStates["color"].(string)
		if !normal {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not parse blockStates[\"color\"]; blockStates = %#v", name, blockStates)
		}
		switch color {
		case "white":
			return 0, nil
		case "orange":
			return 1, nil
		case "magenta":
			return 2, nil
		case "light_blue":
			return 3, nil
		case "yellow":
			return 4, nil
		case "lime":
			return 5, nil
		case "pink":
			return 6, nil
		case "gray":
			return 7, nil
		case "silver":
			return 8, nil
		case "cyan":
			return 9, nil
		case "purple":
			return 10, nil
		case "blue":
			return 11, nil
		case "brown":
			return 12, nil
		case "green":
			return 13, nil
		case "red":
			return 14, nil
		case "black":
			return 15, nil
		}
	}
	// 染色的潜影盒
	if name == "lectern" {
		_, ok := blockStates["direction"]
		if !ok {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not find blockStates[\"direction\"]; blockStates = %#v", name, blockStates)
		}
		direction, normal := blockStates["direction"].(int32)
		if !normal {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not parse blockStates[\"direction\"]; blockStates = %#v", name, blockStates)
		}
		_, ok = blockStates["powered_bit"]
		if !ok {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not find blockStates[\"powered_bit\"]; blockStates = %#v", name, blockStates)
		}
		got, normal := blockStates["powered_bit"].(string)
		if !normal {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not parse blockStates[\"powered_bit\"]; blockStates = %#v", name, blockStates)
		}
		var powered_bit uint8 = 0
		if got == "true" {
			powered_bit = 1
		} else if got != "false" {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could unexpected blockStates[\"powered_bit\"]; blockStates = %#v", name, blockStates)
		}
		if powered_bit == 0 {
			return uint16(direction), nil
		} else {
			return uint16(direction) + 4, nil
		}
	}
	// 讲台
	if name == "dropper" || name == "dispenser" {
		_, ok := blockStates["facing_direction"]
		if !ok {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not find blockStates[\"facing_direction\"]; blockStates = %#v", name, blockStates)
		}
		facing_direction, normal := blockStates["facing_direction"].(int32)
		if !normal {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not parse blockStates[\"facing_direction\"]; blockStates = %#v", name, blockStates)
		}
		_, ok = blockStates["triggered_bit"]
		if !ok {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not find blockStates[\"triggered_bit\"]; blockStates = %#v", name, blockStates)
		}
		got, normal := blockStates["triggered_bit"].(string)
		if !normal {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not parse blockStates[\"triggered_bit\"]; blockStates = %#v", name, blockStates)
		}
		var triggered_bit uint8 = 0
		if got == "true" {
			triggered_bit = 1
		} else if got != "false" {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could unexpected blockStates[\"triggered_bit\"]; blockStates = %#v", name, blockStates)
		}
		if triggered_bit == 0 {
			return uint16(facing_direction), nil
		} else {
			return uint16(facing_direction) + 8, nil
		}
	}
	// 投掷器, 发射器
	if name == "brewing_stand" {
		_, ok := blockStates["brewing_stand_slot_a_bit"]
		if !ok {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not find blockStates[\"brewing_stand_slot_a_bit\"]; blockStates = %#v", name, blockStates)
		}
		got, normal := blockStates["brewing_stand_slot_a_bit"].(string)
		if !normal {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not parse blockStates[\"brewing_stand_slot_a_bit\"]; blockStates = %#v", name, blockStates)
		}
		var brewing_stand_slot_a_bit uint8 = 0
		if got == "true" {
			brewing_stand_slot_a_bit = 1
		} else if got != "false" {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could unexpected blockStates[\"brewing_stand_slot_a_bit\"]; blockStates = %#v", name, blockStates)
		}
		_, ok = blockStates["brewing_stand_slot_b_bit"]
		if !ok {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not find blockStates[\"brewing_stand_slot_b_bit\"]; blockStates = %#v", name, blockStates)
		}
		got, normal = blockStates["brewing_stand_slot_b_bit"].(string)
		if !normal {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not parse blockStates[\"brewing_stand_slot_b_bit\"]; blockStates = %#v", name, blockStates)
		}
		var brewing_stand_slot_b_bit uint8 = 0
		if got == "true" {
			brewing_stand_slot_b_bit = 1
		} else if got != "false" {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could unexpected blockStates[\"brewing_stand_slot_b_bit\"]; blockStates = %#v", name, blockStates)
		}
		_, ok = blockStates["brewing_stand_slot_c_bit"]
		if !ok {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not find blockStates[\"brewing_stand_slot_c_bit\"]; blockStates = %#v", name, blockStates)
		}
		got, normal = blockStates["brewing_stand_slot_c_bit"].(string)
		if !normal {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could not parse blockStates[\"brewing_stand_slot_c_bit\"]; blockStates = %#v", name, blockStates)
		}
		var brewing_stand_slot_c_bit uint8 = 0
		if got == "true" {
			brewing_stand_slot_c_bit = 1
		} else if got != "false" {
			return 0, fmt.Errorf("containerBlockStatesToBlockData: Crashed in %v, because of could unexpected blockStates[\"brewing_stand_slot_c_bit\"]; blockStates = %#v", name, blockStates)
		}
		slotSit := [3]uint8{brewing_stand_slot_c_bit, brewing_stand_slot_b_bit, brewing_stand_slot_a_bit}
		return uint16(slotSit[2]*4 + slotSit[1]*2 + slotSit[0]), nil
	}
	// 酿造台
	return 0, fmt.Errorf("containerBlockStatesToBlockData: %v is not a supported container; blockStates = %#v", name, blockStates)
}

func openContainer(
	Environment *environment.PBEnvironment,
	MainHandItemInfo *types.ChestSlot,
	ContainerBlockName *string,
	ContainerBlockStates *string,
	ContainerPos [3]int32,
) error {
	ContainerOpenData = nil
	got, err := containerBlockStatesToBlockData(*ContainerBlockName, *ContainerBlockStates)
	if err != nil {
		return fmt.Errorf("openContainer: %v", err)
	}
	blockRuntimeID, ok1 := ContainerDataToBlockRunTimeId[SingleContainer{*ContainerBlockName, got}]
	networkID, ok2 := ItemRunTimeID[MainHandItemInfo.Name]
	if ok1 && ok2 {
		Environment.Connection.(*minecraft.Conn).WritePacket(&packet.InventoryTransaction{
			LegacyRequestID:    0,
			LegacySetItemSlots: []protocol.LegacySetItemSlot(nil),
			Actions:            []protocol.InventoryAction{},
			TransactionData: &protocol.UseItemTransactionData{
				LegacyRequestID:    0,
				LegacySetItemSlots: []protocol.LegacySetItemSlot(nil),
				Actions:            []protocol.InventoryAction(nil),
				ActionType:         0,
				BlockPosition:      protocol.BlockPos{ContainerPos[0], ContainerPos[1], ContainerPos[2]},
				HotBarSlot:         0,
				HeldItem: protocol.ItemInstance{
					StackNetworkID: 0,
					Stack: protocol.ItemStack{
						ItemType: protocol.ItemType{
							NetworkID:     int32(networkID),
							MetadataValue: uint32(MainHandItemInfo.Damage),
						},
						BlockRuntimeID: 0,
						Count:          1,
						CanBePlacedOn:  MainHandItemInfo.CanPlaceOn,
						CanBreak:       MainHandItemInfo.CanDestroy,
						HasNetworkID:   false,
					},
				},
				BlockRuntimeID: blockRuntimeID,
			},
		})
	}
	return nil
}

func closeContainer(
	Environment *environment.PBEnvironment,
	WindowID byte,
) {
	Environment.Connection.(*minecraft.Conn).WritePacket(&packet.ContainerClose{
		WindowID:   WindowID,
		ServerSide: false,
	})
}

func requestStackNetworkID(
	Environment *environment.PBEnvironment,
) error {
	Container_Hotbar_0_StackNetworkID = 0
	cmdsender := Environment.CommandSender.(*commands.CommandSender)
	_, err := cmdsender.SendWSCommandWithResponce("replaceitem entity @s slot.hotbar 1 air")
	if err != nil {
		return fmt.Errorf("requestStackNetworkID: %v", err)
	}
	if Container_Hotbar_0_StackNetworkID == 0 {
		return fmt.Errorf("requestStackNetworkID: Failed to get the StackNetworkID")
	}
	return nil
}

func putItemIntoContainer(
	Environment *environment.PBEnvironment,
	ItemInfo *types.ChestSlot,
	ContainerID byte,
) error {
	err := requestStackNetworkID(Environment)
	if err != nil {
		return fmt.Errorf("putItemIntoContainer: %v", err)
	}
	PlaceStackRequestAction := protocol.PlaceStackRequestAction{}
	PlaceStackRequestAction.Count = ItemInfo.Count
	PlaceStackRequestAction.Source = protocol.StackRequestSlotInfo{
		ContainerID:    ContainerCombinedHotBarAndInventory,
		Slot:           0,
		StackNetworkID: Container_Hotbar_0_StackNetworkID,
	}
	Container_Hotbar_0_StackNetworkID = 0
	PlaceStackRequestAction.Destination = protocol.StackRequestSlotInfo{
		ContainerID:    ContainerID,
		Slot:           ItemInfo.Slot,
		StackNetworkID: 0,
	}
	request := &packet.ItemStackRequest{
		Requests: []protocol.ItemStackRequest{
			{
				RequestID: -1,
				Actions: []protocol.StackRequestAction{
					&PlaceStackRequestAction,
				},
				FilterStrings: []string{},
			},
		},
	}
	Environment.Connection.(*minecraft.Conn).WritePacket(request)
	return nil
}

/*
将 EnchItemList 列表中的物品放入物品，主要是为了支持附魔性物品

需要占用槽位 slot.hotbar 0

支持的容器请见当前文件夹下 container.go 中的 ContainerIdIndexMap
*/
func PutItemIntoContainerRun(
	Environment *environment.PBEnvironment,
	Input EnchItemList,
) error {
	var containerOpenData packet.ContainerOpen
	err := ReplaceitemAndEnchant(Environment, &Input[0].WantPutItem)
	if err != nil {
		return fmt.Errorf("PutItemIntoContainerRun: %v", err)
	}
	err = openContainer(Environment, &Input[0].WantPutItem, Input[0].ContainerInfo.Block.Name, &Input[0].ContainerInfo.Block.BlockStates, [3]int32{int32(Input[0].ContainerInfo.Point.X), int32(Input[0].ContainerInfo.Point.Y), int32(Input[0].ContainerInfo.Point.Z)})
	if err != nil {
		return fmt.Errorf("PutItemIntoContainerRun: %v", err)
	}
	failedCount := 0
	for {
		if failedCount > 100 {
			return fmt.Errorf("PutItemIntoContainerRun: Failed to open the container, please check the target area is loaded")
		}
		if ContainerOpenData != nil {
			containerOpenData = ContainerOpenData.(packet.ContainerOpen)
			break
		}
		failedCount++
		time.Sleep(50 * time.Millisecond)
	}
	for key, value := range Input {
		if key > 0 {
			err := ReplaceitemAndEnchant(Environment, &value.WantPutItem)
			if err != nil {
				return fmt.Errorf("PutItemIntoContainerRun: %v", err)
			}
		}
		containerID := ContainerIdIndexMap[*value.ContainerInfo.Block.Name]
		err = putItemIntoContainer(Environment, &value.WantPutItem, byte(containerID))
		if err != nil {
			return fmt.Errorf("PutItemIntoContainerRun: %v", err)
		}
	}
	closeContainer(Environment, containerOpenData.WindowID)
	ContainerOpenData = nil
	return nil
}
