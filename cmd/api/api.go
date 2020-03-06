package main

import (
	"fmt"
	"github.com/YonghoChoi/slack-interact-api/model/packet"
	"github.com/labstack/echo"
	"net/http"
)

func Approve(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	var u map[string]interface{}
	if err := c.Bind(&u); err != nil {
		resp.Code = "500"
		resp.Message = "invalid data"
		fmt.Println(err.Error())
		return nil
	}

	fmt.Println(u)
	resp.Data = u
	return nil
}
