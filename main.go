package go_simple_api_gateway

import (
	"github.com/mkideal/cli"
	"github.com/Soontao/go-simple-api-gateway/server"
)

type cliArgs struct {
	cli.Helper
	ConnectionStr string `cli:"*c,*conn" usage:"mysql connection str" dft:"$API_CONN_STR"`
	ListenAddress string `cli:"*l,*listen" usage:"listen host and port" dft:"$API_HOST_LS"`
}

func main() {
	cli.Run(new(cliArgs), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*cliArgs)
		s := server.GateWayServer{}
		return nil
	})
}
