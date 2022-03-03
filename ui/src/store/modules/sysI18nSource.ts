import { Module } from 'vuex';
import { SysI18nSourceState, RootStateTypes } from '/@/store/interface/index';

const sysI18nSourceStateModule: Module<SysI18nSourceState, RootStateTypes> = {
	namespaced: true,
	state: {
		sysI18nSource: {},
	},
	mutations: {
		// 后端i18n资源文件
		getSysI18nSource(state: any, data: object) {
			state.sysI18nSource = data;
		},
	},
	actions: {
		// 后端i18n资源文件
		setSysI18nSource({ commit }, data: object) {
			commit('getSysI18nSource', data);
		},
	},
};

export default sysI18nSourceStateModule;