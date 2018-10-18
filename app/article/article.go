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

package article

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"aahframe.work"
	"aahframe.work/essentials"
	"aahframe.work/log"

	"github.com/russross/blackfriday"
	"thumbai.app/website/app/models"
)

var (
	articleCache = make(map[string]*models.Article)

	markdownHTMLFlags = 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_DASHES |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	markdownExtensions = 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_AUTO_HEADER_IDS |
		blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
		blackfriday.EXTENSION_DEFINITION_LISTS

	markdownOptions = blackfriday.Options{Extensions: markdownExtensions}

	isCacheEnabled bool

	mu              = &sync.Mutex{}
	contentPreparer *strings.Replacer
)

// Get method returns the parsed markdown content for given URL path.
func Get(mdPath string) (*models.Article, bool) {
	if isCacheEnabled {
		if article, found := articleCache[mdPath]; found {
			return article, true
		}
	}
	article := loadArticle(mdPath)
	if article == nil {
		return nil, false
	}
	if len(article.Content) > 0 && isCacheEnabled {
		mu.Lock()
		articleCache[mdPath] = article
		mu.Unlock()
	}
	return article, len(article.Content) > 0
}

// Parse method parsed the markdown content into html using blackfriday library
// and create Article object.
func Parse(lines []string) *models.Article {
	article := &models.Article{TOC: make([]models.TocLink, 0)}
	pos := lineNumByDelimiter(lines, "---")
	if pos > 0 {
		for _, v := range lines[:pos] {
			idx := strings.IndexByte(v, ':')
			if idx == -1 {
				continue
			}
			switch v[:idx] {
			case "Title":
				article.Title = strings.TrimSpace(v[idx+1:])
			case "Desc":
				article.Desc = strings.TrimSpace(v[idx+1:])
			case "Keywords":
				article.Keywords = strings.TrimSpace(v[idx+1:])
			}
		}
	}

	tocPos := lineNumByDelimiter(lines[pos+1:], "---")
	tocPos += pos + 1
	if pos > 0 {
		for _, v := range lines[pos+1 : tocPos] {
			parts := strings.Split(v, "|")
			if len(parts) == 1 {
				article.TOC = append(article.TOC, models.TocLink{Name: parts[0]})
			} else {
				article.TOC = append(article.TOC, models.TocLink{Name: parts[0], AnchorTag: parts[1]})
			}
		}
	}

	fileContent := strings.Join(lines[tocPos+1:], "\n")
	htmlRender := blackfriday.HtmlRenderer(markdownHTMLFlags, "", "")
	content := string(blackfriday.MarkdownOptions([]byte(fileContent), htmlRender, markdownOptions))
	content = contentPreparer.Replace(content)
	article.Content = content
	return article
}

// LoadCache methods loads the markdown into cache for given base path.
func LoadCache(docBasePath string) {
	var files []string
	excludes := ess.Excludes{".git", "LICENSE", "README.md"}
	_ = ess.Walk(docBasePath, func(srcPath string, info os.FileInfo, err error) error {
		if excludes.Match(filepath.Base(srcPath)) {
			if info.IsDir() {
				// excluding directory
				return filepath.SkipDir
			}
			// excluding file
			return nil
		}

		if !info.IsDir() {
			files = append(files, srcPath)
		}
		return nil
	})

	for _, md := range files {
		RefreshCacheByFile(md)
	}
}

// RefreshCacheByFile method refereshes the Markdown cache by file.
func RefreshCacheByFile(mdPath string) {
	article := loadArticle(mdPath)
	if article != nil && len(article.Content) > 0 {
		mu.Lock()
		articleCache[mdPath] = article
		mu.Unlock()
		log.Infof("Refreshed file: %s", mdPath)
	} else {
		log.Warnf("Referesh: File not found: %s", mdPath)
	}
}

// RemoveCacheByFile method removes the Markdown cache by file.
func RemoveCacheByFile(mdPath string) {
	if _, found := articleCache[mdPath]; found {
		mu.Lock()
		delete(articleCache, mdPath)
		mu.Unlock()
		log.Infof("Removed from cache: %s", mdPath)
	} else {
		log.Warnf("Remove: File not found: %s", mdPath)
	}
}

// ClearCache method clears the Markdown cache.
func ClearCache(_ *aah.Event) {
	if len(articleCache) > 0 {
		aah.AppLog().Info("Clearing cache")
	}
	mu.Lock()
	articleCache = make(map[string]*models.Article)
	mu.Unlock()
}

// LoadConfig gets markdown related config value
func LoadConfig(_ *aah.Event) {
	cfg := aah.AppConfig()
	isCacheEnabled = cfg.BoolDefault("markdown.cache", false)

	// Dynamica URL placement
	contentPreparer = strings.NewReplacer(
		"<table>", `<table class="table table-bordered table-hover">`,
	)
}

func loadArticle(mdPath string) *models.Article {
	f, err := os.Open(mdPath)
	if err != nil {
		return nil
	}

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	article := Parse(lines)
	article.File = mdPath
	return article
}

func lineNumByDelimiter(lines []string, delim string) int {
	pos := 0
	for i, l := range lines {
		if strings.TrimSpace(l) == delim {
			return i
		}
	}
	return pos
}

