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
	"net/http"

	"aahframe.work"
)

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Before interceptor
func (c *AppController) Before() {
	c.AddViewArg("IsHome", false)
}

// Index method is application's home page.
func (c *AppController) Index() {
	c.Reply().Ok().HTML(aah.Data{
		"IsHome": true,
	})
}

// ComingSoon display banner.
func (c *AppController) ComingSoon() {
	c.Reply().Ok().HTML(aah.Data{
		"IsComingSoon": true,
	})
}

// LatestRelease method to provides information on latest release.
func (c *AppController) LatestRelease() {
	c.Reply().RedirectWithStatus("https://github.com/thumbai/thumbai/releases/latest", http.StatusFound)
}

// About serves the about page of THUMBAI app.
func (c *AppController) About() {
	c.Reply().Ok()
}
