package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"

	"xorm.io/builder"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
)

const RECOMMOND_TYPE = 5
const NORMAL_TYPE = 0
const IMAGE_STATUS_COMMIT = 0
const IMAGE_STATUS_SUCCESS = 1
const IMAGE_STATUS_Failed = 2

type ApplyStatus int

const NoneApply ApplyStatus = 1
const Applying ApplyStatus = 2
const OKApply ApplyStatus = 3
const FailApply ApplyStatus = 4

type TrainTypeNum int

const (
	Notebook TrainTypeNum = 1 << iota
	TrainJob
	All = Notebook | TrainJob
)

var trainTypeNames = map[TrainTypeNum]string{
	Notebook: "Notebook",
	TrainJob: "TrainJob",
}

var nameToTrainType = map[string]TrainTypeNum{
	"Notebook": Notebook,
	"TrainJob": TrainJob,
}

type Image struct {
	ID                     int64       `xorm:"pk autoincr" json:"id"`
	ImageID                string      `json:"image_id"`
	Type                   int         `xorm:"INDEX NOT NULL" json:"type"`           //0 normal  5官方推荐,中间值保留为后续扩展
	CloudbrainType         int         `xorm:"INDEX NOT NULL" json:"cloudbrainType"` //0 云脑一   1云脑二
	UID                    int64       `xorm:"INDEX NOT NULL" json:"uid"`
	IsPrivate              bool        `xorm:"INDEX NOT NULL" json:"isPrivate"`
	Tag                    string      `xorm:"varchar(100) UNIQUE" json:"tag"`
	Description            string      `xorm:"varchar(1000)" json:"description"`
	Topics                 []string    `xorm:"TEXT JSON" json:"topics"`
	Place                  string      `xorm:"INDEX varchar(300)" json:"place"`
	NumStars               int         `xorm:"NOT NULL DEFAULT 0" json:"numStars"`
	ComputeResource        string      `xorm:"varchar(30)" json:"compute_resource"`
	IsStar                 bool        `xorm:"-" json:"isStar"`
	UserName               string      `xorm:"-" json:"userName"`
	RelAvatarLink          string      `xorm:"-" json:"relAvatarLink"`
	Status                 int         `xorm:"INDEX NOT NULL DEFAULT 0" json:"status"` //0代表正在提交，1提交完成，2提交失败
	ApplyStatus            ApplyStatus `xorm:"DEFAULT 1" json:"apply_status"`
	Message                string      `xorm:"varchar(500)" json:"message"`
	UseCount               int64       `xorm:"NOT NULL DEFAULT 0" json:"useCount"`
	Framework              string      `xorm:"varchar(100)" json:"framework" `
	FrameworkVersion       string      `xorm:"varchar(50)" json:"frameworkVersion" `
	CudaVersion            string      `xorm:"varchar(50)" json:"cudaVersion" `
	CannVersion            string      `xorm:"varchar(50)" json:"cannVersion" `
	DTKVersion             string      `xorm:"varchar(50)" json:"dtkVersion" `
	PythonVersion          string      `xorm:"varchar(50)" json:"pythonVersion" `
	OperationSystem        string      `xorm:"varchar(100)" json:"operationSystem" `
	OperationSystemVersion string      `xorm:"varchar(50)" json:"operationSystemVersion" `
	ThirdPackages          string      `xorm:"varchar(1000)" json:"thirdPackages" `
	AiCenterImages         `gorm:"type:json;comment:智算中心镜像"`
	GrampusBaseImage       int                `xorm:"" json:"grampusBaseImage"`                     //为1，表示分中心基础镜像
	TrainType              string             `xorm:"varchar(1024)" json:"trainType"`               //	镜像适用调试还是训练
	TrainTypeNum           TrainTypeNum       `xorm:"INDEX NOT NULL DEFAULT 0" json:"trainTypeNum"` //	镜像适用调试还是训练(INT值类型)
	HuJingId               string             `xorm:"varchar(50)" json:"huJingId" `                 //虎鲸侧镜像ID，用于同步镜像信息
	ProcessorType          string             `xorm:"varchar(50)" json:"processorType" `            //虎鲸侧资源类型，用于同步镜像信息
	AccCardType            string             `xorm:"varchar(50)" json:"accCardType" `              //虎鲸侧镜像卡类型，从AiCenterImages解释出来，用于根据资源规格过滤
	CreatedUnix            timeutil.TimeStamp `xorm:"INDEX created" json:"createdUnix"`
	UpdatedUnix            timeutil.TimeStamp `xorm:"INDEX updated" json:"updatedUnix"`
}

type AiCenterImages []AiCenterImage

