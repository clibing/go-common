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