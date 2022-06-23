/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"dipankardas011/accounts/database"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// creditCmd represents the credit command
var creditCmd = &cobra.Command{
	Use:   "credit",
	Short: "Create a credit transaction",
	Long: `
	This command creates a credit trnsaction for a particular user
	Usage: accounts credit <username> --amount=<amount> --narration=<narration>.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Username not specified")
		}
		username := args[0]
		user, err := database.FindOrCreateUser(username)
		if err != nil {
			log.Fatal(err)
		}
		user.Balance = user.Balance + creditAmount
		creditTransaction := database.Transaction{Amount: creditAmount, Type: "credit", Narration: creditNarration}
		user.Transactions = append(user.Transactions, creditTransaction)
		database.UpdateUser(user)
		fmt.Println("Transaction created successfully")
	},
}

var (
	creditNarration string
	creditAmount    int64
)

func init() {
	rootCmd.AddCommand(creditCmd)
	creditCmd.Flags().StringVarP(&creditNarration, "narration", "n", "", "Narration for this credit transaction")
	creditCmd.Flags().Int64VarP(&creditAmount, "amount", "a", 0, "Amount to be credited")
	creditCmd.MarkFlagRequired("narration")
	creditCmd.MarkFlagRequired("amount")
}