type AiCenterImage struct {
	AiCenterId string `json:"aiCenterId"`

	TrainJobUrl    string   `json:"trainJobUrl"`
	NotebookUrl    string   `json:"notebookUrl"`
	AccDeviceModel string   `json:"accDeviceModel"`
	PoolIds        []string `json:"poolIds"`
}

func (r AiCenterImages) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *AiCenterImages) Scan(input interface{}) error {
	switch v := input.(type) {
	case []byte:
		return json.Unmarshal(input.([]byte), r)
	default:
		return fmt.Errorf("cannot Scan() from: %#v", v)
	}
}

type ImageList []*Image

type ImageStar struct {
	ID          int64              `xorm:"pk autoincr"`
	UID         int64              `xorm:"UNIQUE(s)"`
	ImageID     int64              `xorm:"UNIQUE(s)"`
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
}

type ImageTopic struct {
	ID          int64  `xorm:"pk autoincr"`
	Name        string `xorm:"UNIQUE VARCHAR(105)"`
	ImageCount  int
	CreatedUnix timeutil.TimeStamp `xorm:"INDEX created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"INDEX updated"`
}

type ImageVersionList struct {
}

type ImageTopicRelation struct {
	ImageID int64 `xorm:"UNIQUE(s)"`
	TopicID int64 `xorm:"UNIQUE(s)"`
}
type SearchAvailableValueOptions struct {
	Column           string
	Framework        string
	FrameworkVersion string
	PythonVersion    string
	ComputeResource  string
	OnlyRecommend    bool
	IncludeOwnerOnly bool
	IncludeStarByMe  bool
	UID              int64
	AiCenterIds      []string
}

type SearchImageOptions struct {
	Keyword                string
	UID                    int64
	Status                 int
	IncludePublicOnly      bool
	IncludeOfficialOnly    bool
	IncludePrivateOnly     bool
	IncludeStarByMe        bool
	IncludeCustom          bool
	IncludeOwnerOnly       bool
	Topics                 string
	ApplyStatus            int
	CloudbrainType         int
	Framework              string
	FrameworkVersion       string
	CudaVersion            string
	CannVersion            string
	DTKVersion             string
	PythonVersion          string
	OperationSystem        string
	OperationSystemVersion string
	ThirdPackages          string
	AiCenterIds            []string
	TrainType              string
	TrainTypeNum           TrainTypeNum
	ComputeResource        string
	AccCardType            string
	OnlyAccCardType        bool
	OnlyOpenIImage         bool
	ListOptions
	SearchOrderBy
}

type GrampusImageStatusResult struct {
	GrampusResult

	Image struct {
		CreatedAt      int    `json:"createdAt"`
		Id             string `json:"id"`
		ImageAddr      string `json:"imageAddr"`
		TrainJobUrl    string `json:"trainJobUrl"` // 训练类型
		NotebookUrl    string `json:"notebookUrl"` // 调试类型
		ImageDesc      string `json:"imageDesc"`
		ImageFullAddr  string `json:"imageFullAddr"`
		TrainType      string `json:"trainType"`
		ImageName      string `json:"imageName"`
		AccDeviceModel string `json:"accDeviceModel"`
		ImageStatus    int    `json:"imageStatus"`
		ImageVersion   string `json:"imageVersion"`
		SourceType     int    `json:"sourceType"`
		SpaceId        string `json:"spaceId"`
		UpdatedAt      int    `json:"updatedAt"`
		UserId         string `json:"userId"`
		Username       string `json:"username"`
	} `json:"image"`
}
type CommitImageGrampusParams struct {
	Description  string `json:"description"`
	ImageName    string `json:"imageName"`
	ImageVersion string `json:"imageVersion"`
	TaskName     string `json:"taskName"`
	TrainType    string `json:"trainType"`
}
type CommitGrampusImageParams struct {
	CommitImageGrampusParams
	IsPrivate              bool
	Topics                 []string
	CloudBrainType         int
	UID                    int64
	Place                  string
	Type                   int
	Framework              string
	FrameworkVersion       string
	CudaVersion            string
	CannVersion            string
	DTKVersion             string
	PythonVersion          string
	OperationSystem        string
	OperationSystemVersion string
	ThirdPackages          string
	ComputeResource        string
}

type CommitGrampusImageResult struct {
	GrampusResult
	ImageId string `json:"imageId"`
	Id      string `json:"id"`
}
type ErrorImageTagExist struct {
	Tag string
}

type ErrorImageCommitting struct {
	Tag string
}

type ImagesPageResult struct {
	Count  int64          `json:"count"`
	Images []*ImageResult `json:"images"`
}

type ImageResult struct {
	*Image
	AiCenterName string `json:"aiCenterName"`
}

func (err ErrorImageTagExist) Error() string {
	return fmt.Sprintf("Image already exists [tag: %s]", err.Tag)
}

