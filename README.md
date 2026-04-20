# Kubernetes Provider for Devsy

[![Join us on Slack!](docs/static/media/slack.svg)](https://slack.devsy.sh/) [![Open in Devsy!](https://devsy.sh/assets/open-in-devsy.svg)](https://devsy.sh/open#https://github.com/devsy-org/devsy-provider-kubernetes)

## Getting started

The provider is available for auto-installation using

```sh
devsy provider add kubernetes
devsy provider use kubernetes
```

Follow the on-screen instructions to complete the setup.

### Creating your first devsy env with kubernetes

After the initial setup, just use:

```sh
devsy up .
```

You'll need to wait for the pod and environment setup.


## Testing locally
1. Build the new version in a dev mode with some version tag (e.g. 0.0.1-dev)
```sh
chmod +x ./hack/build.sh
RELEASE_VERSION=0.0.1-dev ./hack/build.sh --dev
```
2. Remove the old provider from your devsy installation (make sure you delete all workspaces using the provider).
```sh
devsy provider delete kubernetes
```
3. Install the new provider from the local build
```sh
devsy provider add --name kubernetes --use ./release/provider.yaml
```
4. Test your provider, e.g. with `devsy up` command. Make sure you have a valid kubeconfig file in your home directory.
```sh
devsy up <repository-url> --provider kubernetes --debug
```
