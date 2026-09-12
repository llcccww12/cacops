package git

import (
	"bufio"
	"bytes"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	Log "code.gitea.io/gitea/modules/log"
)

type RepoKPIStats struct {
	Contributors        int64
	KeyContributors     int64
	DevelopAge          int64
	ContributorsAdded   int64
	Commits             int64
	CommitsAdded        int64
	CommitLinesModified int64
	WikiPages           int64
	Authors             []*UserKPITypeStats
}

type UserKPIStats struct {
	Name        string
	Email       string
	Commits     int64
	CommitLines int64
}
type UserKPITypeStats struct {
	UserKPIStats
	isNewContributor bool //是否是4个月内的新增贡献者
}

func SetDevelopAge(repoPath string, stats *RepoKPIStats, fromTime time.Time) error {
	since := fromTime.Format(time.RFC3339)
	args := []string{"log", "--no-merges", "--branches=*", "--format=%cd", "--date=short", fmt.Sprintf("--since='%s'", since)}
	stdout, err := NewCommand(args...).RunInDirBytes(repoPath)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(bytes.NewReader(stdout))
	scanner.Split(bufio.ScanLines)
	developMonth := make(map[string]struct{})
	for scanner.Scan() {
		l := strings.TrimSpace(scanner.Text())
		month := l[0:strings.LastIndex(l, "-")]
		if _, ok := developMonth[month]; !ok {
			developMonth[month] = struct{}{}
		}
	}

	stats.DevelopAge = int64(len(developMonth))
	return nil
}

func GetUserKPIStats(repoPath string, startTime time.Time, endTime time.Time) (map[string]*UserKPIStats, error) {

	after := startTime.Format(time.RFC3339)
	until := endTime.Format(time.RFC3339)
	args := []string{"log", "--numstat", "--no-merges", "--branches=*", "--pretty=format:---%n%h%n%an%n%ae%n", "--date=iso", fmt.Sprintf("--after='%s'", after), fmt.Sprintf("--until='%s'", until)}
	stdout, err := NewCommand(args...).RunInDirBytes(repoPath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(bytes.NewReader(stdout))
	scanner.Split(bufio.ScanLines)
	usersKPIStatses := make(map[string]*UserKPIStats)
	var author string
	p := 0
	var email string
	for scanner.Scan() {
		l := strings.TrimSpace(scanner.Text())
		if l == "---" {
			p = 1
		} else if p == 0 {
			continue
		} else {
			p++
		}
		if p > 4 && len(l) == 0 {
			continue
		}
		switch p {
		case 1: // Separator
		case 2: // Commit sha-1
		case 3: // Author
			author = l
		case 4: // E-mail
			email = strings.ToLower(l)
			if _, ok := usersKPIStatses[email]; !ok {
				usersKPIStatses[email] = &UserKPIStats{
					Name:        author,
					Email:       email,
					Commits:     0,
					CommitLines: 0,
				}
			}
			usersKPIStatses[email].Commits++
		default: // Changed file
			//Log.Info("code commit file=" + l)
			if parts := strings.Fields(l); len(parts) >= 3 {
				if isCodeFile(parts[2]) {
					if parts[0] != "-" {
						if c, err := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64); err == nil {
							usersKPIStatses[email].CommitLines += c
						}
					}
					if parts[1] != "-" {
						if c, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64); err == nil {
							usersKPIStatses[email].CommitLines += c
						}
					}
					if usersKPIStatses[email].CommitLines > 100000 {
						Log.Info("count code more than 100000, email=" + email + " lines=" + fmt.Sprint(usersKPIStatses[email].CommitLines) + " repo=" + repoPath)
					}
				}
			}
		}
	}

	return usersKPIStatses, nil

}

var (
	codePostfix []string = []string{".py", ".go", ".js", ".html", ".css", ".c", ".cpp", ".h", ".hpp", ".java", ".sql", ".php", ".rb", ".vue", ".yml", ".swift", ".kt", ".sh", ".pl", ".rs", ".lua", ".ts", ".m"}
)

func isCodeFile(name string) bool {
	lowerName := strings.ToLower(name)
	for _, v := range codePostfix {
		if strings.HasSuffix(lowerName, v) {
			return true
		}
	}
	return false
}

