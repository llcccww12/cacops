(function () {
  const menuDataInit = {
    getWorkspaceMenu: function(isOrgMode, orgName) {
      if (!isOrgMode) {
        // 个人工作台菜单 (原有逻辑)
        return [{
          name: '我的工作台',
          name_en: 'My Workspace',
          children: [
            { name: '概览', name_en: 'Overview', icon: 'overview', url: '/dashboard' },
            { name: '项目', name_en: 'Repository', icon: 'repo', url: '/repositories' },
            { name: '计算任务', name_en: 'Computing task', icon: 'task', url: '/cloudbrains' },
            { name: '计算任务模板', name_en: 'Computing Task Template', icon: 'aitasktmpl', url: '/ai_task_tmpl/list_my' },
            { name: '数据集', name_en: 'Dataset', icon: 'dataset', url: '/explore/datasets_my' },
            { name: '模型', name_en: 'Model', icon: 'model', url: '/explore/models_my' },
            { name: '镜像', name_en: 'Image', icon: 'image', url: '/explore/images_my' },
            { name: '组织', name_en: 'Organization', icon: 'org', url: '/explore/org_my' }
          ]
        }, {
          name: '资源使用情况',
          name_en: 'Resource Usage',
          children: [
            { name: '算力积分', name_en: 'Calculation points', icon: 'point', url: '/reward/point' },
            { name: '存储管理', name_en: 'Storage management', icon: 'storage', url: '/storages' }
          ]
        }];
      } else {
        // 组织工作台菜单 (新逻辑)
        // 注意：url 中需要替换 :orgName 为实际的 orgName
        const orgPrefix = `/org/${orgName}`;
        return [{
          name: `组织工作台`, // 根据需求，显示组织名称
          name_en: `Org Workspace`,
          children: [
            { name: '组织主页', name_en: 'Org Home', icon: 'overview', url: `/${orgName}` },
            // { name: '概览', name_en: 'Overview', icon: 'overview', url: `${orgPrefix}/dashboard` }, // 假设有组织概览页
            { name: '项目', name_en: 'Repositories', icon: 'repo', url: `${orgPrefix}/repositories` },
            { name: '数据集', name_en: 'Dataset', icon: 'dataset', url: `${orgPrefix}/datasets` }, // 假设路径
            { name: '模型', name_en: 'Model', icon: 'model', url: `${orgPrefix}/models` }, // 假设路径
            // { name: '团队', name_en: 'Teams', icon: 'team', url: `${orgPrefix}/teams` },
            // { name: '成员', name_en: 'Members', icon: 'member', url: `${orgPrefix}/members` },
            
          ]
        }, {
          name: '资源使用情况',
          name_en: 'Resource Usage',
          children: [
            {
              name: '存储管理',
              name_en: 'Storage management',
              icon: 'storage',
              url: `${orgPrefix}/storages`,
            }
          ]
        }];
      }
    },
    // 其他菜单（modelbase, resources, computingpower）保持不变
    modelbase: [{
      name: '大模型基地',
      name_en: 'Large Model Base',
      url: '/modelbase'
    }],
    agent: [{
      name: '智能体',
      name_en: 'Agent',
      url: '/agent'
    }],
    resources: [{
      name: '公开资源',
      name_en: 'Open Resources',
      children: [{
        name: '数据集',
        name_en: 'Dataset',
        icon: 'dataset',
        url: '/explore/datasets',
      }, {
        name: '模型',
        name_en: 'Model',
        icon: 'model',
        url: '/explore/models',
      }, {
        name: '项目',
        name_en: 'Repository',
        icon: 'repo',
        url: '/explore/repos/square',
      }, {
        name: '镜像',
        name_en: 'Image',
        icon: 'image',
        url: '/explore/images',
      }, {
        name: '计算任务模板',
        name_en: 'Computing Task Template',
        icon: 'aitasktmpl',
        url: '/ai_task_tmpl/list',
      }]
    }],
    computingpower: [{
      name: '普惠算力',
      name_en: 'Computing Power',
      children: [{
        name: '算力资源',
        name_en: 'Computing resources',
        icon: '',
        url: '/computingpower/demand',
      }, {
        name: '国产算力',
        name_en: 'Domestic computing power',
        icon: '',
        url: '/computingpower/domestic',
      }, {
        name: '算力中心',
        name_en: 'Computing power center',
        icon: '',
        url: '/explore/c2net_map',
      }]
    }]
  }
  
  const lang = document.querySelector('html').getAttribute('lang');
  const loginEl = document.querySelector('meta[name="_uid"]')
  const userName = loginEl && loginEl.getAttribute("content-ext");
  const { pathname } = window.location;

  
  // 改进的parseRoute函数
  function parseRoute(path) {
    // 匹配 /org/:orgname/xxx 格式
    const orgMatch = path.match(/^\/org\/([^\/]+)(\/.*)?$/);
    if (orgMatch) {
      return {
        mode: 'org',
        orgName: orgMatch[1],
        fullPath: path,
        restPath: orgMatch[2] || '/',
        avatar_url: `/user/avatar/${orgMatch[1]}/-1?`
      };
    }
    return {
      mode: 'personal',
      orgName: userName,
      fullPath: path,
      restPath: path,
      avatar_url: `/user/avatar/${userName}/-1?`
    };
  }
  const routeInfo = parseRoute(pathname);
  console.log("routeInfo",routeInfo)
  
  
  // 3. 根据当前模式动态生成最终的菜单数据对象
  const menuData  = {
    workspace: menuDataInit.getWorkspaceMenu( routeInfo.mode === 'org', routeInfo.orgName),
    modelbase: menuDataInit.modelbase,
    resources: menuDataInit.resources,
    computingpower: menuDataInit.computingpower,
    agent: menuDataInit.agent
  };
  let matched = false;
  let mainMenuKey = '';
  let subMenu = null;
  let activeSubMenuUrl = ''

  const run = (list, _mainMenuKey) => {
    for (let i = 0, iLen = list.length; i < iLen; i++) {
      const listI = list[i];
      // 改进的匹配逻辑：支持精确匹配和子路径匹配
      if (listI.url) {
        // 检查是否是精确匹配
        if (listI.url === pathname) {
          mainMenuKey = _mainMenuKey;
          subMenu = listI;
          activeSubMenuUrl = listI.url;
          matched = true;
          break;
        }
        
        // 检查是否是子路径
        // 例如：/dashboard 匹配 /dashboard/123/xxx
        if (pathname.startsWith(listI.url + '/')) {
          mainMenuKey = _mainMenuKey;
          subMenu = listI;
          activeSubMenuUrl = listI.url;
          matched = true;
          break;
        }
        
        // 特殊处理：对于根路径，也检查是否包含（比如 /cloudbrains/create 应该匹配 /cloudbrains）
        // 但避免错误匹配，如 /cloudbrain 不应该匹配 /cloudbrains
        if (listI.url === '/' && pathname !== '/') {
          continue; // 根路径特殊处理
        }
      }
      
      if (listI.children && listI.children.length) {
        run(listI.children, _mainMenuKey);
      }
    }
  };

  for (let mainMenuKey in menuData) {
    const subMenuDataList = menuData[mainMenuKey];
    if (!matched) {
      run(subMenuDataList, mainMenuKey);
    }
  }
  // 从meta标签获取用户信息
  function getUserInfoFromMeta() {
    const uidMeta = document.querySelector('meta[name="_uid"]');
    if (!uidMeta) return null;

    const userId = uidMeta.getAttribute("content");
    const userName = uidMeta.getAttribute("content-ext");

    return {
      id: userId,
      name: userName,
      avatar_url: `/user/avatar/${userName}/-1?`,
    };
  }
  // 6. 将配置暴露给全局
  window.MENU_CONFIG = {
    lang: lang == 'zh-CN' ? 'CN' : 'US',
    menuData: menuData, // 使用动态数据
    showLeftMenu: matched && mainMenuKey != 'modelbase',
    collapseLeftMenu: false,
    activeTopMenu: mainMenuKey,
    activeLeftMenu: subMenu,
    activeSubMenuUrl,
    // 暴露当前工作台模式和组织信息，供页面其他部分（如下拉列表）使用
    workspaceMode: routeInfo.mode,
    routeInfo: routeInfo,
    userInfo: getUserInfoFromMeta()
  };
  console.log(window.MENU_CONFIG)
})();
