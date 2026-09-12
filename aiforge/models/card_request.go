package models

import (
	"errors"
	"strings"

	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

const CARD_REQUEST_COMMIT = 1
const CARD_REQUEST_AGREE = 2
const CARD_REQEST_DISAGREE = 3

const RESOURCE_TYPE_SHARE = 1     //共享
const RESOURCE_TYPE_EXCLUSIVE = 2 //独占

const OrderByIDDesc = "card_request.id desc"
const OrderByStatus = "card_request.status asc,card_request.id desc"

type CardRequest struct {
	ID                int64 `xorm:"pk autoincr"`
	ComputeResource   string
	UID               int64
	UserName          string `xorm:"-"`
	CardType          string
	AccCardsNum       string
	DiskCapacity      int64
	ResourceType      int
	BeginDate         string
	BeginUnix         int64 `xorm:"INDEX"`
	EndDate           string
	EndUnix           int64 `xorm:"INDEX"`
	Contact           string
	PhoneNumber       string
	EmailAddress      string
	Wechat            string
	IsResearchProject bool
	InstitutionName   string
	ProjectName       string
	ProjectCode       string
	Org               string `xorm:"varchar(500)"`
	Description       string `xorm:"varchar(3000)"`
	Status            int
	Review            string `xorm:"varchar(3000)"`
	CreatedUnix       int64  `xorm:"INDEX created"`
	UpdatedUnix       int64  `xorm:"INDEX updated"`
	DeleteUnix        int64  `xorm:"deleted"`
}

type CardRequestSpecRes struct {
	ID                int64
	ComputeResource   string
	UID               int64
	UserName          string
	CardType          string
	AccCardsNum       string
	DiskCapacity      int64
	ResourceType      int
	BeginDate         string
	BeginUnix         int64
	EndDate           string
	EndUnix           int64
	Contact           string
	PhoneNumber       string
	EmailAddress      string
	Wechat            string
	IsResearchProject bool
	InstitutionName   string
	ProjectName       string
	ProjectCode       string
	Org               string
	Description       string
	Status            int
	Review            string
	CreatedUnix       int64
	UpdatedUnix       int64
	DeleteUnix        int64
	Specs             []RequestSpecInfo
}

func (CardRequestSpecRes) TableName() string {
	return "card_request"
}

type CardRequestSpec struct {
	ID          int64              `xorm:"pk autoincr"`
	RequestId   int64              `xorm:"unique(idx_request_spec)"`
	SpecId      int64              `xorm:"unique(idx_request_spec)"`
	CreatedTime timeutil.TimeStamp `xorm:"created"`
}

type CardRequestOptions struct {
	ListOptions
	UserID  int64
	OrderBy string
	Keyword string

	AiCenterCode      string
	QueueId           int64
	ComputeResource   string
	AccCardType       string
	Cluster           string
	ResourceType      int
	UseBeginTime      int64
	UseEndTime        int64
	BeginTimeUnix     int64
	EndTimeUnix       int64
	NeedSpec          bool
	IsResearchProject int
}
type CardRequestShowList struct {
	Total           int64                  `json:"total"`
	CardRequestList []*CardRequestSpecShow `json:"cardRequestList"`
}

type CardRequestSpecShow struct {
	ID                int64             `json:"id"`
	ComputeResource   string            `json:"compute_resource"`
	CardType          string            `json:"card_type"`
	AccCardsNum       string            `json:"acc_cards_num"`
	BeginDate         string            `json:"begin_date"`
	EndDate           string            `json:"end_date"`
	ResourceType      int               `json:"resource_type"`
	DiskCapacity      int64             `json:"disk_capacity"`
	UID               int64             `json:"uid"`
	UserName          string            `json:"user_name"`
	TargetCenter      []string          `json:"target_center"`
	Contact           string            `json:"contact"`
	PhoneNumber       string            `json:"phone_number"`
	EmailAddress      string            `json:"email_address"`
	Wechat            string            `json:"wechat"`
	IsResearchProject bool              `json:"is_research_project"`
	InstitutionName   string            `json:"institution_name"`
	ProjectName       string            `json:"project_name"`
	ProjectCode       string            `json:"project_code"`
	Org               string            `json:"org"`
	Description       string            `json:"description"`
	Status            int               `json:"status"`
	Review            string            `json:"review"`
	CreatedUnix       int64             `json:"created_unix"`
	Specs             []RequestSpecInfo `json:"specs"`
}

type RequestSpecInfo struct {
	ID           int64
	SourceSpecId string
	AccCardsNum  int
	CpuCores     int
	MemGiB       float32
	GPUMemGiB    float32
	ShareMemGiB  float32
	UnitPrice    float64
	Status       int
	UpdatedTime  timeutil.TimeStamp
	RequestId    int64
	//queue
	Cluster         string
	AiCenterCode    string
	AiCenterName    string
	QueueCode       string
	QueueName       string
	QueueType       string
	QueueId         int64
	ComputeResource string
	AccCardType     string
	HasInternet     int
}

func (RequestSpecInfo) TableName() string {
	return "resource_specification"
}

type CardRequestReview struct {
	ID      int64
	Review  string
	SpecIds []int64
}

func AgreeCardRequest(r CardRequestReview) error {
	sess := x.NewSession()
	var err error
	defer func() {
		if err != nil {
			sess.Rollback()
		}
		sess.Close()
	}()

	// find old scene
	old := CardRequest{}
	if has, _ := sess.ID(r.ID).Get(&old); !has {
		return errors.New("CardRequest not exist")
	}
	//check specification
	specs := make([]ResourceSpecification, 0)
	cond := builder.In("id", r.SpecIds)
	if err := sess.Where(cond).Find(&specs); err != nil {
		return err
	}
	if len(specs) < len(r.SpecIds) {
		return errors.New("specIds not correct")
	}

	rs := CardRequest{
		Status: CARD_REQUEST_AGREE,
		Review: "",
	}
	if _, err = sess.ID(r.ID).Cols("status", "review").Update(&rs); err != nil {
		return err
	}

	//delete scene spec relation
	if _, err = sess.Where("request_id = ? ", r.ID).Delete(&CardRequestSpec{}); err != nil {
		sess.Rollback()
		return err
	}

	if len(r.SpecIds) == 0 {
		return sess.Commit()
	}
	//build new scene spec relation
	rss := make([]CardRequestSpec, len(r.SpecIds))
	for i, v := range r.SpecIds {
		rss[i] = CardRequestSpec{
			RequestId: r.ID,
			SpecId:    v,
		}
	}
	if _, err = sess.Insert(&rss); err != nil {
		sess.Rollback()
		return err
	}

	return sess.Commit()
}

func DisagreeCardRequest(r CardRequestReview) error {
	sess := x.NewSession()
	var err error
	defer func() {
		if err != nil {
			sess.Rollback()
		}
		sess.Close()
	}()

	// find old scene
	old := CardRequest{}
	if has, _ := sess.ID(r.ID).Get(&old); !has {
		return errors.New("CardRequest not exist")
	}
	//update review_message
	rs := CardRequest{
		Status: CARD_REQEST_DISAGREE,
		Review: r.Review,
	}
	if _, err = sess.ID(r.ID).Update(&rs); err != nil {
		return err
	}

	//delete scene spec relation
	if _, err = sess.Where("request_id = ? ", r.ID).Delete(&CardRequestSpec{}); err != nil {
		sess.Rollback()
		return err
	}

	return sess.Commit()
}

func CreateCardRequest(cardRequest *CardRequest) error {
	_, err := x.Insert(cardRequest)
	return err
}
func GetCardRequestById(id int64) (*CardRequest, error) {
	rel := new(CardRequest)
	has, err := x.
		ID(id).
		Get(rel)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrNotExist{id}
	}

	return rel, nil
}

