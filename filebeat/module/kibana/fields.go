// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package kibana

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "kibana", asset.ModuleFieldsPri, AssetKibana); err != nil {
		panic(err)
	}
}

// AssetKibana returns asset data.
// This is the base64 encoded gzipped contents of module/kibana.
func AssetKibana() string {
	return "eJzMlEHO2yAQhfc+xejfhwN4UanKsuoZookZY2oM7jCoyu0rG6dyCGnaKIuf5Uz43stjxgcY6dLCaM/osQEQK45a+MiFjwZAU+zYzmKDb+FLAwDbr+F70MlRA9Bbcjq2a+8AHifaEZcjl5laMBzSvFUq1FvOnuWC+VOrwR4C8/mG/YgLBJz1FNWuWSruVQVNvGlcpUe6/Aqsi95fDKwmcmYuGGO9WeGqKhsFhd6ne0zM5CVjIfSbkbr2RIJV6XD+QZ0UrVw81c1V+XkolAtGLVKK6acaCDVxVEw9MXFVHp3F8i1mlKGFQWReMImiZMQ9Y7KGMecjnOifnX0+R9esUiQ+oCFfPskzc8vF03pRBbbGenRvC2sKQl+1Zor1tXnsKobEHSmsXn7VUOLyrz0Nh907U1k2LsVj0PVtfjpBcQ4+0oY5dfecV0xNJEMoPyH/NdJVQmnldwAAAP//jzOwrA=="
}
