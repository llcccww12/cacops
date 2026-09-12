import Vue from 'vue';
import Router from 'vue-router';

const originalPush = Router.prototype.push;
Router.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err);
}

Vue.use(Router);

import Layout from '../layout/Layout.vue';

const router = new Router({
  base: 'agent',
  mode: 'history',
  routes: [{
    path: '/',
    component: Layout,
    redirect: 'list',
    children: [{
      path: 'list',
      name: 'AgentList',
      meta: { menu: true, icon: 'ri-robot-line', label: 'agentPortal.agentSquare', ignoreAuth: true },
      component: () => import(/* webpackChunkName: "agent-portal-list" */ '../views/list/AgentList.vue'),
    }],
  }, {
    path: '*',
    redirect: 'list',
  }]
});

router.auth = {
  isInit: false,
  isLogin: !!document.querySelector('meta[name="_uid"]'),
}

router.beforeEach(async (to, from, next) => {
  if (to.meta.ignoreAuth) {
    next();
  } else {
    if (router.auth.isLogin) {
      next();
    } else {
      window.location.href = `/user/login?redirect_to=${encodeURIComponent(window.location.origin + '/' + router.options.base + to.fullPath)}`;
    }
  }
});

export default router;