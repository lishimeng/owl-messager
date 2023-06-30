<template>
  <div class="layout-padding-auto">
      <TableSearch :search="state.tableData.search" @search="onSearch" />
      <Table
              ref="tableRef"
              v-bind="state.tableData"
              class="table-demo"
      >

      </Table>
  </div>
  <div class="layout-padding-auto">
      <JsEditor code="" @on-change="onCodeChange" />
  </div>
</template>

<script setup lang="ts" name="smsSender">
import {defineAsyncComponent, onMounted, reactive, ref} from "vue"
const Table = defineAsyncComponent(() => import('/@/components/table/index.vue'))
const TableSearch = defineAsyncComponent(() => import('/@/views/make/tableDemo/search.vue'))
const JsEditor = defineAsyncComponent(() => import('/@/components/js/index.vue'))

const tableRef = ref<RefType>()
const state = reactive<TableDemoState>({
    tableData: {
        // 列表数据（必传）
        data: [],
        // 表头内容（必传，注意格式）
        header: [
            { key: 'name', colWidth: '', title: '名称', type: 'text', isCheck: false },
            { key: 'vendor', colWidth: '', title: '平台', type: 'text', isCheck: true },
        ],
        // 配置项（必传）
        config: {
            total: 0, // 列表总数
            loading: true, // loading 加载
            isBorder: false, // 是否显示表格边框
            isSerialNo: true, // 是否显示表格序号
            isSelection: true, // 是否显示表格多选
            isOperate: true, // 是否显示表格操作栏
        },
        // 搜索表单，动态生成（传空数组时，将不显示搜索，注意格式）
        search: [
            { label: '名称', prop: 'name', placeholder: '名称', required: true, type: 'input' },
            {
                label: '平台',
                prop: 'vendor',
                placeholder: '请选择',
                required: false,
                type: 'select',
                options: [
                    { label: '阿里', value: "ali" },
                    { label: '腾讯', value: "tencent" },
                ],
            },
        ],
        // 搜索参数（不用传，用于分页、搜索时传给后台的值，`getTableData` 中使用）
        param: {
            pageNum: 1,
            pageSize: 10,
        },
    },
})

// 初始化列表数据
const getTableData = () => {
    state.tableData.config.loading = true
    state.tableData.data = []
    state.tableData.data.push({
        id: `1234567891`,
        name: `高大上发顺丰`,
        vendor: `阿里`,
    })
    state.tableData.data.push({
        id: `1234567892`,
        name: `都给我发`,
        vendor: `腾讯`,
    })
    // 数据总数（模拟，真实从接口取）
    state.tableData.config.total = state.tableData.data.length
    setTimeout(() => {
        state.tableData.config.loading = false
    }, 500)
}

const onSearch = () => {

}

const onCodeChange = (content: String) => {
    console.log("editor:" + content)
}

// 页面加载时
onMounted(() => {
    getTableData();
})

</script>

<style scoped>

</style>