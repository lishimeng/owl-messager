# Message Sender约束

一般原则：

同一组织，同一发送平台下保留一个发送账号生效，采用优先级原则。

ORG -> APP_ID -> CATEGORY -> SENDER
> 组织1, App1,  短信,  默认1

1. 如果对接多个平台，按照账号优先级，总是有一个生效，依照具体算法排序
2. API在调用时不再提供sender_id。 template id依旧需要明确在参数中。