func (err ErrorImageCommitting) Error() string {
	return fmt.Sprintf("Image already exists [tag: %s]", err.Tag)
}

type ErrImageNotExist struct {
	ID  int64
	Tag string
}

func (err ErrImageNotExist) Error() string {
	return fmt.Sprintf("Image does not exist [id: %d] [tag: %s]", err.ID, err.Tag)
}

func IsErrorImageCommitting(err error) bool {
	_, ok := err.(ErrorImageCommitting)
	return ok
}

func IsErrImageNotExist(err error) bool {
	_, ok := err.(ErrImageNotExist)
	return ok
}

func IsErrImageTagExist(err error) bool {
	_, ok := err.(ErrorImageTagExist)
	return ok
}

func IsImageExist(tag string) (bool, error) {
	return x.Exist(&Image{
		Tag: tag,
	})
}

func IsImageExistByUser(tag string, uid int64) (bool, error) {
	return x.Where("cloudbrain_type=?", TypeCloudBrainOne).Exist(&Image{
		Tag:    tag,
		UID:    uid,
		Status: IMAGE_STATUS_SUCCESS,
	})
}

type FindImageTopicOptions struct {
	ListOptions
	ImageID int64
	Keyword string
}

func (opts *FindImageTopicOptions) toConds() builder.Cond {
	var cond = builder.NewCond()
	if opts.ImageID > 0 {
		cond = cond.And(builder.Eq{"image_topic_relation.image_id": opts.ImageID})
	}

	if opts.Keyword != "" {
		cond = cond.And(builder.Like{"image_topic.name", strings.ToLower(opts.Keyword)})
	}

	return cond
}

func GetImageByID(id int64) (*Image, error) {
	rel := new(Image)
	has, err := x.
		ID(id).
		Get(rel)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrImageNotExist{ID: id}
	}

	return rel, nil
}

func GetImageByTag(tag string) (*Image, error) {

	image := &Image{Tag: tag}
	has, err := x.
		Get(image)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrImageNotExist{Tag: tag}
	}

	return image, nil
}

func GetImageByPlace(place string) (*Image, error) {

	image := &Image{Place: place}
	has, err := x.
		Get(image)
	if err != nil {
		return image, err
	} else if !has {
		return image, errors.New("place is not find image")
	}

	return image, nil
}

func GetImageByImageId(imageId string) (*Image, error) {
	if imageId == "" {
		return &Image{}, nil
	}

	image := &Image{ImageID: imageId}
	_, err := x.Get(image)
	if err != nil {
		return &Image{}, err
	}

	return image, nil
}

func GetImageByComputeResource(computeResource string) ([]Image, error) {
	images := make([]Image, 0)
	err := x.Where("status=1 and compute_resource=?", computeResource).Find(&images)
	if err != nil {
		return nil, err
	}
	return images, nil
}

func GetImageByTagAndCloudbrainType(tag string, cloudbrainType int) (*Image, error) {
	image := &Image{Tag: tag}
	has, err := x.
		Where("cloudbrain_type=?", cloudbrainType).
		Get(image)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrImageNotExist{Tag: tag}
	}

	return image, nil
}

func SanitizeAndValidateImageTopics(topics []string) (validTopics []string, invalidTopics []string) {
	validTopics = make([]string, 0)
	mValidTopics := make(map[string]struct{})
	invalidTopics = make([]string, 0)

	for _, topic := range topics {
		topic = strings.TrimSpace(strings.ToLower(topic))
		// ignore empty string
		if len(topic) == 0 {
			continue
		}
		// ignore same topic twice
		if _, ok := mValidTopics[topic]; ok {
			continue
		}
		if utf8.RuneCountInString(topic) <= 35 {
			validTopics = append(validTopics, topic)
			mValidTopics[topic] = struct{}{}
		} else {
			invalidTopics = append(invalidTopics, topic)
		}
	}

	return validTopics, invalidTopics
}
func FindImageTopics(opts *FindImageTopicOptions) (topics []*ImageTopic, err error) {
	sess := x.Select("image_topic.*").Where(opts.toConds())
	if opts.ImageID > 0 {
		sess.Join("INNER", "image_topic_relation", "image_topic_relation.topic_id = image_topic.id")
	}
	if opts.PageSize != 0 && opts.Page != 0 {
		sess = opts.setSessionPagination(sess)
	}
	return topics, sess.Desc("image_topic.image_count").Find(&topics)
}

