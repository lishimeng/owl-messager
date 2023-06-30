<template>
	<div class="layout-footer pb15">
		<div class="layout-footer-warp">
			<div>{{ state.description }}</div>
			<div class="mt5">{{ state.copyRight }}</div>
		</div>
	</div>
</template>

<script setup lang="ts" name="layoutFooter">
// 此处需有内容（注释也得），否则缓存将失败
// 定义变量内容
import {reactive} from "vue";

const state = reactive({
	description: 'Owl Messager, message push provider',
	copyRight: 'githug.com/lishimeng版权所有'
})
</script>

<style scoped lang="scss">
.layout-footer {
	width: 100%;
	display: flex;
	&-warp {
		margin: auto;
		color: var(--el-text-color-secondary);
		text-align: center;
		animation: error-num 0.3s ease;
	}
}
</style>
