<template>
  <div>
    <div class="ui container">
      <div class="ui placeholder segment bgtask-none">
        <div class="bgtask-content-header">
          <h2 class="wx-title">微信扫码认证</h2>
          <p v-if="scanType == 'modify'" class="wx-desc-top">请绑定新的微信</p>
          <p v-else class="wx-desc-top">请绑定微信，然后再使用启智算力环境</p>
        </div>
        <div class="wx-login">
          <div class="qrcode">
            <img class="wx-qrcode" :src="wx_qrcodeImg" v-show="!!wx_qrcodeImg">
            <span class="el-icon-loading" v-show="wxLoading"></span>
            <img class="wx-qrcode-reload" v-if="qrCodeReload" @click="getWxQrcode(true)"
              src="https://files.wondercv.com/auth/qrcode_reload.png">
          </div>
          <div class="wx-desc-bottom" style="color:#919191">微信扫码关注公众号即可完成绑定</div>
        </div>
        <div class="user-agreement">
          <i class="ri-information-line orange"></i>绑定微信代表已阅读并接受
          <a href="/home/term/">OpenI启智社区AI协作平台使用协议</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
const { _AppSubUrl, _StaticUrlPrefix, csrf } = window.config;
export default {
  components: {
  },
  data() {
    return {
      scanType: '', // modify or other
      wx_qrcodeImg: '',
      wxLoading: false,
      isLogin: false,
      SceneStr: '',
      status: '',
      qrCodeReload: false,
      timer: null,
    };
  },
  methods: {
    getWxQrcode(reloadFlag) {
      if (reloadFlag) {
        this.qrCodeReload = false
      }
      this.wxLoading = true
      this.timer && clearInterval(this.timer)
      this.$axios.get('/authentication/wechat/qrCode4Bind').then((res) => {
        let ticket = res.data.data.Ticket
        this.SceneStr = res.data.data.SceneStr
        this.wx_qrcodeImg = 'https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=' + ticket
        this.wxLoading = false
        this.WxqrCheck()
      })
    },
    WxqrCheck() {
      this.timer && clearInterval(this.timer)
      this.timer = setInterval(() => {
        this.$axios.get(`/authentication/wechat/bindStatus?sceneStr=${this.SceneStr}`).then((res) => {
          this.status = res.data.data.status
          if (this.status === 0) { // scaning

          } else if (this.status === 9) { // expired
            this.qrCodeReload = true
            this.timer && clearInterval(this.timer)
          } else if (this.status === 2) { // success
            clearInterval(this.timer)
            $('.alert').html('绑定成功!').removeClass('alert-danger').addClass('alert-success').show().delay(1500).fadeOut();
            window.location.href = this.$Cookies.get('redirect_to')
          }
        })
      }, 1500);
    }
  },
  mounted() {
    const urlSearchParams = new URLSearchParams(window.location.search)
    this.scanType = urlSearchParams.get('type')
    this.getWxQrcode(false)
  },
  beforeDestroy() {
    this.timer && clearInterval(this.timer)
  }
};
</script>

<style scoped>
.el-icon-loading {
  position: absolute;
  font-size: 32px;
  color: #bcbcbc;
  top: calc(50% - 16px);
  left: calc(50% - 16px);
  animation: rotating 2s linear infinite;
}

.wx-title {
  font-family: SourceHanSansSC-medium;
  font-size: 24px;
  color: rgba(16, 16, 16, 100);
}

.wx-desc-top {
  font-size: 14px;
  color: rgba(16, 16, 16, 100);
  font-family: SourceHanSansSC-light;
}

.qrcode {
  width: 200px;
  height: 200px;
  line-height: 180px;
  text-align: center;
  position: relative;
}

.wx-login {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 24px;
}

.wx-qrcode {
  width: 100%;
  height: 100%;
}

.wx-qrcode-reload {
  width: 100%;
  height: 100%;
  left: 0px;
  top: 0px;
  position: absolute;
  cursor: pointer;
}

.wx-desc-bottom {
  color: rgba(145, 145, 145, 100);
  font-size: 14px;
  font-family: SourceHanSansSC-regular;
}

.user-agreement {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-top: 51px;
  color: rgba(16, 16, 16, 100);
  font-size: 14px;
  font-family: SourceHanSansSC-regular;
}

.user-agreement .orange {
  color: #f2711c;
  margin-right: 3px;
}

.user-agreement a {
  margin-left: 2px;
}
</style>
