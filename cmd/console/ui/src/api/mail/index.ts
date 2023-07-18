import {get, post} from '/src/utils/request';

const baseUrl = "/api"
export const createMailSenderConfigApi = (p: object) => post(baseUrl + "/sender/mail/set_default", p);
export const updateMailSenderConfigApi = (p: object) => post(baseUrl + "/sender/mail/up_default", p);
export const getMailSenderInfoApi = (p: object) => get(baseUrl + "/sender/mail/vendor", p);
export const getMailSendersApi = (p: object) => get(baseUrl + "/sender/mail/list/page", p);
export const getSenderInfoByCategoryAPi = (p: object) => get(baseUrl + "/sender/mail/info/category", p);