// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"code.gitea.io/gitea/modules/log"
	"fmt"
	"strings"

	"code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/util"

	"xorm.io/builder"
)

// RepositoryListDefaultPageSize is the default number of repositories
// to load in memory when running administrative tasks on all (or almost
// all) of them.
// The number should be low enough to avoid filling up all RAM with
// repository data...
const RepositoryListDefaultPageSize = 64

// RepositoryList contains a list of repositories
type RepositoryList []*Repository

func (repos RepositoryList) Len() int {
	return len(repos)
}

func (repos RepositoryList) Less(i, j int) bool {
	return repos[i].FullName() < repos[j].FullName()
}

func (repos RepositoryList) Swap(i, j int) {
	repos[i], repos[j] = repos[j], repos[i]
}

// RepositoryListOfMap make list from values of map
func RepositoryListOfMap(repoMap map[int64]*Repository) RepositoryList {
	return RepositoryList(valuesRepository(repoMap))
}

func (repos RepositoryList) loadAttributes(e Engine) error {
	if len(repos) == 0 {
		return nil
	}

	set := make(map[int64]struct{})
	repoIDs := make([]int64, len(repos))
	setCreator := make(map[int64]struct{})
	for i := range repos {
		set[repos[i].OwnerID] = struct{}{}
		repoIDs[i] = repos[i].ID
		setCreator[repos[i].CreatorID] = struct{}{}

	}

	// Load owners.
	users := make(map[int64]*User, len(set))
	if err := e.
		Where("id > 0").
		In("id", keysInt64(set)).
		Find(&users); err != nil {
		return fmt.Errorf("find users: %v", err)
	}
	//Load creator
	creators := make(map[int64]*User, len(set))
	if err := e.
		Where("id > 0").
		In("id", keysInt64(setCreator)).
		Find(&creators); err != nil {
		return fmt.Errorf("find create repo users: %v", err)
	}

	for i := range repos {
		repos[i].Owner = users[repos[i].OwnerID]
		repos[i].Creator = creators[repos[i].CreatorID]
	}

	// Load primary language.
	stats := make(LanguageStatList, 0, len(repos))
	if err := e.
		Where("`is_primary` = ? AND `language` != ?", true, "other").
		In("`repo_id`", repoIDs).
		Find(&stats); err != nil {
		return fmt.Errorf("find primary languages: %v", err)
	}
	stats.loadAttributes()
	for i := range repos {
		for _, st := range stats {
			if st.RepoID == repos[i].ID {
				repos[i].PrimaryLanguage = st
				break
			}
		}
	}

	return nil
}

// LoadAttributes loads the attributes for the given RepositoryList
func (repos RepositoryList) LoadAttributes() error {
	return repos.loadAttributes(x)
}

// MirrorRepositoryList contains the mirror repositories
type MirrorRepositoryList []*Repository

func (repos MirrorRepositoryList) loadAttributes(e Engine) error {
	if len(repos) == 0 {
		return nil
	}

	// Load mirrors.
	repoIDs := make([]int64, 0, len(repos))
	for i := range repos {
		if !repos[i].IsMirror {
			continue
		}

		repoIDs = append(repoIDs, repos[i].ID)
	}
	mirrors := make([]*Mirror, 0, len(repoIDs))
	if err := e.
		Where("id > 0").
		In("repo_id", repoIDs).
		Find(&mirrors); err != nil {
		return fmt.Errorf("find mirrors: %v", err)
	}

	set := make(map[int64]*Mirror)
	for i := range mirrors {
		set[mirrors[i].RepoID] = mirrors[i]
	}
	for i := range repos {
		repos[i].Mirror = set[repos[i].ID]
	}
	return nil
}

// LoadAttributes loads the attributes for the given MirrorRepositoryList
func (repos MirrorRepositoryList) LoadAttributes() error {
	return repos.loadAttributes(x)
}

