<template>
  <div class="content">
    <div class="content-r">
      <div class="title">{{ $t('org.organization') }}</div>
      <div class="condition-a">
        <div class="tab-c">
          <div class="tab-item" v-for="(item, index) in tabList" :class="tabIndex == item.key ? 'focus' : ''"
            :key="item.key">
            {{ item.label }}
          </div>
        </div>
        <div class="condition-b">
          <div class="search-c">
            <el-input class="search-keyword" :placeholder="$t('org.searchOrg')" v-model="conds.q"
              @keyup.enter.native="search">
              <i slot="suffix" class="el-input__icon el-icon-search" @click="search"></i>
            </el-input>
          </div>
          <el-button class="create-btn" type="primary" @click="goCreate">
            <div class="btn-content">
              <i class="ri-add-box-line"></i>
              <span>{{ $t('org.createOrg') }}</span>
            </div>
          </el-button>
        </div>
      </div>
      <div class="table-c">
        <a class="org-box" v-for="item in tableData" :key="item.org_id" :href="'/' + item.org_name">
          <div class="item-b">
            <div class="header-w">
              <div class="left-c">
                <img :src="item.rel_avatar_link" alt=""/>
                <span class="name nowrap" :title="item.org_name">{{ item.org_name }}</span>
              </div>
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" width="20" height="20" fill="none"><defs></defs><g><rect width="48" height="48" fill="white" fill-opacity="0.01"></rect><path d="M9 18V42H39V18L24 6L9 18Z" fill="none"></path><path d="M9 42V18L4 22L24 6L44 22L39 18V42H9Z" stroke="#bbb" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path><path d="M19 29V42H29V29H19Z" fill="none" stroke="#bbb" stroke-width="4" stroke-linejoin="round"></path><path d="M9 42H39" stroke="#bbb" stroke-width="4" stroke-linecap="round"></path></g></svg>
            </div>
            <p class="desc nowrap" :title="item.description"> {{ item.description }} </p>
          </div>
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import { listMyOrgUser } from '~/apis/modules/organization';
export default {
  data() {
    return {
      tabList: [{
        key: 'owned',
        label: this.$t('org.myOrg'),
      }],
      tabIndex: 'owned',
      conds: {
        q: '',
      },
      tableData: [],
      loading: false
    };
  },
  components: { },
  computed:{
    
  },
  methods: {
    search() {
      this.conds.q = this.conds.q.trim();
      this.getListData();
    },
    async getListData(){
      try{
        this.loading = true
        const response = await listMyOrgUser({
          q: this.conds.q
        })
        const res = response.data
        if(res.Code === 0){
          this.tableData  = res.Data
        }else{
          this.$message.error( res.Msg|| '获取组织列表失败')
        }
        console.log(response)
      } catch(error){
        this.$message.error( error|| '获取组织列表失败')
      }finally{
        this.loading = false        
      }
    },
    goCreate(){
      window.location.href = '/org/create';
    },
  },
  beforeMount() {
    
  },
  mounted() { 
    this.search();
  },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.content {
  display: flex;
  min-height: 100%;
  padding-left: 20px;
  .content-r{
    flex: 1;
    width: 0;
    padding: 30px 20px;
    padding-right: 36px;
    .title {
      color: rgb(16, 16, 16);
      font-size: 18px;
      text-align: left;
      font-weight: bold;
      padding: 6px 0 16px 0;
    }
    .condition-a {
      display: flex;
      align-items: center;
      justify-content: space-between;
      margin-bottom: 4px;
      flex-wrap: wrap;

      .tab-c {
        display: flex;
        align-items: center;
        margin-bottom: 16px;

        .tab-item {
          height: 32px;
          display: flex;
          align-items: center;
          margin-right: 6px;
          padding: 0 6px;
          border-bottom: 2px solid rgba(51, 38, 98, 0.3);
          font-size: 16px;
          box-sizing: border-box;
          cursor: pointer;
          color: rgba(16, 16, 16, 0.5);

          &:hover {
            color: rgba(16, 16, 16, 1);
            border-color: rgba(51, 38, 98, 1);
          }

          &.focus {
            color: rgba(0, 102, 255, 1);
            border-bottom: 2px solid rgba(0, 102, 255, 1);
          }
        }
      }

      .create-btn {
        margin-left: 14px;
        display: flex;
        align-items: center;
        height: 36px;
        font-size: 14px;
        background: rgba(22, 132, 252, 0.9);
        border-radius: 4px;

        &:active {
          background: rgb(22, 132, 252, 1);
        }


        &:focus,
        &:hover {
          background: rgba(22, 132, 252, 0.8);
        }

        .btn-content {
          display: flex;
          align-items: center;

          i {
            font-size: 14px;
            margin-right: 10px;
          }
        }
      }
    }


    .condition-b {
      display: flex;
      align-items: center;
      justify-content: space-between;
      margin-bottom: 16px;

      .search-c {
        display: flex;
        align-items: center;

        .el-select {
          margin-right: 15px;

          /deep/.el-input__inner {
            height: 36px;
          }
        }

        .el-input {
          width: 320px;
          margin-right: 15px;

          /deep/.el-input__inner {
            height: 36px;
          }
        }

        /deep/ .el-icon-search {
          cursor: pointer;
          color: rgb(16, 16, 16);
        }
      }
    }

    .table-c {
      display: grid;
      grid-template-columns: repeat(4, 1fr);
      width: 100%;
      gap: 27px;
      .org-box{
        padding: 24px 31px 18px 21px;
        border-radius: 10px;
        background: linear-gradient(166.14deg, #fffbed 0.84%, #ffffff 54.19%);
        height: 100px;
        color: rgba(16,16,16,1);
        box-shadow: 0px 0px 20px 0px rgba(221,221,221,0.5);
        border: 1px solid rgba(255,255,255,1);
        min-width: 0; /* 允许收缩 */
        &:hover{
          border-color: rgba(187,193,246,0.7);
          background: #fff;
          .name{
            color: rgba(0,26,199,0.9) !important;
          }
        }
        .item-b{
          display: flex;
          flex-direction: column;
          height: 100%;
          justify-content: space-between;
          .header-w{
            display: flex;
            align-items: center;
            justify-content: space-between;
            .left-c{
              display: flex;
              align-items: center;
              flex: 1;
              min-width: 0;
              img{
                width: 32px;
                height: 32px;
              }
              .name{
                color: rgba(16,16,16,1);
                font-size: 18px;
                font-weight: 700;
                margin-left: 14px;
              }
            }
          }
          .desc{
            margin-left: 46px;
          }
        }
      }
    }
  }

}
@media only screen and (max-width: 800px) {
  .content {
    padding-left: 0 !important;
    .content-r{
      padding: 20px 16px;
      .condition-a{
        justify-content: center;
      }
      .condition-b{
        display: none;
      }
    }
  }
}
/* 中等屏幕显示3列 */
@media (max-width: 1550px) {
  .table-c {
    grid-template-columns: repeat(3, 1fr) !important;
  }
}

/* 小屏幕显示2列 */
@media (max-width: 1200px) {
  .table-c {
    grid-template-columns: repeat(2, 1fr) !important;
  }
}

/* 超小屏幕显示1列 */
@media (max-width: 768px) {
  .table-c {
    grid-template-columns: 1fr !important;
  }
}
</style>
