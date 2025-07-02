<template>
  <div class="home-container layout-pd" style="margin-top: 10px">
    <el-form :inline="true">
      <el-form-item label="通讯方式">
        <el-select v-model="state.category" @change="chooseCategory" placeholder="请选择通讯方式" clearable style="width: 120px">
          <el-option v-for="(item,index) in state.categoryList" :key="index" :label="item" :value="item">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="跳转到日期">
        <el-date-picker
            v-model="state.selectedDate"
            type="date"
            placeholder="选择日期"
            :clearable="false"
            :disabled-date="disabledDate"
            @change="selectDate"
        />
      </el-form-item>
    </el-form>
    <div style="margin-top: 10px">
      <el-table :data="state.dataList" border style="width: 100%">
        <el-table-column prop="id" label="Id" width="60"/>
        <el-table-column prop="category" label="通讯方式" width="160"/>
        <el-table-column prop="subject" label="主题" show-overflow-tooltip/>
        <el-table-column prop="status" label="状态" width="120">
          <template #default="scope">
            <el-tag :type="getMessageStatusType(scope.row.status)" style="width: 85px;">
              {{ getMessageStatus(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="时间">
          <template #default="scope">
            {{ formatDate(new Date(scope.row.createTime), 'YYYY-mm-dd HH:MM:SS') }}
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="180">
          <template #default="scope">
            <el-button size="small" type="text" @click="showDetail(scope.row)">
              查看详情
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin: 10px 10px 10px 10px;">
        <el-pagination background
                       layout="sizes, prev, pager, next, jumper"
                       v-model:currentPage="state.queryValue.pageNum"
                       v-model:page-size="state.queryValue.pageSize"
                       :page-sizes="[5,10,15,30,50,100]"
                       :page-count="state.queryValue.totalNum"
                       @size-change="onSizeChange"
                       @current-change="onCurrentChange"
        />
      </div>
    </div>
    <el-dialog
        v-model="state.showDetail"
        title="消息详情"
        width="70%"
        center>
      <el-form :model="state.detailForm" label-width="100px">
        <el-form-item label="模板" prop="template_name">
          <el-input v-model="state.detailForm.template_name" readonly></el-input>
        </el-form-item>
        <el-form-item label="收信人" prop="receivers">
          <el-input v-model="state.detailForm.receivers" readonly></el-input>
        </el-form-item>
        <el-form-item label="通讯方式" prop="category">
          <el-input v-model="state.detailForm.category" readonly></el-input>
        </el-form-item>
        <el-form-item label="通讯平台" prop="provider">
          <el-input v-model="state.detailForm.provider" readonly></el-input>
        </el-form-item>
        <el-form-item label="模板参数" prop="params">
          <el-input v-model="state.detailForm.params" readonly></el-input>
        </el-form-item>
        <el-form-item label="内容" prop="content" v-if="state.detailForm.content">
          <el-input
              v-model="state.detailForm.content"
              type="textarea"
              :autosize="{ minRows: 2, maxRows: 6 }"
              readonly
          ></el-input>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">

import {onMounted, reactive} from "vue";
import {formatDate} from "/@/utils/formatTime";
import {getHistoryCountApi, getHistoryListApi, getHistoryOneApi} from "/@/api/history";

const state = reactive({
  category: '',
  categoryList: [
    "mail",
    "sms",
    "im"
  ],
  dataList: [],
  queryValue: {
    pageNum: 1,
    pageSize: 10,
    totalNum: 0,
    category: ''
  },
  selectedDate: new Date(),
  showDetail: false,
  detailForm: {
    template_name: '',
    receivers: '',
    category: '',
    provider: '',
    params: '',
    content: '',
  }
})

const messageStatus = ["Expired", "Init", "Sending", "Success", "Failed"]
const messageStatusType = ["info", "info", "primary", "success", "danger"]
const getMessageStatus = (status: number) => {
  if (status < 0 || status > 4) {
    return messageStatus[0];
  }
  return messageStatus[status];
}
const getMessageStatusType = (status: number) => {
  if (status < 0 || status > 4) {
    return messageStatusType[0];
  }
  return messageStatusType[status];
}

const getHistoryList = () => {
  state.queryValue.category = state.category;
  getHistoryListApi(state.queryValue).then(res => {
    if (res && res.code == 200 && res.items) {
      state.dataList = res.items
      state.queryValue.totalNum = res.totalPage
    } else {
      state.dataList = []
      state.queryValue.totalNum = 0
    }
    // console.log(state.queryValue.totalNum)
  })
}

const onSizeChange = (val: any) => {
  state.queryValue.pageSize = val
  getHistoryList();
}
const onCurrentChange = (val: any) => {
  state.queryValue.pageNum = val
  getHistoryList();
}
const chooseCategory = () => {
  state.queryValue.pageNum = 1
  getHistoryList();
}

const showDetail = (row: any) => {
  getHistoryOneApi({
    message_id: row.id,
    category: row.category,
  }).then(res => {
    state.detailForm = res.item;
    state.showDetail = true;
  })
}

const disabledDate = (time: Date) => {
  return time.getTime() > Date.now()
}

const selectDate = () => {
  getHistoryCountApi({
    category: state.category,
    time: state.selectedDate.toISOString(),
  }).then(res => {
    if (res && res.code == 200) {
      state.queryValue.pageNum = 1 + Math.floor(res.count / state.queryValue.pageSize)
      getHistoryList();
    }
  })
}

onMounted(() => {
  getHistoryList();
})

</script>

<style scoped lang="scss">

</style>