// SearchRepoOptions holds the search options
type SearchRepoOptions struct {
	ListOptions
	Actor           *User
	Keyword         string
	OwnerID         int64
	PriorityOwnerID int64
	OrderBy         SearchOrderBy
	Private         bool // Include private repositories in results
	OnlyPrivate     bool // Include only private repositories in results
	StarredByID     int64
	AllPublic       bool // Include also all public repositories of users and public organisations
	AllLimited      bool // Include also all public repositories of limited organisations
	// None -> include collaborative AND non-collaborative
	// True -> include just collaborative
	// False -> incude just non-collaborative
	Collaborate util.OptionalBool
	// None -> include forks AND non-forks
	// True -> include just forks
	// False -> include just non-forks
	Fork util.OptionalBool
	// None -> include templates AND non-templates
	// True -> include just templates
	// False -> include just non-templates
	Template util.OptionalBool
	// None -> include mirrors AND non-mirrors
	// True -> include just mirrors
	// False -> include just non-mirrors
	Mirror util.OptionalBool
	// None -> include archived AND non-archived
	// True -> include just archived
	// False -> include just non-archived
	Archived util.OptionalBool
	// only search topic name
	TopicOnly bool
	//search by Specific TopicName
	TopicName string
	//Multi topic names.available just when topicName is empty
	MultiTopicNames []string
	//available just when MultiTopicNames is not empty,true for intersect,false for union
	IsTopicIntersect bool
	// include description in keyword search
	IncludeDescription bool
	// None -> include has milestones AND has no milestone
	// True -> include just has milestones
	// False -> include just has no milestone
	HasMilestones util.OptionalBool
	// None -> include all repos
	// True -> include just courses
	// False -> include just no courses
	Course            util.OptionalBool
	OnlySearchPrivate bool

	RepoIds []int64

	// 置顶
	HasPinned bool
}

// SearchOrderBy is used to sort the result
type SearchOrderBy string

func (s SearchOrderBy) String() string {
	return string(s)
}

// PlatformSearchResults 工作台查询结果返回
type PlatformSearchResults struct {
	Data PlatformSearch
}

// PlatformSearchResults 工作台查询结果返回
type PlatformSearch struct {
	Repos    []*Repository4Card
	Total    int64
	Page     int
	PageSize int
}

type FindReposResponse struct {
	Repos    []*Repository4Card
	Page     int
	PageSize int
	Total    int64
	RepoIds  []int64
}

// Strings for sorting result
const (
	SearchOrderByTitleAlphabetically         SearchOrderBy = "title ASC"
	SearchOrderByTitleAlphabeticallyReverse  SearchOrderBy = "title DESC"
	SearchOrderByAlphabetically              SearchOrderBy = "name ASC"
	SearchOrderByAlphabeticallyReverse       SearchOrderBy = "name DESC"
	SearchOrderByLeastUpdated                SearchOrderBy = "updated_unix ASC"
	SearchOrderByRecentUpdated               SearchOrderBy = "updated_unix DESC"
	SearchOrderByOldest                      SearchOrderBy = "created_unix ASC"
	SearchOrderByNewest                      SearchOrderBy = "created_unix DESC"
	SearchOrderBySize                        SearchOrderBy = "size ASC"
	SearchOrderBySizeReverse                 SearchOrderBy = "size DESC"
	SearchOrderByID                          SearchOrderBy = "id ASC"
	SearchOrderByIDReverse                   SearchOrderBy = "id DESC"
	SearchOrderByStars                       SearchOrderBy = "num_stars ASC"
	SearchOrderByStarsReverse                SearchOrderBy = "num_stars DESC"
	SearchOrderByForks                       SearchOrderBy = "num_forks ASC"
	SearchOrderByForksReverse                SearchOrderBy = "num_forks DESC"
	SearchOrderByDownloadTimes               SearchOrderBy = "download_times DESC"
	SearchOrderByUseCount                    SearchOrderBy = "use_count ASC"
	SearchOrderByUseCountReverse             SearchOrderBy = "use_count DESC"
	SearchOrderByHot                         SearchOrderBy = "(num_watches + num_stars + num_forks + clone_cnt) DESC"
	SearchOrderByActive                      SearchOrderBy = "(num_issues + num_pulls + num_commit) DESC"
	SearchOrderByWatches                     SearchOrderBy = "num_watches DESC"
	SearchOrderByDefault                     SearchOrderBy = "recommend desc,num_stars DESC,updated_unix DESC"
	SearchOrderByAiTaskCntReverse            SearchOrderBy = "ai_task_cnt desc"
	SearchOrderByModelCntReverse             SearchOrderBy = "model_cnt desc"
	SearchOrderByDatasetCntReverse           SearchOrderBy = "dataset_cnt desc"
	SearchOrderByLastMonthVisitsReverse      SearchOrderBy = "last_month_visits desc"
	SearchOrderByLastFourMonthCommitsReverse SearchOrderBy = "last_four_month_commits desc"
	SearchOrderByRecentRightUpdated          SearchOrderBy = "right_unix DESC"
)

