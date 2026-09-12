var token;
if(isEmpty(token)){
	var meta = $("meta[name=_uid]");
	if(!isEmpty(meta)){
		token = meta.attr("content");
		console.log("token is uid:" + token);
	}
}
var authorMeta = $("meta[name=author]");
var userName;
if(!isEmpty(authorMeta)){
	userName = authorMeta.attr("content");
	console.log("user name=" + userName);
}
var userType;
if(isEmpty(userType)){
	userType =1;
}
var ip = getIp();

var pageSize = 10;

var tableData;
var tablePageData;

var modelListData;
var dataSetTaskData;

var userInfoData;
var labelPropertyData;
var repoId = $('#repoId').val();


console.log("repoId=" + repoId);

function setDataSetTask(){
	
	//dataset_task_list();
	//display_createdatasetlabel(0);
    //getUser();
    //dislpayUser();
	getLabelPropertyTask();
	displayLabelPropertyTask();
	dataset_sele_Change("");
	$(".ui.dataset.modal").modal("show");
}

function getLabelPropertyTask(){
	 $.ajax({
       type:"GET",
       url:ip + "/api/label-property-task-all/",
       headers: {
          authorization:token,
        },
       dataType:"json",
       async:false,
       success:function(json){
        labelPropertyData = json;
        console.log(json);
       },
	   error:function(response) {
		  redirect(response);
       }
   });
}

function displayLabelPropertyTask(){
  var html="<option value=\"\" selected=\"\">请选择</option>";
  for (var i=0;i<labelPropertyData.length;i++){
        var row = "<option value=\""+labelPropertyData[i].id+
        "\">"+labelPropertyData[i].task_name +
        "</option>";
      
      html=html+row;
  }
  console.log(html);
  document.getElementById('labelpropertytask_dataset').innerHTML=html; 
  document.getElementById('labelpropertytask_auto').innerHTML=html;
}

function dataset_task_list(){
  $.ajax({
     type:"GET",
     url:ip + "/gitea-dataset/",
     headers: {
        authorization:token,
      },
     dataType:"json",
	   data:{
       'repoId':repoId,
       'dateset_type':'[1]'
      },
     async:false,
     success:function(json){
      dataSetTaskData = json;
      console.log(json);
      // return json.token;
    },
	error:function(response) {
		  redirect(response);
    }
 });
}

function countLabel(){
  $.ajax({
     type:"GET",
     url:ip + "/gitea/label-count/",
     headers: {
        authorization:token,
      },
     dataType:"json",
     async:false,
     success:function(json){
         alert("请等待几分钟，服务端正在加紧统计。");
     },
	 error:function(response) {
		redirect(response);
     }
 });
	
}




function setPredictTask(){
  //dataset_task_list();
	//display_createdatasetlabel(1);


	get_model_list();

	displayModelTask();
  dataset_auto_sele_Change("");
  $(".ui.predict.modal").modal("show");
}

function get_model_list(){
  $.ajax({
     type:"GET",
     url:ip + "/api/queryAlgModelForAutoLabel/",
     headers: {
        authorization:token,
      },
     dataType:"json",
     async:false,
     success:function(json){
        modelListData = json;
        console.log(json);
      // return json.token;
     },
	 error:function(response) {
		redirect(response);
     }
 });
}

function displayModelTask(){
  var html="";
  for (var i=0;i<modelListData.length;i++){
    if (i==0){
        var row = "<option value=\""+modelListData[i].id+
              "\" selected=\"\">"+modelListData[i].model_name +
              "</option>";
    }else{
        var row = "<option value=\""+modelListData[i].id+
        "\">"+modelListData[i].model_name+
        "</option>";
      }
    html=html+row;
  }
  console.log(html);

  document.getElementById('model_list').innerHTML=html; 
}


function sele_export_Change(sele){

  var isNeedPicture = $('#isNeedPicture option:selected').val();

  if(isNeedPicture == 3){	
     document.getElementById("maxscore_div").style.display="block";
     document.getElementById("minscore_div").style.display="block";
	 $('#maxscore').val("1.0");
	 $('#minscore').val("0.6");
  }else{
	 document.getElementById("maxscore_div").style.display="none";
     document.getElementById("minscore_div").style.display="none"; 
	 $('#maxscore').val("");
	 $('#minscore').val("");
  }
 
}


