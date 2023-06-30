import {get, post} from '/src/utils/request';

const baseUrl = "/api"
export const createTemplateApi = (p: object) => post(baseUrl + "/template/createTemplate", p);
export const updateTemplateApi = (p: object) => post(baseUrl + "/template/updateTemplate", p);
export const getTemplateListAPi = (p: object) => get(baseUrl + "/template/getTemplateList",p);
export const getTemplateInfoAPi = (p: object) => get(baseUrl + "/template/getTemplateInfo",p);