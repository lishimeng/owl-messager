<template>
  <div class="icons-container">
    <aside>
      <a href="https://panjiachen.github.io/vue-element-admin-site/guide/advanced/icon.html" target="_blank">Add and use
      </a>
    </aside>
    <el-tabs type="border-card">
      <el-tab-pane label="Icons">
        <div class="grid">
          <div v-for="item of state.svgIcons" :key="item" @click="handleClipboard(generateIconCode(item),$event)">
            <el-tooltip placement="top">
              <div slots="content">
                {{ generateIconCode(item) }}
              </div>
              <div class="icon-item">
                <svg-icon :icon-class="item" class-name="disabled" />
                <span>{{ item }}</span>
              </div>
            </el-tooltip>
          </div>
        </div>
      </el-tab-pane>
      <el-tab-pane label="Element-UI Icons">
        <div class="grid">
          <div v-for="item of state.elementIcons" :key="item" @click="handleClipboard(generateElementIconCode(item),$event)">
            <el-tooltip placement="top">
              <div slots="content">
                {{ generateElementIconCode(item) }}
              </div>
              <div class="icon-item">
                <i :class="'el-icon-' + item" />
                <span>{{ item }}</span>
              </div>
            </el-tooltip>
          </div>
        </div>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup>
import { reactive, onMounted, getCurrentInstance } from "vue"
import clipboard from '/@/utils/clipboard.ts'
import svgIcons from './svg-icons'
import elementIcons from './element-icons'
const state = reactive({
  svgIcons,
  elementIcons
})

function generateIconCode(symbol) {
  return `<svg-icon icon-class="${symbol}" />`
}
function generateElementIconCode(symbol) {
  return `<i class="el-icon-${symbol}" />`
}
function handleClipboard(text, event) {
  clipboard(text, event)
}
</script>

<style lang="scss" scoped>
.icons-container {
  margin: 10px 20px 0;
  overflow: hidden;

  .grid {
    position: relative;
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
  }

  .icon-item {
    margin: 20px;
    height: 85px;
    text-align: center;
    width: 100px;
    float: left;
    font-size: 30px;
    color: #24292e;
    cursor: pointer;
  }

  span {
    display: block;
    font-size: 16px;
    margin-top: 10px;
  }

  .disabled {
    pointer-events: none;
  }
}
</style>
