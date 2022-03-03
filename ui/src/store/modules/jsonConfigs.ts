import { Module } from 'vuex';
import { JsonConfigsState, RootStateTypes } from '/@/store/interface/index';

const jsonConfigsStateModule: Module<JsonConfigsState, RootStateTypes> = {
	namespaced: true,
	state: {
		jsonConfigs: {},
	},
	mutations: {
		// 后端i18n资源文件
		getJsonConfigs(state: any, data: object) {
			state.jsonConfigs = data;
		},
	},
	actions: {
		// 后端i18n资源文件
		setJsonConfigs({ commit }, data: object) {
			commit('getJsonConfigs', data);
		},
	},
};

export default jsonConfigsStateModule;