func SearchCardRequest(opts *CardRequestOptions) (int64, []*CardRequestSpecRes, error) {
	var cond = builder.NewCond()
	if opts.Page <= 0 {
		opts.Page = 1
	}

	needJoinSpec := false
	if opts.Keyword != "" {
		lowerKeyWord := strings.ToLower(opts.Keyword)
		cond = cond.And(builder.Or(builder.Like{"LOWER(card_request.contact)", lowerKeyWord},
			builder.Like{"LOWER(card_request.acc_cards_num)", lowerKeyWord},
			builder.Like{"LOWER(card_request.description)", lowerKeyWord}, builder.Like{"LOWER(card_request.description)", lowerKeyWord},
			builder.Like{"LOWER(card_request.phone_number)", lowerKeyWord}, builder.Like{"LOWER(card_request.wechat)", lowerKeyWord},
			builder.Like{"LOWER(card_request.org)", lowerKeyWord}, builder.Like{"LOWER(\"user\".name)", lowerKeyWord}))
	}
	if opts.UserID != 0 {
		cond = cond.And(builder.Eq{"\"user\".id": opts.UserID})
	}

	if opts.ResourceType != 0 {
		cond = cond.And(builder.Eq{"card_request.resource_type": opts.ResourceType})
	}
	if opts.AiCenterCode != "" {
		needJoinSpec = true
		cond = cond.And(builder.Eq{"resource_queue.ai_center_code": opts.AiCenterCode})
	}
	if opts.QueueId > 0 {
		cond = cond.And(builder.Eq{"resource_queue.id": opts.QueueId})
		needJoinSpec = true
	}
	if opts.ComputeResource != "" {
		cond = cond.And(builder.Eq{"card_request.compute_resource": opts.ComputeResource})
	}
	if opts.AccCardType != "" {
		cond = cond.And(builder.Eq{"card_request.card_type": opts.AccCardType})
	}
	if opts.Cluster != "" {
		cond = cond.And(builder.Eq{"resource_queue.cluster": opts.Cluster})
		needJoinSpec = true
	}

	if opts.UseBeginTime != 0 {
		cond = cond.And(builder.Gte{"card_request.begin_unix": opts.UseBeginTime})
	}
	if opts.UseEndTime != 0 {
		cond = cond.And(builder.Lte{"card_request.end_unix": opts.UseEndTime})
	}

	if opts.BeginTimeUnix != 0 {
		cond = cond.And(builder.Gte{"card_request.created_unix": opts.BeginTimeUnix})
	}
	if opts.EndTimeUnix != 0 {
		cond = cond.And(builder.Lte{"card_request.created_unix": opts.EndTimeUnix})
	}
	if opts.OrderBy == "" {
		opts.OrderBy = OrderByIDDesc
	}
	if opts.IsResearchProject == 1 {
		cond = cond.And(builder.Or(builder.Eq{"card_request.is_research_project": false}, builder.IsNull{"card_request.is_research_project"}))
	}
	if opts.IsResearchProject == 2 {
		cond = cond.And(builder.Eq{"card_request.is_research_project": true})
	}

	cond = cond.And(builder.NewCond().Or(builder.Eq{"card_request.delete_unix": 0}).Or(builder.IsNull{"card_request.delete_unix"}))
	cols := []string{"card_request.id", "card_request.compute_resource", "card_request.contact", "card_request.card_type", "card_request.acc_cards_num",
		"card_request.disk_capacity", "card_request.resource_type", "card_request.begin_date", "card_request.end_date", "card_request.uid",
		"card_request.phone_number", "card_request.email_address", "card_request.wechat", "card_request.is_research_project", "card_request.institution_name", "card_request.project_name", "card_request.project_code", "card_request.org", "card_request.description", "card_request.status", "card_request.review",
		"card_request.created_unix"}
	var count int64
	var err error
	if needJoinSpec {
		count, err = x.Where(cond).
			Distinct("card_request.id").
			Join("INNER", "card_request_spec", "card_request_spec.request_id = card_request.id").
			Join("INNER", "user", "\"user\".id = card_request.uid").
			Join("INNER", "resource_specification", "resource_specification.id = card_request_spec.spec_id").
			Join("INNER", "resource_queue", "resource_queue.id = resource_specification.queue_id").
			Count(&CardRequestSpecRes{})
		if err != nil {
			return 0, nil, err
		}
	} else {
		count, err = x.Where(cond).
			Distinct("card_request.id").
			Join("INNER", "user", "\"user\".id = card_request.uid").
			Count(&CardRequestSpecRes{})
	}

	r := make([]*CardRequestSpecRes, 0)
	if needJoinSpec {
		if err = x.Where(cond).Distinct(cols...).
			Join("INNER", "card_request_spec", "card_request_spec.request_id = card_request.id").
			Join("INNER", "user", "\"user\".id = card_request.uid").
			Join("INNER", "resource_specification", "resource_specification.id = card_request_spec.spec_id").
			Join("INNER", "resource_queue", "resource_queue.id = resource_specification.queue_id").
			OrderBy(opts.OrderBy).
			Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).
			Find(&r); err != nil {
			return 0, nil, err
		}
	} else {
		if err = x.Where(cond).Distinct(cols...).
			Join("INNER", "user", "\"user\".id = card_request.uid").
			OrderBy(opts.OrderBy).
			Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).
			Find(&r); err != nil {
			return 0, nil, err
		}
	}

	if len(r) == 0 {
		return 0, r, err
	}

	for _, v := range r {
		user, _ := GetUserByID(v.UID)
		if user != nil {
			v.UserName = user.Name
		}

	}

	//find related specs
	if opts.NeedSpec {
		requestIds := make([]int64, 0, len(r))
		for _, v := range r {
			requestIds = append(requestIds, v.ID)
		}

		specs := make([]RequestSpecInfo, 0)

		if err := x.Cols("resource_specification.id", "resource_specification.source_spec_id",
			"resource_specification.acc_cards_num", "resource_specification.cpu_cores",
			"resource_specification.mem_gi_b", "resource_specification.gpu_mem_gi_b",
			"resource_specification.share_mem_gi_b", "resource_specification.unit_price",
			"resource_specification.status", "resource_specification.updated_time",
			"card_request_spec.request_id", "resource_queue.cluster",
			"resource_queue.ai_center_code", "resource_queue.acc_card_type",
			"resource_queue.id as queue_id", "resource_queue.compute_resource",
			"resource_queue.queue_code", "resource_queue.queue_name",
			"resource_queue.queue_type", "resource_queue.ai_center_name",
			"resource_queue.has_internet",
		).In("card_request_spec.request_id", requestIds).
			Join("INNER", "card_request_spec", "card_request_spec.spec_id = resource_specification.id").
			Join("INNER", "resource_queue", "resource_queue.id = resource_specification.queue_id").
			OrderBy("resource_specification.acc_cards_num").
			Find(&specs); err != nil {
			return 0, nil, err
		}

		specsMap := make(map[int64][]RequestSpecInfo, 0)
		for _, v := range specs {
			if _, ok := specsMap[v.RequestId]; !ok {
				specsMap[v.RequestId] = []RequestSpecInfo{v}
			} else {
				specsMap[v.RequestId] = append(specsMap[v.RequestId], v)
			}
		}

		for i, v := range r {
			s := specsMap[v.ID]
			if s == nil {
				s = make([]RequestSpecInfo, 0)
			}
			r[i].Specs = s
		}
	}

	return count, r, nil
}

func UpdateCardRequest(cardRequest *CardRequest) error {
	_, err := x.ID(cardRequest.ID).Cols("compute_resource", "contact", "card_type", "acc_cards_num", "disk_capacity", "resource_type", "begin_date", "end_date", "phone_number", "wechat", "is_research_project", "institution_name", "project_name", "project_code", "email_address", "org", "description", "begin_unix", "end_unix").Update(cardRequest)
	return err
}
