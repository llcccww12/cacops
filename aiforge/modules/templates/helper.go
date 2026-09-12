// Copyright 2018 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package templates

import (
	"bytes"
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"html/template"
	"mime"
	"net/url"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	texttmpl "text/template"
	"time"
	"unicode"

	"code.gitea.io/gitea/modules/dataset"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/emoji"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/markup"
	"code.gitea.io/gitea/modules/repository"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/modules/util"
	"code.gitea.io/gitea/services/gitdiff"
	mirror_service "code.gitea.io/gitea/services/mirror"

	"github.com/editorconfig/editorconfig-core-go/v2"
)

const (
	REF_HEADS_PREFIX  = "refs/heads/"
	REF_TAGS_PREFIX   = "refs/tags/"
	REF_TYPE_BRANCH   = "branch"
	REF_TYPE_TAG      = "tag"
	REF_TYPE_PATTERN  = "(refs/heads/|refs/tags/)"
	DURATION_STR_ZERO = "00:00:00"
)

// Used from static.go && dynamic.go
var mailSubjectSplit = regexp.MustCompile(`(?m)^-{3,}[\s]*$`)

// NewFuncMap returns functions for injecting to templates
func NewFuncMap() []template.FuncMap {
	return []template.FuncMap{map[string]interface{}{
		"GoVer": func() string {
			return strings.Title(runtime.Version())
		},
		"UseHTTPS": func() bool {
			return strings.HasPrefix(setting.AppURL, "https")
		},
		"AppName": func() string {
			return setting.AppName
		},
		"MLOPS": func() bool {
			return setting.MLOPS
		},
		"MlopsHost": func() string {
			return setting.MlopsHost
		},
		"AppSubUrl": func() string {
			return setting.AppSubURL
		},
		"StaticUrlPrefix": func() string {
			return setting.StaticURLPrefix
		},
		"AppUrl": func() string {
			return setting.AppURL
		},
		"AppVer": func() string {
			return setting.AppVer
		},
		"AppBuiltWith": func() string {
			return setting.AppBuiltWith
		},
		"AppDomain": func() string {
			return setting.Domain
		},
		"DisableGravatar": func() bool {
			return setting.DisableGravatar
		},
		"DefaultShowFullName": func() bool {
			return setting.UI.DefaultShowFullName
		},
		"ShowFooterTemplateLoadTime": func() bool {
			return setting.ShowFooterTemplateLoadTime
		},
		"LoadTimes": func(startTime time.Time) string {
			return fmt.Sprint(time.Since(startTime).Nanoseconds()/1e6) + "ms"
		},
		"AllowedReactions": func() []string {
			return setting.UI.Reactions
		},
		"DebugAttachSize": func() int {
			return setting.DebugAttachSize * 1024 * 1024 * 1024
		},
		"LlmCommonKB": func() string {
			return setting.LLM_CHAT_API.COMMON_KB
		},
		"LlmMaxCounts": func() string {
			return strconv.FormatInt(setting.LLM_CHAT_API.MAX_FREE_TRIES, 10)
		},
		"LlmExpireMinutes": func() string {
			return strconv.FormatInt(setting.LLM_CHAT_API.CHAT_EXPIRED_MINUTES, 10)
		},
		"AvatarLink":                 models.AvatarLink,
		"Safe":                       Safe,
		"SafeJS":                     SafeJS,
		"Str2html":                   Str2html,
		"subOne":                     subOne,
		"addOne":                     addOne,
		"TimeStampNow":               timeutil.TimeStampNow,
		"TimeSince":                  timeutil.TimeSince,
		"TimeSinceUnix":              timeutil.TimeSinceUnix,
		"TimeSinceUnix1":             timeutil.TimeSinceUnix1,
		"AttachmentResourceType":     dataset.GetResourceType,
		"AttachmentStatus":           dataset.GetStatusText,
		"IsShowDataSetOfCurrentRepo": dataset.IsShowDataSetOfCurrentRepo,
		"TimeSinceUnixShort":         timeutil.TimeSinceUnixShort,
		"ConvertDurationToStr":       ConvertDurationToStr,
		"RawTimeSince":               timeutil.RawTimeSince,
		"FileSize":                   base.FileSize,
		"PrettyNumber":               base.PrettyNumber,
		"Subtract":                   base.Subtract,
		"EntryIcon":                  base.EntryIcon,
		"MigrationIcon":              MigrationIcon,
		"Add": func(a, b int) int {
			return a + b
		},
		"ActionIcon": ActionIcon,
		"DateFmtLong": func(t time.Time) string {
			return t.Format(time.RFC1123Z)
		},
		"DateFmtShort": func(t time.Time) string {
			return t.Format("Jan 02, 2006")
		},
		"SizeFmt": base.FileSize,
		"List":    List,
		"SubStr": func(str string, start, length int) string {
			if len(str) == 0 {
				return ""
			}
			end := start + length
			if length == -1 {
				end = len(str)
			}
			if len(str) < end {
				return str
			}
			return str[start:end]
		},
		"EllipsisString":        base.EllipsisString,
		"DiffTypeToStr":         DiffTypeToStr,
		"DiffLineTypeToStr":     DiffLineTypeToStr,
		"Sha1":                  Sha1,
		"ShortSha":              base.ShortSha,
		"MD5":                   base.EncodeMD5,
		"ActionContent2Commits": ActionContent2Commits,
		"PathEscape":            url.PathEscape,
		"EscapePound": func(str string) string {
			return strings.NewReplacer("%", "%25", "#", "%23", " ", "%20", "?", "%3F").Replace(str)
		},
		"IpynbBool": func(str string) bool {
			return strings.Contains(str, ".ipynb")
		},
		"nl2br": func(text string) template.HTML {
			return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
		},
		"PathEscapeSegments":             util.PathEscapeSegments,
		"URLJoin":                        util.URLJoin,
		"RenderCommitMessage":            RenderCommitMessage,
		"RenderCommitMessageLink":        RenderCommitMessageLink,
		"RenderCommitMessageLinkSubject": RenderCommitMessageLinkSubject,
		"RenderCommitBody":               RenderCommitBody,
		"RenderEmoji":                    RenderEmoji,
		"RenderEmojiPlain":               emoji.ReplaceAliases,
		"ReactionToEmoji":                ReactionToEmoji,
		"RenderNote":                     RenderNote,
		"IsMultilineCommitMessage":       IsMultilineCommitMessage,
		"ThemeColorMetaTag": func() string {
			return setting.UI.ThemeColorMetaTag
		},
		"MetaAuthor": func() string {
			return setting.UI.Meta.Author
		},
		"MetaDescription": func() string {
			return setting.UI.Meta.Description
		},
		"MetaKeywords": func() string {
			return setting.UI.Meta.Keywords
		},
		"FilenameIsImage": func(filename string) bool {
			mimeType := mime.TypeByExtension(filepath.Ext(filename))
			return strings.HasPrefix(mimeType, "image/")
		},
		"TabSizeClass": func(ec *editorconfig.Editorconfig, filename string) string {
			if ec != nil {
				def, err := ec.GetDefinitionForFilename(filename)
				if err != nil {
					log.Error("tab size class: getting definition for filename: %v", err)
					return "tab-size-8"
				}
				if def.TabWidth > 0 {
					return fmt.Sprintf("tab-size-%d", def.TabWidth)
				}
			}
			return "tab-size-8"
		},
		"SubJumpablePath": func(str string) []string {
			var path []string
			index := strings.LastIndex(str, "/")
			if index != -1 && index != len(str) {
				path = append(path, str[0:index+1], str[index+1:])
			} else {
				path = append(path, str)
			}
			return path
		},
		"Json": func(in interface{}) string {
			out, err := json.Marshal(in)
			if err != nil {
				return ""
			}
			return string(out)
		},
		"JsonPrettyPrint": func(in string) string {
			var out bytes.Buffer
			err := json.Indent(&out, []byte(in), "", "  ")
			if err != nil {
				return ""
			}
			return out.String()
		},
		"DisableGitHooks": func() bool {
			return setting.DisableGitHooks
		},
		"DisableImportLocal": func() bool {
			return !setting.ImportLocalPaths
		},
		"TrN": TrN,
		"Dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, errors.New("invalid dict call")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
		"Printf":   fmt.Sprintf,
		"ToLower":  strings.ToLower,
		"Escape":   Escape,
		"Sec2Time": models.SecToTime,
		"ParseDeadline": func(deadline string) []string {
			return strings.Split(deadline, "|")
		},
		"DefaultTheme": func() string {
			return setting.UI.DefaultTheme
		},
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values) == 0 {
				return nil, errors.New("invalid dict call")
			}

			dict := make(map[string]interface{})

			for i := 0; i < len(values); i++ {
				switch key := values[i].(type) {
				case string:
					i++
					if i == len(values) {
						return nil, errors.New("specify the key for non array values")
					}
					dict[key] = values[i]
				case map[string]interface{}:
					m := values[i].(map[string]interface{})
					for i, v := range m {
						dict[i] = v
					}
				default:
					return nil, errors.New("dict values must be maps")
				}
			}
			return dict, nil
		},
		"percentage": func(n int, values ...int) float32 {
			var sum = 0
			for i := 0; i < len(values); i++ {
				sum += values[i]
			}
			return float32(n) * 100 / float32(sum)
		},
		"CommentMustAsDiff": gitdiff.CommentMustAsDiff,
		"MirrorAddress":     mirror_service.Address,
		"MirrorFullAddress": mirror_service.AddressNoCredentials,
		"MirrorUserName":    mirror_service.Username,
		"MirrorPassword":    mirror_service.Password,
		"CommitType": func(commit interface{}) string {
			switch commit.(type) {
			case models.SignCommitWithStatuses:
				return "SignCommitWithStatuses"
			case models.SignCommit:
				return "SignCommit"
			case models.UserCommit:
				return "UserCommit"
			default:
				return ""
			}
		},
		"NotificationSettings": func() map[string]int {
			return map[string]int{
				"MinTimeout":            int(setting.UI.Notification.MinTimeout / time.Millisecond),
				"TimeoutStep":           int(setting.UI.Notification.TimeoutStep / time.Millisecond),
				"MaxTimeout":            int(setting.UI.Notification.MaxTimeout / time.Millisecond),
				"EventSourceUpdateTime": int(setting.UI.Notification.EventSourceUpdateTime / time.Millisecond),
			}
		},
		"contain": func(s []int64, id int64) bool {
			for i := 0; i < len(s); i++ {
				if s[i] == id {
					return true
				}
			}
			return false
		},
		"svg": func(icon string, size int) template.HTML {
			return template.HTML(fmt.Sprintf(`<svg class="svg %s" width="%d" height="%d" aria-hidden="true"><use xlink:href="#%s" /></svg>`, icon, size, size, icon))
		},
		"tasks":      tasks,
		"categories": categories,
		"licenses":   licenses,
		"DatasetPathJoin": func(arr []string, index int, seq string) string {
			return strings.Join(arr[1:index+1], seq)
		},
		"GetRefType": GetRefType,
		"GetRefName": GetRefName,
		"MB2GB":      MB2GB,
		"encodeURIComponent": encodeURIComponent,

	}}
}

