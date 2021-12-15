package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github/Shitomo/sample-cli/postgres"
)

func timeTripCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "timetrip",
		Short: "time trip",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}

			firstArg := args[0]

			dbClient, err := postgres.NewDbClient(postgres.DbConfig{
				Host: "localhost",
				Port: 5432,
				UserName: "mnd_chat",
				Password: "mnd_chat",
				DbName: "mnd_chat",
			})
			defer postgres.Close(dbClient)
			if err != nil {
				fmt.Printf("error while database connecting, caused by %s", err)
				return nil
			}
			ctx := context.Background()
			name, err := postgres.SelectNameById(dbClient)(ctx, firstArg)
			if err != nil {
				fmt.Printf("error while selecting, caused by %s", err)
				return nil
			}

			fmt.Printf("get %s", name)

			return nil
		},
	}

	return cmd
}
