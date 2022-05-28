/*

Copyright © 2022 Batuhan ÇİTOĞLU <batuhancitoglu@gmail.com>

*/
package cmd

import (
	"EXCHANGETURKEY/parser"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "fiyat",
	Short: "Dolar, Euro ve Sterlin Fiyat Öğrenme",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		isDolar, _ := cmd.Flags().GetBool("dolar")
		isEuro, _ := cmd.Flags().GetBool("euro")
		isSterlin, _ := cmd.Flags().GetBool("sterlin")

		if isDolar {
			getDolar()
		} else if isEuro {
			getEuro()
		} else if isSterlin {
			getSterlin()
		} else {
			getDolar()
		}
	},
}

func getDolar() {
	var exchange, exchangeFlt string
	exchange = parser.ParseWeb("DOLAR")
	exchangeFlt = strings.Replace(exchange, ",", ".", 1)
	current_time := time.Now()
	fmt.Printf("Dolar [%s] : %s₺", current_time.Format("2006-01-02 15:04:05"), exchangeFlt)
	fmt.Println()
}
func getEuro() {
	var exchange, exchangeFlt string
	exchange = parser.ParseWeb("EURO")
	exchangeFlt = strings.Replace(exchange, ",", ".", 1)

	current_time := time.Now()
	fmt.Printf("Euro [%s] : %s€", current_time.Format("2006-01-02 15:04:05"), exchangeFlt)
	fmt.Println()
}

func getSterlin() {
	var exchange, exchangeFlt string
	exchange = parser.ParseWeb("STERLİN")
	exchangeFlt = strings.Replace(exchange, ",", ".", 1)

	current_time := time.Now()
	fmt.Printf("Sterlin [%s] : %s£", current_time.Format("2006-01-02 15:04:05"), exchangeFlt)
	fmt.Println()
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().BoolP("dolar", "d", false, "Get Dolar Currency")
	getCmd.Flags().BoolP("euro", "e", false, "Get Euro Currency")
	getCmd.Flags().BoolP("sterlin", "s", false, "Get Sterlin Currency")
}