// NewTextFuncMap returns functions for injecting to text templates
// It's a subset of those used for HTML and other templates
func NewTextFuncMap() []texttmpl.FuncMap {
	return []texttmpl.FuncMap{map[string]interface{}{
		"GoVer": func() string {
			return strings.Title(runtime.Version())
		},
		"AppName": func() string {
			return setting.AppName
		},
		"AppSubUrl": func() string {
			return setting.AppSubURL
		},
		"AppUrl": func() string {
			return setting.AppURL
		},
		"AppVer": func() string {
			return setting.AppVer
		},
		"AppBuiltWith": func() string {
			return setting.AppBuiltWith
		},
		"AppDomain": func() string {
			return setting.Domain
		},
		"TimeStampNow":               timeutil.TimeStampNow,
		"TimeSince":                  timeutil.TimeSince,
		"TimeSinceUnix":              timeutil.TimeSinceUnix,
		"TimeSinceUnix1":             timeutil.TimeSinceUnix1,
		"TimeSinceUnixShort":         timeutil.TimeSinceUnixShort,
		"ConvertDurationToStr":       ConvertDurationToStr,
		"RawTimeSince":               timeutil.RawTimeSince,
		"AttachmentResourceType":     dataset.GetResourceType,
		"AttachmentStatus":           dataset.GetStatusText,
		"IsShowDataSetOfCurrentRepo": dataset.IsShowDataSetOfCurrentRepo,

		"DateFmtLong": func(t time.Time) string {
			return t.Format(time.RFC1123Z)
		},
		"DateFmtShort": func(t time.Time) string {
			return t.Format("Jan 02, 2006")
		},
		"List": List,
		"SubStr": func(str string, start, length int) string {
			if len(str) == 0 {
				return ""
			}
			end := start + length
			if length == -1 {
				end = len(str)
			}
			if len(str) < end {
				return str
			}
			return str[start:end]
		},
		"EllipsisString": base.EllipsisString,
		"URLJoin":        util.URLJoin,
		"Dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, errors.New("invalid dict call")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
		"Printf":   fmt.Sprintf,
		"ToLower":  strings.ToLower,
		"Escape":   Escape,
		"Sec2Time": models.SecToTime,
		"ParseDeadline": func(deadline string) []string {
			return strings.Split(deadline, "|")
		},
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values) == 0 {
				return nil, errors.New("invalid dict call")
			}

			dict := make(map[string]interface{})

			for i := 0; i < len(values); i++ {
				switch key := values[i].(type) {
				case string:
					i++
					if i == len(values) {
						return nil, errors.New("specify the key for non array values")
					}
					dict[key] = values[i]
				case map[string]interface{}:
					m := values[i].(map[string]interface{})
					for i, v := range m {
						dict[i] = v
					}
				default:
					return nil, errors.New("dict values must be maps")
				}
			}
			return dict, nil
		},
		"percentage": func(n int, values ...int) float32 {
			var sum = 0
			for i := 0; i < len(values); i++ {
				sum += values[i]
			}
			return float32(n) * 100 / float32(sum)
		},
	}}
}

