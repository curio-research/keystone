package server

import "github.com/gin-gonic/gin"

func DecodeRequestBody[T any](c *gin.Context) (T, error) {
	var res KeystoneTx[T]
	if err := c.ShouldBindJSON(&res); err != nil {
		return res.Data, err
	}

	return res.Data, nil
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
