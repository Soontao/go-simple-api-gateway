package main

import (
	"github.com/mkideal/cli"
	"github.com/Soontao/go-simple-api-gateway/server"
)

type cliArgs struct {
	cli.Helper
	ConnectionStr string `cli:"*c,*conn" usage:"mysql connection str" dft:"$GATEWAY_CONN_STR"`
	ListenAddress string `cli:"*l,*listen" usage:"listen host and port" dft:"$GATEWAY_LS"`
}

func main() {
	cli.Run(new(cliArgs), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*cliArgs)
		server.NewAuthServer(argv.ConnectionStr).Start(argv.ListenAddress)
		return nil
	})
}