// Safe render raw as HTML
func Safe(raw string) template.HTML {
	return template.HTML(raw)
}

// SafeJS renders raw as JS
func SafeJS(raw string) template.JS {
	return template.JS(raw)
}

// Str2html render Markdown text to HTML
func Str2html(raw string) template.HTML {
	return template.HTML(markup.Sanitize(raw))
}

//
func subOne(length int) int {
	return length - 1
}
func addOne(length int64) int64 {
	return length + 1
}

// Escape escapes a HTML string
func Escape(raw string) string {
	return html.EscapeString(raw)
}

// List traversings the list
func List(l *list.List) chan interface{} {
	e := l.Front()
	c := make(chan interface{})
	go func() {
		for e != nil {
			c <- e.Value
			e = e.Next()
		}
		close(c)
	}()
	return c
}

// Sha1 returns sha1 sum of string
func Sha1(str string) string {
	return base.EncodeSha1(str)
}

// RenderCommitMessage renders commit message with XSS-safe and special links.
func RenderCommitMessage(msg, urlPrefix string, metas map[string]string) template.HTML {
	return RenderCommitMessageLink(msg, urlPrefix, "", metas)
}

// RenderCommitMessageLink renders commit message as a XXS-safe link to the provided
// default url, handling for special links.
func RenderCommitMessageLink(msg, urlPrefix, urlDefault string, metas map[string]string) template.HTML {
	cleanMsg := template.HTMLEscapeString(msg)
	// we can safely assume that it will not return any error, since there
	// shouldn't be any special HTML.
	fullMessage, err := markup.RenderCommitMessage([]byte(cleanMsg), urlPrefix, urlDefault, metas)
	if err != nil {
		log.Error("RenderCommitMessage: %v", err)
		return ""
	}
	msgLines := strings.Split(strings.TrimSpace(string(fullMessage)), "\n")
	if len(msgLines) == 0 {
		return template.HTML("")
	}
	return template.HTML(msgLines[0])
}

