<template>
  <div>
    <TopHeader :menu="'my_view'"></TopHeader>
    <div class="ui container">
      <div class="table-container">
        <div class="table-title">科技2030项目</div>
        <div class="table-wrap">
          <el-table ref="tableRef" border :data="tableData" style="width:100%;" v-loading="loading" stripe row-key="id"  max-height="650">
            <el-table-column label="序号" type="index" align="center" header-align="center" width="50"
              fixed></el-table-column>
            <el-table-column label="启智项目名称" align="center" header-align="center" fixed min-width="140">
              <template slot-scope="scope">
                <span v-if="scope.row.status == 5">
                  {{ `${scope.row.url.split('/')[3]}/${scope.row.url.split('/')[4]}` }}
                </span>
                <span v-else-if="scope.row.status == 3 || scope.row.status == 4">
                  {{ `${scope.row.url}` }}
                </span>
                <a v-else target="_blank" :href="`/${scope.row.repo_owner_name}/${scope.row.repo_name}`">
                  {{ `${scope.row.repo_owner_name}/${scope.row.repo_name}` }}
                </a>
              </template>
            </el-table-column>
            <el-table-column prop="name" label="科技项目名称" align="center" header-align="center" fixed
              min-width="200"></el-table-column>
            <el-table-column prop="status" label="展示状态" align="center" header-align="center" fixed>
              <template slot-scope="scope">
                <span style="color:rgb(255, 37, 37)"
                  :style="scope.row.status == 1 ? { color: 'rgb(56, 158, 13)' } : ''">{{
                  statusMap[scope.row.status] }}</span>
              </template>
            </el-table-column>
            <el-table-column prop="no" label="项目立项编号" align="center" header-align="center"
              min-width="120"></el-table-column>
            <el-table-column prop="institution" label="项目承担单位" align="center" header-align="center"
              min-width="120"></el-table-column>
            <el-table-column prop="all_institution" label="所有参与单位" align="center" header-align="center"
              min-width="280"></el-table-column>
            <el-table-column prop="contribution_institution" label="成果贡献单位" align="center" header-align="center"
              min-width="280"></el-table-column>
            <el-table-column prop="contribution_institution" label="申请时间" align="center" header-align="center"
              min-width="120">
              <template slot-scope="scope">
                <span>{{ dateFormat(scope.row.created_unix) }}</span>
              </template>
            </el-table-column>
          </el-table>
        </div>
        <div class="center">
          <el-pagination ref="paginationRef" background @current-change="currentChange" @size-change="sizeChange"
            :current-page.sync="page" :page-sizes="pageSizes" :page-size.sync="pageSize"
            layout="total, sizes, prev, pager, next, jumper" :total="total">
          </el-pagination>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import dayjs from 'dayjs';
import TopHeader from '../components/TopHeader.vue';
import { getTechMyList } from '~/apis/modules/tech';

export default {
  data() {
    return {
      loading: false,
      tableData: [],

      page: 1,
      pageSizes: [15, 30, 50],
      pageSize: 15,
      total: 0,

      statusMap: {
        '1': '已展示',
        '2': '未展示',
        '3': '项目迁移中',
        '4': '项目迁移失败',
        '5': '项目不存在',
      },
    };
  },
  components: { TopHeader },
  methods: {
    getData() {
      this.loading = true;
      getTechMyList({
        page: this.page,
        pageSize: this.pageSize,
      }).then(res => {
        this.loading = false;
        const { total, data } = res.data;
        this.tableData = data.map((item, index) => {
          return { ...item, }
        });
        this.total = total;
      }).catch(err => {
        this.loading = false;
        console.log(err);
      });
    },
    currentChange(page) {
      this.page = page;
      this.getData();
    },
    sizeChange(pageSize) {
      this.pageSize = pageSize;
      this.getData();
    },
    dateFormat(unix) {
      return dayjs(unix * 1000).format('YYYY-MM-DD HH:mm');
    }
  },
  beforeMount() {
    this.getData();
  },
  mounted() { },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.table-container {
  margin-top: 30px;

  .table-title {
    height: 43px;
    font-size: 16px;
    font-weight: 700;
    padding: 15px;
    color: rgb(16, 16, 16);
    border-color: rgb(212, 212, 213);
    border-width: 1px;
    border-style: solid;
    border-radius: 5px 5px 0px 0px;
    background: rgb(240, 240, 240);
    display: flex;
    align-items: center;
  }

  .table-wrap {
    overflow-x: auto;
    margin-bottom: 20px;

    /deep/ .el-table__header {
      th {
        background: rgb(249, 249, 249);
        font-size: 12px;
        color: rgb(136, 136, 136);
        font-weight: normal;
      }
    }

    /deep/ .el-table__body {
      td {
        font-size: 12px;
      }
    }

    /deep/ .el-radio__label {
      display: none;
    }
  }
}
</style>
