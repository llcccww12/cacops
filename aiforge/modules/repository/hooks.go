// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repository

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"github.com/unknwon/com"

	"xorm.io/builder"
)

const (
	SIZE_LIMIT_SCRIPT_NAME = "size_limit"
)

func getHookTemplates() (hookNames, hookTpls, giteaHookTpls, sizeLimitTpls []string) {
	hookNames = []string{"pre-receive", "update", "post-receive"}
	hookTpls = []string{
		fmt.Sprintf("#!/usr/bin/env %s\ndata=$(cat)\nexitcodes=\"\"\nhookname=$(basename $0)\nGIT_DIR=${GIT_DIR:-$(dirname $0)}\n\nfor hook in ${GIT_DIR}/hooks/${hookname}.d/*; do\ntest -x \"${hook}\" && test -f \"${hook}\" || continue\necho \"${data}\" | \"${hook}\"\nexitcodes=\"${exitcodes} $?\"\ndone\n\nfor i in ${exitcodes}; do\n[ ${i} -eq 0 ] || exit ${i}\ndone\n", setting.ScriptType),
		fmt.Sprintf("#!/usr/bin/env %s\nexitcodes=\"\"\nhookname=$(basename $0)\nGIT_DIR=${GIT_DIR:-$(dirname $0)}\n\nfor hook in ${GIT_DIR}/hooks/${hookname}.d/*; do\ntest -x \"${hook}\" && test -f \"${hook}\" || continue\n\"${hook}\" $1 $2 $3\nexitcodes=\"${exitcodes} $?\"\ndone\n\nfor i in ${exitcodes}; do\n[ ${i} -eq 0 ] || exit ${i}\ndone\n", setting.ScriptType),
		fmt.Sprintf("#!/usr/bin/env %s\ndata=$(cat)\nexitcodes=\"\"\nhookname=$(basename $0)\nGIT_DIR=${GIT_DIR:-$(dirname $0)}\n\nfor hook in ${GIT_DIR}/hooks/${hookname}.d/*; do\ntest -x \"${hook}\" && test -f \"${hook}\" || continue\necho \"${data}\" | \"${hook}\"\nexitcodes=\"${exitcodes} $?\"\ndone\n\nfor i in ${exitcodes}; do\n[ ${i} -eq 0 ] || exit ${i}\ndone\n", setting.ScriptType),
	}
	giteaHookTpls = []string{
		fmt.Sprintf("#!/usr/bin/env %s\n\"%s\" hook --config='%s' pre-receive\n", setting.ScriptType, setting.AppPath, setting.CustomConf),
		fmt.Sprintf("#!/usr/bin/env %s\n\"%s\" hook --config='%s' update $1 $2 $3\n", setting.ScriptType, setting.AppPath, setting.CustomConf),
		fmt.Sprintf("#!/usr/bin/env %s\n\"%s\" hook --config='%s' post-receive\n", setting.ScriptType, setting.AppPath, setting.CustomConf),
	}
	sizeLimitTpls = []string{
		fmt.Sprintf("#!/usr/bin/env %s\n\n\nset -o pipefail\n\nreadonly DEFAULT_FILE_MAXSIZE_MB=\"30\" \nreadonly CONFIG_NAME=\"hooks.maxfilesize\"\nreadonly NULLSHA=\"0000000000000000000000000000000000000000\"\nreadonly EXIT_SUCCESS=0\nreadonly EXIT_FAILURE=1\nreadonly DEFAULT_REPO_MAXSIZE_MB=\"1024\" \nreadonly CHECK_FLAG_ON=1\n\n\nstatus=\"$EXIT_SUCCESS\"\n\n# skip this hook entirely if shell check is not open\ncheck_flag=${PUSH_SIZE_CHECK_FLAG}\nif [[ $check_flag != $CHECK_FLAG_ON ]]; then\nexit $EXIT_SUCCESS\nfi\n\n\n#######################################\n# check the file max size limit\n#######################################\n\n# get maximum filesize (from repository-specific config)\nmaxsize_mb=\"${REPO_MAX_FILE_SIZE}\"\n\nif [[ \"$?\" != $EXIT_SUCCESS ]]; then\necho \"failed to get ${CONFIG_NAME} from config\"\nexit \"$EXIT_FAILURE\"\nfi\n\npush_size=\"0\"\n# read lines from stdin (format: \"<oldref> <newref> <refname>\\n\")\nwhile read oldref newref refname; do\n# skip branch deletions\nif [[ \"$newref\" == \"$NULLSHA\" ]]; then\n  continue\nfi\n\n# find large objects\n# check all objects from $oldref (possible $NULLSHA) to $newref, but\n# skip all objects that have already been accepted (i.e. are referenced by\n# another branch or tag).\n\nnew_branch_flag=0\nif [[ \"$oldref\" == \"$NULLSHA\" ]]; then\n  target=\"$newref\"\n  new_branch_flag=1\n  echo \"You are creating a new remote branch,openI will check all files in commit history to find oversize files\"\nelse\n  target=\"${oldref}..${newref}\"\nfi\nmaxsize=`expr $maxsize_mb \\* 1048576` \n\n# find objects in this push_size\n# print like:\n# 08da8e2ab9ae4095bf94dd71ac913132b880b463 commit 214\n# 43e993b768ede5740e8c65de2ed6edec25053ea1 tree 185\n# 4476971d76569039df7569af1b8d03c288f6b193 blob 20167318 b0417e6593a1.zip\nfiles=\"$(git rev-list --objects \"$target\" | \\\n  git cat-file $'--batch-check=%%(objectname) %%(objecttype) %%(objectsize) %%(rest)' | \\\n      awk -F ' ' -v maxbytes=\"$maxsize\" 'BEGIN {totalIn=0} {if( $3 > maxbytes && $2 == \"blob\") { totalIn+=$3; print $4} else { totalIn+=$3}} END { printf (\"totalIn=\\t%%s\",totalIn)}' )\"\n  \nif [[ \"$?\" != $EXIT_SUCCESS ]]; then\n  echo \"failed to check for large files in ref ${refname}\"\n  continue\nfi\n\nIFS=$'\\n'\n# rewrite IFS to seperate line in $files\nfor file in $files; do\n  # if don't unset IFS,temp_array=(${file}) will get error answer\n \n  if [[ ${file} == totalIn=* ]]; then\n\tIFS=$'\\t'\n\ttemp_array=(${file})\n\tpush_size=${temp_array[1]}\n\tcontinue\n  fi\n\tunset IFS\n  if [[ \"$status\" == $EXIT_SUCCESS ]]; then\n\t\techo -e \"Error: Your push was rejected because it contains files larger than $(numfmt --to=iec \"$maxsize_mb\") Mb\"\n\t\techo \"help document -- https://openi.pcl.ac.cn/zeizei/OpenI_Learning/src/branch/master/docs/git/repository_capacity_help.md\"\n\t\techo \"oversize files:\"\n\t\tstatus=\"$EXIT_FAILURE\"\t\n  fi\n  echo -e \"\\033[31m- ${file}\\033[0m \"\ndone\n\nif [[ \"$status\" != $EXIT_SUCCESS ]]; then\n\texit \"$status\"\nfi\n\ndone\n\n#######################################\n# check the repo max size limit\n#######################################\nif [[ $push_size -eq \"0\" ]]; then\n\texit $EXIT_SUCCESS\nfi\n\n# if create new branch or tag,use count-objects -v to get pack size\nif [[ $new_branch_flag -eq 1 ]]; then\n    size_kb=`git count-objects -v | grep 'size-pack' | sed  's/.*\\(size-pack:\\).//'`\n    size_pack_kb=`git count-objects -v | grep 'size:' | sed  's/.*\\(size:\\).//'`\n\ttotal_kb=`expr $size_kb + $size_pack_kb`\n\tlet push_size=$total_kb*1024\nfi\n\nsizelimit_mb=\"${REPO_MAX_SIZE}\"\nlet sizelimit_b=$sizelimit_mb*1024*1024\n\n# repo size at here means the size of repo directory in server \nreposize_b=${REPO_CURRENT_SIZE}\n\ntotal=`expr $push_size + $reposize_b`\n\nif [ $total -gt $sizelimit_b ]; then\n    echo \"Error: Your push was rejected because the repository size is large than $sizelimit_mb Mb\"\n    echo \"see the help document--https://openi.pcl.ac.cn/zeizei/OpenI_Learning/src/branch/master/docs/git/repository_capacity_help.md\"\n    exit $EXIT_FAILURE\nfi\n\n\nexit $EXIT_SUCCESS", setting.ScriptType),
		fmt.Sprintf(""),
		fmt.Sprintf(""),
	}
	return
}