// RenderCommitMessageLinkSubject renders commit message as a XXS-safe link to
// the provided default url, handling for special links without email to links.
func RenderCommitMessageLinkSubject(msg, urlPrefix, urlDefault string, metas map[string]string) template.HTML {
	msgLine := strings.TrimLeftFunc(msg, unicode.IsSpace)
	lineEnd := strings.IndexByte(msgLine, '\n')
	if lineEnd > 0 {
		msgLine = msgLine[:lineEnd]
	}
	msgLine = strings.TrimRightFunc(msgLine, unicode.IsSpace)
	if len(msgLine) == 0 {
		return template.HTML("")
	}

	// we can safely assume that it will not return any error, since there
	// shouldn't be any special HTML.
	renderedMessage, err := markup.RenderCommitMessageSubject([]byte(template.HTMLEscapeString(msgLine)), urlPrefix, urlDefault, metas)
	if err != nil {
		log.Error("RenderCommitMessageSubject: %v", err)
		return template.HTML("")
	}
	return template.HTML(renderedMessage)
}

// RenderCommitBody extracts the body of a commit message without its title.
func RenderCommitBody(msg, urlPrefix string, metas map[string]string) template.HTML {
	msgLine := strings.TrimRightFunc(msg, unicode.IsSpace)
	lineEnd := strings.IndexByte(msgLine, '\n')
	if lineEnd > 0 {
		msgLine = msgLine[lineEnd+1:]
	} else {
		return template.HTML("")
	}
	msgLine = strings.TrimLeftFunc(msgLine, unicode.IsSpace)
	if len(msgLine) == 0 {
		return template.HTML("")
	}

	renderedMessage, err := markup.RenderCommitMessage([]byte(template.HTMLEscapeString(msgLine)), urlPrefix, "", metas)
	if err != nil {
		log.Error("RenderCommitMessage: %v", err)
		return ""
	}
	return template.HTML(renderedMessage)
}

