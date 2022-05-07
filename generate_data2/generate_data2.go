package generate_data2

import (
	"generate_data/pkg"
	"github.com/spf13/cobra"
)

func Generatedata() *cobra.Command {
	c := &cobra.Command{
		Use:   "generatedata1",
		Short: "Generatedata for Xinao_test1",
		Run: func(cmd *cobra.Command, args []string) {
			var startTime, endTime, step, number, flag int64
			startTime = 1577808000 //  2020.1.1 0:00
			endTime = 1640966400   //   2022.1.1 0:00
			step = 60              // per 60s
			flag = 1
			for i := startTime; i <= endTime; i = i + step {
				number = 208
				pkg.OutputData(step, i, number) // 1分钟指标有10w个，一分钟一条
				flag = flag + 1
				if flag%15 == 0 {
					number = 146
					pkg.OutputData(step, i, number) // 15分钟指标有7w个，每15分钟一条
				}
				if flag%60 == 0 {
					number = 417
					pkg.OutputData(step, i, number) // 小时指标有20w条记录，每小时一条
				}
				if flag%1440 == 0 {
					number = 834
					pkg.OutputData(step, i, number) // 天级指标有40w个，每天一条
				}
				if flag%43200 == 0 {
					number = 1063
					pkg.OutputData(step, i, number) // 月级指标有51w个，每月一条
				}
				if flag%15768000 == 0 {
					number = 834
					pkg.OutputData(step, i, number) // 年级指标有40w个，每年一条
				}
			}
		},
	}
	return c
}
