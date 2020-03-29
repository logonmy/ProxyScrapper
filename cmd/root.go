package cmd

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:"proxybroker",
	Short:"ProxyBroker is a tool to gather and check proxies",
	Long:"",
}

func Execute() error{
	return rootCmd.Execute()
}

func init(){
	cobra.OnInitialize(initConfig)
	//rootCmd.PersistentFlags().StringP("find","f","","gather and check proxies according to filters")
	//rootCmd.PersistentFlags().StringP("gather","g","","gather proxies")
}

func initConfig(){
	home, err := homedir.Dir()
	if err !=nil{
		log.Println(err)
	}
	//fmt.Println(home)

	f,err := os.Create(home+string(os.PathSeparator)+"proxyBroker.log")
	if err !=nil{
		log.Println(err)
	}
	log.SetOutput(f)
}