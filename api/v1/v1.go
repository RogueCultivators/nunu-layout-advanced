package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Total   int64       `json:"total"`
}

func HandleSuccess(ctx *gin.Context, data interface{}, total int64) {
	//if data == nil {
	//	data = map[string]interface{}{}
	//}
	resp := Response{Code: errorCodeMap[ErrSuccess], Message: ErrSuccess.Error(), Data: data, Total: total}
	if _, ok := errorCodeMap[ErrSuccess]; !ok {
		resp = Response{Code: 0, Message: "", Data: data, Total: total}
	}
	ctx.JSON(http.StatusOK, resp)
}

func HandleError(ctx *gin.Context, err error) {

	resp := Response{Code: errorCodeMap[err], Message: err.Error(), Data: nil}
	if _, ok := errorCodeMap[err]; !ok {
		resp = Response{Code: 500, Message: "unknown error", Data: nil}
	}
	ctx.JSON(http.StatusOK, resp)
}

type Error struct {
	Code    int
	Message string
}

var errorCodeMap = map[error]int{}

func newError(code int, msg string) error {
	err := errors.New(msg)
	errorCodeMap[err] = code
	return err
}
func (e Error) Error() string {
	return e.Message
}
