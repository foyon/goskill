package main

import (
	"fmt"
	"time"
)

//标准格式话
var timeFormat = "2006-01-02 15:04:05"

func main() {
	// time.Time类型
	// 2021-01-23 21:00:06.165904 +0800 CST m=+0.000115433
	fmt.Printf("cur time is %s\n", time.Now())
	// time.Time
	//2021-01-23 13:08:25.835864 +0000 UTC
	fmt.Printf("utc: %s\n", time.Now().UTC())

	//int64 1970.1.1.00000 到现在的秒数
	//1611407091 len(10)
	fmt.Printf("timeStamp: %d\n", time.Now().Unix())
	//纳秒 len(19)
	fmt.Printf("timeStamp nano: %d\n", time.Now().UnixNano())
	//string 2021-01-23 21:19:17
	fmt.Printf("format string: %s\n", time.Now().Format(timeFormat))
	t1 := time.Now()
	t2 := time.Now()
	//float64
	fmt.Printf("time diff: %v\n", t2.Sub(t1).Seconds())
	//false
	fmt.Printf("time equal %v\n", t2.Equal(t1))
}
