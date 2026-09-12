// Copyright 2015 The Gogs Authors. All rights reserved.
// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/go-git/go-git/v5/plumbing"
)

// BranchPrefix base dir of the branch information file store on git
const BranchPrefix = "refs/heads/"

// IsReferenceExist returns true if given reference exists in the repository.
func IsReferenceExist(repoPath, name string) bool {
	_, err := NewCommand("show-ref", "--verify", "--", name).RunInDir(repoPath)
	return err == nil
}

// IsBranchExist returns true if given branch exists in the repository.
func IsBranchExist(repoPath, name string) bool {
	return IsReferenceExist(repoPath, BranchPrefix+name)
}

// IsBranchExist returns true if given branch exists in current repository.
func (repo *Repository) IsBranchExist(name string) bool {
	if name == "" {
		return false
	}
	reference, err := repo.gogitRepo.Reference(plumbing.ReferenceName(BranchPrefix+name), true)
	if err != nil {
		return false
	}
	return reference.Type() != plumbing.InvalidReference
}

// Branch represents a Git branch.
type Branch struct {
	Name string
	Path string

	gitRepo *Repository
}

// GetHEADBranch returns corresponding branch of HEAD.
func (repo *Repository) GetHEADBranch() (*Branch, error) {
	if repo == nil {
		return nil, fmt.Errorf("nil repo")
	}
	stdout, err := NewCommand("symbolic-ref", "HEAD").RunInDir(repo.Path)
	if err != nil {
		return nil, err
	}
	stdout = strings.TrimSpace(stdout)

	if !strings.HasPrefix(stdout, BranchPrefix) {
		return nil, fmt.Errorf("invalid HEAD branch: %v", stdout)
	}

	return &Branch{
		Name:    stdout[len(BranchPrefix):],
		Path:    stdout,
		gitRepo: repo,
	}, nil
}

// SetDefaultBranch sets default branch of repository.
func (repo *Repository) SetDefaultBranch(name string) error {
	_, err := NewCommand("symbolic-ref", "HEAD", BranchPrefix+name).RunInDir(repo.Path)
	return err
}

// GetBranch returns a branch by it's name
func (repo *Repository) GetBranch(branch string) (*Branch, error) {
	if !repo.IsBranchExist(branch) {
		return nil, ErrBranchNotExist{branch}
	}
	return &Branch{
		Path:    repo.Path,
		Name:    branch,
		gitRepo: repo,
	}, nil
}

// GetBranchesByPath returns a branch by it's path
func GetBranchesByPath(path string, skip, limit int) ([]*Branch, int, error) {
	gitRepo, err := OpenRepository(path)
	if err != nil {
		return nil, 0, err
	}
	defer gitRepo.Close()

	brs, countAll, err := gitRepo.GetBranchNames(skip, limit)
	if err != nil {
		return nil, 0, err
	}

	branches := make([]*Branch, len(brs))
	for i := range brs {
		branches[i] = &Branch{
			Path:    path,
			Name:    brs[i],
			gitRepo: gitRepo,
		}
	}

	return branches, countAll, nil
}

// DeleteBranchOptions Option(s) for delete branch
type DeleteBranchOptions struct {
	Force bool
}

// DeleteBranch delete a branch by name on repository.
func (repo *Repository) DeleteBranch(name string, opts DeleteBranchOptions) error {
	cmd := NewCommand("branch")

	if opts.Force {
		cmd.AddArguments("-D")
	} else {
		cmd.AddArguments("-d")
	}

	cmd.AddArguments("--", name)
	_, err := cmd.RunInDir(repo.Path)

	return err
}

// CreateBranch create a new branch
func (repo *Repository) CreateBranch(branch, oldbranchOrCommit string) error {
	cmd := NewCommand("branch")
	cmd.AddArguments("--", branch, oldbranchOrCommit)

	_, err := cmd.RunInDir(repo.Path)

	return err
}

// AddRemote adds a new remote to repository.
func (repo *Repository) AddRemote(name, url string, fetch bool) error {
	cmd := NewCommand("remote", "add")
	if fetch {
		cmd.AddArguments("-f")
	}
	cmd.AddArguments(name, url)

	_, err := cmd.RunInDir(repo.Path)
	return err
}

// RemoveRemote removes a remote from repository.
func (repo *Repository) RemoveRemote(name string) error {
	_, err := NewCommand("remote", "rm", name).RunInDir(repo.Path)
	return err
}

// GetCommit returns the head commit of a branch
func (branch *Branch) GetCommit() (*Commit, error) {
	return branch.gitRepo.GetBranchCommit(branch.Name)
}

// GetBranches returns branches from the repository, skipping skip initial branches and
// returning at most limit branches, or all branches if limit is 0.
func (repo *Repository) GetBranches(skip, limit int) ([]string, int, error) {
	return callShowRef(repo.Path, BranchPrefix, "--heads", skip, limit)
}

