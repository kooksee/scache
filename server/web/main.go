package web

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/kooksee/scache/types"
)

func Run(port string) error {
	e := initMiddles(echo.New())
	e.POST("/", func(c echo.Context) error {
		req := &types.RPCRequest{}
		if err := c.Bind(req); err != nil {
			return types.ErrWithMsg("Web RPC Request Decode Error", err)
		}

		switch req.Method {
		default:
			return c.JSON(
				http.StatusNotFound,
				types.RPCResponse{
					Code: "404",
					Msg:  types.F("%s方法不存在", req.Method),
				},
			)

		case "get":
		case "set":

		case "peer.ls":
		case "peer.rm":
		case "peer.add":
		}

		return nil

	})
	return e.Start(types.F(":%s", port))
}
