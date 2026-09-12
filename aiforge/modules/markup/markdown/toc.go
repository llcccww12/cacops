// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package markdown

import (
	"fmt"
	"net/url"

	"github.com/yuin/goldmark/ast"
)

func createTOCNode(toc []Header, lang string) ast.Node {
	sidebar := ast.NewDocument()
	sidebar.SetAttributeString("class", []byte("markdown_catalog"))
	scrollContainer := ast.NewDocument()
	scrollContainer.SetAttributeString("class", []byte("scroll-container"))
	toggleContainer := ast.NewDocument()
	toggleContainer.SetAttributeString("class", []byte("toggle-container"))
	toggleIcon := ast.NewDocument()
	toggleIcon.SetAttributeString("class", []byte("icon ri-arrow-drop-left-line"))
	container := ast.NewDocument()
	container.SetAttributeString("class", []byte("container"))
	ul := ast.NewList('-')
	ul.SetAttributeString("class", []byte("markdown_toc"))
	topul:=ul

	currentLevel := 6
	for _, header := range toc {
		if header.Level < currentLevel {
			currentLevel = header.Level
		}
	}
	for _, header := range toc {
		for currentLevel > header.Level {
			ul = ul.Parent().Parent().(*ast.List)
			currentLevel--
		}
		for currentLevel < header.Level {
			newLi := ast.NewListItem(currentLevel * 2)
			newLi.SetAttributeString("class", []byte("no-catalog-li"))
			newL := ast.NewList('-')
			newL.SetAttributeString("class", []byte("markdown_toc"))
			newLi.AppendChild(newLi, newL)
			ul.AppendChild(ul, newLi)
			currentLevel++
			ul = newL
		}
		li := ast.NewListItem(currentLevel * 2)
		li.SetAttributeString("class", []byte("catalog-li"))
		a := ast.NewLink()
		a.Destination = []byte(fmt.Sprintf("#%s", url.PathEscape(header.ID)))
		a.AppendChild(a, ast.NewString([]byte(header.Text)))
		a.SetAttributeString("title", []byte(header.Text))
		li.AppendChild(li, a)
		ul.AppendChild(ul, li)
	}
	container.AppendChild(container,topul)
	scrollContainer.AppendChild(scrollContainer,container)
	toggleContainer.AppendChild(toggleContainer,toggleIcon)
	sidebar.AppendChild(sidebar, toggleContainer)
	sidebar.AppendChild(sidebar, scrollContainer)

	//return details
	return sidebar
}
