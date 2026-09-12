
import "./features/letteravatar.js";
import { i18nVue } from "./features/i18nVue.js";
import './features/globalModalDlg.js';
import './features/ad.js';
import { gsap } from "gsap";

import { ScrollTrigger } from "gsap/ScrollTrigger";

// gsap.registerPlugin(ScrollTrigger);

// let panels = gsap.utils.toArray(".pannel_z");
// const tl = gsap.timeline({
//     scrollTrigger: {
//       trigger: "._hm-repo-container",
//       start: "top+=1px top",
//       end: "+=1100", // 明确设置足够的滚动距离
//       scrub: 1,
//       pin: true,
//     //   pinSpacing: false,
//       markers: {startColor: "green", endColor: "red", fontSize: "12px"},
//     }
// });
// tl.to(panels[1], {
//     y: "-135%",
//     duration: 2
// }, 0.5);
// tl.to(panels[2], {
//     y: "-155%",
//     duration: 2
// }, 1);
// tl.to(panels[0], {
//     scale: 0.93,
//     duration: 2
// }, 1);
// tl.to(panels[1], {
//     scale: 0.96,
//     duration: 2
// },1.5);
// 封装动画初始化逻辑
function initAnimations() {
    
    // 检测屏幕宽度是否符合桌面端条件
    if (window.innerWidth > 767) {
      gsap.registerPlugin(ScrollTrigger);

      const panels = gsap.utils.toArray(".pannel_z");
      const container = document.querySelector("._hm-repo-container");

      // 确保容器存在
      if (!container) return;

      const tl = gsap.timeline({
        scrollTrigger: {
          trigger: container,
          start: "top+=1px top",
          end: "+=2400",
          scrub: 1,
          pin: true,
          //   markers: { startColor: "green", endColor: "red", fontSize: "12px" }
        }
      });

      // 组织动画配置
      const animations = [
        // 面板1 (索引0): 在滚动中期开始缩放
        { element: panels[0], props: { scale: 0.90, duration: 3 }, position: 1.0 }, // 2→3秒
        // 面板1
        { element: panels[1], props: { y: "-71%", duration: 3 }, position: 0.8 },    // 2→3秒
        { element: panels[1], props: { scale: 0.93, duration: 3 }, position: 2.2 },  // 2→3秒
        // 面板2
        { element: panels[2], props: { y: "-260%", duration: 3 }, position: 2.2 },   // 2→3秒
        { element: panels[2], props: { scale: 0.95, duration: 3 }, position: 3.0 },  // 2→3秒
        // 面板3
        { element: panels[3], props: { y: "-253%", duration: 3 }, position: 2.8 },   // 2→3秒
        { element: panels[3], props: { scale: 0.97, duration: 3 }, position: 3.5 }   // 2→3秒
      ];

      // 批量添加动画
      animations.forEach(({ element, props, position }) => {
        tl.to(element, props, position);
      });

      // 返回时间轴实例以便后续操作
      return tl;
    }
}

// 初始化动画
let timeline = initAnimations();

// 动态响应窗口变化 (可选)
window.addEventListener('resize', () => {
    // 如果已有时间轴实例且切换到移动端
    if (timeline && window.innerWidth <= 767) {
        timeline.scrollTrigger.kill();
        timeline.kill();
        // 重置元素样式
        gsap.utils.toArray(".pannel_z").forEach(panel => {
        gsap.set(panel, { clearProps: "all" });
        });

        timeline = null;
    }
    // 如果是桌面端且未初始化
    else if (!timeline && window.innerWidth > 767) {
        timeline = initAnimations();
    }
});

$.fn.tab.settings.silent = true;
function isEmpty(str){
    if(typeof str == "undefined" || str == null || str == ""){
       return true;
    }
    return false;
}

var token;
if(isEmpty(token)){
	var meta = $("meta[name=_uid]");
	if(!isEmpty(meta)){
        token = meta.attr("content");
	}
}


$(".link-action").on("click", linkAction);
$('.ui.dropdown').dropdown();

$.get(`${window.config.StaticUrlPrefix}/img/svg/icons.svg`, (data) => {
    const div = document.createElement("div");
    div.style.display = "none";
    div.innerHTML = new XMLSerializer().serializeToString(data.documentElement);
    document.body.insertBefore(div, document.body.childNodes[0]);
});

function linkAction(e) {
  e.preventDefault();
  const $this = $(this);
  const redirect = $this.data("redirect");
  $.post($this.data("url"), {
    _csrf: token,
  }).done((data) => {
    if (data.redirect) {
      window.location.href = data.redirect;
    } else if (redirect) {
      window.location.href = redirect;
    } else {
      window.location.reload();
    }
  });
}



// 全局变量
let maxSize = 20;
let html = document.documentElement;
let lang = html.attributes["lang"]
let isZh = true;
if(lang != null && lang.nodeValue =="en-US" ){
    isZh=false;
}

window.i18n = i18nVue[isZh ? 'CN' : 'US'];

