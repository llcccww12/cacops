import Vue from 'vue'
import Router from 'vue-router'
import IdeProject from '../components/IdeProject.vue'

const originalPush = Router.prototype.push

Router.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}
Vue.use(Router)

export default new Router({
  mode: 'history',
  base: '/',    //添加根目录
  scrollBehavior: () => ({ y: 0 }),
  routes: [
    {
      path: '/ide/project',
      name: 'ide',
      component: IdeProject,
    }
  ],
})
