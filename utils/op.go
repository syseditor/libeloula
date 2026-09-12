package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/syseditor/libeloula/connection"
)

const opFile = "ops.txt" //can be changed or derived from a settings file

var ops []string

func LoadOps() {
	if _, err := os.Stat(opFile); err != nil {
		_, err = os.Create(opFile)
		Check(err)
	}

	f, err := os.OpenFile(opFile, os.O_CREATE|os.O_RDWR, 0665)
	if err != nil {
		Check(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			Check(err)
		}

		ops = append(ops, strings.ReplaceAll(line, "\n", ""))
	}

	fmt.Println("Loaded ops:", ops)
}

func IsOp(name string) bool {
	return slices.Contains(ops, name)
}

func AddOp(name string) bool {
	if IsOp(name) {
		return false
	}

	f, err := os.OpenFile(opFile, os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		Check(err)
	}

	defer f.Close()

	fmt.Fprintf(f, "%s\n", name)
	ops = append(ops, name)

	return true
}

func RemoveOp(name string) bool {
	if !IsOp(name) {
		return false
	}

	f, err := os.OpenFile(opFile, os.O_RDWR, 0660)
	if err != nil {
		Check(err)
	}

	defer f.Close()

	removeOp(name)
	for _, op := range ops {
		fmt.Fprintf(f, "%s\n", op)
	}

	return true
}

func removeOp(name string) {
	var index int
	for i, op := range ops {
		if op == name {
			index = i
			break
		}
	}

	ops = append(ops[:index], ops[index+1:]...)
}

func SendOperatorAbilityPacket(handle *world.EntityHandle) {
	pk := &packet.UpdateAbilities{
			AbilityData: protocol.AbilityData{
				EntityUniqueID: 0,
				PlayerPermissions: 2,
				CommandPermissions: 2,
				Layers: []protocol.AbilityLayer{
					{
						Type: protocol.AbilityLayerTypeBase,
						Abilities: protocol.AbilityAttackMobs | protocol.AbilityAttackPlayers | protocol.AbilityWorldBuilder | protocol.AbilityOperatorCommands | protocol.AbilityMine | protocol.AbilityDoorsAndSwitches,
						Values: protocol.AbilityAttackMobs | protocol.AbilityMine | protocol.AbilityWorldBuilder | protocol.AbilityDoorsAndSwitches,
						FlySpeed: protocol.AbilityBaseFlySpeed,
						VerticalFlySpeed: protocol.AbilityBaseVerticalFlySpeed,
						WalkSpeed: protocol.AbilityBaseWalkSpeed,
					},
				},
			},
		}

	err := connection.GetConn(handle.UUID().String()).WritePacket(pk)
	Check(err)
}