let refreshInterval = null;
const refreshIntervalTime = 30000; // 10秒刷新一次
let isPolling = false; // 标记当前是否在使用轮询方式
let swiperNewMessage = null; // Swiper 实例
let isFirstBannerActive = true; // 第一个 banner 是否活跃
let messageQueue = [];
let output = ''
document.onreadystatechange = function () {
    if(document.readyState != "complete"){
        return;
    }
    console.log("Start to open WebSocket." + document.readyState);
    initSwiper()
    output = document.getElementById("newmessage");
    // 首先检查是否支持WebSocket
    checkNotificationMethod();

    // 注册页面可见性变化事件
    if (document.addEventListener) {
        document.addEventListener('visibilitychange', handleVisibilityChange, false);
    }
    
    // 注册页面卸载事件
    window.addEventListener('beforeunload', cleanupTimers);
    window.addEventListener('pagehide', cleanupTimers);
}
// 初始化 Swiper
function initSwiper() {
    if (!swiperNewMessage && $('.newslist').length) {
        try {
            swiperNewMessage = new Swiper(".newslist", {
                direction: "vertical",
                slidesPerView: 7,
                loop: true,
                spaceBetween: 5,
                autoplay: isFirstBannerActive ? {
                    delay: 2500,
                    disableOnInteraction: false,
                } : false,
            });
            if (swiperNewMessage.updateSlides) {
                swiperNewMessage.updateSlides();
            }
            console.log("Swiper initialized successfully");
        } catch (error) {
            console.error("Failed to initialize Swiper:", error);
        }
    }
}
function toggleSwiperAutoplay(enable) {
    if (swiperNewMessage) {
        if (enable) {
            swiperNewMessage.autoplay.start();
        } else {
            swiperNewMessage.autoplay.stop();
        }
    }
}
// 页面可见性变化处理
function handleVisibilityChange() {
    if (document.hidden) {
        // 页面隐藏，暂停轮询
        pausePolling();
    } else {
        // 页面显示，恢复轮询
        resumePolling();
    }
}
// 恢复轮询
function resumePolling() {
    if (isPolling && !refreshInterval) {
        // 立即获取一次数据
        fetchLatestActions();
        
        // 重新设置定时器
        refreshInterval = setInterval(fetchLatestActions, refreshIntervalTime);
        console.log("Notification polling resumed (page visible)");
    }
}
// 暂停轮询
function pausePolling() {
    if (isPolling && refreshInterval) {
        clearInterval(refreshInterval);
        refreshInterval = null;
        console.log("Notification polling paused (page hidden)");
    }
}
// 清理定时器
function cleanupTimers() {
    if (refreshInterval) {
        clearInterval(refreshInterval);
        refreshInterval = null;
        console.log("Notification timers cleaned up");
    }
    
    // 移除事件监听器
    if (document.removeEventListener) {
        document.removeEventListener('visibilitychange', handleVisibilityChange, false);
    }
    window.removeEventListener('beforeunload', cleanupTimers);
    window.removeEventListener('pagehide', cleanupTimers);
}
// 获取最新动态
function fetchLatestActions() {
    
    fetch('/action/latest_actions')
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            if (data.actions && data.actions.length > 0) {
                console.log("Received latest actions:", data.actions);
                processActions(data.actions);
            }
        })
        .catch(error => {
            console.error("Failed to fetch latest actions:", error);
        });
}
function checkNotificationMethod() {
    // 调用接口检查通知方式
    fetch('/action/latest_actions')
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            if (data.socket === true) {
                // 使用WebSocket方式
                console.log("Using WebSocket for notifications");
                isPolling = false;
                initWebSocket();
            } else {
                // 使用轮询方式
                console.log("Using polling for notifications");
                isPolling = true;
                initPolling(data.actions);
            }
        })
        .catch(error => {
            console.error("Failed to check notification method:", error);
            // 如果接口调用失败，默认使用轮询方式
            isPolling = true;
            initPolling([]);
        });
}
function initWebSocket() {
    
    let url = "ws://" + document.location.host + "/action/notification";
    
    if(document.location.host === "git.openi.org.cn" || document.URL.startsWith("https")){
       url = "wss://" + document.location.host + "/action/notification";
    }
    
    const socket = new WebSocket(url);
    
    socket.onopen = function () {
        messageQueue = [];
        console.log("WebSocket connected.");
    };

    socket.onmessage = function (e) {
        handleNotificationMessage(e.data);
    };
    
    socket.onerror = function (error) {
        console.error("WebSocket error:", error);
        // WebSocket连接失败，切换到轮询方式
        socket.close();
        isPolling = true;
        initPolling([]);
    };
    
    socket.onclose = function () {
        console.log("WebSocket disconnected, switching to polling");
        // WebSocket断开，切换到轮询方式
        isPolling = true;
        initPolling([]);
    };
    
    // 页面离开时关闭WebSocket
    var originalBeforeUnload = window.onbeforeunload;
    window.onbeforeunload = function() {
        if (socket && socket.readyState === WebSocket.OPEN) {
            socket.close();
        }
        if (originalBeforeUnload) {
            return originalBeforeUnload.apply(this, arguments);
        }
    };
}
function initPolling(initialActions) {
    
    // 清除已有的定时器
    cleanupTimers();
    
    // 先处理初始数据
    if (initialActions && initialActions.length > 0) {
        processActions(initialActions);
    }
    
    // 设置定时轮询
    refreshInterval = setInterval(fetchLatestActions, refreshIntervalTime);
    console.log("Notification polling started, refreshing every " + refreshIntervalTime + "ms");
}
function processActions(actions) {
    if (!output) return;
    let msgHtml = ''
    for (let i = 0; i < actions.length; i++) {
        const record = actions[i];
        
        if (messageQueue.length > maxSize) {
            messageQueue.splice(0, 1);
        }
        messageQueue.push(record);
    }
    const currentTime = new Date().getTime();
    for(let i = 0; i < messageQueue.length; i++) {
        const record = messageQueue[i];
        
        // 检查必要的函数是否存在
        if (typeof getAction !== 'function') {
            console.error("getAction function is not defined");
            return;
        }
        
        let actionName = getAction(record.OpType, isZh);
        
        if(record.ActUser == null) {
            console.log("receive action type=" + record.OpType + " name=" + actionName + " but user is null.");
            continue;
        }
        
        if(record.OpType === "24") {
            if(record.Content.indexOf("true") !== -1){
                continue;
            }
        }
        
        const recordPrefix = getMsg(record);
        msgHtml += formatNotificationItem(record, actionName, recordPrefix, currentTime);
    }
    
    if (msgHtml) {
        output.innerHTML = msgHtml;
        
        // 确保元素存在再操作
        if ($('#homenews p').length > 0) {
            $('#homenews p').show();
        }
        
        // 检查并更新 Swiper
        updateSwiper();
    }
}
// 安全的更新 Swiper
function updateSwiper() {
    try {
        if (swiperNewMessage) {
            if (typeof swiperNewMessage.updateSlides === 'function') {
                swiperNewMessage.updateSlides();
            }
            if (typeof swiperNewMessage.updateProgress === 'function') {
                swiperNewMessage.updateProgress();
            }
        } else {
            // 如果 Swiper 未初始化，尝试初始化
            initSwiper();
        }
    } catch (error) {
        console.error("Error updating Swiper:", error);
    }
}
function formatNotificationItem(record, actionName, recordPrefix, currentTime) {
    let html = "";
    let branch = "";
    
    if(record.OpType == "6" || record.OpType == "10" || record.OpType == "12" || record.OpType == "13") {
        html += recordPrefix + actionName;
        html += " <a href=\"" + getIssueLink(record) + "\" rel=\"nofollow\">" + getIssueText(record) + "</a>"
    }
    else if(record.OpType == "7" || record.OpType == "11" || record.OpType == "14" || record.OpType == "15" || record.OpType == "22"
    || record.OpType == "23") {
        html += recordPrefix + actionName;
        html += " <a href=\"" + getPRLink(record) + "\" rel=\"nofollow\">" + getPRText(record) + "</a>"
    }
    else if(record.OpType == "1") {
        html += recordPrefix + actionName;
        html += " <a href=\"" + getRepoLink(record) + "\" rel=\"nofollow\">" + getRepotext(record) + "</a>"
    }
    else if(record.OpType == "5") {
        branch = "<a href=\"" + getRepoLink(record) + "/src/branch/" + encodeURI(record.RefName) +  "\" class=\"special\" rel=\"nofollow\">" + record.RefName + "</a>"
        actionName = actionName.replace("{branch}", branch);
        html += recordPrefix + actionName;
        html += " <a href=\"" + getRepoLink(record) + "\" rel=\"nofollow\">" + getRepotext(record) + "</a>"
    } else if(record.OpType == "9") {
        branch = "<a href=\"" + getRepoLink(record) + "/src/tag/" + encodeURI(record.RefName) +  "\" class=\"special\" rel=\"nofollow\">" + record.RefName + "</a>"
        actionName = actionName.replace("{branch}", branch);
        html += recordPrefix + actionName;
        html += " <a href=\"" + getRepoLink(record) + "\" rel=\"nofollow\">" + getRepotext(record) + "</a>"
    }
    else if(record.OpType == "17") {
        actionName = actionName.replace("{deleteBranchName}", record.RefName);
        var repoLink = "<a href=\"" + getRepoLink(record) + "\" rel=\"nofollow\">" + getRepotext(record) + "</a>"
        actionName = actionName.replace("{repoName}", repoLink);
        html += recordPrefix + actionName;
    }
    else if(record.OpType == "2") {
        actionName = actionName.replace("{oldRepoName}", record.Content);
        html += recordPrefix + actionName;
        html += " <a href=\"" + getRepoLink(record) + "\" rel=\"nofollow\">" + getRepotext(record) + "</a>"
    }
    else if(record.OpType == "25" || record.OpType == "29" || record.OpType == "39" || record.OpType == "40" || record.OpType == "41"
        || record.OpType == "43"|| record.OpType == "44"|| record.OpType == "45"|| record.OpType == "46"|| record.OpType == "47"
        || record.OpType == "48" || record.OpType == "49" || record.OpType == "53" || record.OpType == "54" || record.OpType == '55'
        || record.OpType == '56' || record.OpType == '58' || record.OpType == '59' || record.OpType == '60' || record.OpType == '61'
        || record.OpType == '65' || record.OpType == '66'  ||  record.OpType == "26" || record.OpType == "27" || record.OpType == "28" || record.OpType == "50" || record.OpType == "51"
        || record.OpType == "30" || record.OpType == "31" || record.OpType == "32" || record.OpType == "33" || record.OpType == "42" || record.OpType == "44" || record.OpType == "57"
        ) {
        html += recordPrefix + actionName;
        const taskLink = getTaskLink(record);
        if (taskLink) {
            html += " <a href=\"" + taskLink + "\" rel=\"nofollow\">" + record.RefName + "</a>"
        } else {
            html += " <span style=\"color: rgba(0,0,0,0.3)\">" + record.RefName + "</span>"
        }
        let time = getTime(record.CreatedUnix, currentTime);
        html += " " + time;
    }
    else if (record.OpType == "35") {
        let datasetLink
        if (record.RepoID > 0) {
            datasetLink = "<a href=\"" + getRepoLink(record) + "/datasets" +  "\" rel=\"nofollow\">" + record.Content.split('|')[1] + "</a>";
        } else {
            datasetLink = `<a target="_blank" href="/datasets/detail/${record.Dataset.Owner.Name}/${record.Dataset.Name}">${record.Dataset.Alias || record.Dataset.Name} </a>`
        }
        actionName = actionName.replace('{dataset}', datasetLink);
        html += recordPrefix + actionName;
        let time = getTime(record.CreatedUnix, currentTime);
        html += " " + time;
    } else if (record.OpType == "62") {
        let datasetLink = `<a target="_blank" href="/datasets/detail/${record.Dataset.Owner.Name}/${record.Dataset.Name}">${record.Dataset.Alias || record.Dataset.Name} </a>`
        actionName = actionName.replace('{dataset}', datasetLink);
        html += recordPrefix + actionName;
        let time = getTime(record.CreatedUnix, currentTime);
        html += " " + time;
    } else if (record.OpType == "63") {
        let datasetLink = `<span style="color: rgba(0,0,0,0.3)">${record.Content.split('|')[1]}</span>`
        actionName = actionName.replace('{dataset}', datasetLink);
        html += recordPrefix + actionName;
        let time = getTime(record.CreatedUnix, currentTime);
        html += " " + time;
    } else if (record.OpType == "67") {
        let datasetLink = `<a target="_blank" href="/models/detail/${record.Aimodel.Owner.Name}/${record.Aimodel.name}">${record.Aimodel.Alias || record.Aimodel.name} </a>`
        actionName = actionName.replace('{aimodel}', datasetLink);
        html += recordPrefix + actionName;
        let time = getTime(record.CreatedUnix, currentTime);
        html += " " + time;
    } else if (record.OpType == "68") {
        let datasetLink = `<span style="color: rgba(0,0,0,0.3)">${record.Content.split('|')[1]}</span>`
        actionName = actionName.replace('{aimodel}', datasetLink);
        html += recordPrefix + actionName;
        let time = getTime(record.CreatedUnix, currentTime);
        html += " " + time;
    } else if (record.OpType == "70") {
        let datasetLink = `<a target="_blank" href="/models/detail/${record.Aimodel.Owner.Name}/${record.Aimodel.name}">${record.Aimodel.Alias || record.Aimodel.name} </a>`
        actionName = actionName.replace('{aimodel}', datasetLink);
        html += recordPrefix + actionName;
        let time = getTime(record.CreatedUnix, currentTime);
        html += " " + time;
    } else {
        return ""; // 跳过不支持的类型
    }
    
    if(record.Repo != null) {
        var time = getTime(record.CreatedUnix, currentTime);
        html += " " + time;
    }
    
    html += "</div></div>";
    html += "</div>";
    return html;
}
function handleNotificationMessage(data) {
    try {
        const parsedData = JSON.parse(data);
        if (parsedData.actions) {
            // 新接口格式
            processActions(parsedData.actions);
        } else {
            // 旧的单个消息格式
            processActions([parsedData]);
        }
    } catch (e) {
        console.error("Failed to parse notification data:", e, data);
    }
}
function getTaskLink(record) {
    let re = ''
     if (record.OpType == 55 || record.OpType == 59) {
        if (record.Cloudbrain) {
            re = '/modelbase/experience/detail/' + record.Cloudbrain.ID;
        } else {
            re = '';
        }
    } else if (record.OpType == 56 || record.OpType == 58) {
        if (record.Cloudbrain) {
            re = '/modelbase/nlp/sft/detail/' + record.Cloudbrain.ID;
        } else {
            re = '';
        }
    } else if (record.OpType == 60) {
        if (record.Cloudbrain) {
            re = '/modelbase/cv/sft/detail/' + record.Cloudbrain.ID;
        } else {
            re = '';
        }
    }
    else if (record.OpType == 61) {
        if (record.Cloudbrain) {
            re = '/modelbase/cv/comfyui/detail/' + record.Cloudbrain.ID;
        } else {
            re = '';
        }
    }
    else if (record.OpType == 66) {
        if (record.Cloudbrain) {
            re = '/modelbase/eval/evaluate/detail/' + record.Cloudbrain.ID;
        } else {
            re = '';
        }   
    } else {
        if (record.Cloudbrain) {
            re = '/cloudbrains/detail/' + record.Cloudbrain.ID;
        } else {
            re = '';
        } 
    }
    
    re = encodeURI(re);
    return re;
}

