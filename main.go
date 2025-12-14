package main

import (
	"fmt"
	"os"

	"github.com/o0n1x/aggreGator/internal/cli"
	"github.com/o0n1x/aggreGator/internal/config"
)

func main() {
	cnfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	state := cli.State{
		State: &cnfg,
	}

	commands := cli.Commands{
		Commands: make(map[string]func(*cli.State, cli.Command) error),
	}

	commands.Register("login", cli.HandlerLogin)

	if len(os.Args) < 2 {
		fmt.Printf("Error: %v\n", "Invalid input No arguments")
		os.Exit(1)
	}

	err = commands.Run(&state, cli.Command{Name: os.Args[1], Args: os.Args[2:]})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// cnfg, err = config.Read()
	// if err != nil {
	// 	fmt.Printf("Error:%v", err)
	// 	return
	// }
	// fmt.Printf("db_url : %v , name: %v \n", cnfg.DB_URL, cnfg.CurrentUserName)
}