function dataset_sele_Change(sele){
  var dataset_listName = $('#dataset_list option:selected').text();
  console.log("select dataset_list =" + dataset_listName);
  $("#datasetlabeltaskname").val(dataset_listName+"-人工标注");
}

function dataset_auto_sele_Change(sele){
  var dataset_listName = $('#dataset_list_auto option:selected').text();
  console.log("select dataset_list_auto =" + dataset_listName);
  $("#autolabeltaskname").val(dataset_listName+"-自动标注");
}


var createsucced;

function submit_datasettask(){
  console.log($('#datasetlabeltaskname').val());
  var task_name = $('#datasetlabeltaskname').val();
  if (isEmpty(task_name) || task_name.length > 32){
       alert("人工标注任务名称不能为空或者不能超过32个字符。");
       return;
   }
  var assign_user_id = $('#assign_user option:selected').val();
  if(isEmpty(assign_user_id)){
		assign_user_id = token;
  }
  var relate_task_id = $('#dataset_list option:selected').val();
  if(isEmpty(relate_task_id)){
	   alert("数据集对象不能为空。");
       return;
  }
  var labelpropertytaskid = $('#labelpropertytask_dataset option:selected').val();
  createsucced = true;
  label_task_create(task_name, relate_task_id, 2,assign_user_id,labelpropertytaskid,-1);
  if(createsucced){
    $(".ui.dataset.modal").modal("hide");
	 //$("#labelDataModal").modal('hide');
  }
  page(0,pageSize);
}

function submit_autolabeltask(){
  console.log($('#autolabeltaskname').val());
  var task_name = $('#autolabeltaskname').val();
  if (isEmpty(task_name) || task_name.length > 32){
       alert("自动标注任务名称不能为空或者不能超过32个字符。");
       return;
  }
  var model_id = $('#model_list option:selected').val();
  if(isEmpty(model_id)){
	   alert("标注模型不能为空。");
       return;
  }
  var relate_task_id = $('#dataset_list_auto option:selected').val();
  if(isEmpty(relate_task_id)){
	   alert("数据集对象不能为空。");
       return;
  }
  var assign_user_id = $('#label_assign_user option:selected').val();
  if(isEmpty(assign_user_id)){
		assign_user_id = token;
  }
  var labelpropertytaskid = $('#labelpropertytask_auto option:selected').val();
  createsucced = true;
  label_task_create(task_name, relate_task_id, 1,assign_user_id,labelpropertytaskid,model_id);
  if(createsucced){
	  $("#labelModal").modal('hide');
  }
  page(0,pageSize);
}


function label_task_create(task_name, relate_task_id, taskType,assign_user_id,labelpropertytaskid,model_id){
	
    
  var task_flow_type = $('#task_flow_type option:selected').val();
	var relate_other_label_task = [];
	
	if(task_flow_type == 2){
		var items = document.getElementsByName("category");                 
        for (i = 0; i < items.length; i++) {                    
            if (items[i].checked) {                        
                relate_other_label_task.push(items[i].value);                    
            }                
        }      
	}
	relate_other_label_task_jsonstr = JSON.stringify(relate_other_label_task);
	console.log("relate_task_id=" + relate_task_id);
	
    $.ajax({
       type:"POST",
       contentType:'application/json',
       url:ip + "/gitea/label-task/",
       dataType:"json",
       async:false,
       headers: {
          authorization:token,
        },
	   beforeSend: function (xhr) {
　　　　　　xhr.withCredentials = true
　　　　},
       data:JSON.stringify({'task_name':task_name,
                            'assign_user_id':assign_user_id,
                            'task_flow_type':task_flow_type,
                            'relate_task_id':relate_task_id,//task id
							              'relate_other_label_task': relate_other_label_task_jsonstr,
                            "taskType": taskType,
							              "appid": repoId,
							              "createUserName":userName,
							              "labelPropertyTaskId":labelpropertytaskid,
                            "modelId":model_id
           }),
       success:function(res){
		   console.log(res);
		   if(res.code == 0){
			  alert("标注任务创建成功!");
			  createsucced = true;
		   }
		   else{
			  alert("创建标注任务失败，" + res.message);
			  createsucced = false;
		   }
      },
	    error:function(response) {
		  redirect(response);
        }
   });
}

