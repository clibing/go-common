package util

import (
	"fmt"
	"testing"

	"github.com/clibing/go-common/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	k1 := int64(1024)
	k1_b1 := int64(1025)
	k2 := int64(1024 * 2)
	k2_b100 := int64(k2 + 100)
	m1 := int64(1024 * 1024)
	m1_k100 := int64(m1 + 1024*100)
	g1 := int64(1024 * 1024 * 1024)
	g1_m10 := g1 + int64(m1*10)
	t1 := g1 * 1024

	assert.Equal(t, util.ByteFormat(k1), "1.00KB", "")
	assert.Equal(t, util.ByteFormat(k1_b1), "1.00KB", "")
	assert.Equal(t, util.ByteFormat(k2), "2.00KB", "")
	assert.Equal(t, util.ByteFormat(int64(k2_b100)), "2.10KB", "")
	assert.Equal(t, util.ByteFormat(m1), "1.00MB", "")
	assert.Equal(t, util.ByteFormat(m1_k100), "1.10MB", "")
	assert.Equal(t, util.ByteFormat(g1), "1.00GB", "")
	assert.Equal(t, util.ByteFormat(g1_m10), "1.01GB", "")
	assert.Equal(t, util.ByteFormat(t1), "1.00TB", "")

	i := util.ReverseByteFormat("1.00TB")
	fmt.Println(i)
	// assert.Equal(t, , t1, , "")

}

func TestBeautifyTime(t *testing.T) {
	fmt.Printf("900ms(<1s)->%s\n", util.BeautifyTime(900))
	fmt.Printf("1500ms(1.50s)->%s\n", util.BeautifyTime(1500))
	fmt.Printf("180s(3.00)->%s\n", util.BeautifyTime(1000*180))
	fmt.Printf("30m(30.00)->%s\n", util.BeautifyTime(1000*60*30))
	fmt.Printf("1.5h(1.50h)->%s\n", util.BeautifyTime(1000*60*60*1.5))
	fmt.Printf("1.5d(1.50d)->%s\n", util.BeautifyTime(1000*60*60*1.5*24))
	fmt.Printf("1.5m(1.50m)->%s\n", util.BeautifyTime(1000*60*60*1.5*24*30))
	fmt.Printf("1.5y(1.50y)->%s\n", util.BeautifyTime(1000*60*60*1.5*24*30*12))
}