// RenderEmoji renders html text with emoji post processors
func RenderEmoji(text string) template.HTML {
	renderedText, err := markup.RenderEmoji([]byte(template.HTMLEscapeString(text)))
	if err != nil {
		log.Error("RenderEmoji: %v", err)
		return template.HTML("")
	}
	return template.HTML(renderedText)
}

//ReactionToEmoji renders emoji for use in reactions
func ReactionToEmoji(reaction string) template.HTML {
	val := emoji.FromCode(reaction)
	if val != nil {
		return template.HTML(val.Emoji)
	}
	val = emoji.FromAlias(reaction)
	if val != nil {
		return template.HTML(val.Emoji)
	}
	return template.HTML(fmt.Sprintf(`<img src=%s/img/emoji/%s.png></img>`, setting.StaticURLPrefix, reaction))
}

// RenderNote renders the contents of a git-notes file as a commit message.
func RenderNote(msg, urlPrefix string, metas map[string]string) template.HTML {
	cleanMsg := template.HTMLEscapeString(msg)
	fullMessage, err := markup.RenderCommitMessage([]byte(cleanMsg), urlPrefix, "", metas)
	if err != nil {
		log.Error("RenderNote: %v", err)
		return ""
	}
	return template.HTML(string(fullMessage))
}

// IsMultilineCommitMessage checks to see if a commit message contains multiple lines.
func IsMultilineCommitMessage(msg string) bool {
	return strings.Count(strings.TrimSpace(msg), "\n") >= 1
}

