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

package casbin

import (
	"context"
	"github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/biz/dal/mysql"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	casbin "github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/biz/model/casbin"
)

// CreateRole .
// @router /v1/role/create/ [POST]
func CreateRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req casbin.CreateRoleRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(casbin.CreateRoleResponse)
	resp.Code = 0
	resp.Msg = "ok"

	roleReq := casbin.Role{
		Name: req.Name,
	}

	cRole, err := mysql.QueryRoleByName(req.Name)
	if err != nil {
		resp.Code = 2
		resp.Msg = "create failed"
		c.JSON(consts.StatusOK, resp)
		return
	}

	if cRole.ID > 0 {
		resp.Code = 3
		resp.Msg = "Role data already exists"
		c.JSON(consts.StatusOK, resp)
		return
	}

	err = mysql.CreateRole(&roleReq)
	if err != nil {
		resp.Code = 1
		resp.Msg = "create failed"
		c.JSON(consts.StatusOK, resp)
		return
	}
	resp.Role = &casbin.Role{
		Name: req.Name,
	}

	c.JSON(consts.StatusOK, resp)
}

// BindRole .
// @router /v1/role/bind/ [POST]
func BindRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req casbin.BindRoleRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(casbin.BindRoleResponse)

	c.JSON(consts.StatusOK, resp)
}
