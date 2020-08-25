package gfUtils

import (
	"github.com/gogf/gf/os/gtime"
	"math"
	"math/rand"
	"strconv"
)

//生成流水号
//06123xxxxx
//sum 最少10位,sum 表示全部单号位数
func GeneSn(sum int) string {
	//年
	strs := gtime.Now().Format("YmdHis")

	//剩余随机数
	sum = sum - 14
	if sum < 1 {
		sum = 5
	}
	//0~9999999的随机数
	pow := math.Pow(10, float64(sum)) - 1
	//fmt.Println("sum=>", sum)
	//fmt.Println("pow=>", pow)

	rand.Seed(gtime.Now().Unix())
	result := strconv.Itoa(rand.Intn(int(pow)))

	//fmt.Println("result=>", result)

	//组合
	strs += result
	return strs
}