function list(current,pageSize){
      $.ajax({
       type:"GET",
       url:ip + "/gitea/label-task-page/",
       headers: {
          authorization:token,
        },
       dataType:"json",
       data:{'startPage':current,
       'pageSize':pageSize},
       async:false,
       success:function(json){
        tablePageData = json;
        tableData = json.data;
        //console.log(json);
      },
	    error:function(response) {
		  redirect(response);
        }
   });
}

var otherUserLabelTaskData;

function flow_type_sele_Change(sele){
	
	var task_flow_type = $('#task_flow_type option:selected').val();
	if(task_flow_type == 2){
		  var datasetid = $('#dataset_list option:selected').val();
          getOtherUserLabelTaskByDataSetId(datasetid);
		  var html = "<p>请选择该数据集要审核的标注任务</p>";
		  for(var i = 0; i < otherUserLabelTaskData.length; i++){
			  html += "<p><input type=\"checkbox\" name=\"category\" value=\"" + otherUserLabelTaskData[i].id + "\"/>" + otherUserLabelTaskData[i].task_name + "(标注人：" + otherUserLabelTaskData[i].assign_user + ")"  + "</p>";
		  }
		  document.getElementById('related_task_list').innerHTML=html; 
	}else{
		document.getElementById('related_task_list').innerHTML= ""; 
	}
}

function getOtherUserLabelTaskByDataSetId(datasetid){
    $.ajax({
       type:"GET",
       url:ip + "/gitea/label-related-task/" + datasetid + "/",
       headers: {
          authorization:token,
        },
       dataType:"json",
       async:false,
       success:function(json){
        otherUserLabelTaskData = json;
        console.log(json);
      },
	    error:function(response) {
		  redirect(response);
        }
   });
}

function dislpayUser(){
  var html="<option value=\"\" selected=\"\">请选择</option>";
	
  for (var i=0;i<userInfoData.length;i++){
        var row = "<option value=\""+userInfoData[i].id+
        "\">"+userInfoData[i].username+
        "</option>";
      
      html=html+row;
  }
  console.log(html);
  document.getElementById('assign_user').innerHTML=html; 
  document.getElementById('label_assign_user').innerHTML=html;
}


function getUser(){
    $.ajax({
       type:"GET",
       url:ip + "/api/queryAllUser/",
       headers: {
          authorization:token,
        },
       dataType:"json",
       async:false,
       success:function(json){
        userInfoData = json;
        console.log(json);
      },
	    error:function(response) {
		  redirect(response);
        }
   });
}


function delete_labeltask(){
  var stop = del();
  if (stop){
    return;
  }
  var Check = $("table[id='label_task_list'] input[type=checkbox]:checked");//在table中找input下类型为checkbox属性为选中状态的数据
  Check.each(function () {//遍历
    var row = $(this).parent("td").parent("tr");//获取选中行
    var id = row.find("[id='labeltask_id']").html();//获取name='Sid'的值
    delete_labeltask_byid(id);
    });
  page(0,pageSize);
}

function del(){
    if($("table[id='label_task_list'] input[type=checkbox]").is(":checked")) {
        if (confirm("确实要删除吗？")) {
            // alert("已经删除！");
            return false;
        } else {
            // alert("已经取消了删除操作");
            return true;
        }
    }else if($("table[id='label_task_list']").find("input").length=="0"){
        alert("暂无可删的数据！");
        return true;
    }else{
        alert("请先选择需要删除的选项！");
        return true;
    }
}

function delete_labeltask_byid(label_task_id){
  $.ajax({
    type:"DELETE",
    url:ip + "/gitea/label-task/",
    headers: {
       authorization:token,
     },
    dataType:"json",
    async:false,
    data:{'label_task_id': label_task_id},
    success:function(json){
      console.log(json);
    },
	error:function(response) {
		  redirect(response);
    }
 });
}

