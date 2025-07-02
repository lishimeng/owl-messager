import {get} from '/src/utils/request';

const baseUrl = "/api"

export const getHistoryListApi = (p: object) => get(baseUrl + "/history", p);
export const getHistoryOneApi = (p: object) => get(baseUrl + "/history/one", p);
export const getHistoryCountApi = (p: object) => get(baseUrl + "/history/count", p);
