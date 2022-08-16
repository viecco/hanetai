package main

import (
	"context"

	"github.com/alecthomas/kong"
)

var cli struct {
	AccessToken string `kong:"required,env='HANET_ACCESS_TOKEN'"`
	JSON        bool   `kong:"optional,name='json',default:false"`
	NoHeader    bool   `kong:"optional,name='no-header',default:false"`
	Person      struct {
		Ls             PersonLsCmd        `cmd:"" help:"List person at the place."`
		Rm             PersonRmCmd        `cmd:"" help:"Remove a person using their ID."`
		RmByPlaceAlias PersonRmByAliasCmd `cmd:"" help:"Remove a person from the place"`
	} `cmd:""`
	Device struct {
		Status DeviceConnectionStatusCmd `cmd:"" help:"Get device connection status."`
	} `cmd:""`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&CliContext{
		AccessToken: cli.AccessToken,
		Context:     context.Background(),
		Debug:       false,
		JSON:        cli.JSON,
		NoHeader:    cli.NoHeader,
	})
	ctx.FatalIfErrorf(err)
}
