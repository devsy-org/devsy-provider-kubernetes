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
| ARCH_DETECTION_POD_MANIFEST_TEMPLATE | false | Pod manifest template file path used as template to build the architecture detection pod. | |
| KUBERNETES_PULL_SECRETS_ENABLED | false | If true, Devsy will try to use the pull secrets from the current context. | true |
| HELPER_IMAGE | false | The image Devsy will use to find out the cluster architecture. | alpine |
| HELPER_RESOURCES | false | The resources to use for the workspace init container. E.g. requests.cpu=100m,limits.memory=1Gi | |
| POD_TIMEOUT | false | How long the provider waits for the workspace pod to come up. Examples: 10m, 1h | 10m |
| PVC_ANNOTATIONS | false | If defined, Devsy will add the given annotations to the main workspace pvc. | |
| LABELS | false | The labels to use for the workspace pod. E.g. devsy.sh/example=value,devsy.sh/example2=value2 | |
| WORKSPACE_VOLUME_MOUNT | false | Sets the path of the workspace volume mount. | |
| DOCKERLESS_IMAGE | false | The dockerless image to use. | |
| DOCKERLESS_DISABLED | false | If dockerless should be disabled. | |
| STRICT_SECURITY | false | EXPERIMENTAL! Removes the default security context so you can define your own via POD_MANIFEST_TEMPLATE. | false |
| DANGEROUSLY_OVERRIDE_IMAGE | false | Only set this if you know what you're doing! Overrides the dev container image. | |
