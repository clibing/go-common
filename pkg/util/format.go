package util

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	SECOND = 1000
	MIUNTE = SECOND * 60
	HOUR   = MIUNTE * 60
	DAY    = HOUR * 24
	MONTH  = DAY * 30
	YEAR   = MONTH * 12
)

func BeautifyTime(unixMilli int64) string {
	if unixMilli < SECOND {
		return "<1s"
	} else if unixMilli < MIUNTE {
		return fmt.Sprintf("%.2fs", float64(unixMilli)/float64(SECOND))
	} else if unixMilli < HOUR {
		return fmt.Sprintf("%.2fm", float64(unixMilli)/float64(MIUNTE))
	} else if unixMilli < DAY {
		return fmt.Sprintf("%.2fh", float64(unixMilli)/float64(HOUR))
	} else if unixMilli < MONTH {
		return fmt.Sprintf("%.2fd", float64(unixMilli)/float64(DAY))
	} else if unixMilli < YEAR {
		return fmt.Sprintf("%.2fm", float64(unixMilli)/float64(MONTH))
	} else {
		return fmt.Sprintf("%.2fy", float64(unixMilli)/float64(YEAR))
	}
	return ""
}

/**
 * 字节的单位转换 保留两位小数
 *
 * 1024 -> 1.00KB
 */
func ByteFormat(input int64) string {
	if input < KB {
		return fmt.Sprintf("%.2fB", float64(input)/float64(B))
	} else if input < MB {
		return fmt.Sprintf("%.2fKB", float64(input)/float64(KB))
	} else if input < GB {
		return fmt.Sprintf("%.2fMB", float64(input)/float64(MB))
	} else if input < TB {
		return fmt.Sprintf("%.2fGB", float64(input)/float64(GB))
	} else if input < PB {
		return fmt.Sprintf("%.2fTB", float64(input)/float64(TB))
	} else {
		return fmt.Sprintf("%.2fPB", float64(input)/float64(PB))
	}
}

/**
 * 输入字符串:
 * 类型的大小 转 字节
 *
 */
func ReverseByteFormat(input string) int64 {
	re := regexp.MustCompile(`(?m)[0-9\.]+`)
	unit := string(re.ReplaceAll([]byte(input), []byte("")))
	num, _ := strconv.ParseFloat(strings.Replace(input, unit, "", 1), 32)
	unit = strings.ToUpper(unit)
	var byteNum float64 = 0
	switch unit {
	case "B":
		byteNum = num
	case "KB":
		byteNum = num * KB
	case "MB":
		byteNum = num * MB
	case "GB":
		byteNum = num * GB
	case "TB":
		byteNum = num * TB
	case "PB":
		byteNum = num * PB
	default:
		num = 0
		byteNum = 0
	}
	return int64(byteNum)
}
