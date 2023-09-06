package server

import "github.com/gin-gonic/gin"

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
