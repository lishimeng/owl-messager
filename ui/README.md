
<p align="center">
  <a href="https://github.com/vuejs/vue">
    <img src="https://img.shields.io/badge/vue-2.6.10-brightgreen.svg" alt="vue">
  </a>
  <a href="https://github.com/ElemeFE/element">
    <img src="https://img.shields.io/badge/element--ui-2.7.0-brightgreen.svg" alt="element-ui">
  </a>
</p>

## 简介

本地安装 [node](http://nodejs.org/) 和 [git](https://git-scm.com/)。

技术栈基于 [ES2015+](http://es6.ruanyifeng.com/)、[vue](https://cn.vuejs.org/index.html)、[vuex](https://vuex.vuejs.org/zh-cn/)、[vue-router](https://router.vuejs.org/zh-cn/) 、[vue-cli](https://github.com/vuejs/vue-cli) 、[axios](https://github.com/axios/axios) 和 [element-ui](https://github.com/ElemeFE/element)，所有的请求数据都使用[Mock.js](https://github.com/nuysoft/Mock)进行模拟。

[![Edit on CodeSandbox](https://codesandbox.io/static/img/play-codesandbox.svg)](https://codesandbox.io/s/github/PanJiaChen/vue-element-admin/tree/CodeSandbox)


## 开发

```bash

# 安装依赖
yarn install

# 启动服务
yarn dev
```

浏览器访问 http://localhost:9527

## 发布

```bash
# 构建测试环境
yarn build:stage

# 构建生产环境
yarn build:prod
```

## 其它

```bash
# 预览发布环境效果
yarn preview

# 预览发布环境效果 + 静态资源分析
yarn preview --report

# 代码格式检查
yarn lint

# 代码格式检查并自动修复
yarn lint --fix
```