function getMsg(record){
    var html ="";
    html += "<div class=\"swiper-slide item\">";
    var name = "";
    if(record.ActUser != null){
        name = record.ActUser.Name;
    }else{
        console.log("act user is null.");
    }
    html += "<div class=\"content-c\"><img class=\"ui avatar image\" src=\"/user/avatar/" + name + "/-1\" alt=\"\">"
    html += " <div class=\"middle aligned content nowrap\">"
    html += " <a href=\"/" + encodeURI(name) + "\" title=\"\">" + name + "</a>"
    return html;
}

function getRepotext(record){
    if(record.Repo.Alias){
        return   record.Repo.OwnerName + "/" + record.Repo.Alias;
    }else{
        return   record.Repo.OwnerName + "/" + record.Repo.Name;
    }
}

function getRepoLink(record) {
    return encodeURI(record.Repo.OwnerName + "/" + record.Repo.Name);

}

function getTime(UpdatedUnix,currentTime){
    UpdatedUnix = UpdatedUnix;
    currentTime = currentTime / 1000;
    var timeEscSecond = currentTime - UpdatedUnix;
    if( timeEscSecond < 0){
        timeEscSecond = 1;
    }

    var hours= Math.floor(timeEscSecond / 3600);
    //计算相差分钟数
    var leave2 = Math.floor(timeEscSecond % (3600)); //计算小时数后剩余的秒数
    var minutes= Math.floor(leave2 / 60);//计算相差分钟数

    var leave3=Math.floor(leave2 % 60); //计算分钟数后剩余的秒数
    var seconds= leave3;

    if(hours == 0 && minutes == 0){
        return seconds + getRepoOrOrg(6,isZh,seconds);
    }else{
        if(hours > 0){
            return hours + getRepoOrOrg(4,isZh,hours);
        }else{
            return minutes + getRepoOrOrg(5,isZh,minutes);
        }
    }
}

function getPRLink(record){
    return encodeURI("/" + record.Repo.OwnerName + "/" + record.Repo.Name + "/pulls/" + getIssueId(record));
}
function getPRText(record){
    if(record.Repo.Alias){
        return   record.Repo.OwnerName + "/" + record.Repo.Alias + "#" + getIssueId(record);
    }else{
        return   record.Repo.OwnerName + "/" + record.Repo.Name + "#" + getIssueId(record);
    }

}

function getIssueLink(record){

    return encodeURI("/" + record.Repo.OwnerName + "/" + record.Repo.Name + "/issues/" + getIssueId(record));
}

function getIssueId(record){
    var Id = "1";
    if(!isEmpty(record.Comment) && !isEmpty(record.Comment.Issue)){
        Id = record.Comment.Issue.Index;
    }else{
        if(!isEmpty(record.Content)){
            var content  = record.Content;
            var index = content.indexOf("|");
            if(index != -1){
                Id = content.substring(0,index);
            }
        }
    }
    return Id;
}