// CreateDelegateHooks creates all the hooks scripts for the repo
func CreateDelegateHooks(repoPath string) error {
	return createDelegateHooks(repoPath)
}

// createDelegateHooks creates all the hooks scripts for the repo
func createDelegateHooks(repoPath string) (err error) {
	hookNames, hookTpls, giteaHookTpls, sizeLimitTpls := getHookTemplates()
	hookDir := filepath.Join(repoPath, "hooks")

	for i, hookName := range hookNames {
		oldHookPath := filepath.Join(hookDir, hookName)
		newHookPath := filepath.Join(hookDir, hookName+".d", "gitea")

		if err := os.MkdirAll(filepath.Join(hookDir, hookName+".d"), os.ModePerm); err != nil {
			return fmt.Errorf("create hooks dir '%s': %v", filepath.Join(hookDir, hookName+".d"), err)
		}

		// WARNING: This will override all old server-side hooks
		if err = os.Remove(oldHookPath); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("unable to pre-remove old hook file '%s' prior to rewriting: %v ", oldHookPath, err)
		}
		if err = ioutil.WriteFile(oldHookPath, []byte(hookTpls[i]), 0777); err != nil {
			return fmt.Errorf("write old hook file '%s': %v", oldHookPath, err)
		}

		if err = ensureExecutable(oldHookPath); err != nil {
			return fmt.Errorf("Unable to set %s executable. Error %v", oldHookPath, err)
		}

		if err = os.Remove(newHookPath); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("unable to pre-remove new hook file '%s' prior to rewriting: %v", newHookPath, err)
		}
		if err = ioutil.WriteFile(newHookPath, []byte(giteaHookTpls[i]), 0777); err != nil {
			return fmt.Errorf("write new hook file '%s': %v", newHookPath, err)
		}

		if err = ensureExecutable(newHookPath); err != nil {
			return fmt.Errorf("Unable to set %s executable. Error %v", oldHookPath, err)
		}

		if err = writeHookTpl(generateHookScriptPath(hookDir, hookName, SIZE_LIMIT_SCRIPT_NAME), sizeLimitTpls[i]); err != nil {
			return err
		}
	}

	return nil
}

