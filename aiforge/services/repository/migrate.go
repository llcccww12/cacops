package repository

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/auth"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/migrations"
	"code.gitea.io/gitea/modules/setting"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/task"
	"code.gitea.io/gitea/modules/util"
	"code.gitea.io/gitea/routers/response"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func MigrateSubmit(currentUser *models.User, form auth.MigrateRepoForm) *response.BizError {
	log.Info("receive MigrateSubmit request")
	ctxUser, bizErr := checkMigrateUser(currentUser, form.UID)
	if bizErr != nil {
		return bizErr
	}

	remoteAddr, err := form.ParseRemoteAddr(currentUser)
	if err != nil {
		if models.IsErrInvalidCloneAddr(err) {
			addrErr := err.(models.ErrInvalidCloneAddr)
			switch {
			case addrErr.IsURLError:
				return response.PARAM_ERROR
			case addrErr.IsPermissionDenied:
				return response.INSUFFICIENT_PERMISSION
			case addrErr.IsInvalidPath:
				return response.PARAM_ERROR
			default:
			}
		}
		return response.NewBizError(err)
	}

	var gitServiceType = api.PlainGitService
	u, err := url.Parse(form.CloneAddr)
	if err == nil && strings.EqualFold(u.Host, "github.com") {
		gitServiceType = api.GithubService
	}

	var opts = migrations.MigrateOptions{
		OriginalURL:    form.CloneAddr,
		GitServiceType: gitServiceType,
		CloneAddr:      remoteAddr,
		RepoName:       form.RepoName,
		Alias:          form.Alias,
		Description:    form.Description,
		Private:        form.Private || setting.Repository.ForcePrivate,
		Mirror:         form.Mirror,
		AuthUsername:   form.AuthUsername,
		AuthPassword:   form.AuthPassword,
		Wiki:           form.Wiki,
		Issues:         form.Issues,
		Milestones:     form.Milestones,
		Labels:         form.Labels,
		Comments:       true,
		PullRequests:   form.PullRequests,
		Releases:       form.Releases,
	}
	if opts.Mirror {
		opts.Issues = false
		opts.Milestones = false
		opts.Labels = false
		opts.Comments = false
		opts.PullRequests = false
		opts.Releases = false
	}

	err = models.CheckCreateRepository(currentUser, ctxUser, opts.RepoName, opts.Alias)
	if err != nil {
		return handleMigrateError4Api(ctxUser, remoteAddr, err)
	}

	err = task.MigrateRepository(currentUser, ctxUser, opts)
	if err != nil {
		return handleMigrateError4Api(ctxUser, remoteAddr, err)
	}
	return nil
}

func checkMigrateUser(currentUser *models.User, uid int64) (*models.User, *response.BizError) {
	if uid == currentUser.ID || uid == 0 {
		return currentUser, nil
	}

	org, err := models.GetUserByID(uid)
	if models.IsErrUserNotExist(err) {
		return currentUser, nil
	}

	if err != nil {
		return nil, response.SYSTEM_ERROR
	}

	// Check ownership of organization.
	if !org.IsOrganization() {
		return nil, nil
	}
	if !currentUser.IsAdmin {
		canCreate, err := org.CanCreateOrgRepo(currentUser.ID)
		if err != nil {
			return nil, response.NewBizError(err)
		} else if !canCreate {
			return nil, response.INSUFFICIENT_PERMISSION
		}
	}
	return org, nil
}

func handleMigrateError4Api(repoOwner *models.User, remoteAddr string, err error) *response.BizError {
	switch {
	case models.IsErrRepoAlreadyExist(err):
		return response.BuildBizError(3, "The repository with the same name already exists.")
	case migrations.IsRateLimitError(err):
		return response.NewBizError(errors.New("Remote visit addressed rate limitation."))
	case migrations.IsTwoFactorAuthError(err):
		return response.NewBizError(errors.New("Remote visit required two factors authentication."))
	case models.IsErrReachLimitOfRepo(err):
		return response.NewBizError(errors.New(fmt.Sprintf("You have already reached your limit of %d repositories.", repoOwner.MaxCreationLimit())))
	case models.IsErrNameReserved(err):
		return response.NewBizError(errors.New(fmt.Sprintf("The username '%s' is reserved.", err.(models.ErrNameReserved).Name)))
	case models.IsErrNameCharsNotAllowed(err):
		return response.NewBizError(errors.New(fmt.Sprintf("The username '%s' contains invalid characters.", err.(models.ErrNameCharsNotAllowed).Name)))
	case models.IsErrNamePatternNotAllowed(err):
		return response.NewBizError(errors.New(fmt.Sprintf("The pattern '%s' is not allowed in a username.", err.(models.ErrNamePatternNotAllowed).Pattern)))

	default:
		err = util.URLSanitizedError(err, remoteAddr)
		if strings.Contains(err.Error(), "Authentication failed") ||
			strings.Contains(err.Error(), "Bad credentials") ||
			strings.Contains(err.Error(), "could not read Username") {
			return response.NewBizError(errors.New((fmt.Sprintf("Authentication failed: %v.", err))))
		} else if strings.Contains(err.Error(), "fatal:") {
			return response.NewBizError(errors.New((fmt.Sprintf("Migration failed: %v.", err))))
		}
	}
	return response.NewBizError(err)
}
