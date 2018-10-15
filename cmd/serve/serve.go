package serve

import (
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/percona/pmm-agent/app"
)

// New returns `serve` command.
func New(ctx context.Context, app *app.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve accepts incoming connections and starts supervisor api.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.RunServer(ctx)
		},
	}

	app.Server.Flags(cmd)
	app.Config.Bind(cmd)
	return cmd
}
