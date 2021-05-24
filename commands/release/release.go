package release

import (
	"github.com/profclems/glab/commands/cmdutils"
	releaseCreateCmd "github.com/profclems/glab/commands/release/create"
	releaseDeleteCmd "github.com/profclems/glab/commands/release/delete"
	releaseListCmd "github.com/profclems/glab/commands/release/list"
	releaseUploadCmd "github.com/profclems/glab/commands/release/upload"
	"github.com/spf13/cobra"
)

func NewCmdRelease(f *cmdutils.Factory) *cobra.Command {
	var releaseCmd = &cobra.Command{
		Use:   "release <command> [flags]",
		Short: `Manage GitLab releases`,
		Long:  ``,
	}

	cmdutils.EnableRepoOverride(releaseCmd, f)

	releaseCmd.AddCommand(releaseListCmd.NewCmdReleaseList(f))
	releaseCmd.AddCommand(releaseCreateCmd.NewCmdCreate(f, nil))
	releaseCmd.AddCommand(releaseUploadCmd.NewCmdUpload(f, nil))
	releaseCmd.AddCommand(releaseDeleteCmd.NewCmdDelete(f, nil))

	return releaseCmd
}
