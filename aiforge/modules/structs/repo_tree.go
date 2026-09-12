// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package structs

// GitEntry represents a git tree
type GitEntry struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	Type string `json:"type"`
	Size int64  `json:"size"`
	SHA  string `json:"sha"`
	URL  string `json:"url"`
}

// GitEntryWithChildren represents a git tree entry with children
type GitEntryWithChildren struct {
	GitEntry
	Children []GitEntryWithChildren `json:"children"`
	Name     string                 `json:"name"`
	Suffix   string                 `json:"suffix"`
}

// GitTreeResponse returns a git tree
type GitTreeResponse struct {
	SHA        string     `json:"sha"`
	URL        string     `json:"url"`
	Entries    []GitEntry `json:"tree"`
	Truncated  bool       `json:"truncated"`
	Page       int        `json:"page"`
	TotalCount int        `json:"total_count"`
}

// GitTreeWithChildrenResponse returns a git tree with children
type GitTreeWithChildrenResponse struct {
	SHA        string                 `json:"sha"`
	URL        string                 `json:"url"`
	Entries    []GitEntryWithChildren `json:"tree"`
	Truncated  bool                   `json:"truncated"`
	Page       int                    `json:"page"`
	TotalCount int                    `json:"total_count"`
}