// SearchRepositoryCondition creates a query condition according search repository options
func SearchRepositoryCondition(opts *SearchRepoOptions) builder.Cond {
	var cond = builder.NewCond()
	if opts.OnlySearchPrivate {
		cond = cond.And(builder.Eq{"is_private": true})
	}
	if opts.Private {
		if opts.Actor != nil && !opts.Actor.IsAdmin && opts.Actor.ID != opts.OwnerID {
			// OK we're in the context of a User
			cond = cond.And(accessibleRepositoryCondition(opts.Actor))
		}
	} else {
		// Not looking at private organisations
		// We should be able to see all non-private repositories that
		// isn't in a private or limited organisation.
		cond = cond.And(
			builder.Eq{"is_private": false})
	}

	if opts.OnlyPrivate {
		cond = cond.And(
			builder.Or(
				builder.Eq{"is_private": true}))
	}

	if opts.Template != util.OptionalBoolNone {
		cond = cond.And(builder.Eq{"is_template": opts.Template == util.OptionalBoolTrue})
	}

	// Restrict to starred repositories
	if opts.StarredByID > 0 {
		cond = cond.And(builder.In("id", builder.Select("repo_id").From("star").Where(builder.Eq{"uid": opts.StarredByID})))
	}
	if len(opts.RepoIds) > 0 {
		const MaxINItems = 1000
		if len(opts.RepoIds) <= MaxINItems {
			cond = cond.And(builder.In("id", opts.RepoIds))
		} else {
			repoIdsMap := make(map[int][]int64, 0)
			i := 0

			for j := 0; j < len(opts.RepoIds); j++ {
				if _, ok := repoIdsMap[i]; !ok {
					repoIdsMap[i] = make([]int64, 0)
				}
				repoIdsMap[i] = append(repoIdsMap[i], opts.RepoIds[j])
				if (j+1)%MaxINItems == 0 {
					i += 1
				}

			}
			subCond := builder.NewCond()
			for _, repoSplit := range repoIdsMap {
				subCond = subCond.Or(builder.In("id", repoSplit))
			}
			cond = cond.And(subCond)

		}

	}

	// Restrict repositories to those the OwnerID owns or contributes to as per opts.Collaborate
	if opts.OwnerID > 0 {
		var accessCond = builder.NewCond()
		if opts.Collaborate != util.OptionalBoolTrue {
			accessCond = builder.Eq{"owner_id": opts.OwnerID}
		}

		if opts.Collaborate != util.OptionalBoolFalse {
			// A Collaboration is:
			collaborateCond := builder.And(
				// 1. Repository we don't own
				builder.Neq{"owner_id": opts.OwnerID},
				// 2. But we can see because of:
				builder.Or(
					// A. We have access
					builder.In("`repository`.id",
						builder.Select("`access`.repo_id").
							From("access").
							Where(builder.Eq{"`access`.user_id": opts.OwnerID})),
					// B. We are in a team for
					builder.In("`repository`.id", builder.Select("`team_repo`.repo_id").
						From("team_repo").
						Where(builder.Eq{"`team_user`.uid": opts.OwnerID}).
						Join("INNER", "team_user", "`team_user`.team_id = `team_repo`.team_id")),
				),
			)
			if !opts.Private {
				collaborateCond = collaborateCond.And(builder.Expr("owner_id NOT IN (SELECT org_id FROM org_user WHERE org_user.uid = ? AND org_user.is_public = ?)", opts.OwnerID, false))
			}

			accessCond = accessCond.Or(collaborateCond)
		}

		if opts.AllPublic {
			accessCond = accessCond.Or(builder.Eq{"is_private": false})
		}

		cond = cond.And(accessCond)
	}

	if opts.Keyword != "" {
		// separate keyword
		var subQueryCond = builder.NewCond()
		for _, v := range strings.Split(opts.Keyword, ",") {
			if opts.TopicOnly {
				subQueryCond = subQueryCond.Or(builder.Eq{"topic.name": strings.ToLower(v)})
			} else {
				subQueryCond = subQueryCond.Or(builder.Like{"topic.name", strings.ToLower(v)})
			}
		}
		subQuery := builder.Select("repo_topic.repo_id").From("repo_topic").
			Join("INNER", "topic", "topic.id = repo_topic.topic_id").
			Where(subQueryCond).
			GroupBy("repo_topic.repo_id")

		var keywordCond = builder.In("id", subQuery)
		if !opts.TopicOnly {
			var likes = builder.NewCond()
			for _, v := range strings.Split(opts.Keyword, ",") {
				likes = likes.Or(builder.Like{"lower_alias", strings.ToLower(v)})
				likes = likes.Or(builder.Like{"alias", v})
				if opts.IncludeDescription {
					likes = likes.Or(builder.Like{"LOWER(description)", strings.ToLower(v)})
				}
			}
			keywordCond = keywordCond.Or(likes)
		}
		cond = cond.And(keywordCond)
	}
	if opts.TopicName != "" {
		var subQueryCond = builder.NewCond()
		subQueryCond = subQueryCond.Or(builder.Eq{"topic.name": strings.ToLower(opts.TopicName)})
		subQuery := builder.Select("repo_topic.repo_id").From("repo_topic").
			Join("INNER", "topic", "topic.id = repo_topic.topic_id").
			Where(subQueryCond).
			GroupBy("repo_topic.repo_id")

		var topicNameCond = builder.In("id", subQuery)
		cond = cond.And(topicNameCond)

	} else if opts.MultiTopicNames != nil && len(opts.MultiTopicNames) > 0 {
		var subQueryCond = builder.NewCond()
		subQueryCond = subQueryCond.And(builder.In("topic.name", opts.MultiTopicNames))
		subQuery := builder.Select("repo_topic.repo_id").From("repo_topic").
			Join("INNER", "topic", "topic.id = repo_topic.topic_id").
			Where(subQueryCond).
			GroupBy("repo_topic.repo_id")
		if opts.IsTopicIntersect {
			subQuery = subQuery.Having(fmt.Sprintf("count(repo_topic.repo_id) >= %d", len(opts.MultiTopicNames)))
		}
		var topicNameCond = builder.In("id", subQuery)
		cond = cond.And(topicNameCond)
	}

	if opts.Fork != util.OptionalBoolNone {
		cond = cond.And(builder.Eq{"is_fork": opts.Fork == util.OptionalBoolTrue})
	}

	if opts.Mirror != util.OptionalBoolNone {
		cond = cond.And(builder.Eq{"is_mirror": opts.Mirror == util.OptionalBoolTrue})
	}

	if opts.Course == util.OptionalBoolTrue {
		cond = cond.And(builder.Eq{"repo_type": RepoCourse})
	}

	if opts.Actor != nil && opts.Actor.IsRestricted {
		cond = cond.And(accessibleRepositoryCondition(opts.Actor))
	}

	if opts.Archived != util.OptionalBoolNone {
		cond = cond.And(builder.Eq{"is_archived": opts.Archived == util.OptionalBoolTrue})
	}

	switch opts.HasMilestones {
	case util.OptionalBoolTrue:
		cond = cond.And(builder.Gt{"num_milestones": 0})
	case util.OptionalBoolFalse:
		cond = cond.And(builder.Eq{"num_milestones": 0}.Or(builder.IsNull{"num_milestones"}))
	}

	return cond
}

