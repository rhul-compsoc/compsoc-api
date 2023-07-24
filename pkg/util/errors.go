package util

import (
	"errors"
	"log"
)

var (
	ErrTaskConv             = errors.New("error converting interface to task")
	ErrQueryResp            = errors.New("error getting query response")
	ErrUndefinedRouteMethod = errors.New("error, this route method is undefined")
	ErrFileNotExist         = errors.New("error, tried writting to a file that does not exist")
)

func ErrOut(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrFat(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func ErrMsg(err error, msg string) {
	if err != nil {
		log.Println(err, msg)
	}
}

func ErrLog(err error) {
	if err != nil {
		log.Println(err)
	}
}