function getTaskTypeDesc(task_type){
	if(task_type == 1){
		return "图片-自动标注";
	}else if(task_type == 2){
		return "图片";
	}else if(task_type == 3){
		return "CT影像";
	}else if(task_type == 4){
		return "视频";
	}
	return "其它";
}

function getLabelDesc(task_flow_type){
	if(task_flow_type == 2){
		return "审核";
	}else{
		return "人工"
	}
}

function getTaskSataus(task_status,task_status_desc){
	if(task_status == 0){
		return "标注中：" + task_status_desc;
	}else if(task_status == 1){
		return "审核中：" + task_status_desc;
  }
  else if(task_status == -1){
    return "关联的数据集已经被删除。"
  }
  else if(task_status == 20){
    return "自动标注进行中:" + task_status_desc;
  }else if(task_status == 21){
    return task_status_desc;
  }
  return "";
}

function getVerify(task_status,id,task_type){
	console.log("task_status=" + task_status + " userType=" + userType);
	if(task_status == 0 && (userType == 1 || userType == 0) ){
		return "<a onclick=\"startToVerify(\'"+id+"\',\'" + task_type +"\');\" class=\"btn btn-xs btn-success\">转审核</a>&nbsp;&nbsp;&nbsp;";
	}else if(task_status == 1 && userType == 2){
		return "<a onclick=\"goVerify(\'"+id+"\',\'" + task_type +"\');\" class=\"btn btn-xs btn-success\">进入审核</a>&nbsp;&nbsp;&nbsp;" + "<a onclick=\"startToLabel(\'"+id+"\',\'" + task_type +"\');\" class=\"btn btn-xs btn-success\">转标注</a>&nbsp;&nbsp;&nbsp;";
	}else{
		return "";
	}
}

function getLabel(task_status,id,task_type,task_flow_type){
	if((task_status == 0 || task_status == 21) && (userType == 1 || userType == 0)){
		return "<a onclick=\"personLabel(\'" + id + "\'," + task_type + ")\"><b>" + getLabelDesc(task_flow_type) + "标注</b></a><br>";
	}else{
		return "";
	}
}


function display_list(){

  var html="<tr>\
            <th></th>\
            <th id=\"labeltask_head\"></th>\
            <th>标注任务名称</th>\
            <th>关联的数据集名称</th>\
            <th>数据类型</th>\
            <th>标注人员</th>\
			<th>审核人员</th>\
            <th>任务开始时间</th>\
            <th>任务状态</th>\
			<th>总标注数量</th>\
            <th>操作</th>\
            </tr>";
 for (var i=0;i<tableData.length;i++){
    var row = "<tr>\
            <td><input type=\"checkbox\" class=\"flat-grey list-child\"/></td>\
            <td id=\"labeltask_id\">"+tableData[i].id+"</td>\
            <td>"+tableData[i].task_name+"</td>\
            <td>"+tableData[i].relate_task_name+"</td>\
            <td>"+ getTaskTypeDesc(tableData[i].task_type) +"</td>\
            <td>"+tableData[i].assign_user+"</td>\
			<td>"+tableData[i].verify_user+"</td>\
            <td>"+tableData[i].task_add_time+"</td>\
            <td>"+getTaskSataus(tableData[i].task_status,tableData[i].task_status_desc)+"</td>\
			<td>"+tableData[i].total_label+"</td>\
            <td>" + 
			getLabel(tableData[i].task_status,tableData[i].id,tableData[i].task_type,tableData[i].task_flow_type)  + "<a onclick=\"setTaskId(\'"+tableData[i].id+"\');\"><b>导出标注</b></a>" 
			+ 
			"</td>\
            </tr>";

    html=html+row;
  }
  //console.log(html);
  document.getElementById('label_task_list').innerHTML=html;

  $('#label_task_list tr').find('td:eq(1)').hide();
  $('#label_task_list tr').find('th:eq(1)').hide();
  isNeedToRefresh = false;
  taskNeedIntervalToRefresh();
}

var isNeedToRefresh;
var refreshTimeId = [];
var refreshCount;
var refreshMaxTime;

