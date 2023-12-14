package util

import (
	"fmt"
	"testing"

	"github.com/clibing/go-common/pkg/util"
)

func TestRandom(t *testing.T) {
	s1 := util.StringRandom(10)
	fmt.Printf("s1: %v\n", s1)

	s2 := util.NumberRandom(8)
	fmt.Printf("s2: %v\n", s2)

	s3 := util.Random(16)
	fmt.Printf("s3: %v\n", s3)
}