// SearchRepository returns repositories based on search options,
// it returns results in given range and number of total results.
func SearchRepository(opts *SearchRepoOptions) (RepositoryList, int64, error) {
	if opts.Keyword != "" {
		opts.Keyword = strings.ReplaceAll(opts.Keyword, "'", "")
	}
	cond := SearchRepositoryCondition(opts)
	return SearchRepositoryByCondition(opts, cond, true)
}

// SearchRepositoryByCondition search repositories by condition
func SearchRepositoryByCondition(opts *SearchRepoOptions, cond builder.Cond, loadAttributes bool) (RepositoryList, int64, error) {
	if opts.Page <= 0 {
		opts.Page = 1
	}

	if len(opts.OrderBy) == 0 {
		opts.OrderBy = SearchOrderByAlphabetically
	}

	if opts.PriorityOwnerID > 0 {
		opts.OrderBy = SearchOrderBy(fmt.Sprintf("CASE WHEN owner_id = %d THEN 0 ELSE owner_id END, %s", opts.PriorityOwnerID, opts.OrderBy))
	}
	if opts.Keyword != "" {
		opts.OrderBy = SearchOrderBy(fmt.Sprintf("CASE WHEN alias='%s' THEN 0 ELSE 1 END, %s", opts.Keyword, opts.OrderBy))
	}

	sess := x.NewSession()
	defer sess.Close()

	count, err := sess.Where(cond).Count(new(Repository))

	if err != nil {
		log.Error("Repository Count err:  %+v", err)
		return nil, 0, fmt.Errorf("Count: %v", err)
	}

	repos := make(RepositoryList, 0, opts.PageSize)

	if opts.PageSize > 0 {
		sess.Limit(opts.PageSize, (opts.Page-1)*opts.PageSize)
	}

	// 条件设置
	// 当前用户是否存在置顶
	if opts.HasPinned && opts.Actor != nil {
		// 使用子查询避免 JOIN 导致的列名歧义
		// 需要限定 user_id 避免同一 repository_id 对应多条记录导致子查询返回多行
		// pinned_unix > 0 才表示有效置顶，pinned_unix = 0 表示已取消置顶，应参与普通排序
		opts.OrderBy = SearchOrderBy(fmt.Sprintf(
			"CASE WHEN EXISTS (SELECT 1 FROM repository_pinned_record WHERE repository_pinned_record.repository_id = `repository`.id AND repository_pinned_record.user_id = %d AND repository_pinned_record.pinned_unix > 0) THEN 0 ELSE 1 END, "+
				"(SELECT pinned_unix FROM repository_pinned_record WHERE repository_pinned_record.repository_id = `repository`.id AND repository_pinned_record.user_id = %d AND repository_pinned_record.pinned_unix > 0) DESC, %s", opts.Actor.ID, opts.Actor.ID, opts.OrderBy))
	}

	sess.Where(cond).OrderBy(opts.OrderBy.String())

	if err = sess.Find(&repos); err != nil {
		return nil, 0, fmt.Errorf("Repo: %v", err)
	}

	if loadAttributes {
		if err = repos.loadAttributes(sess); err != nil {
			return nil, 0, fmt.Errorf("LoadAttributes: %v", err)
		}
	}

	return repos, count, nil
}

