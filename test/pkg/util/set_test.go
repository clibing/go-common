package util

import (
	// "fmt"
	"fmt"
	"testing"
	"time"

	"github.com/clibing/go-common/pkg/util"
)

type Auth[T any] interface {
	Login(name, secret string) (t *T, err error)
}

type U struct {
	Name string `json:"name,omitempty" default:"admin" comment:"user name"`
	Pass string `json:"pass,omitempty" default:"123456" comment:"user password"`
}

func (u *U) Login(name, secret string) (token *Token, err error) {
	td, _ := time.ParseDuration("30d")
	token = &Token{
		Value:    "62b21148-4066-e9e8-ffac-7a3afcf9e9d4",
		ExpireAt: time.Now().Add(td),
	}
	return
}

type Token struct {
	Value    string
	ExpireAt time.Time
}

func TestAuth(t *testing.T) {
	u := &U{
		Name: "admin",
		Pass: "123456",
	}
	v, e := u.Login("admin", "123456")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(v.Value, v.ExpireAt)
}

func TestSet(t *testing.T) {
	v := util.NewSet[string](16)
	v.Add("admin")

}
