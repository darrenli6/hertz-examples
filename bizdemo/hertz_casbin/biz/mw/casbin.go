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

package mw

import (
	"github.com/casbin/casbin"
	xormadapter "github.com/casbin/xorm-adapter"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_casbin/pkg/consts"
)

var AuthEnforcer *casbin.Enforcer

func InitCasbin() {
	adapter := xormadapter.NewAdapter("mysql", consts.MysqlDSN, true)

	enforcer := casbin.NewEnforcer("conf/auth_model.conf", adapter)

	AuthEnforcer = enforcer
}

func Authorize(rvals ...interface{}) (result bool, err error) {
	// casbin enforce
	res, err1 := AuthEnforcer.EnforceSafe(rvals[0], rvals[1], rvals[2])
	return res, err1
}
