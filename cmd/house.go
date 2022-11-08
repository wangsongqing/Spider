package cmd

import (
	"Spider/app/controllers"
	"Spider/helpers"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// 可以根据参数名称--传参
func init() {
	rootCmd.AddCommand(houseCmd)
	houseCmd.Flags().StringP("city", "c", "", "CityName")
}

// 运行项目命令 go run main.go house -c cd | all

var houseCmd = &cobra.Command{
	Use:     "house",
	Short:   "",
	Long:    ``,
	Example: "go run main.go house -c bj | all",
	Run: func(cmd *cobra.Command, args []string) {
		city, _ := cmd.Flags().GetString("city")
		if len(city) == 0 {
			color.Redln("-c 参数必填")
			return
		}

		if _, err := helpers.CityName[city]; err == false && city != "all" {
			color.Redln("-c 参数错误")
			return
		}

		house := controllers.House{}
		house.Start(city)
	},
}
