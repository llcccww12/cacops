package response

var RESOURCE_QUEUE_NOT_AVAILABLE = &BizError{Code: 1001, DefaultMsg: "resource queue not available"}
var SPECIFICATION_NOT_EXIST = &BizError{Code: 1002, DefaultMsg: "specification not exist"}
var SPECIFICATION_NOT_AVAILABLE = &BizError{Code: 1003, DefaultMsg: "specification not available"}

var CATEGORY_STILL_HAS_BADGES = &BizError{Code: 1004, DefaultMsg: "Please delete badges in the category first"}
var BADGES_STILL_HAS_USERS = &BizError{Code: 1005, DefaultMsg: "Please delete users of badge first"}

// common response
var SYSTEM_ERROR = &BizError{Code: 9009, DefaultMsg: "System error.Please try again later", TrCode: "common_error.system_error"}
var INSUFFICIENT_PERMISSION = &BizError{Code: 9003, DefaultMsg: "You do not have permission to perform this operation", TrCode: "common_error.insufficient_permission"}
var PARAM_ERROR = &BizError{Code: 9001, DefaultMsg: "param error", TrCode: "common_error.param_error"}
var WECHAT_NOT_BIND = &BizError{Code: 9002, DefaultMsg: "Please scan the code and bind to wechat first", TrCode: "common_error.wechat_not_bind"}
var NOT_EXISTS_OR_NO_PERMISSION = &BizError{Code: 9004, DefaultMsg: "The requested data does not exist or you do not have permission to access it.", TrCode: "common_error.not_exists_or_no_permission"}
var USER_NOT_EXISTS = &BizError{Code: 9005, DefaultMsg: "User not exists", TrCode: "common_error.user_not_exists"}

// 云脑任务相关错误
var AI_TASK_NOT_EXISTS = &BizError{Code: 2001, DefaultMsg: "AI task not exists", TrCode: "ai_task.task_not_exists"}
var AI_TASK_NOT_FINISHED = &BizError{Code: 2002, DefaultMsg: "AItask not finished", TrCode: "ai_task.task_not_finished"}
var SPEC_NOT_AVAILABLE = &BizError{Code: 2003, DefaultMsg: "Specification not available", TrCode: "ai_task.spec_not_available"}
var MULTI_TASK = &BizError{Code: 2004, DefaultMsg: "You have already a running or waiting task, can not create more", TrCode: "ai_task.multi_task"}
var MULTI_MODEL_TASK = &BizError{Code: 2044, DefaultMsg: "You have already a running or waiting task with the model, can not create more", TrCode: "ai_task.multi_model_task"}
var JOB_NAME_ALREADY_USED = &BizError{Code: 2005, DefaultMsg: "The job name did already exist", TrCode: "ai_task.job_name_already_used"}
var INSUFFICIENT_POINT_BALANCE = &BizError{Code: 2006, DefaultMsg: "Insufficient point balance", TrCode: "ai_task.insufficient_point_balance"}
var DATASET_NOT_EXISTS = &BizError{Code: 2007, DefaultMsg: "The part of datasets in the task does not exist or has been deleted, please create a new debug job.", TrCode: "repo.debug.manage.dataset_not_exist"}
var MODEL_NOT_EXISTS = &BizError{Code: 2008, DefaultMsg: "The model in the task does not exist or has been deleted", TrCode: "ai_task.model_not_exist"}
var RESULT_CLEARD = &BizError{Code: 2009, DefaultMsg: "The files of the task have been cleared, can not restart or retrain any more, please create a new task instead.", TrCode: "ai_task.result_cleared"}
var CREATE_FAILED = &BizError{Code: 2010, DefaultMsg: "Create AI task failed", TrCode: "ai_task.create_failed"}
var RESTART_FAILED = &BizError{Code: 2011, DefaultMsg: "Restart AI task failed", TrCode: "ai_task.restart_failed"}
var STOP_FAILED = &BizError{Code: 2012, DefaultMsg: "Stop AI task failed", TrCode: "ai_task.stop_failed"}
var DATASET_SIZE_OVER_LIMIT = &BizError{Code: 2013, DefaultMsg: "The size of dataset exceeds limitation", TrCode: "ai_task.dataset_size_over_limit"}
var BOOT_FILE_NOT_EXIST = &BizError{Code: 2014, DefaultMsg: "The boot file not exist", TrCode: "ai_task.boot_file_not_exist"}
var BOOT_FILE_MUST_BE_PYTHON = &BizError{Code: 2015, DefaultMsg: "The boot file must be a python file", TrCode: "ai_task.boot_file_must_python"}
var NO_NODE_RIGHR = &BizError{Code: 2016, DefaultMsg: "The boot file must be a python file", TrCode: "repo.modelarts.no_node_right"}
var DATASET_SELECT_ERROR = &BizError{Code: 2017, DefaultMsg: "Dataset select error: the count exceed the limit or has same name", TrCode: "cloudbrain.error.dataset_select"}
var PARTIAL_DATASETS_NOT_AVAILABLE = &BizError{Code: 2018, DefaultMsg: "There are non-existent or deleted files in the selected dataset file, please select again", TrCode: "cloudbrain.error.partial_datasets_not_available"}
var LOAD_CODE_FAILED = &BizError{Code: 2019, DefaultMsg: "Fail to load code.", TrCode: "cloudbrain.load_code_failed"}
var BRANCH_NOT_EXISTS = &BizError{Code: 2020, DefaultMsg: "The branch does not exist", TrCode: "ai_task.branch_not_exists"}
var MODEL_NUM_OVER_LIMIT = &BizError{Code: 2021, DefaultMsg: "The number of models exceeds the limit of 30", TrCode: "repo.debug.manage.model_num_over_limit"}
var DATASET_NUMBER_OVER_LIMIT = &BizError{Code: 2022, DefaultMsg: "The dataset count exceed the limit", TrCode: "ai_task.dataset_number_over_limit"}
var NOTEBOOK_EXCEED_MAX_NUM = &BizError{Code: 2023, DefaultMsg: "You can have up to 5 Debug Tasks, please try again after delete some tasks. ", TrCode: "ai_task.too_many_notebook"}
var CAN_NOT_STOP_CREATING_JOB = &BizError{Code: 2024, DefaultMsg: "AI task is creating, can not be stopped", TrCode: "ai_task.can_not_stop_creating_job"}
var NO_CENTER_MATCH = &BizError{Code: 2024, DefaultMsg: "Can not match an AI center, please select other specification.", TrCode: "ai_task.no_center_match"}
var MODEL_NUMBER_OVER_LIMIT = &BizError{Code: 2025, DefaultMsg: "The model count exceed the limit", TrCode: "ai_task.model_number_over_limit"}
var MODEL_SIZE_OVER_LIMIT = &BizError{Code: 2026, DefaultMsg: "The size of model exceeds limitation", TrCode: "ai_task.model_size_over_limit"}
var IMAGE_NOT_AVAILABLE = &BizError{Code: 2027, DefaultMsg: "The image is not available", TrCode: "ai_task.image_not_available"}
var REPO_CAN_NOT_CREATE_AI_TEMPLATE = &BizError{Code: 2028, DefaultMsg: "The current repository does not support the creation of cloudbrain task template", TrCode: "ai_task.repo_can_not_create_ai_template"}

