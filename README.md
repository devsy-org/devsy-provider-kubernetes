# Kubernetes Provider for Devsy

[![Open in Devsy!](https://img.shields.io/badge/open_in_devsy-8A2BE2?style=for-the-badge)](https://devsy.sh/open#https://github.com/devsy-org/devsy-provider-kubernetes)

## Getting started

The provider is available for auto-installation using:

```sh
devsy provider add kubernetes
devsy provider use kubernetes
```

Follow the on-screen instructions to complete the setup.

### Creating your first workspace with kubernetes

After the initial setup, just use:

```sh
devsy up .
```

You'll need to wait for the pod and workspace setup.

### Customize the Provider

This provider has the following options:

| NAME                                  | REQUIRED | DESCRIPTION                                                                                     | DEFAULT   |
|---------------------------------------|----------|-------------------------------------------------------------------------------------------------|-----------|
| KUBERNETES_NAMESPACE                  | false    | The kubernetes namespace to use.                                                                | devsy     |
| DISK_SIZE                             | false    | The default size for the persistent volume to use.                                              | 10Gi      |
| KUBERNETES_CONTEXT                    | false    | The kubernetes context to use. E.g. my-kube-context                                             |           |
| KUBERNETES_CONFIG                     | false    | The kubernetes config to use. E.g. /path/to/my/kube/config.yaml                                 |           |
| CREATE_NAMESPACE                      | false    | If true, Devsy will try to create the namespace.                                                | true      |
| CLUSTER_ROLE                          | false    | If defined, Devsy will create a role binding for the given cluster role.                         |           |
| SERVICE_ACCOUNT                       | false    | If defined, Devsy will use the given service account for the dev container.                      |           |
| KUBECTL_PATH                          | false    | The path where to find the kubectl binary.                                                      | kubectl   |
| INACTIVITY_TIMEOUT                    | false    | If defined, will automatically stop the pod after the inactivity period. E.g. 10m, 1h           |           |
| STORAGE_CLASS                         | false    | If defined, Devsy will use the given storage class to create the persistent volume claim.        |           |
| NODE_SELECTOR                         | false    | The node selector to use for the workspace pod. E.g. my-label=value                             |           |
| RESOURCES                             | false    | The resources to use for the workspace container. E.g. requests.cpu=500m,limits.memory=5Gi      |           |
| POD_MANIFEST_TEMPLATE                 | false    | Pod manifest template file path used as template to build the devsy pod.                        |           |
| DOCKERLESS_DISABLED                   | false    | If dockerless should be disabled.                                                               | false     |
