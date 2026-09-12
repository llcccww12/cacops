package auth

import (
	"gitea.com/macaron/binding"
	"gitea.com/macaron/macaron"
)

// CreateDatasetForm form for dataset page
type CreateDatasetForm struct {
	Title       string `binding:"Required"`
	Category    string `binding:"Required"`
	Description string `binding:"Required"`
	License     string `binding:"MaxSize(64)"`
	Task        string `binding:"Required;MaxSize(64)"`
	ReleaseID   int64  `xorm:"INDEX"`
	Files       []string
}

func (f *CreateDatasetForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return validate(errs, ctx.Data, f, ctx.Locale)
}

type EditDatasetForm struct {
	ID          int64  `binding:"Required"`
	Title       string `binding:"Required"`
	Category    string `binding:"Required"`
	Description string `binding:"Required"`
	License     string `binding:"Required;MaxSize(64)"`
	Task        string `binding:"Required;MaxSize(64)"`
	ReleaseID   int64  `xorm:"INDEX"`
	Files       []string
	Type        string `binding:"Required"`
}

func (f *EditDatasetForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return validate(errs, ctx.Data, f, ctx.Locale)
}

type EditAttachmentForm struct {
	ID          int64 `binding:"Required"`
	Description string
}

func (f *EditAttachmentForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return validate(errs, ctx.Data, f, ctx.Locale)
}

type ReferenceDatasetForm struct {
	DatasetID []int64 `binding:"Required"`
}

func (f *ReferenceDatasetForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return validate(errs, ctx.Data, f, ctx.Locale)
}