var ENDPOINT_NOT_EQAUL_TASK = &BizError{Code: 2029, DefaultMsg: "Can not create as the url has been used.", TrCode: "ai_task.endpoint_been_used"}

var ENDPOINT_AND_PORT = &BizError{Code: 2030, DefaultMsg: "Custom path and port must be filled in", TrCode: "ai_task.endpoint_and_port_need"}
var PORT_RANGE = &BizError{Code: 2031, DefaultMsg: "The port range is 8000 to 8800", TrCode: "ai_task.port_range"}
var ENDPOINT_NOT_START_SLASH = &BizError{Code: 2032, DefaultMsg: "Custom paths need to comply with URL specifications", TrCode: "ai_task.endpoint_not_start_slash"}
var ENDPOINT_MUST_BE_VALID = &BizError{Code: 2033, DefaultMsg: "Custom paths can only include strings, numbers, and /", TrCode: "ai_task.endpoint_must_be_valid"}
var CAN_NOT_STOP_SAVING_IMAGE_JOB = &BizError{Code: 2034, DefaultMsg: "Cannot stop the AI task while it is saving the image.", TrCode: "ai_task.can_not_stop_saving_image_job"}
var CAN_NOT_FINETUNE_EXPERIENCE = &BizError{Code: 2035, DefaultMsg: "The task can not be exepirenced online.", TrCode: "ai_task.can_not_finetune_experience"}
var CAN_NOT_Eval = &BizError{Code: 2036, DefaultMsg: "The task can not be evaluated online.", TrCode: "ai_task.can_not_eval"}

var DATASET_CANNOT_READ = &BizError{Code: 2037, DefaultMsg: "Dataset select error: select dataset what you can read.", TrCode: "ai_task.dataset_cannot_read"}
var MODEL_CANNOT_READ = &BizError{Code: 2038, DefaultMsg: "Model select error: select model what you can read.", TrCode: "ai_task.model_cannot_read"}
var AI_TAKS_TEMPLATE_NAME_INVALID = &BizError{Code: 2039, DefaultMsg: "Name is invalid", TrCode: "ai_task_template.name_invalid"}
var REPOSITORY_NOT_EXISTS = &BizError{Code: 2040, DefaultMsg: "Repository not exists", TrCode: "ai_task_template.repo_not_exists"}
var PRIVATE_DATASET_IN_PUBLIC_TEMPLATE = &BizError{Code: 2041, DefaultMsg: "Exists private dataset in public ai task template", TrCode: "ai_task_template.private_dataset_in_public_template"}
var PRIVATE_MODEL_IN_PUBLIC_TEMPLATE = &BizError{Code: 2042, DefaultMsg: "Exists private model in public ai task template", TrCode: "ai_task_template.private_model_in_public_template"}
var PRIVATE_REPOSITORY_IN_PUBLIC_TEMPLATE = &BizError{Code: 2043, DefaultMsg: "Exists private repository in public ai task template", TrCode: "ai_task_template.private_repo_in_public_template"}
var RESOURCE_USAGE_NOT_SUPPORT = &BizError{Code: 6003, DefaultMsg: "Not support query resource usage", TrCode: "ai_task.resource_usage_not_support"}

