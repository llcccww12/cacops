const { AppSubUrl, AppVer, csrf, NotificationSettings } = window.config;

export function initNotificationsTable() {
  $('#notification_table .button').on('click', async function () {
    const data = await updateNotification(
      $(this).data('url'),
      $(this).data('status'),
      $(this).data('page'),
      $(this).data('q'),
      $(this).data('notification-id'),
    );

    $('#notification_div').replaceWith(data);
    initNotificationsTable();
    await updateNotificationCount();

    return false;
  });
  const notificationDiv = $('#notice-container-json');
  if (!notificationDiv.length) {
    return;
  } else {
    getNotice()
  }
}

export function initNotificationCount() {
  const notificationCount = $('.notification_count');
  if (!notificationCount.length) {
    return;
  }
  if (NotificationSettings.EventSourceUpdateTime > 0 && !!window.EventSource && !!window.SharedWorker) {
    // Try to connect to the event source via the shared worker first
    const worker = new SharedWorker(`${AppSubUrl}/static/eventsource.sharedworker.js?v=${AppVer}`, 'notification-worker');
    worker.addEventListener('error', (event) => {
      console.error('worker error', event);
    });
    worker.port.addEventListener('messageerror', () => {
      console.error('unable to deserialize message');
    });
    worker.port.postMessage({
      type: 'start',
      url: `${window.location.origin}${AppSubUrl}/user/events`,
    });
    worker.port.addEventListener('message', async (event) => {
      if (!event.data || !event.data.type) {
        console.error('unknown worker message event', event);
        return;
      }
      if (event.data.type === 'notification-count') {
        try {
          const data = JSON.parse(event.data.data);
          const notificationCount = $('.notification_count');
          if (data.Count == 0) {
            notificationCount.addClass('hidden');
          } else {
            notificationCount.removeClass('hidden');
          }
          notificationCount.text(`${data.Count}`);
          await updateNotificationTable();
        } catch (error) {
          console.error(error);
        }
      }
      if (event.data.type === 'reward-operation') {
        try {
          const data = JSON.parse(event.data.data);
          if (data.OperateType !== 'INCREASE') {
            return;
          }
          const notice = $(`
            <div class="__notice-wrap" style="display:none;z-index:99;">
              <style>
              .__notice-container {
                position: fixed;
                width:400px;
                height:36px;
                right: 20px;
                top: 70px;
                animation:__notice 2s;
                -webkit-animation:__notice 2s;
                text-align:right;
                display: flex;
                justify-content: flex-end;
              }
              .__notice-content {
                background: rgb(252, 202, 0);
                padding: 4px 16px;    
                box-shadow: rgb(135 56 0 / 20%) 0px 2px 6px 0px;
                border-radius: 4px;
                font-size: 14px;
                display: flex;
                color: #101010;
                align-items: center;
              }
              @keyframes __notice {
                0%, 20%, 53%, 80%, 100% {
                  animation-timing-function: cubic-bezier(0.215, 0.61, 0.355, 1);
                  transform: translate3d(0px, 0px, 0px);
                }
                40%, 43% {
                  animation-timing-function: cubic-bezier(0.755, 0.05, 0.855, 0.06);
                  transform: translate3d(0px, -30px, 0px);
                }
                70% {
                  animation-timing-function: cubic-bezier(0.755, 0.05, 0.855, 0.06);
                  transform: translate3d(0px, -15px, 0px);
                }
                90% {
                  transform: translate3d(0px, -4px, 0px);
                }
              }
              </style>
              <div class="__notice-container">
                <div class="__notice-content">
                <i class="ri-copper-diamond-line" style="padding-right:8px;color:white;font-size:18px;"></i>
                  <span class="__notice-text">${data.OperateType === 'INCREASE' ? ('奖励积分 + ' + data.Amount.toFixed(2)) : ('扣减积分 ' + data.Amount.toFixed(2))}</span>
                </div>
              </div>
            </div>
          `);
          $('body').append(notice);
          notice.fadeIn();
          setTimeout(() => {
            notice.fadeOut();
          }, 3000);
          setTimeout(() => {
            notice.remove();
          }, 5000);
        } catch (error) {
          console.error(error);
        }
      }
      if (event.data.type === 'logout') {
        if (event.data.data !== 'here') {
          return;
        }
        worker.port.postMessage({
          type: 'close',
        });
        worker.port.close();
        window.location.href = AppSubUrl;
      }
    });
    worker.port.addEventListener('error', (e) => {
      console.error('worker port error', e);
    });
    worker.port.start();
    window.addEventListener('beforeunload', () => {
      worker.port.postMessage({
        type: 'close',
      });
      worker.port.close();
    });
    return;
  }

  if (NotificationSettings.MinTimeout <= 0) {
    return;
  }

  const fn = (timeout, lastCount) => {
    setTimeout(async () => {
      await updateNotificationCountWithCallback(fn, timeout, lastCount);
    }, timeout);
  };

  fn(NotificationSettings.MinTimeout, notificationCount.text());
}
function getNotice() {
  $('.notice-container .ui.dimmer').css('display', 'block')
  $.ajax({
    type: "GET",
    url: "/dashboard/invitation",
    dataType: "json",
    data: {
      filename: 'notice/notice.json',
    },
    success: function (data) {
      $('.notice-container .ui.dimmer').css('display', 'none')
      const isZh = document.documentElement.getAttribute('lang') == 'zh-CN';
      if (!data) return;
      try {
        const noticeList = JSON.parse(data).Notices || [];
        const noticeEl = $('.notice-container');
        for (let i = 0, iLen = noticeList.length; i < iLen; i++) {
          let noticeObj = noticeList[i];
          noticeEl.append(`<div class="notice-row">
            <div class="notice-title">
                <a class="_hm-notice" href="${noticeObj.Link}">
                <i class="ri-arrow-right-s-line" style="vertical-align:-2px;"></i>
                <span>${isZh ? noticeObj.Title : (noticeObj.Title_en || noticeObj.Title)}</span>
              </a>
            </div>
            <div class="notice-time">${noticeObj.Date || ''}</div>
          </div>`);
        }
      } catch (e) {
        console.info(e);
      }
    },
    error: function (err) {
      $('.notice-container .ui.dimmer').css('display', 'none')
      console.info(err);
    }
  });
}