func SaveImageTopics(imageID int64, topicNames ...string) error {
	topics, err := FindImageTopics(&FindImageTopicOptions{
		ImageID: imageID,
	})
	if err != nil {
		return err
	}

	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return err
	}

	var addedTopicNames []string
	for _, topicName := range topicNames {
		if strings.TrimSpace(topicName) == "" {
			continue
		}

		var found bool
		for _, t := range topics {
			if strings.EqualFold(topicName, t.Name) {
				found = true
				break
			}
		}
		if !found {
			addedTopicNames = append(addedTopicNames, topicName)
		}
	}

	var removeTopics []*ImageTopic
	for _, t := range topics {
		var found bool
		for _, topicName := range topicNames {
			if strings.EqualFold(topicName, t.Name) {
				found = true
				break
			}
		}
		if !found {
			removeTopics = append(removeTopics, t)
		}
	}

	for _, topicName := range addedTopicNames {
		_, err := addTopicByNameToImage(sess, imageID, topicName)
		if err != nil {
			return err
		}
	}

	for _, topic := range removeTopics {
		err := removeTopicFromImage(sess, imageID, topic)
		if err != nil {
			return err
		}
	}

	topicNames = make([]string, 0, 25)
	if err := sess.Table("image_topic").Cols("name").
		Join("INNER", "image_topic_relation", "image_topic_relation.topic_id = image_topic.id").
		Where("image_topic_relation.image_id = ?", imageID).Desc("image_topic.image_count").Find(&topicNames); err != nil {
		return err
	}

	if _, err := sess.ID(imageID).Cols("topics").Update(&Image{
		Topics: topicNames,
	}); err != nil {
		return err
	}

	return sess.Commit()
}

func addTopicByNameToImage(e Engine, imageID int64, topicName string) (*ImageTopic, error) {
	var topic ImageTopic
	has, err := e.Where("name = ?", topicName).Get(&topic)
	if err != nil {
		return nil, err
	}
	if !has {
		topic.Name = topicName
		topic.ImageCount = 1
		if _, err := e.Insert(&topic); err != nil {
			return nil, err
		}
	} else {
		topic.ImageCount++
		if _, err := e.ID(topic.ID).Cols("image_count").Update(&topic); err != nil {
			return nil, err
		}
	}

	if _, err := e.Insert(&ImageTopicRelation{
		ImageID: imageID,
		TopicID: topic.ID,
	}); err != nil {
		return nil, err
	}

	return &topic, nil
}

func removeTopicFromImage(e Engine, imageId int64, topic *ImageTopic) error {
	topic.ImageCount--
	if _, err := e.ID(topic.ID).Cols("image_count").Update(topic); err != nil {
		return err
	}

	if _, err := e.Delete(&ImageTopicRelation{
		ImageID: imageId,
		TopicID: topic.ID,
	}); err != nil {
		return err
	}

	return nil
}

func GetImageAvailableColumnValues(opts *SearchAvailableValueOptions) []string {
	var cond = builder.NewCond()

	if opts.Framework != "" {
		cond = cond.And(builder.Eq{"framework": opts.Framework})
	}
	if opts.FrameworkVersion != "" {
		cond = cond.And(builder.Eq{"framework_version": opts.FrameworkVersion})
	}
	if opts.PythonVersion != "" {
		cond = cond.And(builder.Eq{"python_version": opts.PythonVersion})
	}
	if opts.ComputeResource != "" {
		cond = cond.And(builder.Eq{"compute_resource": opts.ComputeResource})
	}
	if opts.OnlyRecommend {
		cond = cond.And(builder.Eq{"type": RECOMMOND_TYPE})
	}
	if opts.IncludeOwnerOnly {

		cond = cond.And(builder.Eq{"uid": opts.UID})
	}
	if len(opts.AiCenterIds) > 0 {
		tmpCond := builder.NewCond()
		for _, aiCenter := range opts.AiCenterIds {
			tmpCond = tmpCond.Or(builder.Like{"ai_center_images", "\"aiCenterId\":\"" + aiCenter + "\""})
		}
		cond = cond.And(tmpCond)
	}

	if opts.IncludeStarByMe {

		subQuery := builder.Select("image_id").From("image_star").
			Where(builder.Eq{"uid": opts.UID})
		var starCond = builder.In("id", subQuery)
		cond = cond.And(starCond)

	}

	cond = cond.And(builder.NotNull{opts.Column})

	var columnValues []string
	var re []string
	_ = x.Table("image").Where(cond).Distinct(opts.Column).Desc(opts.Column).Cols(opts.Column).Find(&columnValues)
	if columnValues == nil {
		return []string{}
	} else {
		for _, tmp := range columnValues {
			if tmp != "" {
				re = append(re, tmp)
			}
		}
	}
	return re
}

func SearchImage(opts *SearchImageOptions) (ImageList, int64, error) {
	cond := SearchImageCondition(opts)
	return SearchImageByCondition(opts, cond)
}

