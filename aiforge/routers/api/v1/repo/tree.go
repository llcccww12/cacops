// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repo

import (
	"code.gitea.io/gitea/modules/git"
	api "code.gitea.io/gitea/modules/structs"
	"net/http"
	"regexp"
	"strings"

	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/repofiles"
)

// GetTree get the tree of a repository.
func GetTree(ctx *context.APIContext) {
	// swagger:operation GET /repos/{owner}/{repo}/git/trees/{sha} repository GetTree
	// ---
	// summary: Gets the tree of a repository.
	// produces:
	// - application/json
	// parameters:
	// - name: owner
	//   in: path
	//   description: owner of the repo
	//   type: string
	//   required: true
	// - name: repo
	//   in: path
	//   description: name of the repo
	//   type: string
	//   required: true
	// - name: sha
	//   in: path
	//   description: sha of the commit
	//   type: string
	//   required: true
	// - name: recursive
	//   in: query
	//   description: show all directories and files
	//   required: false
	//   type: boolean
	// - name: page
	//   in: query
	//   description: page number; the 'truncated' field in the response will be true if there are still more items after this page, false if the last page
	//   required: false
	//   type: integer
	// - name: per_page
	//   in: query
	//   description: number of items per page; default is 1000 or what is set in app.ini as DEFAULT_GIT_TREES_PER_PAGE
	//   required: false
	//   type: integer
	// responses:
	//   "200":
	//     "$ref": "#/responses/GitTreeResponse"
	//   "400":
	//     "$ref": "#/responses/error"

	sha := ctx.Params(":sha")
	if len(sha) == 0 {
		ctx.Error(http.StatusBadRequest, "", "sha not provided")
		return
	}
	if tree, err := repofiles.GetTreeBySHA(ctx.Repo.Repository, sha, ctx.QueryInt("page"), ctx.QueryInt("per_page"), ctx.QueryBool("recursive"), true); err != nil {
		ctx.Error(http.StatusBadRequest, "", err.Error())
	} else {
		ctx.JSON(http.StatusOK, tree)
	}
}

// GetTree get the tree of a repository.
func GetCodes(ctx *context.APIContext) {
	// swagger:operation GET /repos/{owner}/{repo}/git/codes/{branch} repository GetTree
	// ---
	// summary: Gets the tree of a repository.
	// produces:
	// - application/json
	// parameters:
	// - name: owner
	//   in: path
	//   description: owner of the repo
	//   type: string
	//   required: true
	// - name: repo
	//   in: path
	//   description: name of the repo
	//   type: string
	//   required: true
	// - name: recursive
	//   in: query
	//   description: show all directories and files
	//   required: false
	//   type: boolean
	// - name: page
	//   in: query
	//   description: page number; the 'truncated' field in the response will be true if there are still more items after this page, false if the last page
	//   required: false
	//   type: integer
	// - name: per_page
	//   in: query
	//   description: number of items per page; default is 1000 or what is set in app.ini as DEFAULT_GIT_TREES_PER_PAGE
	//   required: false
	//   type: integer
	// responses:
	//   "200":
	//     "$ref": "#/responses/GitTreeResponse"
	//   "400":
	//     "$ref": "#/responses/error"
	var currentBranch = ctx.Params("currentBranch")
	repository := ctx.Repo.Repository
	repository.CurrentBranch = currentBranch

	// 获取分支的头部commit
	gitRepo, _ := git.OpenRepository(repository.RepoPath())
	branch, _ := gitRepo.GetBranch(repository.CurrentBranch)
	commit, _ := branch.GetCommit()
	defer gitRepo.Close()
	sha := commit.ID.String()

	if tree, err := repofiles.GetTreeBySHA(repository, sha, ctx.QueryInt("page"), ctx.QueryInt("per_page"), true, false); err != nil {
		ctx.Error(http.StatusBadRequest, "", err.Error())
	} else {
		// 第一步，反序
		var entries = tree.Entries
		entriesWithChildren := make([]api.GitEntryWithChildren, len(entries))
		slashRegex := regexp.MustCompile(`.*/`)
		for i, j := len(entries)-1, 0; i >= 0; i-- {
			entriesWithChildren[j].Children = make([]api.GitEntryWithChildren, 0)
			entriesWithChildren[j].GitEntry = entries[i]
			entriesWithChildren[j].Name = slashRegex.ReplaceAllString(entriesWithChildren[j].Path, "")
			tempList := strings.Split(entriesWithChildren[j].Name, ".")
			if len(tempList) > 1 {
				entriesWithChildren[j].Suffix = tempList[len(tempList)-1]
			} else {
				entriesWithChildren[j].Suffix = ""
			}
			j++
		}

		// 第二步，对包含分隔符的节点，向前寻找路径等于当前路径去除当前文件名的节点，挂到其children下
		for i := 0; i < len(entriesWithChildren); i++ {
			if strings.Contains(entriesWithChildren[i].Path, "/") {
				parentPath := entriesWithChildren[i].Path[0:strings.LastIndex(entriesWithChildren[i].Path, "/")]
				for j := i; j < len(entriesWithChildren); j++ {
					if entriesWithChildren[j].Path == parentPath && entriesWithChildren[j].Type == "tree" {
						entriesWithChildren[j].Children = append(entriesWithChildren[j].Children, entriesWithChildren[i])
					}
				}
			}
		}

		// 第三步，找到第一层的节点，返回
		var treeWithChildren = new(api.GitTreeWithChildrenResponse)
		entriesWithChildrenForResponse := make([]api.GitEntryWithChildren, 0)
		for i := len(entriesWithChildren) - 1; i >= 0; i-- {
			if !strings.Contains(entriesWithChildren[i].Path, "/") {
				entriesWithChildrenForResponse = append(entriesWithChildrenForResponse, entriesWithChildren[i])
			}
		}

		response := make([]api.GitEntryWithChildren, 0)

		for _,entry := range entriesWithChildrenForResponse{
			if entry.Type == "tree" {
				response = append(response, entry)
			}

		}

		for _,entry := range entriesWithChildrenForResponse{
			if entry.Type == "blob" {
				response = append(response, entry)
			}
		}

		//treeWithChildren.Entries = entriesWithChildrenForResponse
		treeWithChildren.Entries = response
		treeWithChildren.URL = tree.URL
		treeWithChildren.SHA = tree.SHA
		treeWithChildren.TotalCount = tree.TotalCount

		ctx.JSON(http.StatusOK, treeWithChildren)
	}
}
