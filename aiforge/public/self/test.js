
function displayDir(uuid){
	  console.log('uuid 1=' + uuid);
	  
	  var html="<tr>\
	            <th></th>\
	            <th id=\"dataset_head\"></th>\
	            <th>数据集名称</th>\
	            <th>数据集类型</th>\
				<th>数据集描述</th>\
				<th>数据集创建者</th>\
	            </tr>";
	  
	   for (var i=0;i<1;i++){
		       var row = "<tr>\
		               <td><input type=\"checkbox\"/></td>\
		               <td id=\"dataset_id\">"+uuid+"</td>\
		               <td>" + uuid +"</td>\
		               <td>测试</td>\
		   			<td>测试</td>\
		   			<td>测试</td>\
		               </tr>";
		       html=html+row;
		     }
	  
	  document.getElementById('dataset-files-table').innerHTML=html;
	  console.log('uuid 2=' + uuid);
}
