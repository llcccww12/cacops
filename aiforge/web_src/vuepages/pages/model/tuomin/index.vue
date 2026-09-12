<template>
  <div class="ui container">
    <div class="tuomin-title">
      <h2>{{ $t("dataDesensitizationModelExperience") }}</h2>
      <p>
        {{ $t("dataDesensitizationModelDesc") }} &nbsp;<a
          href="https://openi.pcl.ac.cn/tengxiao/tuomin"
          target="_blank"
          >tengxiao / tuomin</a
        >
      </p>
    </div>
    <el-row :gutter="12" style="margin-top: 33px">
      <el-col :xs="24" :span="12">
        <div class="tuomin-content-image">
          <div class="tuomin-icon">
            <i
              class="ri-image-line"
              style="font-size: 16px; margin-right: 2px"
            ></i>
            <span style="font-size: 12px">img</span>
          </div>
          <div style="height: 230px; width: 96%; margin: 0 auto">
            <el-upload
              action="#"
              accept=".jpg,.jpeg,.png,.JPG,.JPEG,.PNG"
              :show-file-list="false"
              :on-change="handleChangePicture"
              list-type="picture-card"
              :file-list="fileList"
              :style="{ display: ImageUrl ? 'none' : 'block' }"
              :auto-upload="false"
              drag
            >
              <div class="el-upload__text">
                <span>
                  <span>{{ $t("dragThePictureHere") }}</span><span style="color: rgba(136, 136, 136, 0.87)">{{ $t("or") }}</span><span>{{ $t("clickUpload") }}</span>
                </span>
              </div>
            </el-upload>

            <img class="preview-image" v-if="ImageUrl" :src="ImageUrl" alt="" />
          </div>
          <div>
            <div class="tuomin-radio-model">
              <label class="radio-label"
                >{{ $t("desensitizationObject") }}：</label
              >
              <div>
                <el-radio-group v-model="radio">
                  <el-radio :label="2">{{ $t("all") }}</el-radio>
                  <el-radio :label="1">{{ $t("onlyFace") }}</el-radio>
                  <el-radio :label="0">{{ $t("onlyLicensePlate") }}</el-radio>
                </el-radio-group>
              </div>
            </div>
            <div class="tuomin-button-model">
              <el-button @click="cancelUpload">{{ $t("cancel") }}</el-button>
              <el-button
                :disabled="fileList.length < 1"
                @click="startTranform"
                type="primary"
                >{{ $t("startDesensitization") }}</el-button
              >
            </div>
          </div>
        </div>
      </el-col>
      <el-col :xs="24" :span="12">
        <div class="tuomin-content-image" v-loading="tranformImageLoading">
          <div class="tuomin-icon">
            <i
              class="ri-image-line"
              style="font-size: 16px; margin-right: 2px"
            ></i>
            <span style="font-size: 12px">output</span>
          </div>
          <div
            v-if="resultImgSrc"
            class="tuomin-icon-download"
            @click="downImg"
          >
            <i
              class="ri-download-2-line"
              style="font-size: 16px; margin-right: 2px"
            ></i>
            <span style="font-size: 14px">下载</span>
          </div>
          <div style="height: 358px">
            <el-image
              style="height: 100%; width: 100%"
              :src="resultImgSrc"
              :preview-src-list="[resultImgSrc]"
            >
              <div slot="error" class="image-slot">
                <i style="font-size: 0" class="el-icon-picture-outline"></i>
              </div>
            </el-image>
          </div>
        </div>
      </el-col>
    </el-row>

    <div style="margin: 39px 0 21px 0">
      <span>{{ $t("example") }}：</span>
    </div>

    <div class="table-container">
      <div>
        <el-table border :data="tableData1" style="width: 100%">
          <el-table-column :label="$t('originPicture')" align="center">
            <template slot-scope="scope">
              <div style="width: 100%; height: 200px">
                <el-image
                  style="height: 100%; width: 100%"
                  :src="scope.row.imgSrc1"
                  :preview-src-list="[scope.row.imgSrc1]"
                >
                  <div slot="error" class="image-slot">
                    <i style="font-size: 0" class="el-icon-picture-outline"></i>
                  </div>
                </el-image>
              </div>
            </template>
          </el-table-column>
          <el-table-column :label="$t('desensitizationPicture')" align="center">
            <template slot-scope="scope">
              <div style="width: 100%; height: 200px">
                <el-image
                  style="height: 100%; width: 100%"
                  :src="scope.row.imgSrc2"
                  :preview-src-list="[scope.row.imgSrc2]"
                >
                  <div slot="error" class="image-slot">
                    <i style="font-size: 0" class="el-icon-picture-outline"></i>
                  </div>
                </el-image>
              </div>
            </template>
          </el-table-column>
          <el-table-column
            prop="mode"
            :label="$t('desensitizationObject')"
            align="center"
            header-align="center"
          >
          </el-table-column>
        </el-table>
      </div>
    </div>
  </div>
</template>

