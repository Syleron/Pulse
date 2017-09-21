package commands

import (
	"github.com/mitchellh/cli"
	"strings"
	"flag"
	"google.golang.org/grpc"
	"github.com/Syleron/PulseHA/proto"
	"context"
)

type CreateCommand struct {
	Ui cli.Ui
}

/**
 *
 */
func (c *CreateCommand) Help() string {
	helpText := `
Usage: pulseha create [options] ...
  Tells the PulseHA daemon to configure a new cluster.
Options:
  -bind-addr Pulse daemon bind address and port
`
	return strings.TrimSpace(helpText)
}

/**
 *
 */
func (c *CreateCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("create", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }

	// Get the bind-addr value
	bindAddr := cmdFlags.String("bind-addr", "127.0.0.1", "Bind address for local Pulse daemon")

	// Make sure we have cmd args
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	// If we have the default.. which we don't want.. error out.
	if *bindAddr == "127.0.0.1" {
		c.Ui.Error("Please specify a bind address.\n")
		c.Ui.Output(c.Help())
		return 1
	}

	bindAddrString := strings.Split(*bindAddr, ":")

	if len(bindAddrString) < 2 {
		c.Ui.Error("Please provide an IP and Port for PulseHA to bind on")
		c.Ui.Output(c.Help())
		return 1
	}

	connection, err := grpc.Dial("127.0.0.1:9443", grpc.WithInsecure())

	if err != nil {
		c.Ui.Error("GRPC client connection error")
		c.Ui.Error(err.Error())
		return 1
	}

	defer connection.Close()

	client := proto.NewRequesterClient(connection)

	r, err := client.Create(context.Background(), &proto.PulseCreate{
		BindIp:   bindAddrString[0],
		BindPort: bindAddrString[1],
	})

	if err != nil {
		c.Ui.Output("PulseHA CLI connection error. Is the PulseHA service running?")
		return 1
	} else {
		if r.Success {
			c.Ui.Output("\n[\u2713] " + r.Message + "\n")
		} else {
			c.Ui.Output("\n[x] " + r.Message + "\n")
		}
	}

	return 0
}

/**
 *
 */
func (c *CreateCommand) Synopsis() string {
	return "Tell Pulse to create new HA cluster"
}
