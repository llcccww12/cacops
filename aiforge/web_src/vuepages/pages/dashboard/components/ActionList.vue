<template>
  <div class="action-container">
    <div class="action-wrap">
      <el-skeleton style="height:100%" :loading="isLoading" animated>
        <template slot="template">
          <div v-for="item in skeletonList" :key="item" class="skeleton-item-wrap">
            <el-skeleton-item variant="text" style="height: 100%; width: 90px; margin-right: 10px;" />
            <el-skeleton-item variant="text" style="flex:1; height: 100%;" />
          </div>
        </template>
        <template>
          <div class="swiper-container1">
            <div class="swiper-wrapper">
              <div class="swiper-slide" v-for="(item, index) in notificationList" :key="`${item.opType}-${index}`">
                <div class="content-item">
                  <img class="avatar-c" :src="item.avatar" />
                  <div class="content nowrap" v-html="item.formattedContent">
                    <a :href="item.userName">{{ item.userName }}</a>
                  </div>
                  <span class="time">{{ item.timeAgo }} </span>
                </div>

              </div>
            </div>
          </div>
        </template>
      </el-skeleton>
    </div>
  </div>
</template>

<script>
import { getAction } from "~/apis/modules/common";
import { actionNameZH, actionNameEN, OP_TYPES, TASK_OP_TYPES } from '../const.js'
import { lang } from '~/langs';
const skeletonList = ['first', 'second', 'third', 'fourth', 'fifth']
export default {
  name: "ActionList",
  data() {
    return {
      isLoading: true,
      isZh: lang == 'zh-CN',
      actionNameZH: actionNameZH,
      actionNameEN: actionNameEN,
      rawNotifications: [],
      notificationList: [],
      swiperHandler: null,
      skeletonList: skeletonList,
    };
  },
  computed: {
    // 动作名称映射
    actionNameMap() {
      return this.isZh ? this.actionNameZH : this.actionNameEN
    }
  },
  mounted() {
    this.loadAction()
  },
  methods: {
    initSwiper(loop = true) {
      // 检查 Swiper 容器是否存在
      if (!document.querySelector('.swiper-container1')) {
        console.warn('Swiper container not found');
        return;
      }

      // 销毁旧的 Swiper 实例（防止重复初始化）
      if (this.swiperHandler && this.swiperHandler.destroy) {
        this.swiperHandler.destroy(true, true);
      }
      let length = this.notificationList.length;
      this.swiperHandler = new Swiper(".swiper-container1", {
        slidesPerView: 5,
        direction: 'vertical',  // 改为垂直方向
        loop: loop && length > 5, // 只有一张图片时不需要 loop
        spaceBetween: 0,
        autoplay: {
          delay: 2500,
          disableOnInteraction: false,
        },
      });
    },
    async loadAction() {
      // 将上述请求修改成下面的的代码
      this.isLoading = true
      try {
        const response = await getAction();
        const res = response.data
        if (res.Code === 0) {
          this.rawNotifications = res.Data || []
          this.processNotifications()
          this.$nextTick(() => {
            this.initSwiper(true);
          });
        } else {
          this.$message.error(res.msg || '获取动态失败')
        }

      } catch (error) {
        this.$message.error(error || '获取动态失败')

      } finally {
        this.isLoading = false
      }
    },
    // 处理通知数据
    processNotifications() {
      const currentTime = Date.now()

      this.notificationList = this.rawNotifications
        .map(record => this.processNotification(record, currentTime))
        .filter(item => item !== null && item.formattedContent) // 过滤无效项

      //.filter(item => item !== null && item.formattedContent) // 过滤无效项
    },// 处理单个通知
    processNotification(record, currentTime) {
      const opType = record.opType.toString()
      const actionTemplate = this.actionNameMap[opType]
      if (!actionTemplate) {
        console.warn(`未识别的操作类型: ${opType}`)
        return null
      }

      const handlers = this.getNotificationHandlers()
      const handler = handlers[opType]
      if (!handler) {
        console.warn(`未找到对应的处理器: ${opType}`)
        return null
      }

      const result = handler.call(this, record, actionTemplate)
      if (!result || !result.formattedContent) {
        return null
      }
      let formattedContent = `<a href="/${record.userName}">${record.userName} </a>${result.formattedContent}`
      return {
        formattedContent: formattedContent,
        timeAgo: this.getTimeAgo(record.createdUnix, currentTime),
        avatar: `/user/avatar/${record.userName}/-1`,
        rawRecord: record
      }
    },
    // 然后在getNotificationHandlers中
    getNotificationHandlers() {
      const handlers = {}

      // 批量注册任务处理器
      TASK_OP_TYPES.forEach(opType => {
        handlers[opType] = this.handleTask
      })

      // 注册其他处理器
      handlers[OP_TYPES.REPO_CREATE] = this.handleCreateRepo
      handlers[OP_TYPES.REPO_RENAME] = this.handleRenameRepo
      handlers[OP_TYPES.REPO_PUSH_BRANCH] = this.handlePushBranch
      handlers[OP_TYPES.REPO_PUSH_TAG] = this.handlePushTag
      handlers[OP_TYPES.REPO_DELETE_BRANCH] = this.handleDeleteBranch

      handlers[OP_TYPES.ISSUE_OPEN] = this.handleIssue
      handlers[OP_TYPES.ISSUE_COMMENT] = this.handleCommentIssue
      handlers[OP_TYPES.ISSUE_CLOSE] = this.handleCloseIssue
      handlers[OP_TYPES.ISSUE_REOPEN] = this.handleReopenIssue

      handlers[OP_TYPES.PR_CREATE] = this.handleCreatePR
      handlers[OP_TYPES.PR_MERGE] = this.handleMergePR
      handlers[OP_TYPES.PR_CLOSE] = this.handleClosePR
      handlers[OP_TYPES.PR_REOPEN] = this.handleReopenPR
      handlers[OP_TYPES.PR_PROPOSE_CHANGES] = this.handleProposeChanges
      handlers[OP_TYPES.PR_COMMENT] = this.handleCommentPR

      handlers[OP_TYPES.DATASET_CREATE] = this.handleCreateDataset
      handlers[OP_TYPES.DATASET_DELETE] = this.handleDeleteDataset
      handlers[OP_TYPES.DATASET_RECOMMEND] = this.handleRecommendDataset

      handlers[OP_TYPES.MODEL_CREATE] = this.handleCreateModel
      handlers[OP_TYPES.MODEL_DELETE] = this.handleDeleteModel
      handlers[OP_TYPES.MODEL_RECOMMEND] = this.handleRecommendModel

      handlers[OP_TYPES.IMAGE_COMMIT] = this.handleCommitImage
      handlers[OP_TYPES.IMAGE_RECOMMEND] = this.handleRecommendImage

      return handlers
    },
    // 通用任务处理器
    handleTask(record, actionTemplate) {
      const link = this.getTaskLink(record)
      const linkText = record.refName || '未知任务'

      return {
        formattedContent: actionTemplate + (link ?
          ` <a href="${link}" rel="nofollow" title="${linkText}"> ${linkText} </a>` :
          ` <span class="disabled" title="${linkText}">${linkText}</span>`)
      }
    },
    handleCreateRepo(record, actionTemplate) {
      return {
        formattedContent: actionTemplate +
          ` <a href="${this.getRepoLink(record)}" rel="nofollow" title="${this.getRepoText(record)}"> ${this.getRepoText(record)} </a>`
      }
    },
    // 重命名仓库
    handleRenameRepo(record, actionTemplate) {
      if (!record.content || !record.repo) return null

      const oldRepoName = record.content
      const newRepoHtml = `<a href="${this.getRepoLink(record)}" rel="nofollow" title="${this.getRepoText(record)}"> ${this.getRepoText(record)} </a>`

      return {
        formattedContent: actionTemplate.replace('{oldRepoName}', oldRepoName) + ' ' + newRepoHtml
      }
    },
    // 推送分支
    handlePushBranch(record, actionTemplate) {
      const branchLink = `${this.getRepoLink(record)}/src/branch/${encodeURIComponent(record.refName)}`
      const branchHtml = `<a href="${branchLink}" class="branch-link" title="${record.refName}"> ${record.refName} </a>`

      return {
        formattedContent: actionTemplate.replace('{branch}', branchHtml) +
          ` <a href="${this.getRepoLink(record)}" rel="nofollow" title="${this.getRepoText(record)}">${this.getRepoText(record)}</a>`
      }
    },
    // 推送标签
    handlePushTag(record, actionTemplate) {
      const tagLink = `${this.getRepoLink(record)}/src/tag/${encodeURIComponent(record.refName)}`
      const tagHtml = `<a href="${tagLink}" class="tag-link" title="${record.refName}"> ${record.refName} </a>`

      return {
        formattedContent: actionTemplate.replace('{branch}', tagHtml) +
          ` <a href="${this.getRepoLink(record)}" rel="nofollow" title="${this.getRepoText(record)}">${this.getRepoText(record)}</a>`
      }
    },
    // 删除分支
    handleDeleteBranch(record, actionTemplate) {
      if (!record.repo || !record.refName) return null

      const repoHtml = `<a href="${this.getRepoLink(record)}" rel="nofollow" title="${this.getRepoText(record)}"> ${this.getRepoText(record)} </a>`
      const formattedAction = actionTemplate
        .replace('{deleteBranchName}', record.refName)
        .replace('{repoName}', repoHtml)

      return {
        formattedContent: formattedAction
      }
    },
    // Issue相关
    handleIssue(record, actionTemplate) {
      return this.handleIssueLike(record, actionTemplate)
    },

    handleCommentIssue(record, actionTemplate) {
      return this.handleIssueLike(record, actionTemplate)
    },

    handleCloseIssue(record, actionTemplate) {
      return this.handleIssueLike(record, actionTemplate)
    },

    handleReopenIssue(record, actionTemplate) {
      return this.handleIssueLike(record, actionTemplate)
    },
    // Issue通用处理
    handleIssueLike(record, actionTemplate) {
      return {
        formattedContent: actionTemplate +
          ` <a href="${this.getIssueLink(record)}" rel="nofollow" title="${this.getIssueText(record)}"> ${this.getIssueText(record)} </a>`
      }
    },
    // PR相关
    handleCreatePR(record, actionTemplate) {
      return this.handlePRLike(record, actionTemplate)
    },

    handleMergePR(record, actionTemplate) {
      return this.handlePRLike(record, actionTemplate)
    },

    handleClosePR(record, actionTemplate) {
      return this.handlePRLike(record, actionTemplate)
    },

    handleReopenPR(record, actionTemplate) {
      return this.handlePRLike(record, actionTemplate)
    },

    handleProposeChanges(record, actionTemplate) {
      return this.handlePRLike(record, actionTemplate)
    },

    handleCommentPR(record, actionTemplate) {
      return this.handlePRLike(record, actionTemplate)
    },

    // PR通用处理
    handlePRLike(record, actionTemplate) {
      return {
        formattedContent: actionTemplate +
          ` <a href="${this.getPRLink(record)}" rel="nofollow" title="${this.getPRText(record)}"> ${this.getPRText(record)} </a>`
      }
    },
    // 创建数据集
    handleCreateDataset(record, actionTemplate) {
      if (!record.dataset) return null

      const datasetLink = `/datasets/detail/${record.dataset.ownerName}/${record.dataset.name}`
      const datasetText = record.dataset.alias || record.dataset.name
      const datasetHtml = `<a href="${datasetLink}" target="_blank" title="${datasetText}"> ${datasetText} </a>`

      return {
        formattedContent: actionTemplate.replace('{dataset}', datasetHtml)
      }
    },

    // 删除数据集
    handleDeleteDataset(record, actionTemplate) {
      if (!record.content) return null

      const parts = record.content.split('|')
      if (parts.length <= 1) return null

      const datasetText = `<span class="deleted" title="${parts[1]}"> ${parts[1]} </span>`

      return {
        formattedContent: actionTemplate.replace('{dataset}', datasetText)
      }
    },

    // 创建模型
    handleCreateModel(record, actionTemplate) {
      if (!record.aimodel) return null

      const modelLink = `/models/detail/${record.aimodel.ownerName}/${record.aimodel.name}`
      const modelText = record.aimodel.alias || record.aimodel.name
      const modelHtml = `<a href="${modelLink}" target="_blank" title="${modelText}"> ${modelText} </a>`

      return {
        formattedContent: actionTemplate.replace('{aimodel}', modelHtml)
      }
    },
    // 推荐数据集
    handleRecommendDataset(record, actionTemplate) {
      let datasetHtml
      if (record.repo && record.repo.id > 0) {
        datasetHtml = `<a href="${this.getRepoLink(record)}/datasets" rel="nofollow">${record.content.split('|')[1]}</a>`
      } else if (record.dataset) {
        datasetHtml = `<a target="_blank" href="/datasets/detail/${record.dataset.ownerName}/${record.dataset.name}" title="${record.dataset.alias || record.dataset.name}">${record.dataset.alias || record.dataset.name}</a>`
      } else {
        return null
      }

      return {
        formattedContent: actionTemplate.replace('{dataset}', datasetHtml)
      }
    },
    // 删除模型
    handleDeleteModel(record, actionTemplate) {
      if (!record.content) return null

      const parts = record.content.split('|')
      if (parts.length <= 1) return null

      const modelText = `<span class="deleted" title="${parts[1]}"> ${parts[1]} </span>`

      return {
        formattedContent: actionTemplate.replace('{aimodel}', modelText)
      }
    },

    // 推荐模型
    handleRecommendModel(record, actionTemplate) {
      if (!record.aimodel) return null

      const modelLink = `/models/detail/${record.aimodel.ownerName}/${record.aimodel.name}`
      const modelText = record.aimodel.alias || record.aimodel.name
      const modelHtml = `<a href="${modelLink}" target="_blank" title="${modelText}"> ${modelText} </a>`

      return {
        formattedContent: actionTemplate.replace('{aimodel}', modelHtml)
      }
    },
    // 提交镜像
    handleCommitImage(record, actionTemplate) {
      if (!record.content) return null

      const parts = record.content.split('|')
      if (parts.length <= 1) return null

      const imageHtml = `<span title="${parts[1]}">${parts[1]}</span>`

      return {
        formattedContent: actionTemplate.replace('{image}', imageHtml)
      }
    },

    // 推荐镜像
    handleRecommendImage(record, actionTemplate) {
      if (!record.content) return null

      const parts = record.content.split('|')
      if (parts.length <= 1) return null

      const imageHtml = `<span> ${parts[1]} </span>`

      return {
        formattedContent: actionTemplate.replace('{image}', imageHtml)
      }
    },
    // 辅助方法
    getRepoLink(record) {
      if (!record.repo || !record.repo.ownerName || !record.repo.name) {
        return ''
      }
      return `/${record.repo.ownerName}/${record.repo.name}`
    },

    getRepoText(record) {
      if (!record.repo) return ''

      if (record.repo.alias) {
        return `${record.repo.ownerName}/${record.repo.alias}`
      }
      return `${record.repo.ownerName}/${record.repo.name}`
    },

    getIssueLink(record) {
      if (!record.repo) return ''
      return `/${record.repo.ownerName}/${record.repo.name}/issues/${this.getIssueId(record)}`
    },

    getIssueId(record) {
      if (record.comment && record.comment.issue) {
        return record.comment.issue.index || '1'
      }
      if (record.content) {
        const index = record.content.indexOf('|')
        if (index !== -1) {
          return record.content.substring(0, index)
        }
      }
      return '1'
    },

    getIssueText(record) {
      const repoText = this.getRepoText(record)
      return repoText ? `${repoText}#${this.getIssueId(record)}` : ''
    },

    getPRLink(record) {
      if (!record.repo) return ''
      return `/${record.repo.ownerName}/${record.repo.name}/pulls/${this.getIssueId(record)}`
    },

    getPRText(record) {
      return this.getIssueText(record)
    },

    getTaskLink(record) {
      if (!record.cloudbrainId) return ''

      let basePath = '/cloudbrains/detail/'
      const taskId = record.cloudbrainId

      const opType = record.opType

      switch (opType) {
        case 55:
        case 59:
          basePath = '/modelbase/experience/detail/'
          break
        case 56:
        case 58:
          basePath = '/modelbase/nlp/sft/detail/'
          break
        case 60:
          basePath = '/modelbase/cv/sft/detail/'
          break
        case 61:
          basePath = '/modelbase/cv/comfyui/detail/'
          break
        case 66:
          basePath = '/modelbase/eval/evaluate/detail/'
          break
        default:
          basePath = '/cloudbrains/detail/'
      }

      return basePath + taskId
    },

    getTimeAgo(timestamp, currentTime) {
      if (!timestamp) return '未知时间'

      const seconds = Math.floor((currentTime / 1000) - timestamp)

      if (seconds < 0) {
        return this.isZh ? '刚刚' : 'just now'
      }

      if (seconds < 60) {
        return `${seconds}${this.isZh ? '秒前' : seconds === 1 ? ' second ago' : ' seconds ago'}`
      }

      const minutes = Math.floor(seconds / 60)
      if (minutes < 60) {
        return `${minutes}${this.isZh ? '分钟前' : minutes === 1 ? ' minute ago' : ' minutes ago'}`
      }

      const hours = Math.floor(minutes / 60)
      if (hours < 24) {
        return `${hours}${this.isZh ? '小时前' : hours === 1 ? ' hour ago' : ' hours ago'}`
      }

      const days = Math.floor(hours / 24)
      return `${days}${this.isZh ? '天前' : days === 1 ? ' day ago' : ' days ago'}`
    }
  },
}
</script>

<style lang="less" scoped>
.action-container {
  flex: 1;
  height: 0;
  overflow: hidden;

  .action-wrap {
    height: 100%; // 确保容器有高度

    .skeleton-item-wrap {
      display: flex;
      height: 23px;
      margin-bottom: 15px;
    }

    /deep/ .swiper-container1 {
      height: 100%;

      .swiper-wrapper {
        .content-item {
          display: flex;
          align-items: center;
          font-size: 12px;
          height: auto !important; // 关键：让高度自适应内容

          .avatar-c {
            width: 24px;
            height: 24px;
            border-radius: 100%;
            margin-right: 6px;
          }

          /deep/ .content {
            color: rgba(64, 64, 64, 1);

            .branch-link {
              color: rgba(255, 98, 0, 1);
            }

            a {
              color: rgba(0, 136, 234, 1);
            }
          }

          .time {
            flex-shrink: 0;
            margin-left: 10px;
          }
        }
      }
    }
  }
}
</style>
