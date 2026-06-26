# Kubernetes Provider for Devsy

## Getting started

The provider is available for auto-installation using:

```sh
devsy provider add kubernetes
devsy provider use kubernetes
```

Follow the on-screen instructions to complete the setup.

Options can be set using `devsy provider set-options`, for example:

```sh
devsy provider set-options kubernetes -o DISK_SIZE=20Gi -o KUBERNETES_NAMESPACE=my-namespace
```

### Creating your first workspace with kubernetes

After the initial setup, just use:

```sh
devsy up .
```

You'll need to wait for the pod and workspace setup.

### Customize the Provider

This provider has the following options:

| NAME | REQUIRED | DESCRIPTION | DEFAULT |
|------|----------|-------------|---------|
| KUBERNETES_NAMESPACE | false | The kubernetes namespace to use | devsy |
| KUBERNETES_CONTEXT | false | The kubernetes context to use. E.g. my-kube-context | |
| KUBERNETES_CONFIG | false | The kubernetes config to use. E.g. /path/to/my/kube/config.yaml | |
| DISK_SIZE | false | The default size for the persistent volume to use. | 10Gi |
| CREATE_NAMESPACE | false | If true, Devsy will try to create the namespace. | true |
| CLUSTER_ROLE | false | If defined, Devsy will create a role binding for the given cluster role. | |
| SERVICE_ACCOUNT | false | If defined, Devsy will use the given service account for the dev container. | |
| KUBECTL_PATH | false | The path where to find the kubectl binary. | kubectl |
| INACTIVITY_TIMEOUT | false | If defined, will automatically stop the pod after the inactivity period. Examples: 10m, 1h | |
| STORAGE_CLASS | false | If defined, Devsy will use the given storage class to create the persistent volume claim. | |
| PVC_ACCESS_MODE | false | If defined, Devsy will use the given access mode to create the persistent volume claim. E.g. RWO or ROX or RWX or RWOP | |
| NODE_SELECTOR | false | The node selector to use for the workspace pod. E.g. my-label=value,my-label-2=value-2 | |
| RESOURCES | false | The resources to use for the workspace container. E.g. requests.cpu=500m,limits.memory=5Gi | |
| POD_MANIFEST_TEMPLATE | false | Pod manifest template file path or inline yaml used as template to build the devsy pod. | |
