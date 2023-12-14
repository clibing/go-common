package util

import (
	"fmt"
	"testing"

	"github.com/clibing/go-common/pkg/util"
)

type User struct {
	Name    string  `json:"name,omitempty" default:"admin" comment:"user name"`
	Pass    string  `json:"pass,omitempty" default:"123456" comment:"user password"`
	Age     int     `json:"age,omitempty" default:"18" comment:"age number"`
	Address string  `json:"address,omitempty" comment:"address value"`
	Value   float32 `json:"value,omitempty" default:"0.15"`
	Score   float64 `json:"score,omitempty" default:"0.151111"`
}

func TestJson(*testing.T) {
	user := &User{
		Name: "root",
	}
	b, _ := util.Marshal(user)
	fmt.Println(string(b))

	b, _ = util.MarshalIndent(user, "", "  ")
	fmt.Println(string(b))
}
