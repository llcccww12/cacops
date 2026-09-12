package entity

type TaskData struct {
	Code          ContainerData
	Dataset       []ContainerData
	PreTrainModel ContainerData
	OutPutPath    ContainerData
}

type ContainerData struct {
	Name              string `json:"name"`
	Bucket            string `json:"bucket"`
	EndPoint          string `json:"endPoint"`
	ObjectKey         string `json:"objectKey"`
	ContainerPath     string `json:"containerPath"`
	RealPath          string `json:"realPath"`
	ReadOnly          bool   `json:"readOnly"`
	IsDir             bool   `json:"isDir"`
	GetBackEndpoint   string `json:"getBackEndpoint"`
	S3DownloadUrl     string `json:"s3DownloadUrl"`
	Size              int64  `json:"size"`
	IsOverwrite       bool   `json:"isOverwrite"`
	IsNeedUnzip       bool   `json:"isNeedUnzip"`
	IsNeedTensorboard bool   `json:"isNeedTensorboard"`
	Id                string `json:"id"`

	StorageType StorageType
	SizeLimit   int `json:"sizeLimit"` // 目录大小限制，单位GB，0表示不限制，仅启智混合集群生效
}

type ContainerDataType string

const (
	ContainerCode             ContainerDataType = "code"
	ContainerDataset          ContainerDataType = "dataset"
	ContainerPreTrainModel    ContainerDataType = "pre_train_model"
	ContainerOutPutPath       ContainerDataType = "output"
	ContainerLogPath          ContainerDataType = "log"
	ContainerFileNoteBookCode ContainerDataType = "file_note_book_code"
	ContainerVisualization    ContainerDataType = "visualization"
	ContainerOutputAsModel    ContainerDataType = "output_as_model"
)

type ContainerBuildOpts struct {
	Disable bool
	//容器内路径
	ContainerPath string
	//在aiforge存储上基于云脑存储目录的相对路径
	StorageRelativePath string
	ReadOnly            bool
	AcceptStorageType   []StorageType
	Uncompressed        bool
	MKDIR               bool
	VolumeFolder        bool
}

func (opts ContainerBuildOpts) IsStorageTypeIn(storageType StorageType) bool {
	for _, s := range opts.AcceptStorageType {
		if string(s) == string(storageType) {
			return true
		}
	}
	return false
}
func (opts ContainerBuildOpts) GetLocalPath() string {
	if opts.StorageRelativePath != "" {
		return opts.StorageRelativePath
	}
	return opts.ContainerPath
}
