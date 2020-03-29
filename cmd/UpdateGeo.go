package cmd

import (
	"fmt"
	"github.com/oucema001/ProxyScrapper/proxy"
	"github.com/spf13/cobra"
)

var updateGeo = &cobra.Command{
	Use : "update-geo",
	Short : "Updates the geo IPS Database location ",
	Long:"",
	Run:func(cmd *cobra.Command, args []string){
		fmt.Println("Updating the Geo database")
		proxy.UpdateGeoDb("WgqfJQ59iOUlCDvZ")
	},
}

func init(){
	rootCmd.AddCommand(updateGeo)
}