async function updateNotificationCountWithCallback(callback, timeout, lastCount) {
  const currentCount = $('.notification_count:eq(0)').text();
  if (lastCount != currentCount) {
    callback(NotificationSettings.MinTimeout, currentCount);
    return;
  }

  const newCount = await updateNotificationCount();
  let needsUpdate = false;

  if (lastCount != newCount) {
    needsUpdate = true;
    timeout = NotificationSettings.MinTimeout;
  } else if (timeout < NotificationSettings.MaxTimeout) {
    timeout += NotificationSettings.TimeoutStep;
  }

  callback(timeout, newCount);
  if (needsUpdate) {
    await updateNotificationTable();
  }
}

async function updateNotificationTable() {
  const notificationDiv = $('#notification_div');
  if (notificationDiv.length > 0) {
    const data = await $.ajax({
      type: 'GET',
      url: `${AppSubUrl}/notifications?${notificationDiv.data('params')}`,
      data: {
        'div-only': true,
      }
    });
    notificationDiv.replaceWith(data);
    initNotificationsTable();
  }
}

async function updateNotificationCount() {
  const data = await $.ajax({
    type: 'GET',
    url: `${AppSubUrl}/api/v1/notifications/new`,
    headers: {
      'X-Csrf-Token': csrf,
    },
  });

  const notificationCount = $('.notification_count');
  if (data.new == 0) {
    notificationCount.addClass('hidden');
  } else {
    notificationCount.removeClass('hidden');
  }

  notificationCount.text(`${data.new}`);

  return `${data.new}`;
}

async function updateNotification(url, status, page, q, notificationID) {
  if (status !== 'pinned') {
    $(`#notification_${notificationID}`).remove();
  }

  return $.ajax({
    type: 'POST',
    url,
    data: {
      _csrf: csrf,
      notification_id: notificationID,
      status,
      page,
      q,
      noredirect: true,
    },
  });
}
