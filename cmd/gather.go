package cmd

import (
	"github.com/oucema001/ProxyScrapper/proxy"
	"github.com/spf13/cobra"
	"log"
	"os"
	"fmt"
)

var cmdGather = &cobra.Command{
	Use:"gather",
	Short :"gathers proxies from different providers without checking",
	Long:"",
	Run:gather,
}

func init(){
	rootCmd.AddCommand(cmdGather)
	cmdGather.Flags().IntP("limit","l",10,"number of proxies to get DEFAULT:10")
	cmdGather.Flags().StringP("format", "f", "", "output format choice JSON,CSV")
	cmdGather.Flags().StringP("output", "o", "", "")
}

func gather (cmd *cobra.Command, args []string){
	limit,_ := cmd.Flags().GetInt("limit")
	format, _ := cmd.Flags().GetString("format")
	output, _ := cmd.Flags().GetString("output")
	res := proxy.Gather(limit)
	if format == "JSON" {
		fmt.Println(proxy.JSONOutput(res))
	} else if format == "CSV" {
		if output =="" {
			dir ,err:= os.Getwd()
			if err!=nil{
				log.Println(err)
			}
			f := dir +string(os.PathSeparator)+ output
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