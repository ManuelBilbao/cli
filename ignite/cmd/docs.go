package ignitecmd

import (
	"github.com/spf13/cobra"

	"github.com/manuelbilbao/cli/v28/docs"
	"github.com/manuelbilbao/cli/v28/ignite/pkg/localfs"
	"github.com/manuelbilbao/cli/v28/ignite/pkg/markdownviewer"
)

func NewDocs() *cobra.Command {
	c := &cobra.Command{
		Use:   "docs",
		Short: "Show Ignite CLI docs",
		Args:  cobra.NoArgs,
		RunE:  docsHandler,
	}
	return c
}

func docsHandler(*cobra.Command, []string) error {
	path, cleanup, err := localfs.SaveTemp(docs.Docs)
	if err != nil {
		return err
	}
	defer cleanup()

	return markdownviewer.View(path)
}