function getIssueText(record){
    if(record.Repo.Alias){
        return   record.Repo.OwnerName + "/" + record.Repo.Alias + "#" + getIssueId(record);
    }else{
        return   record.Repo.OwnerName + "/" + record.Repo.Name + "#" + getIssueId(record);
    }

}
function getRepoOrOrg(key,isZhLang,numbers=1){
    if(numbers > 1){
        key+="1";
    }
    if(isZhLang){
        return repoAndOrgZH[key];
    }else{
        return repoAndOrgEN[key];
    }
}
/*
    ActionCreateRepo         ActionType = iota + 1 // 1
	ActionRenameRepo                               // 2
	ActionStarRepo                                 // 3
	ActionWatchRepo                                // 4
	ActionCommitRepo                               // 5
	ActionCreateIssue                              // 6
	ActionCreatePullRequest                        // 7
	ActionTransferRepo                             // 8
	ActionPushTag                                  // 9
	ActionCommentIssue                             // 10
	ActionMergePullRequest                         // 11
	ActionCloseIssue                               // 12
	ActionReopenIssue                              // 13
	ActionClosePullRequest                         // 14
	ActionReopenPullRequest                        // 15
	ActionDeleteTag                                // 16
	ActionDeleteBranch                             // 17
	ActionMirrorSyncPush                           // 18
	ActionMirrorSyncCreate                         // 19
	ActionMirrorSyncDelete                         // 20
	ActionApprovePullRequest                       // 21
	ActionRejectPullRequest                        // 22
	ActionCommentPull                              // 23
*/

var actionNameZH={
    "1":"创建了项目",
    "2":"重命名项目 {oldRepoName} 为",
    "5":"推送了 {branch} 分支的代码到",
    "6":"创建了任务",
    "7":"创建了合并请求",
    "9":"推送了标签 {branch} 到",
    "10":"评论了任务",
    "11":"合并了合并请求",
    "12":"关闭了任务",
    "13":"重新开启了任务",
    "14":"关闭了合并请求",
    "15":"重新开启了合并请求",
    "17":"从 {repoName} 删除分支 {deleteBranchName}",
    "22":"建议变更",
    "23":"评论了合并请求",
    "24":"上传了数据集文件",
    "25":"创建了CPU/GPU类型调试任务",
    "26":"创建了NPU类型调试任务",
    "27":"创建了NPU类型训练任务",
    "28":"创建了推理任务",
    "50":"创建了GPU类型推理任务",
    "51":"创建了ILUVATAR-GPGPU类型推理任务",
    "29":"创建了评测任务",
    "30":"导入了新模型",
    "31":"创建了CPU/GPU类型训练任务",
    "32":"创建了NPU类型训练任务",
    "33":"创建了CPU/GPU类型训练任务",
    "35":"创建的数据集 {dataset} 被设置为推荐数据集",
    "36":"提交了镜像 {image}",
    "37":"提交的镜像 {image} 被设置为推荐镜像",
    "39":"创建了NPU类型调试任务",
    "40":"创建了CPU/GPU类型调试任务",
    "41":"创建了GCU类型调试任务",
    "42":"创建了GCU类型训练任务",
    "43":"创建了MLU类型调试任务",
    "44":"创建了MLU类型训练任务",
    "45":"创建了GPU类型在线推理任务",
    "46":"创建了DCU类型调试任务",
    "47":"创建了CPU类型超算任务",
    "48":"创建了ILUVATAR-GPGPU类型调试任务",
    "49":"创建了METAX-GPGPU类型调试任务",
    "53":"创建了ILUVATAR-GPGPU类型训练任务",
    "54":"创建了GPU类型通用任务",
    "55":"创建了在线体验任务",
    "56":"创建了NLP微调任务",
    "57":"创建了DCU类型训练任务",
    "58":"创建了NLP微调任务",
    "59":"创建了在线体验任务",
    "60":"创建了CV微调任务",
    "61":"创建了Comfy UI任务",
    "62":"创建了数据集 {dataset}",
    "63":"删除了数据集 {dataset}",
    "65":"创建了METAX-GPGPU类型训练任务",
    "66":"创建模型评测任务",
    "67":"创建了模型 {aimodel}",
    "68":"删除了模型 {aimodel}",
    "70": "创建的模型 {aimodel} 被设置为推荐模型",
    "71":"创建了BIREN-GPU类型调试任务",
    "72":"创建了BIREN-GPU类型训练任务",
};

var actionNameEN={
    "1":" created repository",
    "2":" renamed repository from {oldRepoName} to ",
    "5":" pushed to {branch} at",
    "6":" opened issue",
    "7":" created pull request",
    "9":" pushed tag {branch} to ",
    "10":" commented on issue",
    "11":" merged pull request",
    "12":" closed issue",
    "13":" reopened issue",
    "14":" closed pull request",
    "15":" reopened pull request",
    "17":" deleted branch {deleteBranchName} from {repoName}",
    "22":" proposed changes",
    "23":" commented on pull request",
    "24":" upload dataset ",
    "25":" created CPU/GPU type debugging task ",
    "26":" created NPU type debugging task ",
    "27":" created NPU type training task",
    "28":" created inference task",
    "50":" created GPU type inference task",
    "51":" created ILUVATAR-GPGPU type inference task",
    "29":" created profiling task",
    "30":" created new model",
    "31":" created CPU/GPU type training task",
    "32":" created NPU type training task",
    "33":" created CPU/GPU type training task",
    "35":" created dataset {dataset} was set as recommended dataset",
    "36":" committed image {image}",
    "37":" committed image {image} was set as recommended image",
    "39":" created NPU type debugging task ",
    "40":" created CPU/GPU type debugging task ",
    "41":" created GCU type debugging task ",
    "42":" created GCU type training task ",
    "43":" created MLU type debugging task ",
    "44":" created MLU type training task ",
    "45":" created GPU type online inference task ",
    "46":" created DCU type debugging task ",
    "47":" created CPU type super compute task ",
    "48":" created ILUVATAR-GPGPU type debugging task ",
    "49":" created METAX-GPGPU type debugging task ",
    "53":" created ILUVATAR-GPGPU type training task ",
    "54":" created GPU type general task ",
    "55":" created model experience task",
    "56":" created NLP finetune task",
    "57":" created DCU type training task ",
    "58":" created NLP finetune task",
    "59": "created model experience task",
    "60": "created CV finetune task",
    "61": "created Comfy UI task",
    "62": "created {dataset} dataset",
    "63": "deleted {dataset} dataset",
    "65": "created METAX-GPGPU type training task ",
    "66": "created model evaluate task",
    "67": "created {aimodel} model",
    "68": "deleted {aimodel} model",
    "70": "created model {aimodel} was set as recommended model",
    "71":" created BIREN-GPU type debugging task ",
    "72":" created BIREN-GPU type training task ",
};

var repoAndOrgZH={
    "1":"项目",
    "2":"成员",
    "3":"团队",
    "11":"项目",
    "21":"成员",
    "31":"团队",
    "4":"小时前",
    "5":"分钟前",
    "6":"秒前",
    "41":"小时前",
    "51":"分钟前",
    "61":"秒前"
};

var repoAndOrgEN={
    "1":"Repository",
    "2":"Member ",
    "3":"Team",
    "11":"Repositories",
    "21":"Members ",
    "31":"Teams",
    "4":" hour ago",
    "5":" minute ago",
    "6":" second ago",
    "41":" hours ago",
    "51":" minutes ago",
    "61":" seconds ago"
};


function getAction(opType,isZh){
    if(isZh){
        return actionNameZH[opType]
    }else{
        return actionNameEN[opType]
    }
}

async function queryRecommendData(){
    const recommendData = await $.ajax({
        type:"GET",
        url:"/api/v1/home",
        headers: {
           authorization:token,
         },
        dataType:"json",

    });
    return recommendData
}