func SearchImageCondition(opts *SearchImageOptions) builder.Cond {
	var cond = builder.NewCond()

	if len(opts.Keyword) > 0 {

		var subQueryCond = builder.NewCond()
		for _, v := range strings.Split(opts.Keyword, ",") {

			subQueryCond = subQueryCond.Or(builder.Like{"LOWER(image_topic.name)", strings.ToLower(v)})

		}
		subQuery := builder.Select("image_topic_relation.image_id").From("image_topic_relation").
			Join("INNER", "image_topic", "image_topic.id = image_topic_relation.topic_id").
			Where(subQueryCond).
			GroupBy("image_topic_relation.image_id")
		var keywordCond = builder.In("id", subQuery)

		var likes = builder.NewCond()
		for _, v := range strings.Split(opts.Keyword, ",") {
			likes = likes.Or(builder.Like{"LOWER(tag)", strings.ToLower(v)})

			likes = likes.Or(builder.Like{"LOWER(description)", strings.ToLower(v)})

			likes = likes.Or(builder.Like{"LOWER(third_packages)", strings.ToLower(v)})
			likes = likes.Or(builder.Like{"LOWER(train_type)", strings.ToLower(v)})

			likes = likes.Or(builder.Like{"LOWER(CONCAT(framework,framework_version))", strings.ToLower(v)})
			likes = likes.Or(builder.Like{"LOWER(CONCAT('cuda',cuda_version))", strings.ToLower(v)})
			likes = likes.Or(builder.Like{"LOWER(CONCAT('dtk',dtk_version))", strings.ToLower(v)})
			likes = likes.Or(builder.Like{"LOWER(CONCAT('python',python_version))", strings.ToLower(v)})
			likes = likes.Or(builder.Like{"LOWER(CONCAT(operation_system,operation_system_version))", strings.ToLower(v)})

		}
		keywordCond = keywordCond.Or(likes)

		cond = cond.And(keywordCond)

	}
	if len(opts.Topics) > 0 { //标签精确匹配
		var subQueryCond = builder.NewCond()
		for _, v := range strings.Split(opts.Keyword, ",") {

			subQueryCond = subQueryCond.Or(builder.Eq{"LOWER(image_topic.name)": strings.ToLower(v)})
			subQuery := builder.Select("image_topic_relation.image_id").From("image_topic_relation").
				Join("INNER", "image_topic", "image_topic.id = image_topic_relation.topic_id").
				Where(subQueryCond).
				GroupBy("image_topic_relation.image_id")
			var topicCond = builder.In("id", subQuery)
			cond = cond.And(topicCond)
		}
	}

	if opts.ApplyStatus != 0 {
		cond = cond.And(builder.Eq{"apply_status": opts.ApplyStatus})
	}

	if opts.IncludePublicOnly {
		cond = cond.And(builder.Eq{"is_private": false})
	}

	if opts.IncludePrivateOnly {
		cond = cond.And(builder.Eq{"is_private": true})
	}

	if opts.IncludeOwnerOnly {

		cond = cond.And(builder.Eq{"uid": opts.UID})
	}
	if opts.IncludeOfficialOnly {
		cond = cond.And(builder.Eq{"type": RECOMMOND_TYPE})
	}
	if opts.Status >= 0 {
		cond = cond.And(builder.Eq{"status": opts.Status})
	}
	if opts.Framework != "" {
		cond = cond.And(builder.Eq{"framework": opts.Framework})
	}
	if opts.FrameworkVersion != "" {
		cond = cond.And(builder.Eq{"framework_version": opts.FrameworkVersion})
	}
	if opts.CudaVersion != "" {
		cond = cond.And(builder.Eq{"cuda_version": opts.CudaVersion})
	}
	if opts.CannVersion != "" {
		cond = cond.And(builder.Eq{"cann_version": opts.CannVersion})
	}
	if opts.DTKVersion != "" {
		cond = cond.And(builder.Eq{"dtk_version": opts.DTKVersion})
	}
	if opts.PythonVersion != "" {
		cond = cond.And(builder.Eq{"python_version": opts.PythonVersion})
	}
	if opts.OperationSystem != "" {
		cond = cond.And(builder.Eq{"operation_system": opts.OperationSystem})
	}
	if opts.OperationSystemVersion != "" {
		cond = cond.And(builder.Eq{"operation_system_version": opts.OperationSystemVersion})
	}
	if len(opts.AiCenterIds) != 0 {
		var orCond = builder.NewCond()
		for _, aiCenterId := range opts.AiCenterIds {
			orCond = orCond.Or(builder.Like{"ai_center_images", "\"aiCenterId\":\"" + aiCenterId + "\""})
		}
		cond = cond.And(orCond)
	}

	if opts.TrainTypeNum > 0 && opts.TrainType != "" {
		var orCond = builder.NewCond()
		orCond = orCond.Or(builder.Expr("(train_type_num & ?) = ?", opts.TrainTypeNum, opts.TrainTypeNum))
		orCond = orCond.Or(builder.Eq{"train_type_num": 0})
		log.Info(fmt.Sprintf("opts.TrainTypeNum=%d", opts.TrainTypeNum))
		cond = cond.And(orCond)
	}
	if opts.AccCardType != "" {
		// 如果是需要精确查询
		if opts.OnlyAccCardType {
			var orCond = builder.NewCond()
			orCond = orCond.Or(builder.Eq{"acc_card_type": opts.AccCardType})
			log.Info("opts.AccCardType=" + opts.AccCardType)
			cond = cond.And(orCond)
		} else {
			var orCond = builder.NewCond()
			orCond = orCond.Or(builder.Eq{"acc_card_type": opts.AccCardType})
			orCond = orCond.Or(builder.Eq{"acc_card_type": ""})
			orCond = orCond.Or(builder.IsNull{"acc_card_type"})
			log.Info("opts.AccCardType=" + opts.AccCardType)
			cond = cond.And(orCond)
		}
	}
	if opts.ComputeResource != "" {
		cond = cond.And(builder.Eq{"compute_resource": opts.ComputeResource})
	}
	var third_likes, isValid = builder.NewCond(), false
	for _, v := range strings.Split(opts.ThirdPackages, " ") {
		if v != "" {
			third_likes = third_likes.And(builder.Like{"LOWER(third_packages)", strings.ToLower(v)})
			isValid = true
		}

	}
	if isValid {
		cond = cond.And(third_likes)
	}

	if opts.IncludeStarByMe {

		subQuery := builder.Select("image_id").From("image_star").
			Where(builder.Eq{"uid": opts.UID})
		var starCond = builder.In("id", subQuery)
		cond = cond.And(starCond)

	}

	if opts.CloudbrainType >= 0 {
		cond = cond.And(builder.Eq{"cloudbrain_type": opts.CloudbrainType})
	}
	if opts.OnlyOpenIImage {
		cond = cond.And(builder.Or(builder.Eq{"grampus_base_image": 0}, builder.IsNull{"grampus_base_image"}))
	}

	return cond
}