func (repo *Repository) GetBranchNames(skip, limit int) ([]string, int, error) {
	return callShowRefNew(repo.Path, []string{BranchPrefix}, skip, limit)
}

func callShowRefNew(repoPath string, extraArgs []string, skip, limit int) (branchNames []string, countAll int, err error) {

	countAll, err = walkShowRef(repoPath, extraArgs, skip, limit, func(branchName string) error {
		branchNames = append(branchNames, branchName)
		return nil
	})
	return branchNames, countAll, err
}

func walkShowRef(repoPath string, extraArgs []string, skip, limit int, walkfn func(refname string) error) (countAll int, err error) {
	stdoutReader, stdoutWriter := io.Pipe()
	defer func() {
		_ = stdoutReader.Close()
		_ = stdoutWriter.Close()
	}()

	go func() {
		stderrBuilder := &strings.Builder{}
		args := []string{"for-each-ref", "--format=%(refname:short)"}
		args = append(args, extraArgs...)
		err := NewCommand(args...).RunInDirPipeline(repoPath, stdoutWriter, stderrBuilder)
		if err != nil {
			if stderrBuilder.Len() == 0 {
				_ = stdoutWriter.Close()
				return
			}
			_ = stdoutWriter.CloseWithError(ConcatenateError(err, stderrBuilder.String()))
		} else {
			_ = stdoutWriter.Close()
		}
	}()

	i := 0
	bufReader := bufio.NewReader(stdoutReader)
	for i < skip {
		_, isPrefix, err := bufReader.ReadLine()
		if err == io.EOF {
			return i, nil
		}
		if err != nil {
			return 0, err
		}
		if !isPrefix {
			i++
		}
	}
	for limit == 0 || i < skip+limit {
		// The output of show-ref is simply a list:
		// <sha> SP <ref> LF
		branchName, err := bufReader.ReadString('\n')
		if err == io.EOF {
			// This shouldn't happen... but we'll tolerate it for the sake of peace
			return i, nil
		}
		if err != nil {
			return i, err
		}

		if len(branchName) > 0 {
			branchName = branchName[:len(branchName)-1]
		}

		err = walkfn(branchName)
		if err != nil {
			return i, err
		}
		i++
	}
	// count all refs
	for limit != 0 {
		_, isPrefix, err := bufReader.ReadLine()
		if err == io.EOF {
			return i, nil
		}
		if err != nil {
			return 0, err
		}
		if !isPrefix {
			i++
		}
	}
	return i, nil
}

func isBranchMatch(branchName string, search string) bool {
	return search == "" || (search != "" && strings.Contains(branchName, search))
}

// callShowRef return refs, if limit = 0 it will not limit
func callShowRef(repoPath, prefix, arg string, skip, limit int) (branchNames []string, countAll int, err error) {
	stdoutReader, stdoutWriter := io.Pipe()
	defer func() {
		_ = stdoutReader.Close()
		_ = stdoutWriter.Close()
	}()

	go func() {
		stderrBuilder := &strings.Builder{}
		err := NewCommand("show-ref", arg).RunInDirPipeline(repoPath, stdoutWriter, stderrBuilder)
		if err != nil {
			if stderrBuilder.Len() == 0 {
				_ = stdoutWriter.Close()
				return
			}
			_ = stdoutWriter.CloseWithError(ConcatenateError(err, stderrBuilder.String()))
		} else {
			_ = stdoutWriter.Close()
		}
	}()

	i := 0
	bufReader := bufio.NewReader(stdoutReader)
	for i < skip {
		_, isPrefix, err := bufReader.ReadLine()
		if err == io.EOF {
			return branchNames, i, nil
		}
		if err != nil {
			return nil, 0, err
		}
		if !isPrefix {
			i++
		}
	}
	for limit == 0 || i < skip+limit {
		// The output of show-ref is simply a list:
		// <sha> SP <ref> LF
		_, err := bufReader.ReadSlice(' ')
		for err == bufio.ErrBufferFull {
			// This shouldn't happen but we'll tolerate it for the sake of peace
			_, err = bufReader.ReadSlice(' ')
		}
		if err == io.EOF {
			return branchNames, i, nil
		}
		if err != nil {
			return nil, 0, err
		}

		branchName, err := bufReader.ReadString('\n')
		if err == io.EOF {
			// This shouldn't happen... but we'll tolerate it for the sake of peace
			return branchNames, i, nil
		}
		if err != nil {
			return nil, i, err
		}
		branchName = strings.TrimPrefix(branchName, prefix)
		if len(branchName) > 0 {
			branchName = branchName[:len(branchName)-1]
		}
		branchNames = append(branchNames, branchName)
		i++
	}
	// count all refs
	for limit != 0 {
		_, isPrefix, err := bufReader.ReadLine()
		if err == io.EOF {
			return branchNames, i, nil
		}
		if err != nil {
			return nil, 0, err
		}
		if !isPrefix {
			i++
		}
	}
	return branchNames, i, nil
}
