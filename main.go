package main

import (
	"fmt"
	"generate_data/generate_data1"
	"generate_data/generate_data2"
	"github.com/spf13/cobra"
)

func main() {
	mainCmd := GetCommand()
	
	Generatedata1Cmd := generate_data1.Generatedata()
	mainCmd.AddCommand(Generatedata1Cmd)
	
	Generatedata2Cmd := generate_data2.Generatedata()
	mainCmd.AddCommand(Generatedata2Cmd)
	
	if err := mainCmd.Execute(); err != nil {
		fmt.Printf("Error : %+v\n", err)
	}
}
func GetCommand() *cobra.Command {
	c := &cobra.Command{
		Use:  "generate-data",
		Long: "generate-data testdata",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   true,
			DisableNoDescFlag:   true,
			DisableDescriptions: true},
	}
	return c
}