// CountRepositoryOverView 计算项目的概览页面
func CountRepositoryOverView(userID int64) (totalSum, sourceSum, forkSum, mirrorSum, collaborateSum int64, err error) {
	// 查询用户自建/派生/镜像的sql
	statsSql := `
	SELECT
		COUNT(CASE WHEN is_fork = false AND is_mirror = false THEN 1 END) as source_sum,
		COUNT(CASE WHEN is_fork = true THEN 1 END) as fork_sum,
		COUNT(CASE WHEN is_mirror = true THEN 1 END) as mirror_sum
	FROM "repository"
	WHERE owner_id = ? AND is_archived = false`

	var stats struct {
		SourceSum int64
		ForkSum   int64
		MirrorSum int64
	}

	if _, err = x.SQL(statsSql, userID).Unscoped().Get(&stats); err != nil {
		return
	}

	sourceSum = stats.SourceSum
	forkSum = stats.ForkSum
	mirrorSum = stats.MirrorSum

	// 查询用户的协作
	collaborateSql := `SELECT COUNT(*) FROM "repository" WHERE (owner_id <> ? AND (
        "repository".id IN (
            SELECT "access".repo_id
            FROM access
            WHERE "access".user_id = ?
        )
        OR "repository".id IN (
            SELECT "team_repo".repo_id
            FROM team_repo
            INNER JOIN team_user
                ON "team_user".team_id = "team_repo".team_id
            WHERE "team_user".uid = ?
        )
    )
) AND   is_archived = false`

	if _, err = x.SQL(collaborateSql, userID, userID, userID).Unscoped().Get(&collaborateSum); err != nil {
		return
	}

	totalSum = sourceSum + forkSum + mirrorSum + collaborateSum

	return
}

