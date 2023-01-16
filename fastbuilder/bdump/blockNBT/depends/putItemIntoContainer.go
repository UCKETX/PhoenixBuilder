package blockNBT_depends

import (
	"fmt"
	"phoenixbuilder/fastbuilder/environment"
	"phoenixbuilder/fastbuilder/types"
	"phoenixbuilder/minecraft"
	"phoenixbuilder/minecraft/protocol"
	"phoenixbuilder/minecraft/protocol/packet"
)

// WIP
func OpenContainer(
	Environment *environment.PBEnvironment,
	ItemInfo *types.ChestSlot,
	ContainerPos [3]int32,
) error {
	if protocol.CurrentProtocol == 504 {
		networkID, ok := ItemRunTimeID[ItemInfo.Name]
		if ok {
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
								MetadataValue: uint32(ItemInfo.Damage),
							},
							BlockRuntimeID: 0,
							Count:          1,
							CanBePlacedOn:  ItemInfo.CanPlaceOn,
							CanBreak:       ItemInfo.CanDestroy,
							HasNetworkID:   false,
						},
					},
					//Position:        mgl32.Vec3{float32(posx), float32(posy), float32(posz)},
					//ClickedPosition: mgl32.Vec3{0, 0, 0},
					/*
						BlockRuntimeID: blockRuntimeID,
					*/
				},
			})
		}
	}
	return nil
}

// WIP
func PutItemIntoContainer(
	Environment *environment.PBEnvironment,
	ItemInfo *types.ChestSlot,
) error {
	if protocol.CurrentProtocol == 504 {
		stackNetworkID, ok := ItemRunTimeID[ItemInfo.Name]
		if ok {
			PlaceStackRequestAction := protocol.PlaceStackRequestAction{}
			PlaceStackRequestAction.Count = ItemInfo.Count
			PlaceStackRequestAction.Source = protocol.StackRequestSlotInfo{
				ContainerID:    0xc,
				Slot:           0,
				StackNetworkID: int32(stackNetworkID),
			}
			PlaceStackRequestAction.Destination = protocol.StackRequestSlotInfo{
				ContainerID:    7,
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
			err := Environment.Connection.(*minecraft.Conn).WritePacket(request)
			if err != nil {
				return fmt.Errorf("PutItemIntoContainer: %v", err)
			}
		}
	}
	return nil
}
