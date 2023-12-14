package util

import (
	"math"
	"time"
)

/**
 * Reddit 用于计算分值
 * 参考: https://www.ruanyifeng.com/blog/2012/03/ranking_algorithm_reddit.html
 */
func Reddit(createTime time.Time, ups, downs int) (value float64) {
	duration := createTime.Sub(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC))
	epoch_seconds := duration.Hours()/24*86400 + duration.Seconds() + float64(duration.Microseconds())/1000000
	score := ups - downs
	order := math.Log10(math.Max(math.Abs(float64(score)), 1))

	sign := 0
	if score > 0 {
		sign = 1
	} else if score < 0 {
		sign = -1
	}
	seconds := epoch_seconds - 1134028003
	value = math.Round(order + float64(sign)*seconds/45000)
	return
}