function taskNeedIntervalToRefresh(){
  var isNeedToSetInterval = false;
	refreshMaxTime= 1;
	if(!isEmpty(tableData)){
        for (var i=0;i<tableData.length;i++){
		    if(tableData[i].task_status == 20){
				  console.log("有任务在自动标注中。需要自动刷新。");
				  isNeedToRefresh = true;//有任务在进行中才刷新，否则不刷新。
			    refreshMaxTime = 500;
			  	break;
			  }
	    }
	}
	if(!isEmpty(refreshTimeId)){
		for(var i =0; i < refreshTimeId.length ;i++){
			console.log("清除定时器。refreshTimeId=" + refreshTimeId[i]);
		    window.clearInterval(refreshTimeId[i]);
		}
		refreshTimeId = [];
	}
	if(isNeedToRefresh){
		refreshCount = 0;
		var tmpTimeId = self.setInterval("clockRefresh('" + refreshMaxTime +"')",5000);//5秒刷新
		refreshTimeId.push(tmpTimeId);
		console.log("开始刷新。tmpTimeId=" + tmpTimeId);
	}

}
function clockRefresh(refreshMaxTime){
  refreshCount = refreshCount + 1;
  if(refreshCount > refreshMaxTime){
     for(var i =0; i < refreshTimeId.length ;i++){
       console.log("超过刷新最大次数，清除定时器。timeId=" + refreshTimeId[i]);
       window.clearInterval(refreshTimeId[i]);
     }
     refreshTimeId = [];
     return;
  }
  var current = $('#displayPage1').text();
  console.log("开始刷新。current=" + current);
  if(current >= 1){
      page(current - 1,pageSize);
  }
}

function startToLabel(taskid, task_type){//从审核转回标注，标注人不变。
	 $.ajax({
       type:"PATCH",
       url:ip + "/gitea/label-task-status/",
       dataType:"json",
       async:false,
       headers: {
          authorization:token,
        },
       data:{
		   "label_task_id" : taskid,
		   "verify_user_id" : 0,
           "task_status" : 	0
	    },
       success:function(res){
		  console.log(res);
       },
	   error:function(response) {
		  redirect(response);
       }
   });
   var current =$('#displayPage1').text();
   page(current - 1,pageSize);
}

function startToVerify(taskid, task_type){
	$("#hide_labeltasktoverifyid").val(taskid);
	
	$.ajax({
       type:"GET",
       url:ip + "/api/queryVerifyUser/",
       headers: {
          authorization:token,
        },
       dataType:"json",
       async:false,
       success:function(json){
        console.log(json);
		var html="<option value=\"\" selected=\"\">请选择</option>";
        for (var i=0;i<json.length; i++){
           var row = "<option value=\""+json[i].id+"\">" +  json[i].username  +   "</option>";
           html=html+row;
        }
        document.getElementById('label_verify_user').innerHTML=html; 
      },
	  error:function(response) {
		  redirect(response);
      }
	});
	$("#startToVerify").modal('show');
	
}

function submit_labeltask_toverify(){
	var label_task_id =  $('#hide_labeltasktoverifyid').val();
	console.log("label_task_id=" +label_task_id);
	var verify_user_id =  $('#label_verify_user option:selected').val();
	//修改状态
   $.ajax({
       type:"PATCH",
       url:ip + "/gitea/label-task-status/",
       dataType:"json",
       async:false,
       headers: {
          authorization:token,
        },
       data:{
		   "label_task_id" : label_task_id,
		   "verify_user_id" : verify_user_id,
           "task_status" : 	1   
	    },
       success:function(res){
		  console.log(res);
       },
	   error:function(response) {
		  redirect(response);
       }
   });
   $("#startToVerify").modal('hide');
   var current =$('#displayPage1').text();
   page(current - 1,pageSize);
}
function goVerify(taskid, task_type){
	sessionStorage.setItem('label_task',taskid);
	sessionStorage.setItem('label_task_status',1);//审核
	console.log("task_type=" + task_type);
	if(task_type == 2 || task_type == 1 || task_type == 4){
		window.location.href="labeling.html";
	}else if(task_type == 3){
		window.location.href="labelingDcm.html";
	}
}