function disPlayNotice(data) {
    const homeNotice = document.getElementById("noticemessage");
    if (homeNotice && data) {
        try {
            // 销毁已有实例
            if (homeNotice.swiperInstance) {
                homeNotice.swiperInstance.destroy(true, true);
                delete homeNotice.swiperInstance;
            }
            let html = ''
            data.forEach(element => {
                html += `<div class="swiper-slide activate-cards nowrap" style="width: auto;">
                    <div class="card">
                        <span>${element.Date}</span>
                        <a class="card-wrap" href="${element.Link}">
                            ${isZh ? element.Title : element.TitleEn}
                        </a>
                    </div>
                </div>`
            });
            homeNotice.innerHTML = html //+ html;

            homeNotice.swiperInstance = new Swiper(".notice-list", {
                // direction: 'vertical',
                loop: false,
                // autoplay: {
                //     delay: 3000, // 2秒切换一次
                //     disableOnInteraction: false,
                //     pauseOnMouseEnter: true
                // },
                slidesPerView: 'auto',
                spaceBetween: 50,
                // centeredSlides: true,
            });
            // 动态检测是否需要启用 loop
            function checkLoopNeed() {
                const container = document.querySelector('.notice-list');
                const slides = container.querySelectorAll('.swiper-slide');

                // 计算所有 slides 的总宽度
                let totalSlidesWidth = 0;
                slides.forEach(slide => {
                    totalSlidesWidth += slide.offsetWidth + 50; // 包含 spaceBetween
                });
                // 比较总宽度和容器宽度
                const containerWidth = container.offsetWidth;
                const needLoop = totalSlidesWidth > containerWidth;
                // 动态更新 Swiper 配置
                if (needLoop !== homeNotice.swiperInstance.params.loop) {

                    homeNotice.swiperInstance.destroy(); // 销毁旧实例
                    slides.forEach(slide => {
                        slide.style.width = 'auto';
                    });
                    homeNotice.swiperInstance = new Swiper(".notice-list", {
                        loop: needLoop, // 动态设置 loop
                        observer: true,       // 监听 DOM 变化
                        observeParents: true, // 监听父级 DOM 变化
                        autoplay: needLoop ? { // 只有需要 loop 时才启用 autoplay
                            delay: 3000,
                            disableOnInteraction: false,
                            pauseOnMouseEnter: true
                        } : false,
                        slidesPerView: 'auto',
                        spaceBetween: 50,
                    });
                }
            }

            // 初始化检测 + 监听窗口变化
            checkLoopNeed();
            // window.addEventListener('resize', checkLoopNeed);
            // swiperEvent.updateSlides();

        } catch (e) {
            console.error("Failed to parse notice data:", e);
        }
    }
}

function disPlayModelExperice(data) {
    const homeModelExperience = document.getElementsByClassName("model-experience-content");
    if (homeModelExperience && data) {
        try {
            let html = ''
            data.forEach(element => {
                let btnHtml = ''
                if (element.experienceGpu) {
                    btnHtml += `<a class="btn-i nowrap gpu-btn" href="${element.experienceGpu.url}">
                        ${isZh ? element.experienceGpu.title : element.experienceGpu.titleEn}

                        <i class="ri-play-circle-line gpu-icon"></i>
                    </a>`
                }
                if (element.experienceNpu) {
                    btnHtml += `<a class="btn-i nowrap npu-btn" href="${element.experienceNpu.url}">
                        ${isZh ? element.experienceNpu.title : element.experienceNpu.titleEn}
                        <i class="ri-play-circle-line npu-icon"></i>
                    </a>`
                }
                html += `<div class="model-experience-item">
                    <div class="item-head">
                        <img src="${element.img}" alt="">
                        <a href="${element?.url}" title="${element.name}" class="nowrap">${element.name}</a>
                    </div>
                    <div class="item-btn-group">
                        ${btnHtml}
                    </div>
                </div>`
            });
            homeModelExperience[0].innerHTML = html;
            // homeModelExperience[1].innerHTML = html;
        } catch (e) {
            console.error("Failed to parse notice data:", e);
        }
    }
}
function disPlayRepoList(data) {
    const repoCards = document.getElementById('repo-cards')
    if (repoCards && data) {
        try {
            let html = ''
            data.forEach(element => {
                html += `<div class="card-wrap">
                <div class="card-box">
                    <div>
                        <div class="card-header-repo">
                            ${element.Avatar ? `<img class="header-img" src="${element.Avatar}" />` : `<img class="header-img" avatar="${element.Name}">`}
                            <a href="${element.RepoLink}" class="header-text nowrap" title="${element.Name}">${element.Name}</a>
                        </div>
                        <div class="crad-desc-repo">${element.Description}</div>
                    </div>
                    <div class="card-footer-repo">
                        <div class="footer-item"><i class="ri-eye-line"></i><span>${element.NumWatches}</span></div>
                        <div class="footer-item"><i class="ri-star-line"></i><span>${element.NumStars}</span></div>
                        <div class="footer-item"><i class="ri-git-branch-line"></i><span>${element.NumForks}</span></div>
                    </div>
                </div>
                </div>`
            });
            repoCards.innerHTML = html;
            LetterAvatar && LetterAvatar.transform()
        } catch (e) {
            console.error("Failed to parse notice data:", e);
        }
    }
}
const MODEL_ENGINES = [{ k: 0, v: 'PyTorch' }, { k: 1, v: 'TensorFlow' }, { k: 2, v: 'MindSpore' }, { k: 4, v: 'PaddlePaddle' }, { k: 5, v: 'OneFlow' }, { k: 6, v: 'MXNet' }, { k: 3, v: 'Other' }];
const engineMap = MODEL_ENGINES.reduce((map, item) => {
    map.set(item.k, item.v);
    return map;
  }, new Map());
function disPlayModelList(data) {
    const modelCards = document.getElementById('model-cards')
    if (modelCards && data) {
        try {
            let html = ''
            data.forEach(element => {
                let labelHtml = ''
                if (element.Engine != null) {
                    let label = engineMap.get(element.Engine)
                    labelHtml += `<span class="label">${label}</span>`
                }
                if (element.Label) {
                    element.Label.split(' ').forEach(label => {
                        labelHtml += `<span class="label">${label}</span>`
                    })
                }
                html += `<div class="card-wrap">
                <div class="card-box">
                    <div>
                        <div class="card-header-model">
                            <a href="/models/detail/${element.OwnerName}/${element.Name}" class="header-text nowrap" title="${element.Alias}">${element.Alias}</a>
                        </div>
                        <div class="card-label-model">${labelHtml}</div>
                    </div>
                    <div class="card-footer-model">
                        <img src="${element.Avatar}" alt="">
                        <span>${element.OwnerName}</span>
                    </div>
                </div>
                </div>`
            });
            modelCards.innerHTML = html;
        } catch (e) {
            console.error("Failed to parse notice data:", e);
        }
    }
}
function displayTemplate(data) {
    const templatelCards = document.getElementById('template-cards')
    if (templatelCards && data) {
        try {
            let html = ''
            let sliceData = data.slice(0,4)
            sliceData.forEach(element => {
                let modelstr = '--' 
                let datastr = '--'
                let repostr = element.RepoName || '--'
                let imagestr = element.ImageName || '--'
                if (element.ModelLists.length) {
                    modelstr = element.ModelLists.map(item=>item.ModelName).join(';')
                }
                if (element.DatasetLists.length) {
                    datastr = element.DatasetLists.map(item=>item.DatasetName).join(';')
                }
                html += `<div class="template-wrap">
                <a class="template-box" href="/cloudbrains/create?tmpl=${element.ID}">
                  <div class="title-c nowrap" title="${element.Name}">${element.Name}</div>
                  <div class="labels-c">
                    <div class="label nowrap">${i18n['taskType'][element.JobType]}</div>
                    <div class="label nowrap">${i18n['TaskTypeTitle'][element.ComputeSource]}</div>
                  </div>
                  <div class="infos-c">
                    <p class="nowrap" title="${datastr}">数据集：${datastr}</p>
                    <p class="nowrap" title="${modelstr}">模型：${modelstr}</p>
                    <p class="nowrap" title="${repostr}">项目：${repostr}</p>
                    <p class="nowrap" title="${imagestr}">镜像：${imagestr}</p>
                  </div>
                  <div class="card-footer-template">
                    <img src="${element.Owner.RelAvatarLink}" alt="">
                    <span>${element.Owner.Name}</span>
                  </div>
                </a>
                </div>`
            });
            templatelCards.innerHTML = html;
        } catch (e) {
            console.error("Failed to parse notice data:", e);
        }
    }
}
function displayActivateInfo(data) {
    const activityDiv = document.getElementById("recommendactivity");
    if (activityDiv && data) {
        try {
            let html = ''
            data.forEach(element => {
                html += `<div class="swiper-slide activate-cards">
                    <a class="card-wrap" href="${element.Link}">
                        <img width="318" height="150" class="card-img" data-src="${element.Background}" src="" alt="">
                        <div class="card-title">${isZh ? element.Title : element.TitleEn}</div>
                    </a>
                </div>`
            });
            activityDiv.innerHTML = html;
            // 初始化懒加载监听
            const initLazyLoad = () => {
                const lazyImages = document.querySelectorAll('.card-img[data-src]');

                const observer = new IntersectionObserver((entries) => {
                entries.forEach(entry => {
                    if (entry.isIntersecting) {
                    const img = entry.target;
                    img.src = img.dataset.src;
                    img.removeAttribute('data-src');
                    img.onload = () => img.classList.add('loaded');
                    observer.unobserve(img);
                    }
                });
                }, {
                rootMargin: '200px',  // 提前 200px 加载
                threshold: 0.1
                });

                lazyImages.forEach(img => observer.observe(img));
            };
            new Swiper(".activate-list", {
                slidesPerView: 1,
                spaceBetween: 30,
                autoplay: {
                    delay: 2500,
                    disableOnInteraction: false,
                },
                breakpoints: {
                    768: {
                        slidesPerView: Math.min(2, data.length),
                    },
                    1024: {
                        slidesPerView: Math.min(3, data.length),
                    },
                    1600: {
                        slidesPerView: Math.min(4, data.length),
                    },
                },
            });
            initLazyLoad();
            // swiperEvent.updateSlides();
        } catch (e) {
            console.error("Failed to parse notice data:", e);
        }
    }

}

