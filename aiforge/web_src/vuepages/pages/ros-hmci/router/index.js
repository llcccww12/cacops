import Vue from 'vue'
import Router from 'vue-router'

import HomePage from '../views/HomePage.vue'
import OSSystem from '../views/OSSystem.vue'
import OpenApp from "../views/OpenApp.vue";
import OpenDataset from '../views/OpenDataset.vue'
import OpenModel from '../views/OpenModel.vue'
import CommunitySource from '../views/CommunitySource.vue'
import HelpCenter from '../views/HelpCenter.vue'
import ResourceDetail from '../views/ResourceDetail.vue'

Vue.use(Router)

export default new Router({
    base: '/ros-hmci',
    routes: [

        {
            path: '/',
            name: 'HomePage',
            component: HomePage
        },
        {
            path: '/os-system',
            name: 'OSSystem',
            component: OSSystem
        },
        {
            path: '/open-app',
            name: 'OpenApp',
            component: OpenApp
        },
        {
            path: '/open-data',
            name: 'OpenDataset',
            component: OpenDataset
        },
        {
            path: '/open-model',
            name: 'OpenModel',
            component: OpenModel
        },
        {
            path: '/community-source',
            name: 'CommunitySource',
            component: CommunitySource
        },
        {
            path: '/help-center',
            name: 'HelpCenter',
            component: HelpCenter
        }, {
            path: '/source-detail/:name',
            name: 'ResourceDetail',
            component: ResourceDetail
        }
    ]
})