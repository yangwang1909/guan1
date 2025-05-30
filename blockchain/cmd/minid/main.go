package main

import (
	"fmt"
	"os"

	clienthelpers "cosmossdk.io/client/v2/helpers"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/cosmosregistry/chain-minimal/app"
	"github.com/cosmosregistry/chain-minimal/app/params"
	"github.com/cosmosregistry/chain-minimal/cmd/minid/cmd"
)

func main() {
	params.SetAddressPrefixes()

	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, clienthelpers.EnvPrefix, app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