// accessibleRepositoryCondition takes a user a returns a condition for checking if a repository is accessible
func accessibleRepositoryCondition(user *User) builder.Cond {
	var cond = builder.NewCond()

	if user == nil || !user.IsRestricted || user.ID <= 0 {
		orgVisibilityLimit := []structs.VisibleType{structs.VisibleTypePrivate}
		if user == nil || user.ID <= 0 {
			orgVisibilityLimit = append(orgVisibilityLimit, structs.VisibleTypeLimited)
		}
		// 1. Be able to see all non-private repositories that either:
		cond = cond.Or(builder.And(
			builder.Eq{"`repository`.is_private": false},
			// 2. Aren't in an private organisation or limited organisation if we're not logged in
			builder.NotIn("`repository`.owner_id", builder.Select("id").From("`user`").Where(
				builder.And(
					builder.Eq{"type": UserTypeOrganization},
					builder.In("visibility", orgVisibilityLimit)),
			))))
	}

	if user != nil {
		cond = cond.Or(
			// 2. Be able to see all repositories that we have access to
			builder.In("`repository`.id", builder.Select("repo_id").
				From("`access`").
				Where(builder.And(
					builder.Eq{"user_id": user.ID},
					builder.Gt{"mode": int(AccessModeNone)}))),
			// 3. Repositories that we directly own
			builder.Eq{"`repository`.owner_id": user.ID},
			// 4. Be able to see all repositories that we are in a team
			builder.In("`repository`.id", builder.Select("`team_repo`.repo_id").
				From("team_repo").
				Where(builder.Eq{"`team_user`.uid": user.ID}).
				Join("INNER", "team_user", "`team_user`.team_id = `team_repo`.team_id")),
			// 5. Be able to see all public repos in private organizations that we are an org_user of
			builder.And(builder.Eq{"`repository`.is_private": false},
				builder.In("`repository`.owner_id",
					builder.Select("`org_user`.org_id").
						From("org_user").
						Where(builder.Eq{"`org_user`.uid": user.ID}))))
	}

	return cond
}

// SearchRepositoryByName takes keyword and part of repository name to search,
// it returns results in given range and number of total results.
func SearchRepositoryByName(opts *SearchRepoOptions) (RepositoryList, int64, error) {
	opts.IncludeDescription = false
	return SearchRepository(opts)
}

func GetAccessibleRepoIDs(user *User) ([]int64, error) {
	repoIDs := make([]int64, 0, 10)
	if err := x.
		Table("repository").
		Cols("id").
		Where(accessibleRepositoryCondition(user)).
		Find(&repoIDs); err != nil {
		return nil, fmt.Errorf("GetAccessibleRepoIDs: %v", err)
	}
	return repoIDs, nil
}

// AccessibleRepoIDsQuery queries accessible repository ids. Usable as a subquery wherever repo ids need to be filtered.
func AccessibleRepoIDsQuery(user *User) *builder.Builder {
	// NB: Please note this code needs to still work if user is nil
	return builder.Select("id").From("repository").Where(accessibleRepositoryCondition(user))
}

// FindUserAccessibleRepoIDs find all accessible repositories' ID by user's id
func FindUserAccessibleRepoIDs(user *User) ([]int64, error) {
	repoIDs := make([]int64, 0, 10)
	if err := x.
		Table("repository").
		Cols("id").
		Where(accessibleRepositoryCondition(user)).
		Find(&repoIDs); err != nil {
		return nil, fmt.Errorf("FindUserAccesibleRepoIDs: %v", err)
	}
	return repoIDs, nil
}
