/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unidoc/unipdf/pdf"
)

// passwdCmd represents the passwd command
var passwdCmd = &cobra.Command{
	Use:                   "passwd [FLAG]... INPUT_FILE NEW_OWNER_PASSWORD NEW_USER_PASSWORD",
	Short:                 "Change PDF password",
	Long:                  `A longer description that spans multiple lines and likely contains`,
	Example:               "this is the example",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		newOwnerPassword := args[1]
		newUserPassword := args[2]
		ownerPassword, _ := cmd.Flags().GetString("password")

		// Parse output file.
		outputFile, _ := cmd.Flags().GetString("output-file")
		if outputFile == "" {
			outputFile = inputFile
		}

		// Change input file password.
		err := pdf.Passwd(inputFile, outputFile, ownerPassword, newOwnerPassword, newUserPassword)
		if err != nil {
			fmt.Println("Could not change input file password")
			return
		}

		fmt.Println("Password successfully changed")
	},
}

func init() {
	rootCmd.AddCommand(passwdCmd)

	passwdCmd.Flags().StringP("output-file", "o", "", "Output file")
	passwdCmd.Flags().StringP("password", "p", "", "Input file password")
}
