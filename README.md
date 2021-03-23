owl messager
===============================================

Message notify platform.

Great for:

* Mail
* SMS
* APNS
* ARM
* Raspberry Pi

What is this?
---

owl messager is a platform:

1. Support email.
1. Support SMS.
1. Support APNS.
1. Support Multi-tenant.
1. Support Account list(mail sender).

Build & Install
--------------

To build owl as a application just run:

```bash
go build
```

1. A config file is written to `/etc/owl-messager/config.yaml` and the service is automatically started or restarted.
1. Run server:

```bash
sudo owl-messager
```

Contributing
------------

Please check out our [contributing guide](CONTRIBUTING.md) if you're interested in contributing to K3s.
