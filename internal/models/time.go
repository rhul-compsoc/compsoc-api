package models

import (
	"fmt"
	"strings"

	"github.com/rhul-compsoc/compsoc-api-go/pkg/util"
)

type Time struct {
	Minute string `json:"minute"`
	Hour   string `json:"hour"`
}

func (t *Time) String() string {
	return fmt.Sprintf("%s:%s", t.Minute, t.Hour)
}

func NewTime(s string) (Time, error) {
	t := strings.Split(s, ":")

	if len(t) != 2 {
		return Time{}, util.ErrBadTimeConv
	}

	return Time{
		Minute: t[0],
		Hour:   t[1],
	}, nil
}
