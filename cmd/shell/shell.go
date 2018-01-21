package shell

import (
	"github.com/chrislusf/vasto/client"
	"context"
)

type ShellOption struct {
	// fixed cluster mode options
	FixedCluster *string
	// dynamic cluster mode options
	Master     *string
	DataCenter *string
	Keyspace   *string
	Verbose    *bool
}

type shell struct {
	option *ShellOption

	vastoClient *client.VastoClient
}

func RunShell(option *ShellOption) {
	var b = &shell{
		option: option,
		vastoClient: client.NewClient(
			&client.ClientOption{
				FixedCluster: option.FixedCluster,
				Master:       option.Master,
				DataCenter:   option.DataCenter,
				Keyspace:     option.Keyspace,
				ClientName:   "shell",
			},
		),
	}

	b.vastoClient.StartClient(context.Background())

	if *option.Verbose {
		b.vastoClient.ClusterListener.RegisterShardEventProcessor(b)
		b.vastoClient.ClusterListener.SetVerboseLog(true)
	}

	b.runShell()

}
