# Tencent ses SDK

> 使用腾讯云ses API推送邮件

## 资源准备

### 账号

| 资源       | 名称           |
|----------|--------------|
| secretId | 应用ID         |
| secret   | 秘钥           |
| region   | 国内选香港,国外选新加坡 |
| sender   | 发件账号         |

### 业务资源

| 资源         | 名称        |
|------------|-----------|
| templateId | 模板ID,不是名称 |
| params     | 参数列表,json |

## 用法(Usage)

### 配置

```shell
TencentConfig{
    AppId:  secretId,
    Secret: secret,
    Region: region,
    Sender: sender,
}
```

```shell
package xxx

import (
	"encoding/json"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/messager"
	"testing"
	"time"
)

// 配置
config := model.TencentConfig{
    AppId:  secretId,
    Secret: secret,
    Region: region,
    Sender: sender,
}
bs, err := json.Marshal(config)
if err != nil {
    return err
}
s, err := NewTencent(string(bs))
if err != nil {
    return err
}

param := map[string]interface{}{
    "key1": "value1",
    "key2": "value2",
}
bs, err = json.Marshal(param)
if err != nil {
    return err
}

err = s.Send(messager.MailRequest{
    Subject:   "subject",
    Receivers: []string{"mail_address"},
    Template:  templateId,
    Params:    param,
})

if err != nil {
    return err
}
```

### 注意

Tencent对region支持有特殊要求:
> 国内是香港，国际是新加坡地域
>
> 发送国内邮件需要指定region:ap-hongkong

