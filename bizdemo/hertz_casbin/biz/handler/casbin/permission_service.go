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
	"strconv"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_casbin/biz/dal/mysql"
	casbin "github.com/cloudwego/hertz-examples/bizdemo/hertz_casbin/biz/model/casbin"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_casbin/biz/mw"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreatePermission .
// @router /v1/permission/create/ [POST]
func CreatePermission(ctx context.Context, c *app.RequestContext) {
	var err error
	var req casbin.CreatePermissionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(casbin.CreatePermissionResponse)
	resp.Code = 0
	resp.Msg = "ok"

	rpermission := casbin.Permission{
		V1: req.V1,
		V2: req.V2,
	}
	cPermission, err := mysql.QueryPermissionByV(req.V1, req.V2)
	if err != nil {
		resp.Code = 2
		resp.Msg = "create failed"
		c.JSON(consts.StatusOK, resp)
		return
	}

	if cPermission.ID > 0 {
		resp.Code = 3
		resp.Msg = "Data already exists"
		c.JSON(consts.StatusOK, resp)
		return
	}

	err = mysql.CreatePermission(&rpermission)
	if err != nil {
		resp.Code = 4
		resp.Msg = "create failed"
		c.JSON(consts.StatusOK, resp)
		return
	}
	resp.Permission = &casbin.Permission{
		V1: req.V1,
		V2: req.V2,
	}

	c.JSON(consts.StatusOK, resp)
}

// BindPermissionRole .
// @router /v1/permissionrole/bind/ [POST]
func BindPermissionRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req casbin.BindPermissionRoleRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(casbin.BindPermissionRoleResponse)
	resp.Code = 0
	resp.Msg = "ok"

	rid, _ := strconv.Atoi(req.Rid)
	pid, _ := strconv.Atoi(req.Pid)

	cRole, err := mysql.QueryRoleById(rid)
	if cRole.ID == 0 {
		resp.Code = 2
		resp.Msg = "Role data is null"
		c.JSON(consts.StatusOK, resp)
		return
	}

	cPermission, err := mysql.QueryPermissionById(pid)
	if cPermission.ID == 0 {
		resp.Code = 2
		resp.Msg = "Permission data is null"
		c.JSON(consts.StatusOK, resp)
		return
	}

	permissionRole := mysql.QuerypermissionRoleByIds(pid, rid)
	if len(permissionRole) > 0 {
		resp.Code = 2
		resp.Msg = "Data already exists "
		c.JSON(consts.StatusOK, resp)
		return
	}

	permissionRoleReq := casbin.PermissionRole{
		Rid: int64(rid),
		Pid: int64(pid),
	}

	err = mysql.BindPermissionRole(&permissionRoleReq)
	if err != nil {
		resp.Code = 7
		resp.Msg = "Bind failed "
		c.JSON(consts.StatusOK, resp)
		return
	}

	// add policy
	if ok := mw.AuthEnforcer.AddPolicy(cRole.Name, cPermission.V1, cPermission.V2); !ok {
		resp.Code = 7
		resp.Msg = "Policy insert failed "
		c.JSON(consts.StatusOK, resp)
		return
	} else {
		resp.Code = 0
		resp.Msg = "Policy insert successfully  "
		resp.PermissionRole.Rid = int64(rid)
		resp.PermissionRole.Pid = int64(pid)
		c.JSON(consts.StatusOK, resp)
		return
	}

}