// Actioner describes an action
type Actioner interface {
	GetOpType() models.ActionType
	GetActUserName() string
	GetRepoUserName() string
	GetRepoName() string
	GetRepoPath() string
	GetRepoLink() string
	GetBranch() string
	GetContent() string
	GetCreate() time.Time
	GetIssueInfos() []string
}

// ActionIcon accepts an action operation type and returns an icon class name.
func ActionIcon(opType models.ActionType) string {
	switch opType {
	case models.ActionCreateRepo, models.ActionTransferRepo:
		return "repo"
	case models.ActionCommitRepo, models.ActionPushTag, models.ActionDeleteTag, models.ActionDeleteBranch:
		return "git-commit"
	case models.ActionCreateIssue:
		return "issue-opened"
	case models.ActionCreatePullRequest:
		return "git-pull-request"
	case models.ActionCommentIssue, models.ActionCommentPull:
		return "comment-discussion"
	case models.ActionMergePullRequest:
		return "git-merge"
	case models.ActionCloseIssue, models.ActionClosePullRequest:
		return "issue-closed"
	case models.ActionReopenIssue, models.ActionReopenPullRequest:
		return "issue-reopened"
	case models.ActionMirrorSyncPush, models.ActionMirrorSyncCreate, models.ActionMirrorSyncDelete:
		return "repo-clone"
	case models.ActionApprovePullRequest:
		return "check"
	case models.ActionRejectPullRequest:
		return "request-changes"
	default:
		return "question"
	}
}

// ActionContent2Commits converts action content to push commits
func ActionContent2Commits(act Actioner) *repository.PushCommits {
	push := repository.NewPushCommits()
	if err := json.Unmarshal([]byte(act.GetContent()), push); err != nil {
		log.Error("json.Unmarshal:\n%s\nERROR: %v", act.GetContent(), err)
	}
	return push
}

// DiffTypeToStr returns diff type name
func DiffTypeToStr(diffType int) string {
	diffTypes := map[int]string{
		1: "add", 2: "modify", 3: "del", 4: "rename",
	}
	return diffTypes[diffType]
}

// DiffLineTypeToStr returns diff line type name
func DiffLineTypeToStr(diffType int) string {
	switch diffType {
	case 2:
		return "add"
	case 3:
		return "del"
	case 4:
		return "tag"
	}
	return "same"
}

// Language specific rules for translating plural texts
var trNLangRules = map[string]func(int64) int{
	"en-US": func(cnt int64) int {
		if cnt == 1 {
			return 0
		}
		return 1
	},
	"lv-LV": func(cnt int64) int {
		if cnt%10 == 1 && cnt%100 != 11 {
			return 0
		}
		return 1
	},
	"ru-RU": func(cnt int64) int {
		if cnt%10 == 1 && cnt%100 != 11 {
			return 0
		}
		return 1
	},
	"zh-CN": func(cnt int64) int {
		return 0
	},
	"zh-HK": func(cnt int64) int {
		return 0
	},
	"zh-TW": func(cnt int64) int {
		return 0
	},
	"fr-FR": func(cnt int64) int {
		if cnt > -2 && cnt < 2 {
			return 0
		}
		return 1
	},
}

// TrN returns key to be used for plural text translation
func TrN(lang string, cnt interface{}, key1, keyN string) string {
	var c int64
	if t, ok := cnt.(int); ok {
		c = int64(t)
	} else if t, ok := cnt.(int16); ok {
		c = int64(t)
	} else if t, ok := cnt.(int32); ok {
		c = int64(t)
	} else if t, ok := cnt.(int64); ok {
		c = t
	} else {
		return keyN
	}

	ruleFunc, ok := trNLangRules[lang]
	if !ok {
		ruleFunc = trNLangRules["en-US"]
	}

	if ruleFunc(c) == 0 {
		return key1
	}
	return keyN
}

// MigrationIcon returns a Font Awesome name matching the service an issue/comment was migrated from
func MigrationIcon(hostname string) string {
	switch hostname {
	case "github.com":
		return "fa-github"
	default:
		return "fa-git-alt"
	}
}

