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
	"path/filepath"

	"aahframe.work"
	"aahframe.work/essentials"
	
	"thumbai.app/website/app/article"
)

// DocsController serve the THUMBAI documentation.
type DocsController struct {
	*aah.Context
}

// Before interceptor
func (c *DocsController) Before() {
	c.AddViewArg("IsHome", false)
}

// Index shows docs homepage
func (c *DocsController) Index() {
	c.Reply().Redirect(c.RouteURL("show_doc", "introduction"))
}

// ShowDoc method serves requested documentation content.
func (c *DocsController) ShowDoc(doc string) {
	p := filepath.Join(aah.AppBaseDir(), "docs", doc+".md")	
	if ess.IsFileExists(p) {
		article, _ := article.Get(p)
		c.Reply().HTMLl("docs.html", aah.Data{
			"IsDocs": true,
			"Article": article,
		})
		return 
	}

	c.Reply().HTMLl("docs.html", aah.Data{
		"IsDocs": true,
	})
}
