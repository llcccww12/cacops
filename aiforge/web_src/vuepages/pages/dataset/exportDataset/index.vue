<template>
    <div class="content content-padding">
        <form id="formId" method="POST" class="ui form">
            <div class="ui error message">
            </div>
            <input type="hidden" name="trainTaskCreate" value="true">
            <div class="inline required field">
                <label class="label_color" for="">可用集群</label>
                <div class="ui blue mini menu compact selectcloudbrain" id="adminCommitImage">
                    <a class="active item" data-type="0">
                        <svg class="svg" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="16" height="16">
                            <path fill="none" d="M0 0h24v24H0z"/>
                            <path d="M3 2.992C3 2.444 3.445 2 3.993 2h16.014a1 1 0 0 1 .993.992v18.016a.993.993 0 0 1-.993.992H3.993A1 1 0 0 1 3 21.008V2.992zM19 11V4H5v7h14zm0 2H5v7h14v-7zM9 6h6v2H9V6zm0 9h6v2H9v-2z"/>
                        </svg>
                        启智CPU/GPU
                    </a>
                    <a class="item" data-type="1">
                        <svg class="svg" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="16" height="16">
                            <path fill="none" d="M0 0h24v24H0z"/>
                            <path d="M3 2.992C3 2.444 3.445 2 3.993 2h16.014a1 1 0 0 1 .993.992v18.016a.993.993 0 0 1-.993.992H3.993A1 1 0 0 1 3 21.008V2.992zM19 11V4H5v7h14zm0 2H5v7h14v-7zM9 6h6v2H9V6zm0 9h6v2H9v-2z"/>
                        </svg>
                        智算CPU/GPU
                    </a>
                </div>
            </div>
            <div class="inline required field">
                <label class="label_color" for="">文件描述</label>
                <textarea style="width: 80%;" required id="description" name="description" rows="3" maxlength="255" placeholder="" onchange="this.value=this.value.substring(0, 255)" onkeydown="this.value=this.value.substring(0, 255)" onkeyup="this.value=this.value.substring(0, 255)"></textarea>
            </div>

            <div class="inline field" style="margin-left: 75px;">
                <button  type="button" class="ui create_train_job green button"
                    style="position: absolute;">
                    sdfsdfsdf
                </button>
            </div>
        </form>
        <div class="actions" style="display: inline-block;margin-left: 180px;">
            <button type="button" class="ui button cancel">ccc</button>
        </div>
    </div>


  </template>
  <script>
  import { getModelFile,exportExistDataset } from "~/apis/modules/dataset";
  
  export default {

    data() {
      return {

      }
    },
    computed: {},
    methods: {
        getModelFileLiST(){
            getModelFile().then((res)=>{})
        }
    },
    mounted() {

    },
  }
  </script>
  <style scoped>

  </style>