function displayDataset(data) {
    const datasetCards = document.getElementById("data-cards");
    const svgStrMap = {
        '0': '<svg xmlns="http://www.w3.org/2000/svg" class="svg svg-icon-path-icon fill" viewBox="0 0 32 32" width="32" height="32"><defs data-reactroot=""><linearGradient id="ila93em9ydx6bi61,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs" x1="0" x2="100%" y1="0" y2="0" gradientUnits="userSpaceOnUse"><stop stop-color="#82d1f6" stop-opacity="1" offset="0"></stop><stop stop-color="#29b6f4" stop-opacity="1" offset="0.5"></stop><stop stop-color="#0089cd" stop-opacity="1" offset="0.99"></stop></linearGradient></defs><g><path d="M16 4c7.189 0 13.171 5.173 14.425 12-1.253 6.827-7.236 12-14.425 12s-13.171-5.173-14.425-12c1.253-6.827 7.236-12 14.425-12zM16 25.333c5.682-0.001 10.442-3.949 11.687-9.252l0.016-0.081c-1.265-5.379-6.024-9.322-11.703-9.322s-10.437 3.943-11.686 9.241l-0.016 0.081c1.261 5.384 6.021 9.332 11.703 9.333h0zM16 22c-3.314 0-6-2.686-6-6s2.686-6 6-6v0c3.314 0 6 2.686 6 6s-2.686 6-6 6v0zM16 19.333c1.841 0 3.333-1.492 3.333-3.333s-1.492-3.333-3.333-3.333v0c-1.841 0-3.333 1.492-3.333 3.333s1.492 3.333 3.333 3.333v0z"></path></g></svg>',
        '1': '<svg xmlns="http://www.w3.org/2000/svg" class="svg svg-icon-path-icon fill" viewBox="0 0 32 32" width="32" height="32"><defs data-reactroot=""><linearGradient id="ila93gapd7aoa4d1,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs" x1="0" x2="100%" y1="0" y2="0" gradientUnits="userSpaceOnUse"><stop stop-color="#82d1f6" stop-opacity="1" offset="0"></stop><stop stop-color="#29b6f4" stop-opacity="1" offset="0.5"></stop><stop stop-color="#0089cd" stop-opacity="1" offset="0.99"></stop></linearGradient></defs><g><path d="M2.667 16c0-7.364 5.969-13.333 13.333-13.333s13.333 5.969 13.333 13.333-5.969 13.333-13.333 13.333h-13.333l3.905-3.905c-2.413-2.407-3.905-5.735-3.905-9.411 0-0.006 0-0.012 0-0.018v0.001zM9.104 26.667h6.896c5.891 0 10.667-4.776 10.667-10.667s-4.776-10.667-10.667-10.667c-5.891 0-10.667 4.776-10.667 10.667v0c0 2.869 1.135 5.553 3.124 7.543l1.885 1.885-1.239 1.239zM14.667 8h2.667v16h-2.667v-16zM9.333 12h2.667v8h-2.667v-8zM20 12h2.667v8h-2.667v-8z"></path></g></svg>',
        '2': '<svg xmlns="http://www.w3.org/2000/svg" class="svg svg-icon-path-icon fill" viewBox="0 0 48 48" width="32" height="32"><defs data-reactroot=""><linearGradient id="ila93j5tkotzm421,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs" x1="0" x2="100%" y1="0" y2="0" gradientUnits="userSpaceOnUse"><stop stop-color="#82d1f6" stop-opacity="1" offset="0"></stop><stop stop-color="#29b6f4" stop-opacity="1" offset="0.5"></stop><stop stop-color="#0089cd" stop-opacity="1" offset="0.99"></stop></linearGradient></defs><g><rect width="48" height="48" fill-opacity="0.01" fill="url(#ila93j5tkotzm421,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" storke="none"></rect><path stroke="url(#ila93j5tkotzm421,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" d="M4 24C4 24 10 15 14 15C18 15 22 17 24 17C26 17 30 15 34 15C38 15 44 24 44 24C44 24 34 34 24 34C14 34 4 24 4 24Z" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path><path stroke="url(#ila93j5tkotzm421,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" d="M4 24H44" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path></g></svg>',
        '3': '<svg xmlns="http://www.w3.org/2000/svg" class="svg svg-icon-path-icon fill" viewBox="0 0 48 48" width="32" height="32"><defs data-reactroot=""><linearGradient id="ila93k86om2ayqg1,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs" x1="0" x2="100%" y1="0" y2="0" gradientUnits="userSpaceOnUse"><stop stop-color="#82d1f6" stop-opacity="1" offset="0"></stop><stop stop-color="#29b6f4" stop-opacity="1" offset="0.5"></stop><stop stop-color="#0089cd" stop-opacity="1" offset="0.99"></stop></linearGradient></defs><g><path d="M6 25C6 15.0589 14.0589 7 24 7C30.8669 7 36.8357 10.8453 39.8706 16.5" stroke="url(#ila93k86om2ayqg1,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4"></path><path d="M41 25H43L42 24L41 25Z" stroke="url(#ila93k86om2ayqg1,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4"></path><path d="M31 34L28.375 27L25 18H23L19.625 27L17 34" stroke="url(#ila93k86om2ayqg1,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4"></path><path d="M28.375 27H19.625" stroke="url(#ila93k86om2ayqg1,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4"></path><path d="M7 25H5L6 26L7 25Z" stroke="url(#ila93k86om2ayqg1,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4"></path><path d="M42 25C42 34.9411 33.9411 43 24 43C17.1331 43 11.1643 39.1547 8.12939 33.5" stroke="url(#ila93k86om2ayqg1,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4"></path></g></svg>',
        '4': '<svg xmlns="http://www.w3.org/2000/svg" class="svg svg-icon-path-icon fill" viewBox="0 0 48 48" width="32" height="32"><defs data-reactroot=""><linearGradient id="ila93lcoqm7710i1,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs" x1="0" x2="100%" y1="0" y2="0" gradientUnits="userSpaceOnUse"><stop stop-color="#82d1f6" stop-opacity="1" offset="0"></stop><stop stop-color="#29b6f4" stop-opacity="1" offset="0.5"></stop><stop stop-color="#0089cd" stop-opacity="1" offset="0.99"></stop></linearGradient></defs><g><path d="M33 4.99976H41C42.1046 4.99976 43 5.89519 43 6.99976V14.9998M43 32.9998V40.9998C43 42.1043 42.1046 42.9998 41 42.9998H33M15 42.9998H7C5.89543 42.9998 5 42.1043 5 40.9998V32.9998M5 14.9998V6.99976C5 5.89519 5.89543 4.99976 7 4.99976H15" stroke="url(#ila93lcoqm7710i1,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path><path d="M24 38C30.6274 38 36 31.732 36 24C36 16.268 30.6274 10 24 10C17.3726 10 12 16.268 12 24C12 31.732 17.3726 38 24 38Z" stroke="url(#ila93lcoqm7710i1,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4"></path><path d="M6 24H42" stroke="url(#ila93lcoqm7710i1,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4" stroke-linecap="round"></path><path d="M20.0693 30.1057C21.3372 31.0429 22.6473 31.5115 23.9996 31.5115C25.3519 31.5115 26.698 31.0429 28.0378 30.1057" stroke="url(#ila93lcoqm7710i1,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4" stroke-linecap="round"></path></g></svg>',
        '5': '<svg xmlns="http://www.w3.org/2000/svg" class="svg svg svg-icon-path-icon fill" viewBox="0 0 32 32" width="32" height="32"><defs data-reactroot=""><linearGradient id="ila93mrrwzocn0r1,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs" x1="0" x2="100%" y1="0" y2="0" gradientUnits="userSpaceOnUse"><stop stop-color="#82d1f6" stop-opacity="1" offset="0"></stop><stop stop-color="#29b6f4" stop-opacity="1" offset="0.5"></stop><stop stop-color="#0089cd" stop-opacity="1" offset="0.99"></stop></linearGradient></defs><g><path d="M16 26.667c5.891 0 10.667-4.776 10.667-10.667s-4.776-10.667-10.667-10.667v0c-5.891 0-10.667 4.776-10.667 10.667s4.776 10.667 10.667 10.667v0zM16 29.333c-7.364 0-13.333-5.969-13.333-13.333s5.969-13.333 13.333-13.333 13.333 5.969 13.333 13.333-5.969 13.333-13.333 13.333zM16 21.333c2.946 0 5.333-2.388 5.333-5.333s-2.388-5.333-5.333-5.333v0c-2.946 0-5.333 2.388-5.333 5.333s2.388 5.333 5.333 5.333v0zM16 24c-4.418 0-8-3.582-8-8s3.582-8 8-8v0c4.418 0 8 3.582 8 8s-3.582 8-8 8v0zM16 18.667c-1.473 0-2.667-1.194-2.667-2.667s1.194-2.667 2.667-2.667v0c1.473 0 2.667 1.194 2.667 2.667s-1.194 2.667-2.667 2.667v0z"></path></g></svg>',
        '6': '<svg xmlns="http://www.w3.org/2000/svg" class="svg svg-icon-path-icon fill" viewBox="0 0 48 48" width="32" height="32"><defs data-reactroot=""><linearGradient id="ila93oiqy4t1x861,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs" x1="0" x2="100%" y1="0" y2="0" gradientUnits="userSpaceOnUse"><stop stop-color="#82d1f6" stop-opacity="1" offset="0"></stop><stop stop-color="#29b6f4" stop-opacity="1" offset="0.5"></stop><stop stop-color="#0089cd" stop-opacity="1" offset="0.99"></stop></linearGradient></defs><g><rect width="48" height="48" fill="white" fill-opacity="0.01"></rect><path d="M14.5397 20.0186C12.8522 17.9434 11.2675 17.4979 9.78564 18.6821C7.5629 20.4583 6.92453 26.6496 8.71324 32.1086C10.502 37.5676 13.9801 45.0017 21.0016 45.0017C28.0231 45.0017 29.684 37.5222 32.5485 33.0001C35.413 28.478 36.9285 24.1152 34.1208 18.6821" stroke="url(#ila93oiqy4t1x861,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path><path d="M11 18.0368C9.29707 15.4428 7.96374 13.4306 6.99996 12.0002C5.5543 9.85464 9.25107 7.08164 11 8.96807C12.1659 10.2257 13.7148 12.078 15.6466 14.5249" stroke="url(#ila93oiqy4t1x861,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4" stroke-linecap="round"></path><path d="M15.0236 25.6396C14.5391 19.5759 14.9333 15.6276 16.2062 13.7947C18.1155 11.0455 21.6631 10.0031 25.0035 10.0031C26.9924 10.0031 28.8087 10.8502 30.4525 12.5444" stroke="url(#ila93oiqy4t1x861,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path><path fill-rule="evenodd" clip-rule="evenodd" d="M41.0003 12.6128C41.5858 14.6492 40.6294 16.5097 37.6844 16.931C34.7393 17.3523 32.5313 18.8332 30.9388 20.079C29.3463 21.3248 26.4983 25.1046 25.9361 27.0023C25.3738 28.9 22.1602 27.1547 21.2971 26.3971C20.434 25.6394 19.5855 23.9806 21.2971 22.2457C23.0086 20.5108 22.6383 20.1646 22.6383 18.4052C22.6383 16.6459 32.0003 10.8263 37.2729 10.2941C38.4449 10.2257 40.4147 10.5763 41.0003 12.6128Z" stroke="url(#ila93oiqy4t1x861,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4"></path><path d="M23.0078 4.00014V9.26283" stroke="url(#ila93oiqy4t1x861,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4" stroke-linecap="round"></path><path d="M20.3066 10.7178C17.2888 6.92534 14.8555 4.80868 13.0068 4.36781" stroke="url(#ila93oiqy4t1x861,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4" stroke-linecap="round"></path><path d="M17.0039 7.02894L17.9944 2.96156" stroke="url(#ila93oiqy4t1x861,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4" stroke-linecap="round"></path><path d="M35.6128 10.7174C35.2905 11.8219 35.2905 12.8575 35.6128 13.8241C35.935 14.7908 36.6255 15.8264 37.6842 16.9308" stroke="url(#ila93oiqy4t1x861,1,rs,1,f000f000,f0rsf000,f000,00e6msqtrs,dw4hjuqlrs,ri00exmcrs)" fill="none" stroke-width="4" stroke-linecap="round"></path></g></svg>',
        '7': '<svg xmlns="http://www.w3.org/2000/svg" class="svg svg-icon-path-icon fill" viewBox="0 0 24 24" width="32" height="32"><defs><linearGradient id="ila93k86UiyG0eljrQnk0c1" x1="0" x2="100%" y1="0" y2="0" gradientUnits="userSpaceOnUse"><stop stop-color="#82d1f6" stop-opacity="1" offset="0"></stop><stop stop-color="#29b6f4" stop-opacity="1" offset="0.5"></stop><stop stop-color="#0089cd" stop-opacity="1" offset="0.99"></stop></linearGradient></defs><g><path d="M17 5c.83 0 1.5-.67 1.5-1.5 0-1-1.5-2.7-1.5-2.7s-1.5 1.7-1.5 2.7c0 .83.67 1.5 1.5 1.5zm-5 0c.83 0 1.5-.67 1.5-1.5 0-1-1.5-2.7-1.5-2.7s-1.5 1.7-1.5 2.7c0 .83.67 1.5 1.5 1.5zM7 5c.83 0 1.5-.67 1.5-1.5C8.5 2.5 7 .8 7 .8S5.5 2.5 5.5 3.5C5.5 4.33 6.17 5 7 5zm11.92 3.01C18.72 7.42 18.16 7 17.5 7h-11c-.66 0-1.21.42-1.42 1.01L3 14v8c0 .55.45 1 1 1h1c.55 0 1-.45 1-1v-1h12v1c0 .55.45 1 1 1h1c.55 0 1-.45 1-1v-8l-2.08-5.99zM6.5 18c-.83 0-1.5-.67-1.5-1.5S5.67 15 6.5 15s1.5.67 1.5 1.5S7.33 18 6.5 18zm11 0c-.83 0-1.5-.67-1.5-1.5s.67-1.5 1.5-1.5 1.5.67 1.5 1.5-.67 1.5-1.5 1.5zM5 13l1.5-4.5h11L19 13H5z"></path></g></svg>'
    }
    if (datasetCards && data) {
        try {
            let html = ''
            data.forEach((element,i) => {
                html += `<div class="card-wrap">
                <a href="/explore/datasets?sort=default&q=&tab=&category=&task=${element.task}&license=">
                    ${svgStrMap[i % 8]}
                    <div class="nowrap" style="width: 100%;text-align: center;" title="{i18n[element.task] || element.task}">${i18n[element.task] || element.task}</div>
                    <span>${i18n['about']} ${element.total} ${i18n['count']}</span>
                </a>
                </div>`
            });
            datasetCards.innerHTML = html;
        } catch (e) {
            console.error("Failed to parse notice data:", e);
        }
    }

}

