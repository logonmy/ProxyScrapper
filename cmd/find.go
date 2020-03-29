package cmd

import (
	"fmt"
	"github.com/oucema001/ProxyScrapper/proxy"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var cmdFind = &cobra.Command{
	Use:   "find",
	Short: "Gathers and checks proxies from various providers",
	Long:  "",
	Run:   find,
}

func init() {
	rootCmd.AddCommand(cmdFind)
	cmdFind.Flags().IntP("limit", "l", 10, "number of proxies returned as a result")
	cmdFind.Flags().StringSliceP("types", "t", nil, "types of proxies HTTP,SOCKS5")
	cmdFind.Flags().StringSliceP("countries", "c", nil, "ISO Codes of countries for proxies")
	cmdFind.Flags().BoolP("anonymous", "a", false, "")
	cmdFind.Flags().StringP("format", "f", "", "output format choice JSON,CSV")
	cmdFind.Flags().StringP("output", "o", "", "")
}

func find(cmd *cobra.Command, args []string) {
	limit, _ := cmd.Flags().GetInt("limit")
	types, _ := cmd.Flags().GetStringSlice("types")
	countries, _ := cmd.Flags().GetStringSlice("countries")
	anonymous, _ := cmd.Flags().GetBool("anonymous")
	format, _ := cmd.Flags().GetString("format")
	output, _ := cmd.Flags().GetString("output")
	res := proxy.Find(limit, types, countries, anonymous)

	if format == "JSON" {
		fmt.Println(proxy.JSONOutput(res))
	} else if format == "CSV" {
		if output =="" {
			dir ,err:= os.Getwd()
			if err!=nil{
				log.Println(err)
			}
			f := dir +"\\"+ output
			file ,err := os.Create(f)
			proxy.CSVOutput(res,file)
		} else{
			file ,err := os.Create(output)
			if err!=nil{
				log.Println(err)
			}
			proxy.CSVOutput(res,file)
		}
	}

}


