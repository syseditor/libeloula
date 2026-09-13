package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

func InitBuffer() {
	fmt.Printf("%sConsole Command Buffer: Type any command to %sdirectly%s execute\n", text.ANSI(text.Amethyst), text.ANSI(text.Red), text.ANSI(text.Amethyst))

	reader := bufio.NewReader(os.Stdin)
	console := &ConsoleCommandSource{}

	for {
		fmt.Print(">> ")
		str, _ := reader.ReadString('\n')

		if len(str) == 0 {
			continue
		}

		str = strings.ReplaceAll(str, "\n", "")
		cmdWithArgs := strings.Split(str, " ")

		command, found := cmd.ByAlias(cmdWithArgs[0])

		if !found {
			fmt.Printf("%sCommand not found.", text.ANSI(text.Redstone))
			// testing this, might not be needed if SendCommandOutput works properly
		}

		command.Execute(strings.Join(cmdWithArgs[1:], " "), *console, nil)
	}
}
