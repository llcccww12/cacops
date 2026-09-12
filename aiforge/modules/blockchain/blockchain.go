package blockchain

const (
	Command            = `pip3 install jupyterlab==2.2.5 -i https://pypi.tuna.tsinghua.edu.cn/simple;service ssh stop;jupyter lab --no-browser --ip=0.0.0.0 --allow-root --notebook-dir="/code" --port=80 --LabApp.token="" --LabApp.allow_origin="self https://cloudbrain.pcl.ac.cn"`
	CodeMountPath      = "/code"
	DataSetMountPath   = "/dataset"
	ModelMountPath     = "/model"
	BenchMarkMountPath = "/benchmark"
	TaskInfoName       = "/taskInfo"

	SubTaskName = "task1"
)
