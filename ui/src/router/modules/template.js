/** When your routing table is too long, you can split it into small modules**/

import Layout from '@/layout'

const templateRouter = {
  path: '/template',
  component: Layout,
  redirect: '/template/list',
  name: 'Template',
  meta: {
    title: 'Template',
    icon: 'el-icon-s-help'
  },
  children: [
    {
      path: 'create',
      component: () => import('@/views/template/create'),
      name: 'CreateTemplate',
      meta: { title: 'Create Template', icon: 'edit' }
    },
    {
      path: 'edit/:id(\\d+)',
      component: () => import('@/views/template/edit'),
      name: 'EditTemplate',
      meta: { title: 'Edit Template', noCache: true, activeMenu: '/template/list' },
      hidden: true
    },
    {
      path: 'list',
      component: () => import('@/views/template/list'),
      name: 'TemplateList',
      meta: { title: 'Template List', icon: 'list' }
    }
  ]
}

export default templateRouter
