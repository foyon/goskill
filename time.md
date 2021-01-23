# [Golang 时间日期处理](http://docscn.studygolang.com/pkg/time/)

**在 golang 中使用`time`标准包对时间进行处理**

**time 包下的`Time`类型用来表示时间**

```
var timeFormat = "2006-01-02 15:04:05"
```



```
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
```

## 获取时间

- 使用用 time.Now()获取当前时间 (time.Time)

- 使用用 time.Now().Unix()获取当前时间戳（1970.1.1.00000 到现在的秒数）(int64)

- // 使用time.Now().Format()格式化, 2006-01-02 15:04:05 是go语言诞生时间，格式化必须写这个时间，格式自定义

- 

  ``` 
  
  ```

```go
Copypackage main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Println(now)

	// 分别获取年月日
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
    second := now.Second()

	// 格式化输出
	// 02d表示数字不够两位前面补0
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	// 获取时间戳
	fmt.Println(now.Unix())		// 获取以秒为单位的时间戳
	fmt.Println(now.UnixNano())	// 获取以以纳秒为单位的时间戳
}
```

## 时间戳转换

**时间戳转换为 Time 类型**

```go
Copypackage main

import (
	"fmt"
	"time"
)

func main() {
	timestamp := time.Now().Unix()

	// 时间戳转time, 第二个参数是对纳秒数的精确为长度
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)

	// 分别获取年月日
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	// 格式化输出
	// 02d表示数字不够两位前面补0
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
```

## golang 源码中的常量定义

```go
Copyconst (
   // 纳秒
   Nanosecond  Duration = 1
   // 微秒
   Microsecond          = 1000 * Nanosecond
   // 毫秒
   Millisecond          = 1000 * Microsecond
   // 秒
   Second               = 1000 * Millisecond
   // 分
   Minute               = 60 * Second
   // 时
   Hour                 = 60 * Minute
)
```

### 时间格式化

在其他语言中我们格式化时间是写`YYYY-MM-dd hh:mm:ss`这种表达式来表示，但是在 golang 中这边表达式是不识别的，我们必须要写成`2006-01-02 15:04:05`这个样子的表达式才能被成功的格式化。这个又是什么意思呢？这个时间是 golang 语言诞生的时间，我们必须要严格按照这个时间格式去写，否则就会格式化不成功，头大....。

```go
Copypackage main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	// 使用time.Now().Format()格式化, 2006-01-02 15:04:05 是go语言诞生时间，格式化必须写这个时间，格式自定义
	nowStr := now.Format("2006-01-02 15:04:05")
	fmt.Println(nowStr)
}
```