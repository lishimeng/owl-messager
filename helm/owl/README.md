Owl Messager Helm charts
===============================================

This chart installs [Owl](https://github.com/lishimeng/owl-messager), an open source message provider.


```bash
$ helm repo add owl https://lishimeng.github.io/charts/
$ helm install my-release owl/owl
```

## Installing the Chart

To install the chart with the release name `my-release`:

```bash
$ helm install my-release owl/owl
```

The command deploys Owl on the Kubernetes cluster in the default configuration. The [Parameters](#parameters) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `my-release` deployment:

```bash
$ helm delete my-release
```

The command removes all the Kubernetes components associated with the chart and deletes the release.