function initHomeTopBanner() {
    var homeSlideTimer = null;
    var homeSlideDuration = 4000;

    function getBannerData() {
        $.ajax({
            type: "GET",
            url: "/dashboard/invitation",
            headers: { authorization: token, },
            dataType: "json",
            data: { filename: 'home/banner', },
            success: function (data) {
                try {
                    var banners = JSON.parse(data);
                    var count = 0, bannerList = [];
                    for (var i = 0; i < banners.length; i++) {
                        (function (banner, index) {
                            $.ajax({
                                type: "GET",
                                url: "/dashboard/invitation",
                                headers: { authorization: token, },
                                data: { filename: 'home/banners/' + banner },
                                success: function (data) {
                                    count++;
                                    bannerList[index] = {
                                        banner: banner,
                                        data: data,
                                    };
                                    if (count == banners.length) {
                                        renderBanners(bannerList);
                                    }
                                },
                                error: function (err) {
                                    count++;
                                    if (count == data.length) {
                                        startSlide();
                                    }
                                    console.log(err);
                                }
                            });
                        })(banners[i], i);
                    }
                    if (!banners.length) {
                        startSlide();
                    }
                } catch (err) {
                    console.log(err);
                    startSlide();
                }
            },
            error: function (err) {
                console.log(err);
                startSlide();
            }
        });
    }

    function renderBanners(bannerList) {
        var hmPageC = $('._hm-bg-container ._hm-pg-c');
        var hmPageSlidePaginationC = $('._hm-slide-pagination-c');
        for (var i = 0, iLen = bannerList.length; i < iLen; i++) {
            var banner = bannerList[i];
            if (banner.data) {
                hmPageC.append($(banner.data));
                hmPageSlidePaginationC.append('<div class="_hm-slide-pagination-item"></div>');
            }
        }
        startSlide();
        // 初始化Swiper
        setTimeout(function() {
            initSwiper();
        }, 100);
    }

    function homeSlide(direction, index) {
        var slidePages = $('._hm-pg-c ._hm-pg');
        var currentPage = slidePages.filter('._hm-pg-show');
        var slidePagination = $('._hm-slide-pagination-c ._hm-slide-pagination-item');
        var currentIndex = currentPage.index();
        var next = 0;
        if (direction) {
            next = direction == 'left' ? currentIndex - 1 : currentIndex + 1;
        } else {
            next = index || 0;
        }
        if (next < 0) next = slidePages.length - 1;
        if (next == slidePages.length) next = 0;

        // 记录当前banner索引
        var wasFirstBanner = currentIndex === 0;
        var willBeFirstBanner = next === 0;

        slidePages.removeClass('_hm-pg-show');
        slidePages.eq(next).addClass('_hm-pg-show');
        slidePagination.removeClass('_hm-slide-pagination-item-active');
        slidePagination.eq(next).addClass('_hm-slide-pagination-item-active');
        // 控制Swiper自动轮播
        if (willBeFirstBanner && !wasFirstBanner) {
            // 切换到第一个banner，恢复Swiper自动轮播
            isFirstBannerActive = true;
            setTimeout(function() {
                toggleSwiperAutoplay(true);
            }, 100);
        } else if (!willBeFirstBanner && wasFirstBanner) {
            // 从第一个banner切换走，暂停Swiper自动轮播
            isFirstBannerActive = false;
            toggleSwiperAutoplay(false);
        }
    }

    function startSlide() {
        $('._hm-slide-pagination-c').show();
        homeSlideTimer && clearTimeout(homeSlideTimer);
        homeSlideTimer = setTimeout(function () {
            homeSlide('right');
            startSlide();
        }, homeSlideDuration);
    }

    function stopSlide() {
        homeSlideTimer && clearTimeout(homeSlideTimer);
    }

    function eventInit() {
        $('._hm-slide-btn').on('click', function () {
            if ($(this).hasClass('_hm-slide-btn-left')) {
                homeSlide('left');
            } else {
                homeSlide('right');
            }
            startSlide();
        });
        $('._hm-pg #homenews').on('mouseenter', function () {
            stopSlide();
        }).on('mouseleave', function () {
            startSlide();
        });
        $('._hm-slide-pagination-c').on('click', '._hm-slide-pagination-item', function () {
            var self = $(this);
            if (self.hasClass('_hm-slide-pagination-item-active')) return;
            homeSlide('', self.index());
            startSlide();
        });
    }

    getBannerData();
    eventInit();
}

