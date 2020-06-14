package console

import "fmt"

type Command interface {
    Handle(...string)
    GetSignature() string
}

var Commands map[string]Command = make(map[string]Command)

func Handle(command string, args... string) {
    Commands[command].Handle(args...)
}

func RegisterCommand(commands []Command) {
    for _, c := range commands {
        fmt.Println(c)
        Commands[c.GetSignature()] = c
    }
}
