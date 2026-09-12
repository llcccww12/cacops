; (function () {
  function PhoneVerifyCode(dom, options) {
    if (!dom) return;
    this.countDownNumber = 120;
    this.options = options;
    this.init(dom);
  }

  PhoneVerifyCode.prototype.init = function (dom) {
    if (!dom) return;
    this.dom = $(dom);
    this.isMoving = false;
    this.verifySucess = false;
    this.canSendCode = false;
    this.countDownEnd = false;
    this.imgID = '';
    // 初始化触摸冲突处理
    this.initTouchConflictHandler();
    this.eventInit();
    this.refreshImages();
    var wrap = this.dom.closest('div.use-type');
    var readonly = wrap.attr('readonly');
    if (readonly) {
      this.dom.find('.phone-area-c select').attr('disabled', true);
      this.dom.find('.phone-num-c input').attr('disabled', true);
      this.dom.find('.slide-bar-wrap').hide();
      this.dom.find('.verify-code-c').hide();
      this.dom.find('.modify-phone-number').show();
    }
    var oldPhoneNum = wrap.attr('ophonenumber');
    if (oldPhoneNum) {
      this.dom.find('input.phoneNumber').val(oldPhoneNum);
    }
    var showlabel = wrap.attr('showlabel');
    if (showlabel) {
      this.dom.find('._label-c ').show();
    } else {
      this.dom.find('._label-c ').hide();
    }
    var showNewpwd = wrap.attr('shownewpwd');
    if (showNewpwd) {
      this.dom.find('.new-pass-word-wrap').show();
    } else {
      this.dom.find('.new-pass-word-wrap').remove();
    }
    var autofocus = wrap.attr('autofocus');
    if (autofocus) {
      this.dom.find('input.phoneNumber').focus();
    }
    var verifyCodeNoRequired = wrap.attr('verifycodenorequired');
    if (verifyCodeNoRequired) {
      this.dom.find('input.verifyCode').removeAttr('required');
    }
  };

  PhoneVerifyCode.prototype.eventInit = function () {
    if (!this.dom) return;
    var self = this;
    var clientX = 0, oLeft = 0, imgHideTimer = null;
    this.dom.find('.slide-bar-bg').on('mouseenter touchstart', function (e) {
      if (self.verifySucess) return;
      imgHideTimer && clearTimeout(imgHideTimer);
      self.dom.find('.slide-image-big').slideDown();
    });
    this.dom.find('.slide-bar-bg').on('mouseleave', function (e) {
      imgHideTimer && clearTimeout(imgHideTimer);
      imgHideTimer = setTimeout(function () {
        self.dom.find('.slide-image-big').slideUp();
      }, 200);
    });
    this.dom.find('.slide-image-big').on('mouseenter', function (e) {
      imgHideTimer && clearTimeout(imgHideTimer);
    });
    this.dom.find('.slide-image-big').on('mouseleave', function (e) {
      imgHideTimer && clearTimeout(imgHideTimer);
      imgHideTimer = setTimeout(function () {
        self.dom.find('.slide-image-big').slideUp();
      }, 200);
    });

    function mouseMove(e) {
      var _clientX = e.clientX !== undefined ? e.clientX : e.targetTouches[0].clientX;
      var offset = _clientX - clientX;
      var triggerEl = self.dom.find('.slide-trigger');
      var triggerWidth = triggerEl.width();
      var parentWidth = triggerEl.parent().width();
      var maxLeft = parentWidth - triggerWidth;
      var left = oLeft + offset;
      if (oLeft + offset < 0) left = 0;
      if (oLeft + offset > maxLeft) left = maxLeft;
      triggerEl.css('left', left + 'px');
      self.dom.find('.slide-bar').css('width', left + triggerWidth);
      var imageBigWidth = self.dom.find('.slide-image-big').width();
      var imageSmallWidth = self.dom.find('.slide-image-small').width();
      self.dom.find('.slide-image-small').css('left', left / maxLeft * (imageBigWidth - imageSmallWidth) + 'px');
      self.isMoving = true;
      self.dom.find('.slide-txt').hide();
      imgHideTimer && clearTimeout(imgHideTimer);
    }

    function mouseUp(e) {
      $(document).off('mousemove', mouseMove);
      $(document).off('mouseup', mouseUp);
      $(document).off('touchmove', mouseMove);
      $(document).off('touchend', mouseUp);
      self.isMoving = false;
      $.ajax({
        url: '/verifySlideImage',
        type: 'post',
        dataType: 'json',
        data: {
          slide_id: self.imgID,
          x: parseInt(self.dom.find('.slide-image-small').position().left / self.dom.find('.slide-image-big').attr('scale'))
        },
        success: function (res) {
          if (res && res.Code === 0) {
            self.verifySucess = true;
            self.canSendCode = true;
            self.dom.find('.slide-bar').addClass('sucess');
            self.dom.find('.slide-trigger').addClass('sucess');
            self.dom.find('.slide-trigger .icon').hide();
            self.dom.find('.slide-trigger .icon.check').show();
            self.dom.find('.slide-image-big').slideUp();
            self.dom.find('.verify-code-send-btn').removeClass('__disabled');
          } else {
            self.dom.find('.slide-bar').addClass('error');
            self.dom.find('.slide-trigger').addClass('error');
            self.dom.find('.slide-trigger .icon').hide();
            self.dom.find('.slide-trigger .icon.close').show();
            setTimeout(function () {
              self.refreshImages();
            }, 300);
          }
        },
        error: function (err) {
          self.dom.find('.slide-bar').addClass('error');
          self.dom.find('.slide-trigger').addClass('error');
          setTimeout(function () {
            self.refreshImages();
          }, 300);
        }
      });
    }

    function mouseDown(e) {
      if (self.verifySucess) return;
      clientX = e.clientX !== undefined ? e.clientX : e.targetTouches[0].clientX;
      oLeft = $(this).position().left;
      $(document).on('mousemove', mouseMove);
      $(document).on('mouseup', mouseUp);
      $(document).on('touchmove', mouseMove);
      $(document).on('touchend', mouseUp);
    }

    this.dom.find('.slide-trigger').on('mousedown', mouseDown);
    this.dom.find('.slide-trigger').on('touchstart', mouseDown);

    this.dom.find('.verify-code-send-btn').on('click', function () {
      if (!self.canSendCode) return;
      if (self.countDownEnd) {
        self.refreshImages();
        return;
      }
      var phoneNumber = self.dom.find('.phone-num-c input').val();
      if (!/^1[3-9]\d{9}$/.test(phoneNumber)) {
        self.dom.find('.phone-num-c').addClass('error');
        return;
      } else {
        self.dom.find('.phone-num-c').removeClass('error');
        var useType = self.dom.closest('div.use-type').attr('usetype') || 0;
        $.ajax({
          url: '/sendVerifyCode',
          type: 'post',
          dataType: 'json',
          data: {
            phone_number: phoneNumber,
            mode: useType, // 0注册，1登录 ，2修改手机号，3找回密码
            slide_id: self.imgID,
          },
          success: function (res) {
            if (res && res.Code === 0) {
              self.countDown();
            } else {
              if ($('.ui.negative.message').length) {
                $('.ui.negative.message').eq(0).show().find('p').text(res.Message);
              } else {
                $('body').toast({
                  message: res.Message,
                  showProgress: 'bottom',
                  showIcon: 'warning circle',
                  class: 'warning',
                  position: 'top right',
                });
              }
              self.refreshImages();
            }
          },
          error: function (err) {
            console.log(err);
          }
        });
      }
    });

    this.dom.find('.modify-phone-number a').on('click', function () {
      self.dom.find('.phone-area-c select').attr('disabled', false);
      self.dom.find('.phone-num-c input').attr('disabled', false);
      self.dom.find('.slide-bar-wrap').css('display', 'flex');
      self.dom.find('.verify-code-c').css('display', 'flex');
      self.dom.find('.modify-phone-number').hide();
      self.refreshImages();
    });
  };

  PhoneVerifyCode.prototype.refreshImages = function () {
    this.isMoving = false;
    this.verifySucess = false;
    this.canSendCode = false;
    this.countDownEnd = false;
    this.imgID = '';
    this.dom.find('.slide-bar').removeClass('sucess error').css('width', '30px');
    this.dom.find('.slide-trigger').removeClass('sucess error').css('left', '0px');
    var scale = this.dom.find('.slide-bar-bg').width() / 391;
    this.dom.find('.slide-image-big').css('transform', `scale(${scale})`).attr('scale', scale);
    this.dom.find('.slide-trigger .icon').hide();
    this.dom.find('.slide-trigger .icon.arrow').show();
    this.dom.find('.slide-txt').show();
    this.dom.find('.slide-image-small').css('left', '0');
    this.dom.find('.verify-code-send-btn').addClass('__disabled');
    var self = this;
    $.ajax({
      url: '/slideImage',
      type: 'get',
      success: function (res) {
        if (res && res.Code === 0) {
          self.imgID = res.Message;
          self.dom.find('.slide-image-big').css('background', `url("/slideimage/${res.Message}.png")`);
          self.dom.find('.slide-image-small').css('background', `url("/slideimage/${res.Message}screenshot.png")`);
        }
      },
      error: function (err) {
        console.log(err);
      }
    });
  };

  PhoneVerifyCode.prototype.countDown = function () {
    var self = this;
    var sendBtnEl = this.dom.find('.verify-code-send-btn');
    var count = this.countDownNumber;
    sendBtnEl.addClass('__disabled').text(count + (self.options && self.options.Lang && self.options.Lang.second_resend ? self.options.Lang.second_resend : 'S后重发'));
    this.canSendCode = false;
    this.countDownEnd = false;
    var timer = setInterval(function () {
      count--;
      sendBtnEl.addClass('__disabled').text(count + (self.options && self.options.Lang && self.options.Lang.second_resend ? self.options.Lang.second_resend : 'S后重发'));
      if (count <= 0) {
        sendBtnEl.removeClass('__disabled').text((self.options && self.options.Lang && self.options.Lang.get_verification_code ? self.options.Lang.get_verification_code : '获取验证码'));
        clearInterval(timer);
        self.canSendCode = true;
        self.countDownEnd = true;
        self.refreshImages();
      }
    }, 1000);
  };
  PhoneVerifyCode.prototype.initTouchConflictHandler = function() {
    var self = this;
    var sliderTrigger = this.dom.find('.slide-trigger')[0];
    var sliderBg = this.dom.find('.slide-bar-bg')[0];
    
    if (!sliderTrigger) return;
    
    // 最小化影响的CSS方案
    var style = document.createElement('style');
    style.textContent = `
      /* 只对滑块背景应用，不影响内部元素 */
      .slide-bar-bg {
        touch-action: pan-y;
      }
      
      /* 滑块本身不应用touch-action，避免影响功能 */
      .slide-trigger {
        -webkit-touch-callout: none;
        -webkit-user-select: none;
        user-select: none;
      }
      
      /* 确保所有图片相关元素可以正常交互 */
      .slide-image-big *,
      .slide-image-small * {
        touch-action: auto !important;
        pointer-events: auto !important;
      }
    `;
    document.head.appendChild(style);
    
    // 只在必要的时候阻止默认行为
    var touchStartX = 0;
    var isHorizontalMove = false;
    
    sliderBg.addEventListener('touchstart', function(e) {
      touchStartX = e.touches[0].clientX;
      isHorizontalMove = false;
    }, { passive: true });
    
    sliderBg.addEventListener('touchmove', function(e) {
      if (!e.touches || e.touches.length === 0) return;
      
      var deltaX = e.touches[0].clientX - touchStartX;
      
      // 如果检测到明显的水平滑动，阻止浏览器后退
      if (Math.abs(deltaX) > 10) {
        isHorizontalMove = true;
        e.preventDefault();
      }
    }, { passive: false });
    
    sliderBg.addEventListener('touchend', function() {
      // 短暂延迟后重置状态
      setTimeout(function() {
        isHorizontalMove = false;
      }, 100);
    }, { passive: true });
  };
  window.PhoneVerifyCode = PhoneVerifyCode;
})();