initHomeTopBanner();
const renderRecommend = (res) => {
    if (res.Code === 0) {
        Promise.all([
            disPlayNotice(res.Data.Notice),
            disPlayModelExperice(res.Data.ModelExperience),
            disPlayRepoList(res.Data.Repo),
            displayDataset(res.Data.Dataset),
            disPlayModelList(res.Data.Model),
            displayTemplate(res.Data.AITaskTemplate),
            displayActivateInfo(res.Data.Activity),
        ])

    }

}


queryRecommendData().then((res) => { renderRecommend(res) });

const API_CONFIG = {
    headers: { authorization: token },
    dataType: 'json'
};
new Vue({
    el: "#app",
    delimiters: ["${", "}"],
    data: {
        itemList: [],
        jobTasks: {},
        rotation3D: null, // 增加3D实例引用,
        error: null
    },
    async mounted() {

        await this.fetchData();
        this.$nextTick(() => {
            this.initRotation3D();
            // this.updateTextContent();
        });
    },
    methods: {
        async fetchData() {
            try {
                const [jobs, centers] = await Promise.all([
                    this.queryJobs(),
                    this.queryCenters()
                ]);

                this.processData(jobs, centers);
            } catch (error) {
                console.error('Failed to load data:', error);
                this.error = 'Failed to load data, please try again later.';
            } finally {

            }
        },

        async queryJobs() {
            const response = await $.ajax({
                ...API_CONFIG,
                type: "GET",
                url: "/api/v1/cloudbrain/get_newest_job"
            });
            return response.reduce((acc, job) => {
                acc[job.ai_center_id] = job.job_name;
                return acc;
            }, {});
        },

        async queryCenters() {
            const response = await $.ajax({
                ...API_CONFIG,
                type: "GET",
                url: "/api/v1/cloudbrain/get_center_info"
            });
            return response;
        },

        processData(jobs, centers) {
            this.jobTasks = jobs;
            this.itemList = centers.map(center => ({
                name: center.name,
                type: center.id in jobs ? 'blue' : 'green',
                icon: "",
                content: isZh ? center.content : center.content_en
            }));
        },

        initRotation3D() {
            const container = document.querySelector('#rotation3D');

            if (!container) return;
            if (this.rotation3D) {
                this.rotation3D.destroy();
                this.rotation3D = null;
            }
            requestIdleCallback(() => {
                this.rotation3D = new Rotation3D({
                    id: '#rotation3D',
                    farScale: 0.6,
                    xRadius: 0,
                    yRadius: 130,
                    // items: '.itemList > div' // 根据实际结构调整选择器
                });
            })

        }
    }
});
