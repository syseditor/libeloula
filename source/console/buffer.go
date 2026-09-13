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
	bufferContinue := make(chan bool, 1)
	bufferContinue <- true

	go func() {
		for {
			select {
			case check := <-bufferContinue:
				if check {
					fmt.Printf("%s>> ", text.ANSI(text.Green))
					str, _ := reader.ReadString('\n')
					bufferContinue <- false
					input <- str
				}
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
			str = strings.ReplaceAll(str, "\n", "")

			if len(str) == 0 {
				bufferContinue <- true
				continue
			}

			if strings.EqualFold(str, "clear") {
				fmt.Print("\033[H\033[2J")
				bufferContinue <- true
				continue
			} else if strings.EqualFold(str, "stop") {
				terminationSignal <- syscall.SIGTERM
				utils.Server.Close()
				continue
			}

			cmdWithArgs := strings.Split(str, " ")

			command, found := cmd.ByAlias(cmdWithArgs[0])

			if !found {
				fmt.Printf("%sCommand %s not found.\n", text.ANSI(text.Redstone), cmdWithArgs[0])
				bufferContinue <- true
				continue
			}

			command.Execute(strings.Join(cmdWithArgs[1:], " "), *console, nil)
			bufferContinue <- true
		}
	}

	utils.Server.CloseOnProgramEnd()
}
