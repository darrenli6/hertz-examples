/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by hertz generator.

package Casbin

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_casbin/biz/model/casbin"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_casbin/biz/mw"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_casbin/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"

	"net/http"
)

func rootMw() []app.HandlerFunc {

	fmt.Println("group middleware start  ")
	return []app.HandlerFunc{func(ctx context.Context, c *app.RequestContext) {

		if c.FullPath() == "/v1/login" {
			c.Next(ctx)
			return
		}

		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusInternalServerError, casbin.BasicResponse{
				Code: 2,
				Msg:  "token is not null ",
			})
			c.Abort()
			return
		}

		claim, err := utils.AnalyzeToken(token)

		if err != nil {
			c.JSON(http.StatusInternalServerError, casbin.BasicResponse{
				Code: 2,
				Msg:  "token is not valid ",
			})
			c.Abort()
			return
		}

		roles := claim.Roles

		// casbin enforce
		var isAuth bool = false

		for _, v := range roles {
			res, err := mw.Authorize(v.Name, c.FullPath(), string(c.Request.Header.Method()))
			if err != nil {
				c.JSON(http.StatusInternalServerError, casbin.BasicResponse{
					Code: 3,
					Msg:  "Authorize is error ",
				})
				c.Abort()
				return
			}
			if res {
				isAuth = true
				break
			}
		}

		if isAuth {
			c.Next(ctx)
		} else {
			c.JSON(http.StatusInternalServerError, casbin.BasicResponse{
				Code: 4,
				Msg:  "FORBIDDEN ",
			})
			c.Abort()
			return
		}

	}}

}

func _v1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _roleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _bindMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _bindroleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createroleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _create0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _permissionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createpermissionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _permissionroleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _bindpermissionroleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _bind0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _create1Mw() []app.HandlerFunc {
	// your code...
	return nil
}
