package main

import (
	"github.com/mkideal/cli"
	"github.com/Soontao/go-simple-api-gateway/server"
)

type cliArgs struct {
	cli.Helper
	ConnectionStr     string `cli:"*c,*conn" usage:"mysql connection str" dft:"$GATEWAY_CONN_STR"`
	AuthServerAddress string `cli:"*a,*auth-ls" usage:"auth server listen host and port" dft:"$GATEWAY_AUTH_SERVER_LS"`
	ReverseHost       string  `cli:"*p,*proxy-target" usage:"reverse proxy target server" dft:"$GATEWAY_REVERSE_HOST"`
}

func main() {
	cli.Run(new(cliArgs), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*cliArgs)
		server.NewAuthServer(argv.ConnectionStr, argv.ReverseHost).Start(argv.AuthServerAddress)
		return nil
	})
}
