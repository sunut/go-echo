package error

import (
	model "fake.com/develop/models"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type httpError struct {
	code     int `json:"code"`
	codeData int
	Key      string `json:"error"`
	Message  string `json:"message"`
}

func HTTPError(code int, codeData int, key string, msg string) *httpError {
	return &httpError{
		code:     code,
		codeData: codeData,
		Key:      key,
		Message:  msg,
	}
}

// Error makes it compatible with `error` interface.
func (e *httpError) Error() string {
	return e.Message
}

func HttpErrorHandler(err error, c echo.Context) {
	var (
		code     = http.StatusInternalServerError
		msg      string
		codeData = 1999
	)

	fmt.Println("err")
	fmt.Println(err)
	//fmt.Println(err.(*httpError))

	if he, ok := err.(*httpError); ok {
		//fmt.Println("00")
		code = he.code
		msg = he.Message
		codeData = he.codeData
		fmt.Println(codeData)
	} else {
		fmt.Println("01")

		msg = http.StatusText(code)
	}

	if !c.Response().Committed {
		fmt.Println("A1")
		if c.Request().Method == echo.HEAD {
			fmt.Println("A2")
			err := c.NoContent(code)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			fmt.Println("A3")
			err := c.JSON(code, &model.Status{model.StatusCode{codeData, msg}})
			fmt.Println(err)
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}
