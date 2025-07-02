import {get} from '/src/utils/request';

const baseUrl = "/api"

export const getThemeConfigApi = (p: object) => get(baseUrl + "/theme", p);