func SearchImageByCondition(opts *SearchImageOptions, cond builder.Cond) (ImageList, int64, error) {
	if opts.Page <= 0 {
		opts.Page = 1
	}

	var err error
	sess := x.NewSession()
	defer sess.Close()

	images := make(ImageList, 0, opts.PageSize)
	count, err := sess.Where(cond).Count(new(Image))

	if err != nil {
		return nil, 0, fmt.Errorf("Count: %v", err)
	}

	sess.Where(cond).OrderBy(opts.SearchOrderBy.String())

	if opts.PageSize > 0 {
		sess.Limit(opts.PageSize, (opts.Page-1)*opts.PageSize)
	}
	if err = sess.Find(&images); err != nil {
		return nil, 0, fmt.Errorf("Images: %v", err)
	}

	if err = images.loadAttributes(sess, opts.UID); err != nil {
		return nil, 0, fmt.Errorf("LoadAttributes: %v", err)
	}

	return images, count, nil
}

func (images ImageList) loadAttributes(e Engine, uid int64) error {
	if len(images) == 0 {
		return nil
	}

	set := make(map[int64]struct{})

	for i := range images {
		set[images[i].UID] = struct{}{}
	}

	// Load creators.
	users := make(map[int64]*User, len(set))
	if err := e.Table("\"user\"").
		Cols("name", "lower_name", "avatar", "email").
		Where("id > 0").
		In("id", keysInt64(set)).
		Find(&users); err != nil {
		return fmt.Errorf("find users: %v", err)
	}

	for i := range images {
		if users[images[i].UID] != nil {
			images[i].UserName = users[images[i].UID].Name
			images[i].RelAvatarLink = users[images[i].UID].RelAvatarLink()
		} else {
			images[i].UserName = ""
			images[i].RelAvatarLink = ""
		}
		if uid == -1 {
			images[i].IsStar = false
		} else {
			images[i].IsStar = isImageStaring(e, uid, images[i].ID)
		}
	}

	return nil
}

func GetCommittingImageCount() int {

	total, err := x.Where("status =?", 0).Count(new(Image))

	if err != nil {
		return 0
	}
	return int(total)
}

func CreateLocalImage(image *Image) error {

	_, err := x.Insert(image)
	return err
}

func UpdateLocalImage(image *Image) error {

	_, err := x.ID(image.ID).Cols("description", "is_private", "status", "message", "apply_status", "framework", "framework_version",
		"cuda_version", "cann_version", "python_version", "operation_system", "operation_system_version,third_packages", "dtk_version").Update(image)
	return err
}

