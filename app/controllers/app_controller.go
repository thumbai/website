// Copyright Jeevanandam M. (https://github.com/jeevatkm, jeeva@myjeeva.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
  "aahframe.work"

  "thumbai.app/website/app/models"
)

// AppController struct application controller
type AppController struct {
  *aah.Context
}

// Index method is application's home page.
func (c *AppController) Index() {
  data := aah.Data{
    "Greet": models.Greet{
      Message: "A Go Mod Repository, Go Vanity and Simple Proxy Server",
    },
  }

  c.Reply().Ok().HTML(data)
}

// ComingSoon display banner.
func (c *AppController) ComingSoon() {
  c.Reply().Ok().HTML(aah.Data{
    "IsComingSoon": true,
  })
}
