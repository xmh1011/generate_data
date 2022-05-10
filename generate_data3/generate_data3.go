package generate_data3

import (
	"generate_data/pkg"
	"github.com/spf13/cobra"
	"io"
	"os"
	"time"
)

type Options struct {
	Stderr io.Writer
	Stdout io.Writer

	startTime int64
	endTime   int64
}

func NewOptions() *Options {
	return &Options{
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}
}

var opt = NewOptions()

func Generatedata() *cobra.Command {
	c := &cobra.Command{
		Use:   "generatedata3",
		Short: "Generatedata for Xinao_test3",
		Run: func(cmd *cobra.Command, args []string) {
			var startTime, endTime, step, number int64
			startTime = opt.startTime //  2020.1.1 0:00
			endTime = opt.endTime     //   2022.1.1 0:00
			step = 60                 // per 60s
			// flag = 1
			for i := startTime; i <= endTime; i = i + step {
				number = 208                                        // 企业数
				pkg.OutputData(pkg.Measurement[0], step, i, number) // 1分钟个，一分钟一条,measurement为air
				pkg.OutputData(pkg.Measurement[1], step, i, 10)     // 15分钟指标有7w个，每15分钟一条,measurement water
				pkg.OutputData(pkg.Measurement[2], step, i, 7)      // 小时指标有20w条记录，每小时一条 ,measurement pressure
				pkg.OutputData(pkg.Measurement[3], step, i, 1)      // 天级指标有40w个，每天一条 measurement electric
				pkg.OutputData(pkg.Measurement[3], step, i, 1)      // 月级指标有51w个，每月一条  measurement electric
				pkg.OutputData(pkg.Measurement[3], step, i, 1)      // 年级指标有40w个，每年一条  measurement electric
			}
		},
	}
	c.PersistentFlags().Int64Var(&opt.startTime, "startTime", time.Now().Unix(), "Set start time")
	c.PersistentFlags().Int64Var(&opt.endTime, "endTime", time.Now().Unix(), "Set end time")
	return c
}
