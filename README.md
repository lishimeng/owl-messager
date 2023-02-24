owl messager
===============================================
[![Go project version](https://badge.fury.io/go/github.com%2Flishimeng%2Fowl-messager.svg)](https://badge.fury.io/go/github.com%2Flishimeng%2Fowl-messager)
[![issues](https://img.shields.io/github/issues/lishimeng/owl-messager)](https://github.com/lishimeng/owl-messager)
[![Go Report Card](https://goreportcard.com/badge/github.com/lishimeng%2Fowl-messager?style=flat-square)](https://goreportcard.com/report/github.com/lishimeng%2Fowl-messager)
[![License](https://img.shields.io/github/license/lishimeng/owl-messager)](https://github.com/lishimeng/owl-messager)

Message notify platform.

Great for:

* Mail
* SMS
* APNS
* ARM
* dingding
* WeChat message(template message)
* WeChat app(MA)

Support on Raspberry Pi

What is this?
---

owl messager is a platform:

1. Support email.
1. Support SMS.
1. Support APNS.
1. Support WeChat.
1. Support DingDing.
1. Support Multi-tenant.
1. Support Account list(mail sender).

Usage
--------------
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
