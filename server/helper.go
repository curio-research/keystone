package server

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type DownloadWorldRequest struct {
	Components []string `json:"components"`
}

func DecodeRequestBody(c *gin.Context, req any) any {
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	return req
}

func CreateBasicResponseObject(requestUuid string) interface{} {
	returnObj := struct {
		Uuid   string `json:"uuid"`
		Status string `json:"status"`
	}{
		Uuid:   requestUuid,
		Status: "success",
	}
	return returnObj
}

func CreateBasicResponseObjectWithData(requestUuid string, data any) interface{} {
	returnObj := struct {
		Uuid   string `json:"uuid"`
		Status string `json:"status"`
		Data   any    `json:"data"`
	}{
		Uuid:   requestUuid,
		Status: "success",
		Data:   data,
	}
	return returnObj
}

func CreateResponseWithError(requestUuid string, errorMessage string) interface{} {
	returnObj := struct {
		Uuid    string `json:"uuid"`
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Uuid:    requestUuid,
		Status:  "fail",
		Message: errorMessage,
	}
	return returnObj
}

func GetSuitFromPokerCard(pokerCard string) string {
	parts := strings.Split(pokerCard, "-")
	return parts[0]
}

func GetValueFromPokerCard(pokerCard string) string {
	parts := strings.Split(pokerCard, "-")
	return parts[1]
}

func FormatMilliseconds(ms int) string {
	duration := time.Millisecond * time.Duration(ms)

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	milliseconds := ms % 1000

	return fmt.Sprintf("%02d:%02d:%02d.%03d", hours, minutes, seconds, milliseconds)
}

func ReverseArray(arr []any) []any {
	reversed := make([]any, len(arr))
	for i, j := len(arr)-1, 0; i >= 0; i, j = i-1, j+1 {
		reversed[j] = arr[i]
	}
	return reversed
}

func HashNumbers(num1, num2 int) int {
	// a prime number used in the hashing process
	prime := 31

	// combine the two numbers using bitwise operations
	hashValue := num1*prime + num2

	return hashValue
}
