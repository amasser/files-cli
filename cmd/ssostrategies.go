package cmd

import (
	"github.com/Files-com/files-cli/lib"
	"github.com/spf13/cobra"

	files_sdk "github.com/Files-com/files-sdk-go"

	"fmt"
	"os"

	sso_strategy "github.com/Files-com/files-sdk-go/ssostrategy"
)

var (
	SsoStrategies = &cobra.Command{
		Use:  "sso-strategies [command]",
		Args: cobra.ExactArgs(1),
		Run:  func(cmd *cobra.Command, args []string) {},
	}
)

func SsoStrategiesInit() {
	var fieldsList string
	paramsSsoStrategyList := files_sdk.SsoStrategyListParams{}
	var MaxPagesList int
	cmdList := &cobra.Command{
		Use:   "list",
		Short: "list",
		Long:  `list`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			params := paramsSsoStrategyList
			params.MaxPages = MaxPagesList
			it, err := sso_strategy.List(params)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			lib.JsonMarshalIter(it, fieldsList)
		},
	}
	cmdList.Flags().StringVarP(&paramsSsoStrategyList.Cursor, "cursor", "c", "", "Used for pagination.  Send a cursor value to resume an existing list from the point at which you left off.  Get a cursor from an existing list via the X-Files-Cursor-Next header.")
	cmdList.Flags().IntVarP(&paramsSsoStrategyList.PerPage, "per-page", "p", 0, "Number of records to show per page.  (Max: 10,000, 1,000 or less is recommended).")
	cmdList.Flags().IntVarP(&MaxPagesList, "max-pages", "m", 1, "When per-page is set max-pages limits the total number of pages requested")
	cmdList.Flags().StringVarP(&fieldsList, "fields", "", "", "comma separated list of field names to include in response")
	SsoStrategies.AddCommand(cmdList)
	var fieldsFind string
	paramsSsoStrategyFind := files_sdk.SsoStrategyFindParams{}
	cmdFind := &cobra.Command{
		Use: "find",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := sso_strategy.Find(paramsSsoStrategyFind)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			lib.JsonMarshal(result, fieldsFind)
		},
	}
	cmdFind.Flags().Int64VarP(&paramsSsoStrategyFind.Id, "id", "i", 0, "Sso Strategy ID.")

	cmdFind.Flags().StringVarP(&fieldsFind, "fields", "", "", "comma separated list of field names")
	SsoStrategies.AddCommand(cmdFind)
}