func SetRepoKPIStats(repoPath string, fromTime time.Time, stats *RepoKPIStats, newContributers map[string]struct{}) error {
	since := fromTime.Format(time.RFC3339)
	args := []string{"log", "--numstat", "--no-merges", "HEAD", "--pretty=format:---%n%h%n%an%n%ae%n", "--date=iso", fmt.Sprintf("--since='%s'", since)}

	stdout, err := NewCommand(args...).RunInDirBytes(repoPath)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(bytes.NewReader(stdout))
	scanner.Split(bufio.ScanLines)

	authors := make(map[string]*UserKPITypeStats)

	var author string
	p := 0
	var email string
	for scanner.Scan() {
		l := strings.TrimSpace(scanner.Text())
		if l == "---" {
			p = 1
		} else if p == 0 {
			continue
		} else {
			p++
		}
		if p > 4 && len(l) == 0 {
			continue
		}
		switch p {
		case 1: // Separator
		case 2: // Commit sha-1
			stats.CommitsAdded++
		case 3: // Author
			author = l
		case 4: // E-mail
			email = strings.ToLower(l)
			if _, ok := authors[email]; !ok {
				authors[email] = &UserKPITypeStats{
					UserKPIStats: UserKPIStats{
						Name:        author,
						Email:       email,
						Commits:     0,
						CommitLines: 0,
					},
					isNewContributor: false,
				}
			}
			if _, ok := newContributers[email]; ok {
				authors[email].isNewContributor = true
			}

			authors[email].Commits++
		default: // Changed file
			if parts := strings.Fields(l); len(parts) >= 3 {
				if parts[0] != "-" {
					if c, err := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64); err == nil {
						stats.CommitLinesModified += c
						authors[email].CommitLines += c
					}
				}
				if parts[1] != "-" {
					if c, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64); err == nil {
						stats.CommitLinesModified += c
						authors[email].CommitLines += c
					}
				}

			}
		}
	}

	a := make([]*UserKPITypeStats, 0, len(authors))
	for _, v := range authors {
		a = append(a, v)
	}
	// Sort authors descending depending on commit count
	sort.Slice(a, func(i, j int) bool {
		return a[i].Commits > a[j].Commits
	})

	stats.Authors = a
	return nil

}

func GetContributorsDetail(repoPath string, fromTime time.Time) ([]Contributor, error) {
	since := fromTime.Format(time.RFC3339)
	cmd := NewCommand("shortlog", "-sne", "HEAD", fmt.Sprintf("--since='%s'", since))
	stdout, err := cmd.RunInDir(repoPath)
	if err != nil {
		return nil, err
	}
	stdout = strings.Trim(stdout, "\n")
	contributorRows := strings.Split(stdout, "\n")
	if len(contributorRows) > 0 {
		contributorsInfo := make([]Contributor, len(contributorRows))
		for i := 0; i < len(contributorRows); i++ {
			var oneCount string = strings.Trim(contributorRows[i], " ")
			if strings.Index(oneCount, "\t") < 0 {
				continue
			}
			number := oneCount[0:strings.Index(oneCount, "\t")]
			commitCnt, _ := strconv.Atoi(number)
			committer := oneCount[strings.Index(oneCount, "\t")+1 : strings.LastIndex(oneCount, " ")]
			committer = strings.Trim(committer, " ")
			email := oneCount[strings.Index(oneCount, "<")+1 : strings.Index(oneCount, ">")]
			contributorsInfo[i] = Contributor{
				commitCnt, committer, email,
			}
		}
		return contributorsInfo, nil
	}
	return nil, nil
}

func SetWikiPages(wikiPath string, stats *RepoKPIStats) {
	wikiPages := 0

	if wikiPath == "" {
		stats.WikiPages = int64(wikiPages)
		return
	}

	wikiRepo, commit, err := findWikiRepoCommit(wikiPath)
	if err != nil {
		if !IsErrNotExist(err) {
			Log.Warn("GetBranchCommit", err)
		}
		stats.WikiPages = int64(wikiPages)
		return
	}

	// Get page list.
	entries, err := commit.ListEntries()
	if err != nil {
		if wikiRepo != nil {
			wikiRepo.Close()
		}
		Log.Warn("GetBranchCommit", err)
		stats.WikiPages = int64(wikiPages)
		return

	}

	for _, entry := range entries {
		if !entry.IsRegular() {
			continue
		}

		wikiName, err := filenameToName(entry.Name())
		if err != nil || wikiName == "_Sidebar" || wikiName == "_Footer" {
			continue
		}

		wikiPages += 1

	}
	//确保wikiRepo用完被关闭
	defer func() {
		if wikiRepo != nil {
			wikiRepo.Close()
		}
	}()
	stats.WikiPages = int64(wikiPages)
	return

}

func filenameToName(filename string) (string, error) {
	if !strings.HasSuffix(filename, ".md") {
		return "", fmt.Errorf("invalid file")
	}
	basename := filename[:len(filename)-3]
	unescaped, err := url.QueryUnescape(basename)
	if err != nil {
		return "", err
	}
	return strings.Replace(unescaped, "-", " ", -1), nil
}

func findWikiRepoCommit(wikiPath string) (*Repository, *Commit, error) {
	wikiRepo, err := OpenRepository(wikiPath)
	if err != nil {

		return nil, nil, err
	}

	commit, err := wikiRepo.GetBranchCommit("master")
	if err != nil {
		return wikiRepo, nil, err
	}
	return wikiRepo, commit, nil
}
