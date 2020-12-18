package utils

import (
	"math/rand"
	"fmt"
	"net"
	"strconv"
)

type ErrorType struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CustomHTTPError struct {
	Error ErrorType `json:"error"`
}

type CustomHTTPResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func GenerateError(errorCode int, msg string) (_error CustomHTTPError) {
	_error = CustomHTTPError{
		Error: ErrorType{
			Code:    errorCode,
			Message: msg,
		},
	}
	return
}

func ConvertToUint(str string) uint {
	uintNumber, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0
	}
	return uint(uintNumber)

}

func GetRandomInt(min int, max int) int {
	randomNum := rand.Intn(max-min+1) + min
	return randomNum
}

func GetRandomFloat() float32 {
	randomNum := 1.00 + rand.Float32()*(1000.00-1.00)
	return randomNum
}

func GetUniqueAllotedId(allottedIds map[int]struct{}) int {
	newId, iterator := 0, 0
	for iterator <= 100 {
		newAllottedId := GetRandomInt(1, 10000)
		if _, found := allottedIds[newAllottedId]; !found {
			newId = newAllottedId
			break;
		}
	}
	return newId
}

func IsPortFree(portNumber int) bool {
	if portNumber < 0 || portNumber > 65535 {
		return false
	}
	conn, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", portNumber))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