func UpdateLocalImageStatus(image *Image) error {

	_, err := x.ID(image.ID).Cols("status").Update(image)
	return err
}

func UpdateLocalImageStatusAndPlace(image *Image) error {
	_, err := x.ID(image.ID).Cols("status", "place", "ai_center_images", "train_type", "train_type_num").Update(image)
	return err
}

func UpdateAutoIncrementIndex() {
	x.Exec("SELECT setval('image_id_seq', (SELECT MAX(id) from image))")
}

func DeleteLocalImage(id int64) error {
	image := new(Image)
	_, err := x.ID(id).Delete(image)
	return err
}

// star or unstar Image
func StarImage(userID, imageID int64, star bool) error {
	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return err
	}

	if star {
		if isImageStaring(sess, userID, imageID) {
			return nil
		}

		if _, err := sess.Insert(&ImageStar{UID: userID, ImageID: imageID}); err != nil {
			return err
		}
		if _, err := sess.Exec("UPDATE `image` SET num_stars = num_stars + 1 WHERE id = ?", imageID); err != nil {
			return err
		}
		if _, err := sess.Exec("UPDATE `user` SET num_image_stars = num_image_stars + 1 WHERE id = ?", userID); err != nil {
			return err
		}
	} else {
		if !isImageStaring(sess, userID, imageID) {
			return nil
		}

		if _, err := sess.Delete(&ImageStar{0, userID, imageID, 0}); err != nil {
			return err
		}
		if _, err := sess.Exec("UPDATE `image` SET num_stars = num_stars - 1 WHERE id = ?", imageID); err != nil {
			return err
		}
		if _, err := sess.Exec("UPDATE `user` SET num_image_stars = num_image_stars - 1 WHERE id = ?", userID); err != nil {
			return err
		}
	}

	return sess.Commit()
}

func IsImageStaring(userID, datasetID int64) bool {
	return isImageStaring(x, userID, datasetID)

}

func isImageStaring(e Engine, userID, imageID int64) bool {
	has, _ := e.Get(&ImageStar{0, userID, imageID, 0})
	return has
}
func RecommendImage(imageId int64, recommend bool, applyStatus ApplyStatus, message string) error {

	image := Image{Type: GetRecommondType(recommend), ApplyStatus: applyStatus, Message: message}
	_, err := x.ID(imageId).Cols("type", "apply_status", "message").Update(image)
	return err
}

func GetRecommondType(recommond bool) int {
	if recommond {

		return RECOMMOND_TYPE
	} else {
		return NORMAL_TYPE
	}

}

func GetGrampusAllBaseImage() (ImageList, error) {
	var cond = builder.NewCond()
	cond = cond.And(builder.Eq{"grampus_base_image": 1})
	images := make(ImageList, 0, 100)

	err := x.Table("image").Where(cond).Find(&images)
	if err != nil {
		return nil, err
	}
	return images, nil
}

func getComputeResourceByProcessType(processType string) string {
	tail := strings.LastIndex(processType, "/")
	if tail > 0 {
		return strings.ToUpper(processType[tail+1:])
	}
	return strings.ToUpper(processType)
}

