package generate_data1

import (
	"generate_data/pkg"
	"github.com/spf13/cobra"
)

func Generatedata() *cobra.Command {
	c := &cobra.Command{
		Use:   "generatedata1",
		Short: "Generatedata for Xinao_test1",
		Run: func(cmd *cobra.Command, args []string) {
			var startTime, endTime, step, number int64 // number 为企业数
			startTime = 1643644800                     //  2022.2.1 0:00
			endTime = 164586400                        //   2022.6.25 0:00
			step = 60                                  // per 60s
			number = 90
			for i := startTime; i <= endTime; i = i + step {
				pkg.OutputData(pkg.Measurement[pkg.RandRange(0, 30)], step, i, number)
			}
		},
	}
	return c
}
