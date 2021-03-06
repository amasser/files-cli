package cmd

import (
	"github.com/Files-com/files-cli/lib"
	"github.com/spf13/cobra"

	"fmt"
	"os"

	files_sdk "github.com/Files-com/files-sdk-go"
	"github.com/Files-com/files-sdk-go/session"
)

var (
	Sessions = &cobra.Command{
		Use:  "sessions [command]",
		Args: cobra.ExactArgs(1),
		Run:  func(cmd *cobra.Command, args []string) {},
	}
)

func SessionsInit() {
	var fieldsCreate string
	paramsSessionCreate := files_sdk.SessionCreateParams{}
	cmdCreate := &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := session.Create(paramsSessionCreate)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			lib.JsonMarshal(result, fieldsCreate)
		},
	}
	cmdCreate.Flags().StringVarP(&paramsSessionCreate.Username, "username", "u", "", "Username to sign in as")
	cmdCreate.Flags().StringVarP(&paramsSessionCreate.Password, "password", "a", "", "Password for sign in")
	cmdCreate.Flags().StringVarP(&paramsSessionCreate.Otp, "otp", "o", "", "If this user has a 2FA device, provide its OTP or code here.")
	cmdCreate.Flags().StringVarP(&paramsSessionCreate.PartialSessionId, "partial-session-id", "p", "", "Identifier for a partially-completed login")

	cmdCreate.Flags().StringVarP(&fieldsCreate, "fields", "", "", "comma separated list of field names")
	Sessions.AddCommand(cmdCreate)
	var fieldsDelete string
	paramsSessionDelete := files_sdk.SessionDeleteParams{}
	cmdDelete := &cobra.Command{
		Use: "delete",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := session.Delete(paramsSessionDelete)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			lib.JsonMarshal(result, fieldsDelete)
		},
	}
	cmdDelete.Flags().StringVarP(&paramsSessionDelete.Format, "format", "f", "", "")

	cmdDelete.Flags().StringVarP(&fieldsDelete, "fields", "", "", "comma separated list of field names")
	Sessions.AddCommand(cmdDelete)
}
