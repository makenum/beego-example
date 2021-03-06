// Copyright 2020 beego-dev
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"io/ioutil"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
)

func main() {

	// be careful about the work directory

	// Upload File
	fileReq := httplib.Post("http://beego.me/")
	fileReq.Param("username", "astaxie")
	fileReq.Param("password", "123456")
	fileReq.PostFile("uploadfile", "hello.txt")

	// Bigfile
	bigFileReq := httplib.Post("http://beego.me/")
	bt, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		logs.Error(err)
	}
	bigFileReq.Body(bt)
}
