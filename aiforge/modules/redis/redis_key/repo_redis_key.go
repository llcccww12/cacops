package redis_key

import "fmt"

const REPO_PREFIX = "repo"

func RepoTopNContributors(repoId int64, N int) string {
	return KeyJoin(REPO_PREFIX, fmt.Sprint(repoId), fmt.Sprint(N), "top_n_contributor")
}

func RepoAITemplateCache(repoId int64, branchName string) string {
	return KeyJoin(REPO_PREFIX, fmt.Sprint(repoId), branchName, "ai_template")
}
