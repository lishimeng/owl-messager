## 发送SMS

### 需要设置发信的账号(Sms Sender)
每个平台的账号有专用参数

### Ali
```json
{
  "appKey": "在阿里云平台的应用ID",
  "appSecret": "秘钥",
  "region": "区域ID",
  "signName": "签名"
}
```

### Tencent
腾讯云的region参数会被忽略. 原因是腾讯云暂时只支持广州通道
```json
{
  "appId": "在腾讯云平台的应用ID",
  "appKey": "秘钥",
  "region": "区域ID",
  "signName": "签名"
}
```

### Huawei
```json

```