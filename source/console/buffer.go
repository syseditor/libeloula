package console

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/syseditor/libeloula/utils"
)

func InitBuffer() {
	fmt.Printf("%sConsole Command Buffer: Type any command to %sdirectly%s execute\n", text.ANSI(text.Amethyst), text.ANSI(text.Red), text.ANSI(text.Amethyst))

	reader := bufio.NewReader(os.Stdin)
	console := &ConsoleCommandSource{}

	terminationSignal := make(chan os.Signal, 1)
	signal.Notify(terminationSignal, syscall.SIGTERM, syscall.SIGINT)

	input := make(chan string, 1)
	bufferContinue := make(chan int, 1)

	go func() {
		for {
			fmt.Printf("%s>> ", text.ANSI(text.Green))
			str, _ := reader.ReadString('\n')
			input <- str
		bufferWait:
			select {
			case <-bufferContinue:
				break bufferWait
			default:
				continue
			}
		}
	}()

loop:
	for {
		select {
		case <-terminationSignal:
			fmt.Printf("%s\nTerminating session and server...\n%s", text.ANSI(text.DarkRed), text.ANSI(text.Reset))
			break loop
		case str := <-input:
			bufferContinue <- 0
			str = strings.ReplaceAll(str, "\n", "")

			if len(str) == 0 {
				bufferContinue <- 1
				continue
			}

			cmdWithArgs := strings.Split(str, " ")

			command, found := cmd.ByAlias(cmdWithArgs[0])

			if !found {
				fmt.Printf("%sCommand %s not found.", text.ANSI(text.Redstone), cmdWithArgs[0])
				bufferContinue <- 1
				continue
				// testing this, might not be needed if SendCommandOutput works properly
			}

			command.Execute(strings.Join(cmdWithArgs[1:], " "), *console, nil)
			bufferContinue <- 1
		}

	}

	utils.Server.CloseOnProgramEnd()
}
