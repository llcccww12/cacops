<template>
  <div class="section_8">

    <el-collapse v-model="activeNames" @change="handleChange">

      <div class="coll_tag">
        <span class="self-start text_9 font_3">常见问题</span>
      </div>

      <el-collapse-item title="如何将项目加入到人机协同智能操作系统社区？" name="1">
        <div class="section_10">
          <span class="font_5">给您的项目赋予标签ros-hmci-app(开源应用)，ros-hmci-datasets(社区数据集)，ros-hmci-models(社区智能模型)。</span>
        </div>
      </el-collapse-item>

      <el-collapse-item title="如何使用人机协同智能操作系统？" name="5">
        <div class="section_10">
          <span class="font_5">在操作系统下载列表中选择您需要的操作系统下载至本地，解压zip的包，得到iso镜像文件后进行安装使用。</span>
        </div>
      </el-collapse-item>
      
      <el-collapse-item title="如何发布新的资源？" name="6">
        <div class="section_10">
          <span class="font_5">在社区资源栏中点击资源发布，按格式填写需要发布的资源后提交，等待管理员审核完成后即可在社区资源找到您发布的资源。</span>
        </div>
      </el-collapse-item>
      
      <el-collapse-item title="是否能发布其他平台的资源？" name="7">
        <div class="section_10">
          <span class="font_5">可以。</span>
        </div>
      </el-collapse-item>
      <div class="coll_tag">
        <span class="self-start font_3 text_15">入门指南</span>
      </div>

      <el-collapse-item title="此处为常见问题2标题" name="2">
        <div class="section_10">
          <span class="font_5">此处为常见问题2内容</span>
        </div>
      </el-collapse-item>

      <div class="coll_tag">
        <span class="self-start font_3 text_15">安装手册</span>
      </div>
      <el-collapse-item title="此处为常见问题3标题" name="3">
        <div class="section_10">
          <span class="font_5">此处为常见问题3内容</span>
        </div>
      </el-collapse-item>
      <div class="coll_tag">
        <span class="self-start font_3 text_15">API接口</span>
      </div>
      <el-collapse-item title="此处为常见问题4标题" name="4">
        <div class="section_10">
          <span class="font_5">此处为常见问题4内容</span>
        </div>
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script>
export default {
  data() {
    return {
      activeNames: ['1']
    };
  },
  methods: {
    handleChange(val) {
      // console.log(val);
    }
  }
}
</script>

<style scoped>
@import url('../style/index.css');

.coll_tag {
  margin-top: 25px;
  margin-bottom: 10px;
}

::v-deep .el-collapse-item .el-collapse-item__header .el-icon-arrow-right {
  display: none;
  /* 隐藏默认的三角形图标 */
}

::v-deep .el-collapse-item .el-collapse-item__header::before {
  content: '';
  /* 设置空内容，这样才能在CSS中设置背景图像 */
  display: inline-block;
  width: 16px;
  /* 自定义图标的宽度 */
  height: 16px;
  /* 自定义图标的高度 */
  background-image: url('/img/ros-hmci/展开.png');
  /* 设置自定义图标的路径 */
  background-size: contain;
  background-repeat: no-repeat;
  vertical-align: middle;
  margin-right: 8px;
  /* 设置图标之间的间距 */
  transform: rotate(0deg);
  /* 添加过渡效果 */
  position: relative;
  left: 0;
}

::v-deep .el-collapse-item.is-active .el-collapse-item__header::before {
  background-image: url('/img/ros-hmci/展开-1.png');
}

::v-deep .el-collapse-item__header {
  font-size: 15px;
  border-top: 0;
  color: #575757;
}

::v-deep .el-collapse {
  border-top: 0;
  border-bottom: 0;
}

/* ::v-deep .el-collapse-item__wrap {
  border-bottom: 0;
} */

.text_15 {
  margin-top: 36px;
}

.text_9 {
  margin-left: 2px;
  line-height: 18.5px;
}

.font_3 {
  font-size: 20px;
  font-family: AlibabaPuHuiTi;
  line-height: 18px;
  color: #181818;
}

.font_2 {
  font-size: 16px;
  font-family: AlibabaPuHuiTi;
  line-height: 22px;
  color: #4d4d4d;
}

.section_8 {
  padding: 7px 24px 100px;
  background-color: #ffffff;
  border-radius: 6px;
  box-shadow: 0px 3px 12px #0000000a;
  width: 949px;
  height: auto;
  border: solid 1px #ebf7f6;
}

.section_10 {
  padding: 24px 24px 26px;
  background-color: #f3f4f6;
  border-radius: 6px;
}

.font_5 {
  font-size: 15px;
  font-family: AlibabaPuHuiTi;
  line-height: 18px;
  color: #757c8d;
}
</style>