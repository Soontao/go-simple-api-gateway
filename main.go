package main

import (
	"github.com/Soontao/go-simple-api-gateway/server"
	"github.com/mkideal/cli"
)

type cliArgs struct {
	cli.Helper
	ConnectionStr string `cli:"*c,*conn" usage:"mysql connection str" dft:"$GATEWAY_CONN_STR"`
	ListenAddr    string `cli:"*l,*listen" usage:"gateway listen host and port" dft:"$GATEWAY_LS"`
	ResourceURL   string `cli:"*r,*resource" usage:"gateway resource url" dft:"$GATEWAY_RESOURCE_URL"`
}

func main() {
	cli.Run(new(cliArgs), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*cliArgs)
		server.NewGatewayServer(argv.ConnectionStr, argv.ResourceURL).Start(argv.ListenAddr)
		return nil
	})
}
