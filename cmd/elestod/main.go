package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	elestoapp "github.com/elesto-dao/elesto/v4/app"
	"github.com/elesto-dao/elesto/v4/cmd/elestod/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, elestoapp.DefaultNodeHome); err != nil {
		switch e := err.(type) { // nolint
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