function personLabel(taskid, task_type){
	sessionStorage.setItem('label_task',taskid);
	sessionStorage.setItem('token',token);
  sessionStorage.setItem('userType',1);
  var pathname = window.location.pathname;
  var search = window.location.search
  var url_name = pathname + search
  sessionStorage.setItem('return_url',url_name);
	console.log("task_type=" + task_type);
	if(task_type == 2 || task_type == 1 || task_type == 4){
		window.open("/self/labeling.html?","标注"+new Date().getMilliseconds());
		//window.location.href="/self/labeling.html";
	}else if(task_type == 3){
		window.location.href="labelingDcm.html";
	}
	
}

function setMultiTaskId(){
	var Check = $("table[id='label_task_list'] input[type=checkbox]:checked");//在table中找input下类型为checkbox属性为选中状态的数据
	if(Check.length == 0){
		 alert("请选择一个或者多个标注数据进行导出。");
		 return;
	}
	var taskList = [];
  Check.each(function () {//遍历
    var row = $(this).parent("td").parent("tr");//获取选中行
    var id = row.find("[id='labeltask_id']").html();//获取name='Sid'的值
    taskList.push(id);
    //$('#hide_labeltaskid').val(id);
  });
    
    setTaskId(JSON.stringify(taskList));
}



function setTaskId(labeltaskid){
	$('#hide_labeltaskid').val(labeltaskid);
  console.log("go here1");
	bar.style.width='1%';
	document.getElementById('text-progress').innerHTML="0%";
	document.getElementById("predtask_id").removeAttribute("disabled");
	console.log("go here2");
  $(".ui.export.modal").modal("show");
}

function isBeetween(score_threshhold){
	if(isEmpty(score_threshhold)){
		return true;
	}
	var regPos = /^\d+(\.\d+)?$/; //非负浮点数
    if(!regPos.test(score_threshhold)){
		   return false;
	}else{
		if(score_threshhold >1 || score_threshhold < 0){
			return false;
		}
	}
	return true;
}


function downloadFile(){

	var labeltaskid = $('#hide_labeltaskid').val();
  var isNeedPicture =  $('#isNeedPicture option:selected').val();
  var maxscore = $('#maxscore').val();
	var minscore = $('#minscore').val();
  var exportFormat = $('#exportFormat option:selected').val();
	if(isNeedPicture == 3){
		if(!isBeetween(maxscore)){
			alert("标注得分最大值应该填写0--1.0之间的数值。");
			return;
		}
		if(!isBeetween(minscore)){
			alert("标注得分最小值应该填写0--1.0之间的数值。");
			return;
		}
		if(!isEmpty(maxscore) && !isEmpty(minscore)){
			if(minscore>maxscore){
				alert("标注得分最小值应该小于标注得分最大值。");
			    return;
			}
		}
	}
	document.getElementById("predtask_id").setAttribute("disabled", true);
	var taskreturnid = "";
	$.ajax({
       type:"GET",
       url:ip + "/gitea-label-task-export/",
       headers: {
          authorization:token,
        },
       dataType:"json",
	   data:{
		   "label_task_id" : labeltaskid,
       "needPicture" : 	isNeedPicture,
       "exportFormat":exportFormat,
		   "maxscore":maxscore,
		   "minscore":minscore
	    },
       async:false,
       success:function(json){
        taskreturnid = json.message;
        console.log(json);
       },
	    error:function(response) {
		  redirect(response);
        }
   });
   console.log("taskreturnid=" +taskreturnid);
   if(!isEmpty(taskreturnid)){
	   setIntervalToDo(taskreturnid);
   }
}


var timeId=[];
var count;
var progress;

function setIntervalToDo(taskreturnid){
	count=0;
	var tmpTimeId = self.setInterval("clock('" + taskreturnid +"')",1000);//5秒刷新
	timeId.push(tmpTimeId);
	console.log("开始刷新。timeId=" + tmpTimeId);
}

