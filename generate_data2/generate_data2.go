package generate_data2

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	"time"
)

var once = true

func randRange(left, right int) int {
	if once {
		rand.Seed(time.Now().Unix())
		once = false
	}
	return rand.Intn(right-left+1) + left
}

func randValue(left, right float64) float64 {
	if once {
		rand.Seed(time.Now().Unix())
		once = false
	}
	return rand.Float64()*(right-left) + left
}

func Generatedata() *cobra.Command {
	c := &cobra.Command{
		Use:   "generatedata2",
		Short: "Generatedata for Xinao_test2",
		Run: func(cmd *cobra.Command, args []string) {
			// var measurements = []string{"air", "electric", "water", "pressure"} // 根据以上指标，设计四个不同的业务域，对应不同的四个measurement
			var domain = []string{"NGS.GGS", "GGS.NGS"} // 业务域
			// MEAC（占数据量的70%） MFDA MFAE MRAT
			var deviceType = []string{"MEAC", "MFDA", "MFAE", "MRAT"}
			var equipmk = []string{"VALV", "METE", "WDbsq", "TMMT", "FLWM", "LNGT", "HBIERD", "YLbsq", "FLWM", "COPS", "CPFLDE", "GCA"}
			var file, _ = os.OpenFile("Xinao_test2.txt", os.O_WRONLY|os.O_CREATE, 0777)
			var startTime, endTime, step uint64
			startTime = 1648742400 //  2022.4.1 0:00
			endTime = 1651334400   //   2022.5.1 0:00
			step = 60              // per 60s
			
			fmt.Fprintf(file, "# DDL\n\n")
			fmt.Fprintf(file, "CREATE DATABASE Xinao_test\n\n")
			fmt.Fprintf(file, "# DML\n\n")
			fmt.Fprintf(file, "# CONTEXT-DATABASE: Xinao_test\n\n")
			
			for i := startTime; i <= endTime; i = i + step {
				timestamp := 1e9 * i
				for _, j := range domain {
					fmt.Fprintf(file, "measurement_1min,domain=%s device_type=\"%s\",device_id=\"%d\",system_code=\"%d\", equipmk=\"%s\", staid=\"%d_%d\"，value_1=%f,value_2=%f,value_3=%f,value_4=%f,value_5=%f,value_6=%f,value_7=%f,value_8=%f %d\n",
						j,
						deviceType[0],                     // device_type
						randRange(10000, 900000),          // device_id
						randRange(1000000000, 1000009000), // system_code 10位企业编号 共9000家企业
						equipmk[randRange(0, 11)],         // equipmk
						randRange(1000, 9000),             // staid %d_%d
						randRange(100000, 900000),         // sataid %d_%d
						randValue(0.0, 1000000.0),         // value_1
						randValue(0.0, 1000000.0),         // value_2
						randValue(0.0, 1000000.0),         // value_3
						randValue(0.0, 1000000.0),         // value_4
						randValue(0.0, 1000000.0),         // value_5
						randValue(0.0, 1000000.0),         // value_6
						randValue(0.0, 1000000.0),         // value_7
						randValue(0.0, 1000000.0),         // value_8
						timestamp,                         // 时间戳
					)
					fmt.Fprintf(file, "measurement_15min,domain=%s device_type=\"%s\",device_id=\"%d\",system_code=\"%d\", equipmk=\"%s\", staid=\"%d_%d\"，value_1=%f,value_2=%f,value_3=%f,value_4=%f,value_5=%f,value_6=%f,value_7=%f,value_8=%f %d\n",
						j,
						deviceType[0],                     // device_type
						randRange(10000, 900000),          // device_id
						randRange(1000000000, 1000009000), // system_code 10位企业编号 共9000家企业
						equipmk[randRange(0, 11)],         // equipmk
						randRange(1000, 9000),             // staid %d_%d
						randRange(100000, 900000),         // sataid %d_%d
						randValue(0.0, 1000000.0),         // value_1
						randValue(0.0, 1000000.0),         // value_2
						randValue(0.0, 1000000.0),         // value_3
						randValue(0.0, 1000000.0),         // value_4
						randValue(0.0, 1000000.0),         // value_5
						randValue(0.0, 1000000.0),         // value_6
						randValue(0.0, 1000000.0),         // value_7
						randValue(0.0, 1000000.0),         // value_8
						timestamp,                       // 时间戳
					)
					fmt.Fprintf(file, "measurement_1h,domain=%s，device_type=\"%s\",device_id=\"%d\",system_code=\"%d\", equipmk=\"%s\", staid=\"%d_%d\"，value_1=%f,value_2=%f,value_3=%f,value_4=%f,value_5=%f,value_6=%f,value_7=%f,value_8=%f %d\n",
						j,
						deviceType[randRange(1, 3)],
						randRange(10000, 900000),          // device_id
						randRange(1000000000, 1000009000), // system_code 10位企业编号 共9000家企业
						equipmk[randRange(0, 11)],
						randRange(1000, 9000),     // staid %d_%d
						randRange(100000, 900000), // sataid %d_%d
						randValue(0.0, 1000000.0), // value_1
						randValue(0.0, 1000000.0), // value_2
						randValue(0.0, 1000000.0), // value_3
						randValue(0.0, 1000000.0), // value_4
						randValue(0.0, 1000000.0), // value_5
						randValue(0.0, 1000000.0), // value_6
						randValue(0.0, 1000000.0), // value_7
						randValue(0.0, 1000000.0), // value_8
						timestamp,
					)
					fmt.Fprintf(file, "measurement_1d,domain=%s，device_type=\"%s\",device_id=\"%d\",system_code=\"%d\", equipmk=\"%s\", staid=\"%d_%d\"，value_1=%f,value_2=%f,value_3=%f,value_4=%f,value_5=%f,value_6=%f,value_7=%f,value_8=%f %d\n",
						j,
						deviceType[randRange(1, 3)],
						randRange(10000, 900000),          // device_id
						randRange(1000000000, 1000009000), // system_code 10位企业编号 共9000家企业
						equipmk[randRange(0, 11)],
						randRange(1000, 9000),     // staid %d_%d
						randRange(100000, 900000), // sataid %d_%d
						randValue(0.0, 1000000.0), // value_1
						randValue(0.0, 1000000.0), // value_2
						randValue(0.0, 1000000.0), // value_3
						randValue(0.0, 1000000.0), // value_4
						randValue(0.0, 1000000.0), // value_5
						randValue(0.0, 1000000.0), // value_6
						randValue(0.0, 1000000.0), // value_7
						randValue(0.0, 1000000.0), // value_8
						timestamp,
					)
					fmt.Fprintf(file, "measurement_1year,domain=%s，device_type=\"%s\",device_id=\"%d\",system_code=\"%d\", equipmk=\"%s\", staid=\"%d_%d\"，value_1=%f,value_2=%f,value_3=%f,value_4=%f,value_5=%f,value_6=%f,value_7=%f,value_8=%f %d\n",
						j,
						deviceType[randRange(1, 3)],
						randRange(10000, 900000),          // device_id
						randRange(1000000000, 1000009000), // system_code 10位企业编号 共9000家企业
						equipmk[randRange(0, 11)],
						randRange(1000, 9000),     // staid %d_%d
						randRange(100000, 900000), // sataid %d_%d
						randValue(0.0, 1000000.0), // value_1
						randValue(0.0, 1000000.0), // value_2
						randValue(0.0, 1000000.0), // value_3
						randValue(0.0, 1000000.0), // value_4
						randValue(0.0, 1000000.0), // value_5
						randValue(0.0, 1000000.0), // value_6
						randValue(0.0, 1000000.0), // value_7
						randValue(0.0, 1000000.0), // value_8
						timestamp,
					)
					fmt.Fprintf(file, "measurement_30d,domain=%s，device_type=\"%s\",device_id=\"%d\",system_code=\"%d\", equipmk=\"%s\", staid=\"%d_%d\"，value_1=%f,value_2=%f,value_3=%f,value_4=%f,value_5=%f,value_6=%f,value_7=%f,value_8=%f %d\n",
						j,
						deviceType[randRange(1, 3)],
						randRange(10000, 900000),          // device_id
						randRange(1000000000, 1000009000), // system_code 10位企业编号 共9000家企业
						equipmk[randRange(0, 11)],
						randRange(1000, 9000),     // staid %d_%d
						randRange(100000, 900000), // sataid %d_%d
						randValue(0.0, 1000000.0), // value_1
						randValue(0.0, 1000000.0), // value_2
						randValue(0.0, 1000000.0), // value_3
						randValue(0.0, 1000000.0), // value_4
						randValue(0.0, 1000000.0), // value_5
						randValue(0.0, 1000000.0), // value_6
						randValue(0.0, 1000000.0), // value_7
						randValue(0.0, 1000000.0), // value_8
						timestamp,
					)
				}
			}
			file.Close()
		},
	}
	return c
}