func buildSubjectBodyTemplate(stpl *texttmpl.Template, btpl *template.Template, name string, content []byte) {
	// Split template into subject and body
	var subjectContent []byte
	bodyContent := content
	loc := mailSubjectSplit.FindIndex(content)
	if loc != nil {
		subjectContent = content[0:loc[0]]
		bodyContent = content[loc[1]:]
	}
	if _, err := stpl.New(name).
		Parse(string(subjectContent)); err != nil {
		log.Warn("Failed to parse template [%s/subject]: %v", name, err)
	}
	if _, err := btpl.New(name).
		Parse(string(bodyContent)); err != nil {
		log.Warn("Failed to parse template [%s/body]: %v", name, err)
	}
}

// Dataset categories
func categories() []string {
	return []string{"computer_vision", "natural_language_processing", "speech_processing", "computer_vision_natural_language_processing", "medical_imaging"}
}

func licenses() []string {
	return []string{"MIT", "GPL", "GFDL", "Apache 2.0", "CC-BY", "CC-BY-SA", "CC-BY-NC", "CC-BY-ND", "CC0: Public Domain", "CDLA Permissive 1.0", "Open Database License", "Research or commercial", "Research and commercial", "CC-BY-NC-ND", "CC-BY-NC-SA", "CC-BY 2.0", "CC-BY-SA 2.0", "GNU Free Documentation License", "CC-BY-NC 2.0", "CC-BY-ND 2.0", "CC-BY-NC-ND 2.0", "CC-BY-NC-SA 2.0", "Non-commercial", "CC-BY 3.0", "CC-BY-SA 3.0", "CC-BY-NC 3.0", "CC-BY-ND 3.0", "CC-BY-NC-ND 3.0", "CC-BY-NC-SA 3.0", "CC-BY 4.0", "CC-BY-SA 4.0", "CC-BY-NC 4.0", "CC-BY-ND 4.0", "CC-BY-NC-ND 4.0", "CC-BY-NC-SA 4.0"}
}

// Dataset tasks
func tasks() []string {
	return []string{"machine_translation", "question_answering_system", "information_retrieval", "knowledge_graph", "text_annotation", "text_categorization", "emotion_analysis", "language_modeling", "speech_recognition", "automatic_digest", "information_extraction", "description_generation", "image_classification", "face_recognition", "image_search", "target_detection", "image_description_generation", "vehicle_license_plate_recognition", "medical_image_analysis", "unmanned", "unmanned_security", "drone", "vr_ar", "2_d_vision", "2.5_d_vision", "3_d_reconstruction", "image_processing", "video_processing", "visual_input_system", "speech_coding", "speech_enhancement", "speech_synthesis", "ros_hmci_datasets"}
}

func GetRefType(ref string) string {
	if strings.HasPrefix(ref, REF_HEADS_PREFIX) {
		return REF_TYPE_BRANCH
	}
	if strings.HasPrefix(ref, REF_TAGS_PREFIX) {
		return REF_TYPE_TAG
	}
	return ""
}

func GetRefName(ref string) string {
	reg := regexp.MustCompile(REF_TYPE_PATTERN)
	return reg.ReplaceAllString(ref, "")
}

func MB2GB(size int) string {
	s := strconv.FormatFloat(float64(size)/float64(1024), 'f', 2, 64)
	for strings.HasSuffix(s, "0") {
		s = strings.TrimSuffix(s, "0")
	}
	if strings.HasSuffix(s, ".") {
		s = strings.TrimSuffix(s, ".")
	}
	return s
}
func ConvertDurationToStr(duration int64) string {
	if duration <= 0 {
		return DURATION_STR_ZERO
	}
	return util.AddZero(duration/3600) + ":" + util.AddZero(duration%3600/60) + ":" + util.AddZero(duration%60)
}
func encodeURIComponent(str string) string {
    r := url.QueryEscape(str)
    r = strings.Replace(r, "+", "%20", -1)
    return r
}