## Owl Messager

Message notify platform.

Great for:

* Mail
* SMS
* APNS
* ARM
* Ding Ding Bot
* WeChat message Bot(template message)
* WeChat app(MA)

Support on Raspberry Pi

## What is this?


owl messager is a platform:

1. Support email.
2. Support SMS.
3. Support APNS.
4. Support WeChat.
5. Support DingDing.
6. Support Multi-tenant.
7. Support Account list(mail sender).

## Usage

send an email
```shell
http://localhost/api/v2/messages/mail/
```
Parameter
```json
{
"template":"{tpl_id}",
"params": {
"code":"{code}"
},
"subject": "MFA email code",
"receiver":"{mail_address}"
}
```

send a sms
```shell
http://localhost/api/v2/messages/sms/
```
Parameter
```json
{
"template":"{tpl_id}",
"params": {
"code":"652442"
},
"receiver":"{phone_number}"
}
```

Build & Install
--------------

To build owl as an application just run:

```bash
go build
```

1. A config file is written to `/etc/owl-messager/config.yaml` and the service is automatically started or restarted.
2. Run server:

```bash
sudo owl-messager
```

Contributing
------------

Please check out our [contributing guide](CONTRIBUTING.md) if you're interested in contributing to owl.