// 登录相关错误
var QR_CODE_EXPIRED = &BizError{Code: 3001, DefaultMsg: "It has expired, please scan the QR code again", TrCode: "form.qr_code_expire"}
var WECHAT_NOT_BOUND = &BizError{Code: 3002, DefaultMsg: "Login failed, and your WeChat account is not bound to the OpenI community account.", TrCode: "form.wechat_not_bound"}

// 数据集相关错误
var DATASET_NAME_INVALID = &BizError{Code: 4001, DefaultMsg: "Dataset name is invalid", TrCode: "dataset.title_format_err"}
var DATASET_DESCRIPTION_INVALID = &BizError{Code: 4002, DefaultMsg: "Dataset description is invalid", TrCode: "dataset.description_format_err"}
var DATASET_NAME_EXIST = &BizError{Code: 4003, DefaultMsg: "Dataset name already exists", TrCode: "dataset.dataset_name_exist"}
var OVER_PREVIEW_SIZE = &BizError{Code: 4004, DefaultMsg: "The file size exceeds the preview limit", TrCode: "dataset.over_preview_size"}
var UNSUPPORTED_PREVIEW_FILE_TYPE = &BizError{Code: 4005, DefaultMsg: "Unsupported file type for preview", TrCode: "dataset.unsupported_preview_file_type"}
var DATASET_PATH_INCORRECT = &BizError{Code: 4006, DefaultMsg: "dataset incorrect", TrCode: ""}
var DATASET_EXIST = &BizError{Code: 4007, DefaultMsg: "dataset exists", TrCode: ""}
var DATASET_PATH_EMPTY = &BizError{Code: 4008, DefaultMsg: "dataset path empty", TrCode: ""}
var DATASET_ALIAS_INVALID = &BizError{Code: 4009, DefaultMsg: "Dataset alias is invalid", TrCode: "dataset.alias_format_err"}
var DATASET_ALIAS_EXIST = &BizError{Code: 4010, DefaultMsg: "Dataset alias already exists", TrCode: "dataset.dataset_alias_exist"}
var RESOURCE_IN_USE = &BizError{Code: 4011, DefaultMsg: "Resource is in use, operation not possible, please try again later.", TrCode: "dataset.resource_in_use"}

// 权限相关错误
var ORG_NOT_ALLOWED_TO_BE_COLLABORATOR = &BizError{Code: 5001, DefaultMsg: "Organizations cannot be added as a collaborator", TrCode: "access.org_not_allowed_to_be_collaborator"}
var OWNER_NOT_ALLOWED_TO_BE_COLLABORATOR = &BizError{Code: 5002, DefaultMsg: "Owner cannot be added as a collaborator", TrCode: "access.owner_not_allowed_to_be_collaborator"}
var ADD_COLLABORATOR_DUPLICATE = &BizError{Code: 5003, DefaultMsg: "The collaborator is already added", TrCode: "access.add_collaborator_duplicate"}
var TEAM_NOT_EXIST = &BizError{Code: 5004, DefaultMsg: "The team does not exist", TrCode: "form.team_not_exist"}
var TEAM_NOT_IN_ORGANIZATION = &BizError{Code: 5005, DefaultMsg: "The team and the target subject are not in the same organization", TrCode: "access.team_not_in_organization"}
var ADD_TEAM_DUPLICATE = &BizError{Code: 5006, DefaultMsg: "Team already has the target subject", TrCode: "access.add_team_duplicate"}
var CHANGE_TEAM_ACCESS_NOT_ALLOWED = &BizError{Code: 5006, DefaultMsg: "Changing team access for subject has been restricted to organization owner", TrCode: "access.change_team_access_not_allowed"}

// 模型相关错误
var AIMODEL_NAME_INVALID = &BizError{Code: 4001, DefaultMsg: "Aimodel name is invalid", TrCode: "aimodel.title_format_err"}
var AIMODEL_ALIAS_INVALID = &BizError{Code: 4009, DefaultMsg: "Aimodel alias is invalid", TrCode: "aimodel.alias_format_err"}
var AIMODEL_NAME_EXIST = &BizError{Code: 4003, DefaultMsg: "Aimodel name already exists", TrCode: "aimodel.aimodel_name_exist"}
var AIMODEL_ALIAS_EXIST = &BizError{Code: 4003, DefaultMsg: "Aimodel alias already exists", TrCode: "aimodel.aimodel_alias_exist"}
var AIMODEL_TASK_NOT_FOUND = &BizError{Code: 6001, DefaultMsg: "The model training task was not found", TrCode: "aimodel.aimodel_task_not_found"}
var AIMODEL_IS_LOADING = &BizError{Code: 6002, DefaultMsg: "The service is loading, please try again later.", TrCode: "aimodel.loading"}