func SyncGrampusAllBaseImageToDb(insertList []GrampusImage, updateList []GrampusImage, updateDbList []*Image, deleteList []*Image, doerId int64) error {
	for _, tmp := range insertList {
		log.Info("insert image,image name=" + tmp.Name)

		trainType, trainNum := ConvertTrainType(tmp.TrainType)

		insObj := &Image{
			Type:             5,
			CloudbrainType:   2,
			UID:              doerId,
			IsPrivate:        false,
			Tag:              tmp.Name,
			ProcessorType:    tmp.ProcessorType,
			ComputeResource:  getComputeResourceByProcessType(tmp.ProcessorType),
			Status:           1,
			GrampusBaseImage: 1,
			TrainType:        trainType,
			TrainTypeNum:     trainNum,
			HuJingId:         tmp.ID,
			ImageID:          tmp.ID,
			AiCenterImages:   tmp.AICenterImage,
		}

		// 保证有值，方便在详情页面进行反查
		if tmp.AICenterImage != nil && len(tmp.AICenterImage) > 0 {
			img := tmp.AICenterImage[0]

			if img.NotebookUrl != "" && (img.TrainJobUrl == "" || img.TrainJobUrl == img.NotebookUrl) {
				insObj.Place = img.NotebookUrl
			} else if img.TrainJobUrl != "" {
				insObj.Place = img.TrainJobUrl
			}
		}

		if tmp.AICenterImage != nil && len(tmp.AICenterImage) > 0 {
			insObj.AccCardType = strings.ToUpper(tmp.AICenterImage[0].AccDeviceModel)
		}
		x.Insert(insObj)
	}
	for _, tmp := range deleteList {
		delObj := &Image{
			ID:  tmp.ID,
			Tag: tmp.Tag,
		}
		log.Info("delete image,image name=" + tmp.Tag)
		x.Delete(delObj)
	}

	for i, tmp := range updateList {
		log.Info("update image,image name=" + tmp.Name)
		updateDbList[i].AiCenterImages = tmp.AICenterImage
		updateDbList[i].Tag = tmp.Name
		if tmp.AICenterImage != nil && len(tmp.AICenterImage) > 0 {
			updateDbList[i].AccCardType = strings.ToUpper(tmp.AICenterImage[0].AccDeviceModel)
		}

		updateDbList[i].ComputeResource = getComputeResourceByProcessType(tmp.ProcessorType)

		// 保证有值，方便在详情页面进行反查
		if tmp.AICenterImage != nil && len(tmp.AICenterImage) > 0 {
			img := tmp.AICenterImage[0]

			if img.NotebookUrl != "" && (img.TrainJobUrl == "" || img.TrainJobUrl == img.NotebookUrl) {
				updateDbList[i].Place = img.NotebookUrl
			} else if img.TrainJobUrl != "" {
				updateDbList[i].Place = img.TrainJobUrl
			}
		}

		trainType, trainNum := ConvertTrainType(tmp.TrainType)

		updateDbList[i].TrainType = trainType
		updateDbList[i].TrainTypeNum = trainNum
		x.ID(updateDbList[i].ID).Cols("ai_center_images", "tag", "acc_card_type", "compute_resource", "place", "train_type", "train_type_num").Update(updateDbList[i])
	}

	return nil
}

func GetCommittingImages() (ImageList, error) {
	var cond = builder.NewCond()
	cond = cond.And(builder.Eq{"status": IMAGE_STATUS_COMMIT})
	images := make(ImageList, 0)
	err := x.Table("image").Where(cond).Find(&images)
	if err != nil {
		return nil, err
	}
	return images, nil
}

func GetGrampusSrcAllBaseImage(srcImageUrl string) (ImageList, error) {
	var cond = builder.NewCond()
	cond = cond.And(builder.Eq{"grampus_base_image": 1})
	cond = cond.And(builder.Eq{"train_type": "Notebook"})
	cond = cond.And(builder.Like{"ai_center_images", "\"imageUrl\":\"" + srcImageUrl + "\""})
	images := make(ImageList, 0, 10)
	err := x.Table("image").Where(cond).Find(&images)
	if err != nil {
		return nil, err
	}
	return images, nil
}

func FindImageByImageUrl(imageUrl string) (ImageList, error) {
	var cond = builder.NewCond()
	cond = cond.Or(builder.Eq{"place": imageUrl})
	cond = cond.Or(builder.Like{"ai_center_images", "\"imageUrl\":\"" + imageUrl + "\""})
	images := make(ImageList, 0, 10)
	err := x.Table("image").Where(cond).Find(&images)
	if err != nil {
		return nil, err
	}
	return images, nil
}

func TrainTypeToString(t TrainTypeNum) string {
	if t == -1 {
		return ""
	}

	isAll := t == All || t == 0

	var names []string
	for typ, name := range trainTypeNames {
		if t&typ != 0 || isAll {
			names = append(names, name)
		}
	}
	return strings.Join(names, "&")
}

// ParseTrainType Parse TrainType from the string
func ParseTrainType(s string) TrainTypeNum {
	if s == "" {
		return All
	}

	parts := strings.Split(s, "&")
	var result TrainTypeNum
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if typ, ok := nameToTrainType[part]; ok {
			result |= typ
		}
	}
	return result
}

// SortTrainType Sort train type with & special tag
func SortTrainType(input string) string {

	inputArray := strings.Split(input, "&")
	sort.Slice(inputArray, func(i, j int) bool {
		return nameToTrainType[inputArray[i]] < nameToTrainType[inputArray[j]]
	})

	return strings.Join(inputArray, "&")
}

// ConvertTrainType convert train type with sort, parse to train type number
func ConvertTrainType(s string) (string, TrainTypeNum) {
	if s == "" {
		return TrainTypeToString(All), All
	}

	parts := strings.Split(s, "&")
	var result TrainTypeNum
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if typ, ok := nameToTrainType[part]; ok {
			result |= typ
		}
	}
	return SortTrainType(s), result
}

// GetImageUrlByTrainType 根据镜像类型获取地址(主要是查询的时候)
func GetImageUrlByTrainType(trainType, noteBookUrl, trainJobUrl string) string {
	switch trainType {
	case "Notebook":
		return noteBookUrl
	case "TrainJob":
		return trainJobUrl
	}
	return ""
}
