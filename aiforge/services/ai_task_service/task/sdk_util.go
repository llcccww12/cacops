package task

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"unicode"

	"code.gitea.io/gitea/entity"

	"code.gitea.io/gitea/models"
)

func GenerateSDKCode(opts entity.SDKCodeOpts) string {
	return generateAITaskSDKCode(opts)
}

func GenerateModelDownloadCode(file_name, owner_name, repo_name string) (string, string) {
	code_filename := ""
	cli_filename := ""

	if file_name != "" {
		code_filename = "filename=\"" + file_name + "\", "
		cli_filename = "--filename " + file_name + " "
	}

	code := "from openi import openi_download_file\n" +
		"openi_download_file(" +
		"\"" + owner_name + "/" + repo_name + "\", " +
		"repo_type=\"model\", " +
		code_filename +
		"local_dir=\"./" + repo_name + "\", " +
		"max_workers=10" +
		")\n"

	cli := "openi model download " +
		owner_name + "/" + repo_name + " " +
		cli_filename +
		"--local_dir ./" + repo_name + " " +
		"--max_workers 10"

	return code, cli
}

func GenerateModelFileDownloadCode(model_name, owner_name, repo_name, file_name string) (string, string) {
	code := "import openi\n"
	code += "openi.download_model_file(repo_id=\"" + owner_name + "/" + repo_name + "\"," +
		"model_name=\"" + model_name + "\"," +
		"file=\"" + file_name + "\"," +
		"save_path=\"./" + model_name + "\")\n"

	cli := "openi model download " + owner_name + "/" + repo_name + " " + model_name + " --save_path ./" + model_name + " --filename " + file_name
	return code, cli
}

func GenerateSDKDownloadCode(file_name, owner_name, repo_name string, subject string) (string, string) {
	code_filename := ""
	cli_filename := ""

	if file_name != "" {
		code_filename = "filename=\"" + file_name + "\", "
		cli_filename = "--filename " + file_name + " "
	}

	code := "from openi import openi_download_file\n" +
		"openi_download_file(" +
		"\"" + owner_name + "/" + repo_name + "\", " +
		"repo_type=\"" + subject + "\", " +
		code_filename +
		"local_dir=\"./" + repo_name + "\", " +
		"max_workers=10" +
		")\n"

	cli := "openi " + subject + " download " +
		owner_name + "/" + repo_name + " " +
		cli_filename +
		"--local_dir ./" + repo_name + " " +
		"--max_workers 10"

	return code, cli
}

func GenerateDatasetUploadCode(dataset_file_name, owner_name, repo_name string) (string, string) {
	code := "from openi import openi_download_file\n"
	code += "openi.upload_file(" +
		"repo_id=\"" + owner_name + "/" + repo_name + "\"," +
		"file=\"" + dataset_file_name + "\")\n\n"

	cli := "openi dataset upload " + owner_name + "/" + repo_name + " " + dataset_file_name

	return code, cli
}

func GenerateModelUploadCode(model_name, owner_name, repo_name string) (string, string) {
	code := "import openi\n"
	code += "openi.upload_model(" +
		"repo_id=\"" + owner_name + "/" + repo_name + "\"," +
		"model_name=\"" + model_name + "\"," +
		"folder=\"d:/Meta-Llama3-8B-instruct\"" +
		")\n\n"
	cli := "openi model upload " + owner_name + "/" + repo_name + " " + model_name + " d:/upload"
	return code, cli
}

func GenerateModelFileUploadCode(model_name, owner_name, repo_name string) (string, string) {
	code := "import openi\n"
	code += "openi.upload_model_file(" +
		"repo_id=\"" + owner_name + "/" + repo_name + "\"," +
		"model_name=\"" + model_name + "\"," +
		"file=\"d:/Meta-Llama3-8B-instruct/config.json\"" +
		")\n\n"
	cli := "openi model upload " + owner_name + "/" + repo_name + " " + model_name + " d:/upload/config.json"
	return code, cli
}

