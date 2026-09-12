<template>
  <div class="teams-c" v-if="pageData.IsOrganizationMember">
    <div class="teams-head">
      <div class="teams-head-l"><i class="icon ri-group-line"></i>{{ $t('org.orgTeam') }}</div>
      <a class="teams-head-r" :href="`${pageData.OrgLink}/teams`">
        {{ pageData.NumTeams }} <i class="icon ri-arrow-right-s-line"></i>
      </a>
    </div>
    <div class="teams-list">
      <div class="teams-item" v-for="item in pageData.Teams">
        <div><a :href="`${pageData.OrgLink}/teams/${item.LowerName}`" :title="item.Name">{{ item.Name }}</a></div>
        <p>
          <a :href="`${pageData.OrgLink}/teams/${item.LowerName}`"><strong>{{ item.NumMembers }}</strong>
            {{ item.NumMembers > 1 ? $t('org.numMembers') : $t('org.numMember') }}
          </a> ·
          <a :href="`${pageData.OrgLink}/teams/${item.LowerName}/repositories`"><strong>{{ item.NumRepos }}</strong>
            {{ item.NumRepos > 1 ? $t('org.numRepos') : $t('org.numRepo') }}
          </a>
        </p>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: "Teams",
  props: {
    pageData: { type: Object, default: () => { } },
  },
};
</script>

<style scoped lang="less">
.teams-c {
  .teams-head {
    display: flex;
    align-items: center;
    justify-content: space-between;
    height: 40px;

    .teams-head-l {
      display: flex;
      align-items: center;
      color: rgba(16, 16, 16, 1);
      font-size: 16px;
      font-weight: 700;

      .icon {
        font-size: 24px;
      }
    }

    .teams-head-r {
      display: flex;
      align-items: center;
      color: rgba(136, 136, 136, 1);
      font-size: 14px;

      &:hover {
        color: rgba(136, 136, 136, 0.8);
      }

      .icon {
        font-size: 20px;
      }
    }
  }

  .teams-list {
    margin-bottom: 6px;
    padding: 0 8px;

    .teams-item {
      margin-bottom: 8px;

      >div {
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;

        a {
          color: rgba(3, 102, 214, 1);
          font-size: 16px;
          font-weight: 600;
          line-height: 24px;

          &:hover {
            opacity: 0.8;
          }
        }
      }

      p {
        a {
          color: #101010;
          font-size: 14px;
          line-height: 20px;

          &:hover {
            opacity: 0.8;
          }
        }
      }
    }
  }
}
</style>
