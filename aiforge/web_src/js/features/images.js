// import Images from '../components/images/Images.vue';
// import  adminImages from '../components/images/adminImages.vue';
// import  selectImages from '../components/images/selectImages.vue';
// import  selectGrampusImages from '../components/images/selectGrampusImages.vue';
// import Vue from 'vue';
// export default async function initImage(){
//     function validate() {
//         $("#form_image")
//         .form({
//             on: 'blur',
//             // inline:true,
//             fields: {
//                 tag: {
//                     identifier  : 'tag',
//                     rules: [
//                     {
//                         type: 'regExp[/^[a-zA-Z][\\w|\\-|\\.]*$/]',
//                     }
//                     ]
//                 },
//                 description:{
//                     identifier  : 'description',
//                     rules: [
//                     {
//                         type: 'empty',
                        
//                     }
//                     ]
//                 },
//                 place:{
//                     identifier  : 'place',
//                     rules: [
//                     {
//                         type: 'empty',
                        
//                     }
//                     ]
//                 },
//             }
//         })
//     }
//     function $params(obj) {
//         var str = [];
//         for (var p in obj) {
//             str.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));
//         }
//         return str.join("&");
//     }
//     function initDropdown(){
//         $('#dropdown_image')
//         .dropdown({
//             allowAdditions: true,
//             onChange: function(value, text, $selectedItem) {
//                 $('#course_label_item').empty()
//             }
//         })
//         $('#dropdown_image input.search').bind('input propertychange', function (event) {
//             // $("#dropdown_container").removeAttr("style"); 
//             const query  = $('input.search').val()
//             if(!query){
//                 $('#course_label_item').empty()
//             }else{
//                 $.get(`/api/v1/image/topics/search?q=${query}`,(data)=>{
//                     if(data.topics.length!==0){
//                         let html=''
//                         $('#course_label_item').empty()
//                         data.topics.forEach(element => {
//                             html += `<div class="item" data-value="${element.topic_name}">${element.topic_name}</div>`
//                         });
//                         $('#course_label_item').append(html)
//                     }
//                 })
//             }
//         });
        
//     }
//     validate()
//     initDropdown()
//     let link = $('.submit-image-tmplvalue').data('link')
//     let pageform = $('.submit-image-tmplvalue').data('edit-page') || ''
//     let repoLink = $('.submit-image-tmplvalue').data('repo-link')
//     function postImage(formData) {
//         $("#mask").css({"display":"block","z-index":"999"})
//         $.ajax({
//             url:link,
//             type:'POST',
//             data:formData,
//             success:function(res){
//                 if(res.Code===1){
//                     $('.ui.negative.message').text(res.Message).show().delay(2500).fadeOut();
//                 }else if(res.Code==0){     
//                     if (location.href.indexOf('imageAdmin') !== -1) {
//                         let params = location.href.split('imageAdmin')[1]
//                         location.href = `${window.config.AppSubUrl}/admin/images${params}`
//                     }else{
//                         location.href = `${window.config.AppSubUrl}/explore/images?type=myimage`
//                     }
//                 }
//             },
//             error: function(xhr){
//             // 隐藏 loading
//             // 只有请求不正常（状态码不为200）才会执行
//                 $('.ui.negative.message').html(xhr.responseText).show().delay(1500).fadeOut();
//             },
//             complete:function(xhr){
//                 $("#mask").css({"display":"none","z-index":"1"})
//             }
//         })
//     }
//     $('#adminCommitImage').on('click', function (e) {
//         document.querySelectorAll('#adminCommitImage a').forEach((item) => { 
//             item.classList.remove('active')
//         })
//         e.target.classList.add('active')
//         document.querySelector('input[name="type"]').value=e.target.dataset.type
//     })
//     $('.ui.create_image.green.button').click(()=>{
//         let pattenTag = new RegExp(/^[a-zA-Z][\w|\-|\.]*$/)
//         if(!pattenTag.test($("input[name='tag']").val())){
//             $("input[name='tag']").parent().addClass('error')
//             return false
//         }
//         if(!$("textarea[name='description']").val()){
//             $("textarea[name='description']").parent().addClass('error')
//             return false
//         }
//         if($("input[name='place']").length>0&&!$("input[name='place']").val()){
//             $("input[name='place']").parent().addClass('error')
//             return false
//         }
        
//         const postData = {
//                         _csrf:$("input[name='_csrf']").val(),
//                         tag:$("input[name='tag']").val(),
//                         description:$("textarea[name='description']").val(),
//                         type:$("input[name='type']").val(),
//                         isPrivate:$("input[name='isPrivate']:checked").val(),
//                         topics:$("input[name='topics']").val(),
//                         id:$("input[name='id']").val()
//                         }
//         if($("input[name='place']").val()&&$("input[name='isRecommend']:checked").val()){
//             postData.isRecommend = $("input[name='isRecommend']:checked").val()
//             postData.place = $("input[name='place']").val()
//         }
//         let formData = $params(postData)
//         if($("input[name='edit']").val()=="edit" || $("input[name='type']").val()==="2"){
//             postImage(formData)
//         }
//         else{
//             $.ajax({
//                 url:link+'/check',
//                 type:'POST',
//                 data:formData,
//                 success:function(res){
//                     if(res.Code===1){
//                         $('.ui.modal.image_confirm_submit')
//                         .modal({
//                             onApprove: function() {
//                                 postImage(formData)
//                             },
                            
//                         })
//                         .modal('show')
//                     }else if(res.Code==0){
//                         postImage(formData)
//                     }
//                 },
//                 error: function(xhr){
//                     $('.ui.negative.message').text(xhr.responseText).show().delay(1500).fadeOut();
//                 }
//             })
//         }
//         return false
//     })
//     $('#cancel_submit_image').click(()=>{
//         if(pageform=='submit'){
//             location.href = `${window.config.AppSubUrl}${repoLink}/debugjob?debugListType=all`
//         }else if(pageform=='imageSquare' || pageform=='apply'){
//             location.href = `${window.config.AppSubUrl}/explore/images?type=myimage`
//         }else if(pageform){
//             location.href = `${window.config.AppSubUrl}/admin/images`
//         }
//     })
    
//     function initVueImages() {
//         const el = document.getElementById('images');
//         if (!el) {
//             return;
//         }
        
//         new Vue({
//             el:el,            
//             render: h => h(Images)
//         });
//     }
//     function initVueAdminImages() {
//         const el = document.getElementById('images-admin');
//         if (!el) {
//             return;
//         }
        
//         new Vue({
//             el:el,            
//             render: h => h(adminImages)
//         });
//     }
//     function initVueselectImages() {
//         const el = document.getElementById('images-new-cb');
//         if (!el) {
//             return;
//         }
        
//         new Vue({
//             el:el,            
//             render: h => h(selectImages)
//         });
//     }
//     // function initVueselectGrampusImages() {
//     //     const el = document.getElementById('images-new-grampus');
//     //     if (!el) {
//     //         return;
//     //     }
        
//     //     new Vue({
//     //         el:el,            
//     //         render: h => h(selectGrampusImages)
//     //     });
//     // }
//     initVueImages()
//     initVueAdminImages()
//     initVueselectImages()
//     // initVueselectGrampusImages()
// }