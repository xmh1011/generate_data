package generate_data1

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
		Use:   "generatedata1",
		Short: "Generatedata for Xinao_test1",
		Run: func(cmd *cobra.Command, args []string) {
			//os.Stdout
			//var measurements = []string{"air", "electric", "water", "pressure"} // 根据以上指标，设计四个不同的业务域，对应不同的四个measurement
			var deviceType = []string{"MEAC", "MFDA", "MFAE", "MRAT"}
			var deviceID = []int{1,2,3,4,5,6,7,8}
			var equipmk = []string{"VALV", "METE", "WDbsq", "TMMT", "FLWM", "LNGT", "HBIERD", "YLbsq", "FLWM", "COPS", "CPFLDE", "GCA"}
			var file, _ = os.OpenFile("Xinao_test1.txt", os.O_WRONLY|os.O_CREATE, 0777)
			var startTime, endTime, step uint64
			// fmt.Fprintf(file, "# DDL\n\n")
			// fmt.Fprintf(file, "CREATE DATABASE Xinao_test\n\n")
			// fmt.Fprintf(file, "# DML\n\n")
			// fmt.Fprintf(file, "# CONTEXT-DATABASE: Xinao_test\n\n")
			startTime = 1643644800 //  2022.2.1 0:00
			endTime = 1651334400   //   2022.5.1 0:00
			step = 60              // per 60s
			for k := 1; k <= 1; k = k + 1 {//k为企业数
				for i := startTime; i <= endTime; i = i + step {
					timestamp := 1e9 * i
					timestamp_2 := 1e9 * (i - step/2)
					//for j := 0;j < 8; j = j + 1 {
					    j := randRange(0,7)
						fmt.Fprintf(file, "air,device_type=%s,device_id=%d,system_code=%d,equipmk=%s,staid=%d_%d value_1=%f,value_2=%f,value_3=%f,value_4=%f,value_5=%f,value_6=%f,value_7=%f,value_8=%f,value_9=%f,value_10=%f,value_11=%f,value_12=%f,value_13=%f,value_14=%f,value_15=%f,value_16=%f,value_17=%f,value_18=%f,value_19=%f,value_20=%f,value_21=%f,value_22=%f,value_23=%f,value_24=%f,value_25=%f,value_26=%f,value_27=%f,value_28=%f %d\n",
							deviceType[j/2],        // device_type
							deviceID[j],         // device_id
							k+1000000000, // system_code 10位企业编号 共9000家企业
							equipmk[randRange(0, 11)],         // equipmk
							randRange(1000, 1100),             // staid %d_%d
							randRange(100000, 101000),         // sataid %d_%d
							randValue(0.0, 1000000.0),         // value_1
							randValue(0.0, 1000000.0),         // value_2
							randValue(0.0, 1000000.0),         // value_3
							randValue(0.0, 1000000.0),         // value_4
							randValue(0.0, 1000000.0),         // value_5
							randValue(0.0, 1000000.0),         // value_6
							randValue(0.0, 1000000.0),         // value_7
							randValue(0.0, 1000000.0),         // value_8
							randValue(0.0, 1000000.0),         // value_9
							randValue(0.0, 1000000.0),         // value_10
							randValue(0.0, 1000000.0),         // value_11
							randValue(0.0, 1000000.0),         // value_12
							randValue(0.0, 1000000.0),         // value_13
							randValue(0.0, 1000000.0),         // value_14
							randValue(0.0, 1000000.0),         // value_15
							randValue(0.0, 1000000.0),         // value_16
							randValue(0.0, 1000000.0),         // value_17
							randValue(0.0, 1000000.0),         // value_18
							randValue(0.0, 1000000.0),         // value_19
							randValue(0.0, 1000000.0),         // value_20
							randValue(0.0, 1000000.0),         // value_21
							randValue(0.0, 1000000.0),         // value_22
							randValue(0.0, 1000000.0),         // value_23
							randValue(0.0, 1000000.0),         // value_24
							randValue(0.0, 1000000.0),         // value_25
							randValue(0.0, 1000000.0),         // value_26
							randValue(0.0, 1000000.0),         // value_27
							randValue(0.0, 1000000.0),         // value_28
							timestamp,                         // 时间戳
						)
						fmt.Fprintf(file, "air,device_type=%s,device_id=%d,system_code=%d,equipmk=%s,staid=%d_%d value_1=%f,value_2=%f,value_3=%f,value_4=%f,value_5=%f,value_6=%f,value_7=%f,value_8=%f,value_9=%f,value_10=%f,value_11=%f,value_12=%f,value_13=%f,value_14=%f,value_15=%f,value_16=%f,value_17=%f,value_18=%f,value_19=%f,value_20=%f,value_21=%f,value_22=%f,value_23=%f,value_24=%f,value_25=%f,value_26=%f,value_27=%f,value_28=%f %d\n",
							deviceType[j/2],        // device_type
							deviceID[j],         // device_id
							k+1000000000, // system_code 10位企业编号 共9000家企业
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
							randValue(0.0, 1000000.0),         // value_9
							randValue(0.0, 1000000.0),         // value_10
							randValue(0.0, 1000000.0),         // value_11
							randValue(0.0, 1000000.0),         // value_12
							randValue(0.0, 1000000.0),         // value_13
							randValue(0.0, 1000000.0),         // value_14
							randValue(0.0, 1000000.0),         // value_15
							randValue(0.0, 1000000.0),         // value_16
							randValue(0.0, 1000000.0),         // value_17
							randValue(0.0, 1000000.0),         // value_18
							randValue(0.0, 1000000.0),         // value_19
							randValue(0.0, 1000000.0),         // value_20
							randValue(0.0, 1000000.0),         // value_21
							randValue(0.0, 1000000.0),         // value_22
							randValue(0.0, 1000000.0),         // value_23
							randValue(0.0, 1000000.0),         // value_24
							randValue(0.0, 1000000.0),         // value_25
							randValue(0.0, 1000000.0),         // value_26
							randValue(0.0, 1000000.0),         // value_27
							randValue(0.0, 1000000.0),         // value_28
							timestamp_2,                       // 时间戳
						)
						fmt.Fprintf(file, "water,device_type=%s,device_id=%d,system_code=%d,equipmk=%s,staid=%d_%d value_1=%f,value_2=%f,value_3=%f,value_4=%f,value_5=%f,value_6=%f,value_7=%f,value_8=%f %d\n",
							deviceType[j/2],        // device_type
							deviceID[j],         // device_id
							k+1000000000, // system_code 10位企业编号 共9000家企业
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
						fmt.Fprintf(file, "pressure,device_type=%s,device_id=%d,system_code=%d,equipmk=%s,staid=%d_%d value_1=%f,value_2=%f,value_3=%f,value_4=%f,value_5=%f,value_6=%f,value_7=%f,value_8=%f %d\n",
							deviceType[j/2],        // device_type
							deviceID[j],         // device_id
							k+1000000000, // system_code 10位企业编号 共9000家企业
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
						fmt.Fprintf(file, "electric,device_type=%s,device_id=%d,system_code=%d,equipmk=%s,staid=%d_%d value_1=%f,value_2=%f,value_3=%f,value_4=%f,value_5=%f,value_6=%f,value_7=%f,value_8=%f %d\n",
							deviceType[j/2],        // device_type
							deviceID[j],         // device_id
							k+1000000000, // system_code 10位企业编号 共9000家企业
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
						// fmt.Fprintf(file, "%s,device_type=%s,device_id=%d,system_code=%d,equipmk=%s,staid=%d_%d value_1=%f,value_2=%f,value_3=%f,value_4=%f,value_5=%f,value_6=%f,value_7=%f,value_8=%f %d\n",
						// 	j,
						// 	deviceType[2],        // device_type
						// 	deviceID[5],         // device_id
						// 	k+1000000000, // system_code 10位企业编号 共9000家企业
						// 	equipmk[randRange(0, 11)],
						// 	randRange(1000, 9000),     // staid %d_%d
						// 	randRange(100000, 900000), // sataid %d_%d
						// 	randValue(0.0, 1000000.0), // value_1
						// 	randValue(0.0, 1000000.0), // value_2
						// 	randValue(0.0, 1000000.0), // value_3
						// 	randValue(0.0, 1000000.0), // value_4
						// 	randValue(0.0, 1000000.0), // value_5
						// 	randValue(0.0, 1000000.0), // value_6
						// 	randValue(0.0, 1000000.0), // value_7
						// 	randValue(0.0, 1000000.0), // value_8
						// 	timestamp,
						// )
						// fmt.Fprintf(file, "%s,device_type=%s,device_id=%d,system_code=%d,equipmk=%s,staid=%d_%d value_1=%f,value_2=%f,value_3=%f,value_4=%f,value_5=%f,value_6=%f,value_7=%f,value_8=%f %d\n",
						// 	j,
						// 	deviceType[3],        // device_type
						// 	deviceID[6],         // device_id
						// 	k+1000000000, // system_code 10位企业编号 共9000家企业
						// 	equipmk[randRange(0, 11)],
						// 	randRange(1000, 9000),     // staid %d_%d
						// 	randRange(100000, 900000), // sataid %d_%d
						// 	randValue(0.0, 1000000.0), // value_1
						// 	randValue(0.0, 1000000.0), // value_2
						// 	randValue(0.0, 1000000.0), // value_3
						// 	randValue(0.0, 1000000.0), // value_4
						// 	randValue(0.0, 1000000.0), // value_5
						// 	randValue(0.0, 1000000.0), // value_6
						// 	randValue(0.0, 1000000.0), // value_7
						// 	randValue(0.0, 1000000.0), // value_8
						// 	timestamp,
						// )
						// fmt.Fprintf(file, "%s,device_type=%s,device_id=%d,system_code=%d,equipmk=%s,staid=%d_%d value_1=%f,value_2=%f,value_3=%f,value_4=%f,value_5=%f,value_6=%f,value_7=%f,value_8=%f %d\n",
						// 	j,
						// 	deviceType[3],        // device_type
						// 	deviceID[7],         // device_id
						// 	k+1000000000, // system_code 10位企业编号 共9000家企业
						// 	equipmk[randRange(0, 11)],
						// 	randRange(1000, 9000),     // staid %d_%d
						// 	randRange(100000, 900000), // sataid %d_%d
						// 	randValue(0.0, 1000000.0), // value_1
						// 	randValue(0.0, 1000000.0), // value_2
						// 	randValue(0.0, 1000000.0), // value_3
						// 	randValue(0.0, 1000000.0), // value_4
						// 	randValue(0.0, 1000000.0), // value_5
						// 	randValue(0.0, 1000000.0), // value_6
						// 	randValue(0.0, 1000000.0), // value_7
						// 	randValue(0.0, 1000000.0), // value_8
						// 	timestamp,
						// )
					}
				 }
				//}
			file.Close()
		},
	}
	return c
}
