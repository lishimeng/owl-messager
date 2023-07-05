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

腾讯云的region

| 地域       | 取值           |
|----------|--------------|
| 华北地区(北京) | ap-beijing   |
| 华南地区(广州) | ap-guangzhou |
| 华东地区(南京) | ap-nanjing   |

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