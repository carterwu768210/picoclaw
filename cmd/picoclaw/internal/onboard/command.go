package onboard

import (
	"embed"

	"github.com/spf13/cobra"
)

//go:generate rm -rf templates/onboard_assets && mkdir -p templates/onboard_assets && cp -r ../../../../workspace/. templates/onboard_assets/
//go:embed templates/onboard_assets
var embeddedFiles embed.FS

func NewOnboardCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "onboard",
		Aliases: []string{"o"},
		Short:   "Initialize picoclaw configuration and workspace",
		Run: func(cmd *cobra.Command, args []string) {
			onboard()
		},
	}

	return cmd
}
