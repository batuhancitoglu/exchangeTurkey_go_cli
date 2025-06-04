package cmd

import (
    "EXCHANGETURKEY/weather"
    "fmt"

    "github.com/spf13/cobra"
)

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
    Use:   "hava",
    Short: "Hava durumunu gosterir",
    Run: func(cmd *cobra.Command, args []string) {
        country, _ := cmd.Flags().GetString("country")
        city, _ := cmd.Flags().GetString("city")
        if city == "" {
            fmt.Println("City is required")
            return
        }
        info := weather.GetWeather(country, city)
        fmt.Println(info)
    },
}

func init() {
    rootCmd.AddCommand(weatherCmd)
    weatherCmd.Flags().StringP("country", "c", "", "Ulke adi")
    weatherCmd.Flags().StringP("city", "i", "", "Sehir adi")
}