func writeHookTpl(hookPath, content string) error {
	if content == "" {
		return nil
	}
	if err := ioutil.WriteFile(hookPath, []byte(content), 0777); err != nil {
		return fmt.Errorf("write new hook file '%s': %v", hookPath, err)
	}

	if err := ensureExecutable(hookPath); err != nil {
		return fmt.Errorf("Unable to set %s executable. Error %v", hookPath, err)
	}
	return nil
}

func checkExecutable(filename string) bool {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return (fileInfo.Mode() & 0100) > 0
}

func ensureExecutable(filename string) error {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return err
	}
	if (fileInfo.Mode() & 0100) > 0 {
		return nil
	}
	mode := fileInfo.Mode() | 0100
	return os.Chmod(filename, mode)
}

// CheckDelegateHooks checks the hooks scripts for the repo
func CheckDelegateHooks(repoPath string) ([]string, error) {
	hookNames, hookTpls, giteaHookTpls, sizeLimitTpls := getHookTemplates()

	hookDir := filepath.Join(repoPath, "hooks")
	results := make([]string, 0, 10)

	for i, hookName := range hookNames {
		oldHookPath := filepath.Join(hookDir, hookName)
		newHookPath := filepath.Join(hookDir, hookName+".d", "gitea")

		cont := false
		if !com.IsExist(oldHookPath) {
			results = append(results, fmt.Sprintf("old hook file %s does not exist", oldHookPath))
			cont = true
		}
		if !com.IsExist(oldHookPath + ".d") {
			results = append(results, fmt.Sprintf("hooks directory %s does not exist", oldHookPath+".d"))
			cont = true
		}
		if !com.IsExist(newHookPath) {
			results = append(results, fmt.Sprintf("new hook file %s does not exist", newHookPath))
			cont = true
		}
		if cont {
			continue
		}
		contents, err := ioutil.ReadFile(oldHookPath)
		if err != nil {
			return results, err
		}
		if string(contents) != hookTpls[i] {
			results = append(results, fmt.Sprintf("old hook file %s is out of date", oldHookPath))
		}
		if !checkExecutable(oldHookPath) {
			results = append(results, fmt.Sprintf("old hook file %s is not executable", oldHookPath))
		}
		contents, err = ioutil.ReadFile(newHookPath)
		if err != nil {
			return results, err
		}
		if string(contents) != giteaHookTpls[i] {
			results = append(results, fmt.Sprintf("new hook file %s is out of date", newHookPath))
		}
		if !checkExecutable(newHookPath) {
			results = append(results, fmt.Sprintf("new hook file %s is not executable", newHookPath))
		}
		if err = checkHookFile(generateHookScriptPath(hookDir, hookName, SIZE_LIMIT_SCRIPT_NAME), sizeLimitTpls[i], results); err != nil {
			return results, err
		}
	}
	return results, nil
}

func generateHookScriptPath(hookDir, hookName, fileName string) string {
	return filepath.Join(hookDir, hookName+".d", fileName)
}

func checkHookFile(filePath, tpl string, results []string) error {
	if tpl == "" {
		return nil
	}
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	if string(contents) != tpl {
		results = append(results, fmt.Sprintf("old hook file %s is out of date", filePath))
	}
	if !checkExecutable(filePath) {
		results = append(results, fmt.Sprintf("old hook file %s is not executable", filePath))
	}
	return nil
}

// SyncRepositoryHooks rewrites all repositories' pre-receive, update and post-receive hooks
// to make sure the binary and custom conf path are up-to-date.
func SyncRepositoryHooks(ctx context.Context) error {
	log.Trace("Doing: SyncRepositoryHooks")

	if err := models.Iterate(
		models.DefaultDBContext(),
		new(models.Repository),
		builder.Gt{"id": 0},
		func(idx int, bean interface{}) error {
			repo := bean.(*models.Repository)
			select {
			case <-ctx.Done():
				return models.ErrCancelledf("before sync repository hooks for %s", repo.FullName())
			default:
			}

			if err := createDelegateHooks(repo.RepoPath()); err != nil {
				return fmt.Errorf("SyncRepositoryHook: %v", err)
			}
			if repo.HasWiki() {
				if err := createDelegateHooks(repo.WikiPath()); err != nil {
					return fmt.Errorf("SyncRepositoryHook: %v", err)
				}
			}
			return nil
		},
	); err != nil {
		return err
	}

	log.Trace("Finished: SyncRepositoryHooks")
	return nil
}
