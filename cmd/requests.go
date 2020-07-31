package cmd

import "github.com/spf13/cobra"
import (
	"fmt"
	"github.com/Files-com/files-cli/lib"
	files_sdk "github.com/Files-com/files-sdk-go"
	"github.com/Files-com/files-sdk-go/request"
	"os"
)

var (
	_ = files_sdk.Config{}
	_ = request.Client{}
	_ = lib.OnlyFields
	_ = fmt.Println
	_ = os.Exit
)

var (
	Requests = &cobra.Command{
		Use:  "requests [command]",
		Args: cobra.ExactArgs(1),
		Run:  func(cmd *cobra.Command, args []string) {},
	}
)

func RequestsInit() {
	var fieldsList string
	paramsRequestList := files_sdk.RequestListParams{}
	var MaxPagesList int
	cmdList := &cobra.Command{
		Use:   "list",
		Short: "list",
		Long:  `list`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			params := paramsRequestList
			params.MaxPages = MaxPagesList
			it := request.List(params)

			lib.JsonMarshalIter(it, fieldsList)
		},
	}
	cmdList.Flags().IntVarP(&paramsRequestList.Page, "page", "p", 0, "List Requests")
	cmdList.Flags().IntVarP(&paramsRequestList.PerPage, "per-page", "e", 0, "List Requests")
	cmdList.Flags().StringVarP(&paramsRequestList.Action, "action", "a", "", "List Requests")
	cmdList.Flags().StringVarP(&paramsRequestList.Cursor, "cursor", "c", "", "List Requests")
	cmdList.Flags().StringVarP(&paramsRequestList.Path, "path", "t", "", "List Requests")
	cmdList.Flags().IntVarP(&MaxPagesList, "max-pages", "m", 1, "When per-page is set max-pages limits the total number of pages requested")
	cmdList.Flags().StringVarP(&fieldsList, "fields", "f", "", "comma separated list of field names to include in response")
	Requests.AddCommand(cmdList)
	var fieldsFindFolder string
	paramsRequestFindFolder := files_sdk.RequestFindFolderParams{}
	cmdFindFolder := &cobra.Command{
		Use: "find-folder",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := request.FindFolder(paramsRequestFindFolder)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			lib.JsonMarshal(result, fieldsFindFolder)
		},
	}
	cmdFindFolder.Flags().IntVarP(&paramsRequestFindFolder.Page, "page", "p", 0, "List Requests")
	cmdFindFolder.Flags().IntVarP(&paramsRequestFindFolder.PerPage, "per-page", "e", 0, "List Requests")
	cmdFindFolder.Flags().StringVarP(&paramsRequestFindFolder.Action, "action", "a", "", "List Requests")
	cmdFindFolder.Flags().StringVarP(&paramsRequestFindFolder.Cursor, "cursor", "c", "", "List Requests")
	cmdFindFolder.Flags().StringVarP(&paramsRequestFindFolder.Path, "path", "t", "", "List Requests")
	cmdFindFolder.Flags().StringVarP(&fieldsFindFolder, "fields", "f", "", "comma separated list of field names")
	Requests.AddCommand(cmdFindFolder)
	var fieldsCreate string
	paramsRequestCreate := files_sdk.RequestCreateParams{}
	cmdCreate := &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := request.Create(paramsRequestCreate)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			lib.JsonMarshal(result, fieldsCreate)
		},
	}
	cmdCreate.Flags().StringVarP(&paramsRequestCreate.Path, "path", "p", "", "Create Request")
	cmdCreate.Flags().StringVarP(&paramsRequestCreate.Destination, "destination", "d", "", "Create Request")
	cmdCreate.Flags().StringVarP(&paramsRequestCreate.UserIds, "user-ids", "u", "", "Create Request")
	cmdCreate.Flags().StringVarP(&paramsRequestCreate.GroupIds, "group-ids", "g", "", "Create Request")
	cmdCreate.Flags().StringVarP(&fieldsCreate, "fields", "f", "", "comma separated list of field names")
	Requests.AddCommand(cmdCreate)
	var fieldsDelete string
	paramsRequestDelete := files_sdk.RequestDeleteParams{}
	cmdDelete := &cobra.Command{
		Use: "delete",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := request.Delete(paramsRequestDelete)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			lib.JsonMarshal(result, fieldsDelete)
		},
	}
	cmdDelete.Flags().StringVarP(&fieldsDelete, "fields", "f", "", "comma separated list of field names")
	Requests.AddCommand(cmdDelete)
}