<script>
import { postImgDesensitization } from "~/apis/modules/desensitization";
export default {
  data() {
    return {
      ImageUrl: "",
      radio: 2,
      file: "",
      fileList: [],
      resultImgSrc: "",
      tranformImageLoading: false,
      tableData1: [
        {
          imgSrc1: "/img/origin.png",
          imgSrc2: "/img/tuomin.png",
          mode: this.$t("all"),
        },
      ],
      startFlag: false,
    };
  },
  components: {},
  methods: {
    // 选择文件、移除文件、上传文件成功/失败后，都会触发
    handleChangePicture(file, fileList) {
      let acceptList = ["jpg", "jpeg", "png"];
      // 根据文件名获取文件的后缀名
      let fileType = file.name.split(".").pop().toLowerCase();
      // 判断文件格式是否符合要求
      if (acceptList.indexOf(fileType) === -1) {
        this.$message.error(this.$t("limitFilesUpload"));
        return false;
      }
      // 判断文件大小是否符合要求
      if (file.size / 1024 / 1024 > 20) {
        this.$message.error(this.$t("limitSizeUpload"));
        return false;
      }

      this.ImageUrl = URL.createObjectURL(file.raw);
      this.file = file;
      this.fileList = fileList;
    },
    cancelUpload() {
      this.ImageUrl = "";
      this.file = "";
      this.fileList = [];
    },
    downImg() {
      const a = document.createElement("a");
      a.download = "result.png";
      a.style.display = "none";
      a.href = this.resultImgSrc;
      document.body.appendChild(a);
      a.click();
      document.body.removeChild(a);
    },
    startTranform() {
      if (!this.startFlag) {
        window.location.href = "/user/login";
        return;
      }
      if (!this.file) return;
      let fd = new FormData();
      fd.append("file", this.file.raw);
      this.tranformImageLoading = true;
      postImgDesensitization({ mode: this.radio }, fd)
        .then((res) => {
          const objectURL = URL.createObjectURL(res.data);
          this.resultImgSrc = objectURL;
          this.tranformImageLoading = false;
        })
        .catch((err) => {
          this.tranformImageLoading = false;
          if (err.response.status === 400) {
            const fr = new FileReader();
            fr.onload = (e) => {
              try {
                const jsonResult = JSON.parse(e.target.result);
                this.$message({
                  type: "error",
                  message: jsonResult.Message,
                });
              } catch (e) {
                this.$message({
                  type: "error",
                  message: this.$t("tranformImageFailed"),
                });
              }
            };
            fr.readAsText(err.response.data);
          } else {
            this.$message({
              type: "error",
              message: this.$t("tranformImageFailed"),
            });
          }
        });
    },
  },
  mounted() {
    const signEl = document.querySelector("#isSignd");
    this.startFlag = !!signEl.getAttribute("data-sign");
  },
  beforeDestroy() {},
};
</script>

<style scoped lang="less">
.tuomin-title {
  margin-top: 34px;
  h2 {
    display: flex;
    justify-content: center;
  }
  p {
    display: flex;
    justify-content: center;
    color: rgba(136, 136, 136, 0.87);
  }
}
.tuomin-content-image {
  border: 1px solid rgba(0, 0, 0, 0.1);
  border-radius: 5px;
  min-height: 358px;
  position: relative;
  .tuomin-icon {
    z-index: 99;
    display: flex;
    align-items: center;
    position: absolute;
    left: 0;
    top: 0;
    width: 66px;
    height: 22px;
    justify-content: center;
    color: rgba(136, 136, 136, 0.87);
    border-radius: 5px 0px 10px 0px;
    border: 1px solid rgba(0, 0, 0, 0.1);
  }
  .tuomin-icon-download {
    z-index: 99;
    display: flex;
    align-items: center;
    position: absolute;
    right: 10px;
    bottom: 10px;
    width: 66px;
    height: 30px;
    justify-content: center;
    color: rgba(255, 255, 255, 1);
    background-color: rgba(0, 0, 0, 0);
    background: transparent;
    border: 1px solid rgba(136, 136, 136, 0.5);
    border-radius: 4px;
    cursor: pointer;
  }
  .el-upload__text {
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .preview-image {
    height: 100%;
    width: 100%;
  }
  .tuomin-radio-model {
    display: flex;
    margin: 25px 12px 0 12px;
    .radio-label {
      font-weight: 600;
      margin-right: 4px;
      color: rgba(16, 16, 16, 1);
    }
  }
  .tuomin-button-model {
    text-align: right;
    margin: 22px 12px 0 12px;
  }
}
/deep/ .el-upload--picture-card {
  width: 100%;
  background: #ffff;
  border: none;
  border-bottom: 1px dashed #888;
  border-radius: 0;
  min-height: 230px;
}
/deep/ .el-upload-dragger {
  width: 100%;
  background: #ffff;
  border: none;
  border-radius: 0;
  height: 100%;
}
.table-container {
  margin-bottom: 16px;

  /deep/ .el-table__header {
    th {
      background: rgb(245, 245, 246);
      font-size: 14px;
      color: rgb(36, 36, 36);
      font-weight: 400;
    }
  }

  /deep/ .el-table__body {
    td {
      font-size: 14px;
    }
  }
}
.center {
  display: flex;
  justify-content: center;
}

@media screen and (max-width: 768px) {
  body {
    background-color: black;
  }
}
</style>
