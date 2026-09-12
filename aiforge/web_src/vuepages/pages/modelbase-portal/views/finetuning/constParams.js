export const basicParamers = [{
    parameter: 'num_train_epochs',
    value: 3,
    desc: 'modelFinetune.epochs',
    type:'input'
  }, {
    parameter: 'learning_rate',
    value: 5e-5,
    desc: 'modelFinetune.learningRate',
    type:'input'
  }, {
    parameter: 'per_device_train_batch_size',
    value: 4,
    desc: 'modelFinetune.batchSize',
    type: 'slider',
    min: 1,
    max: 1024,
    step: 1,
  }, {
    parameter: 'val_size',
    value: 0.1,
    desc: 'modelFinetune.valSize',
    type: 'slider',
    min: 0,
    max: 1,
    step: 0.01,
  },{
    parameter: 'eval_steps',
    value: 100,
    desc: 'modelFinetune.evalSteps',
    type: 'slider',
    min: 10,
    max: 5000,
    step: 1,
},{
    parameter: 'save_steps',
    value: 200,
    desc: 'modelFinetune.saveSteps',
    type: 'slider',
    min: 10,
    max: 5000,
    step: 1,
},{
  parameter: 'compute_type',
  value: 'fp16',
  desc: 'modelFinetune.computeType',
  options: [
      { value: 'bf16', label: 'bf16' },
      { value: 'fp16',label: 'fp16'},
      { value: 'fp32', label: 'fp32' },
      { value: 'pure_bf16', label: 'pure_bf16' },
  ],
  type:'select'
}]
  
export const otherPamrams = [{
  parameter: 'max_samples',
  value: 10000,
  desc: 'modelFinetune.maxsamples',
  type:'input'
}, {
  parameter: 'gradient_accumulation_steps',
  value: 8,
  desc: 'modelFinetune.gradientAccumulation',
  type: 'slider',
  min: 1,
  max: 1024,
  step: 1,
}, {
  parameter: 'lr_scheduler_type',
  value: 'cosine',
  desc: 'modelFinetune.lrScheduler',
  options: [
      { value: 'linear', label: 'linear' },
      { value: 'cosine',label: 'cosine'},
      { value: 'cosine_with_restarts', label: 'cosine_with_restarts' },
      { value: 'polynomial', label: 'polynomial' },
      { value: 'constant', label: 'constant' },
      { value: 'constant_with_warmup', label: 'constant_with_warmup' },
      { value: 'inverse_sqrt', label: 'inverse_sqrt' },
      { value: 'reduce_lr_on_plateau', label: 'reduce_lr_on_plateau' },
      { value: 'cosine_with_min_lr', label: 'cosine_with_min_lr' },
      { value: 'warmup_stable_decay', label: 'warmup_stable_decay' },
  ],
  type:'select'
},{
  parameter: 'max_grad_norm',
  value: 1.0,
  desc: 'modelFinetune.maximumGradientNorm',
  type: 'input',
},
{
  parameter: 'cutoff_len',
  value: 1024,
  desc: 'modelFinetune.cutoffLen',
  type: 'slider',
  min: 4,
  max: 65536,
  step: 1,
},
{
  parameter: 'preprocessing_num_workers',
  value: 16,
  desc: 'modelFinetune.preprocessingNumWorkers',
  type:'input'
},{
  parameter: 'logging_steps',
  value: 5,
  desc: 'modelFinetune.loggingSteps',
  type: 'slider',
  min: 1,
  max: 996,
  step: 1,
},{
  parameter: 'warmup_steps',
  value: 0,
  desc: 'modelFinetune.warmupSteps',
  type: 'slider',
  min: 0,
  max: 5000,
  step: 1,
},
// {
//   parameter: 'neftune_noise_alpha',
//   value: 0,
//   desc: '用于处理的进程数',
//   type: 'slider',
//   min: 0,
//   max: 10,
//   step: 0.1,
// },
{
  parameter: 'packing ',
  value: false,
  desc: 'modelFinetune.packing',
  options: [
      { value: false, label: 'no' },
      { value: true,label: 'yes'},
  ],
  type:'select'
},{
  parameter: 'optim',
  value: 'adamw_torch',
  desc: 'modelFinetune.optim',
  type: 'input',
},{
  parameter: 'lora_rank',
  value: 8,
  desc: 'modelFinetune.loraRank',
  type: 'slider',
  min: 1,
  max: 1024,
  step: 1,
},{
  parameter: 'lora_alpha',
  value: 16,
  desc: 'modelFinetune.loraAlpha',
  type: 'slider',
  min: 1,
  max: 2048,
  step: 1,
},{
  parameter: 'lora_dropout',
  value: 0,
  desc: 'modelFinetune.loraDropout',
  type: 'slider',
  min: 0,
  max: 1,
  step: 0.01,
},{
  parameter: 'lora_target',
  value: 'all',
  desc: 'modelFinetune.loraTarget',
  type: 'input',
},{
  parameter: 'additional_target',
  value: '',
  desc: 'modelFinetune.additionalTarget',
  type: 'input',
}]


export const sparkPamrams = [{
  parameter: 'max_epoch',
  parameterType: 'optimization',
  value: 1,
  desc: 'modelFinetune.max_epoch_tips',
  type: 'input',
},{
  parameter: 'io_strategy',
  parameterType: 'checkpoint',
  value: 'greedy_balance',
  desc: 'modelFinetune.io_strategy_tips',
  options: [
      { value: 'greedy_balance', label: 'greedy_balance' },
  ],
  type:'select'
},{
  parameter: 'keep_interval_update',
  parameterType: 'checkpoint',
  value: 1000,
  desc: 'modelFinetune.keep_interval_update_tips',
  type: 'input',
},{
  parameter: 'save_interval_update',
  parameterType: 'checkpoint',
  value: 1000,
  desc: 'modelFinetune.save_interval_update_tips',
  type: 'input'
},{
  parameter: 'lr',
  parameterType: 'optimization',
  value: '[1.6e-6]',
  desc: 'modelFinetune.lr_tips',
  type: 'input'
},{
  parameter: 'vitual_pipeline_model_parallel_size',
  parameterType: 'model_paraller',
  value: 1,
  desc: 'modelFinetune.vitual_pipeline_model_parallel_size_tips',
  type: 'input'
},{
  parameter: 'pipeline_model_parallel_size',
  parameterType: 'model_paraller',
  value: 8,
  desc: 'modelFinetune.pipeline_model_parallel_size_tips',
  type: 'input'
},{
  parameter: 'tensor_model_parallel_size',
  parameterType: 'model_paraller',
  value: 1,
  desc: 'modelFinetune.tensor_model_parallel_size_tips',
  type: 'input'
},{
  parameter: 'num_micro_batch',
  parameterType: 'model_paraller',
  value: 16,
  desc: 'modelFinetune.num_micro_batch_tips',
  type: 'input'
},{
  parameter: 'micro_batch_size',
  parameterType: 'model_paraller',
  value: 1,
  desc: 'modelFinetune.micro_batch_size_tips',
  type: 'input'
},{
  parameter: 'recomputer_granularity',
  parameterType: 'model_paraller',
  value: 'full',
  desc: 'modelFinetune.recomputer_granularity_tips',
  options: [
      { value: 'full', label: 'full' },
      { value: 'none', label: 'none' },
  ],
  type:'select'
},]