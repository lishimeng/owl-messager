import {formatDate} from '/@/utils/formatTime'

/**
 * 生成随机订单号
 * @param type 订单开头
 * @returns 订单号字符串
 */
export function genRandomOrderNo(type: string) {
  let orderNo : string = type + formatDate(new Date(), "YYYYmmddHHMMSS")
  return orderNo;
}