function clock(taskreturnid){
   count++;
   if(count > 600 ){
       for(var i = 0;i < timeId.length; i++){
		    console.log("清除定时器1。exportTimeId=" + timeId[i]);
		    window.clearInterval(timeId[i]);
	   }
	   timeId = [];
     $(".ui.export.modal").modal("hide");
	   return;
   }
   $.ajax({
       type:"GET",
       url:ip + "/api/query-download-progress/",
       headers: {
          authorization:token,
        },
       dataType:"json",
	   data:{'taskId': taskreturnid},
       async:false,
       success:function(json){
        progress = json;
        console.log(json);
       },
	   error:function(response) {
		  progress = null;
		  console.log('query return null.');
		  redirect(response);
       }
   });
   if(!isEmpty(progress)){
	   if(progress.progress >= 100){

		   var iSpeed = progress.progress;
		   bar.style.width=iSpeed+'%';
		   document.getElementById('text-progress').innerHTML=iSpeed+'%' + ",开始下载文件。"
		   
		    for(var i = 0;i < timeId.length; i++){
		       console.log("清除定时器2。exportTimeId=" + timeId[i]);
		       window.clearInterval(timeId[i]);
	        }
	        timeId = [];
		   
		   var url = ip + "/api/label-file-download/";
		   var $iframe = $('<iframe />');
		   var $form = $('<form  method="get" target="_self"/>');
		   $form.attr('action', url); //设置get的url地址

		   $form.append('<input type="hidden"  name="taskId" value="' + taskreturnid + '" />');
		
		   $iframe.append($form);
		   $(document.body).append($iframe);
		   $form[0].submit();//提交表单
		   $iframe.remove();//移除框架
		   $(".ui.export.modal").modal("hide");
	   }else{
		   //更新进度
			var iSpeed = progress.progress;
			bar.style.width=iSpeed+'%';
			document.getElementById('text-progress').innerHTML=iSpeed+'%'
	   }
   }else{
	   count = 600;
   }
}


function page(current,pageSize){
  list(current,pageSize);
  display_list();
  setPage(tablePageData,pageSize);
  sessionStorage.setItem('label_task_page',current);
}

function nextPage(){
   var current = $('#displayPage1').text();
   console.log("current=" + current);
   page(current,pageSize);
}

function prePage(){
  var current =$('#displayPage1').text();
  console.log("current=" + current);
  if(current > 1){
    console.log("current=" + (current - 2));
    page(current - 2,pageSize);
  } 
}


function goPage(){
   var goNum = $('#goNum').val();

    var pageTotal = $("#totalNum").text();
    var pageNum = parseInt(pageTotal/pageSize);
    if(pageTotal%pageSize!=0){
        pageNum += 1;
    }else {
        pageNum = pageNum;
    }
    if (goNum<=0){
      alert("请输入大于0的数值");
    }
    else if(goNum<=pageNum){
        page(goNum - 1,pageSize);
    }
    else{
        alert("不能超出总页码！");
    }
}

$("#goNum").keydown(function (e) {
    if (e.keyCode == 13) {
        goPage();
    }
});



function setPage(pageData,pageSize){
  if (isEmpty(pageData)){
    return;
  }
   var startIndex = pageData.current * pageSize;
  if(pageData.total > 0){
	  startIndex = startIndex + 1;
  }
  $('#startIndex').text(startIndex);
  $('#endIndex').text(pageData.current * pageSize + pageData.data.length);
  $('#totalNum').text(pageData.total);
  $('#displayPage1').text(pageData.current + 1);

  console.log("set prePage status, pageData.current=" + pageData.current);

  if(pageData.current == 0){
    console.log("set prePage disabled.");
    $('#prePage').removeAttr("href");
    $('#prePage').attr('style','color:#f5f5f6;');
  }
  else{
    $('#prePage').attr("href","javascript:prePage()");
    $('#prePage').attr('style','color:#000;');
  }

  if((pageData.current + 1) * pageSize >= pageData.total){
    console.log("set nextPage disabled.");
    $('#nextPage').removeAttr("href");
  }
  else{
    $('#nextPage').attr("href","javascript:nextPage()");
  }

  var pageTotal = pageData.total;
  var pageNum = parseInt(pageTotal/pageSize);
  if(pageTotal%pageSize!=0){
      pageNum += 1;
  }else {
      pageNum = pageNum;
  }
  $("#totalPageNum").text(pageNum);
}
var tmpCurrent = sessionStorage.getItem("label_task_page");
if(isEmpty(tmpCurrent)){
	tmpCurrent = 0;
}
page(tmpCurrent,pageSize);

