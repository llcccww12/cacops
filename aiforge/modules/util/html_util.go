package util

import (
	"strings"

	"golang.org/x/net/html"
)

// StripHTML 去掉 HTML 标签，保留纯文本，并处理常见换行标签
func StripHTML(input string) string {
	if input == "" {
		return ""
	}

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		// 如果解析失败，就直接返回原始字符串
		return input
	}

	var b strings.Builder
	var f func(*html.Node)
	f = func(n *html.Node) {
		switch n.Type {
		case html.TextNode:
			b.WriteString(n.Data)
		case html.ElementNode:
			// 遇到换行类标签，手动补换行
			switch n.Data {
			case "br":
				b.WriteString("\n")
			case "p", "div":
				// 段落和块级元素，用双换行分隔
				b.WriteString("\n\n")
			}
		}
		// 递归遍历子节点
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	// 去掉多余的首尾空白
	return strings.TrimSpace(b.String())
}
