package release

import "github.com/spf13/cobra"

func Release(cmd *cobra.Command, args []string) {
	u, p := Login("localhost")
	if u == "" && p == "" {
		panic(nil)
	}
}
