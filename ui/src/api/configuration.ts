import { get, post, put, del } from '/@/utils/request';

const baseUrl = "/configuration"
const locale = "zh_cn"

// 获取I18n资源
export const getI18nSourceApi = (p: object) => get(baseUrl + "/product/" + p.productId + "/i18n/"+ locale, p);