package models

import (
	"fmt"
	"strings"

	"github.com/rhul-compsoc/compsoc-api-go/pkg/util"
)

type Date struct {
	Day   string `json:"day"`
	Month string `json:"month"`
	Year  string `json:"year"`
}

func (d *Date) String() string {
	return fmt.Sprintf("%s/%s/%s", d.Day, d.Month, d.Year)
}

func NewDate(s string) (Date, error) {
	d := strings.Split(s, "/")

	if len(d) != 3 {
		return Date{}, util.ErrBadDateConv
	}

	return Date{
		Day:   d[0],
		Month: d[1],
		Year:  d[2],
	}, nil
}