func generateAITaskSDKCode(opts entity.SDKCodeOpts) string {
	datasetNames := opts.DatasetNames
	pretrainModelNames := opts.PretrainModelNames
	parameterKeys := opts.ParameterKeys
	jobType := opts.JobType
	var code string
	if len(parameterKeys) > 0 {
		//添加参数解析代码
		code = code + "import argparse\n\nparser = argparse.ArgumentParser(description='忽略超参数不存在的报错问题')\n#添加自定义参数\n"
		for i := 0; i < len(parameterKeys); i++ {
			code = code + fmt.Sprintf("parser.add_argument(\"--%s\")\n", parameterKeys[i])
		}
		code = code + "args, unknown = parser.parse_known_args()\n\n"
	}

	//判断是否需要执行回传方法
	shouldAddUploadCode := shouldAddUpload(jobType)

	//判断是否需要执行prepare方法
	shouldPrepareCtx := true
	if len(datasetNames) == 0 && len(pretrainModelNames) == 0 && !opts.VisualizeRequired && !shouldAddUploadCode {
		shouldPrepareCtx = false
	}

	if !shouldPrepareCtx && !shouldAddUploadCode {
		return code
	}
	// code = code + "import os\nos.system(\"pip install openi\")\n"

	//添加import相关代码
	if !shouldAddUploadCode {
		code = code + "from c2net.context import prepare\n\n"
	} else if !shouldPrepareCtx {
		code = code + "from c2net.context import upload_output\n\n"
	} else {
		code = code + "from c2net.context import prepare,upload_output\n\n"
	}

	//添加prepare()方法
	if shouldPrepareCtx {
		code = code + "#初始化导入数据集和预训练模型到容器内\nc2net_context = prepare()\n"
	}

	//添加数据集相关代码
	var datasetNameMap = make(map[string]string, 0)
	for i := 0; i < len(datasetNames); i++ {
		if i == 0 {
			code = code + "\n#获取数据集路径\n"
		}
		datasetName := strings.TrimSuffix(datasetNames[i], ".zip")
		datasetName = strings.TrimSuffix(datasetName, ".tar.gz")
		validName := makeValidPythonVariableName(datasetName)
		for {
			if _, exists := datasetNameMap[datasetName]; exists {
				datasetName = datasetName + "_" + generateRandomString(3)
			} else {
				break
			}
		}
		datasetNameMap[datasetName] = ""
		pathCode := fmt.Sprintf("%s_path = c2net_context.dataset_path+\"/\"+\"%s\"\n", validName, datasetName)
		code = code + pathCode
	}

	//添加预训练模型相关代码
	var pretrainModelNameMap = make(map[string]string, 0)
	for i := 0; i < len(pretrainModelNames); i++ {
		if i == 0 {
			code = code + "\n#获取预训练模型路径\n"
		}
		modelName := makeValidPythonVariableName(pretrainModelNames[i])
		for {
			if _, exists := pretrainModelNameMap[modelName]; exists {
				modelName = modelName + "_" + generateRandomString(3)
			} else {
				break
			}
		}
		pretrainModelNameMap[modelName] = ""
		pathCode := fmt.Sprintf("%s_path = c2net_context.pretrain_model_path+\"/\"+\"%s\"\n", modelName, pretrainModelNames[i])
		code = code + pathCode
	}

	//添加tensorboard日志输出路径
	if opts.VisualizeRequired {
		code = code + "\n#tensorboard日志输出路径\ntensorboard_log_path = c2net_context.tensorboard_path\n"
	}

	//添加回传相关提示代码
	if opts.JobType != "" {
		code = code + "\n#输出结果必须保存在该目录\nyou_should_save_here = c2net_context.output_path\n"
	}

	//添加回传函数
	if shouldAddUploadCode {
		code = code + "\n#回传结果到openi，只有训练任务才能回传\nupload_output()\n"
	}
	return code
}

// 判断是否需要添加回传代码
func shouldAddUpload(jobType models.JobType) bool {
	if jobType == models.JobTypeTrain || jobType == models.JobTypeDebug {
		return true
	}
	return false
}

func makeValidPythonVariableName(input string) string {
	validName := strings.ToLower(input)
	// 使用正则表达式将非字母数字下划线的字符替换为下划线
	re := regexp.MustCompile(`[^a-zA-Z0-9_]`)
	validName = re.ReplaceAllString(validName, "_")
	re = regexp.MustCompile(`_+`)
	validName = re.ReplaceAllString(validName, "_")
	// 如果变量名以数字开头，添加一个下划线
	if len(validName) > 0 && unicode.IsDigit(rune(validName[0])) {
		validName = "_" + validName
	}
	if validName == "" || validName == "_" {
		validName = "pretrain_model_" + generateRandomString(3)
	}
	return validName
}

// generateRandomString 生成指定长度的随机字母字符串
func generateRandomString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"

	result := make([]byte, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
