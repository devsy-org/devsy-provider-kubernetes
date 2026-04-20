# Changelog

## [1.1.0](https://github.com/devsy-org/devsy-provider-kubernetes/compare/v1.0.0...v1.1.0) (2026-04-20)


### Features

* add dockerless options ([2242f0a](https://github.com/devsy-org/devsy-provider-kubernetes/commit/2242f0af8b2d0c8e3b524062651d58a4afe2c80c))
* add option to specify PVC annotations ([d60cc0e](https://github.com/devsy-org/devsy-provider-kubernetes/commit/d60cc0ed7972fb9c3011c04e6ad683ca10f69013))
* add POD_TIMEOUT option to provider; default to 10 minutes (old behaviour) ([4732a79](https://github.com/devsy-org/devsy-provider-kubernetes/commit/4732a79f2b86bdde44743e7e165dd46533ec2322))
* add POD_TIMEOUT option to provider; default to 10 minutes (old behaviour) ([f5ed8fc](https://github.com/devsy-org/devsy-provider-kubernetes/commit/f5ed8fc1acfa2f8cc71db4d67fc6cc034bb62ec7))
* add release-please for automated releases ([#3](https://github.com/devsy-org/devsy-provider-kubernetes/issues/3)) ([a2620e7](https://github.com/devsy-org/devsy-provider-kubernetes/commit/a2620e75da2e523dba325feca7cf03e460dfc67b))
* add run option UID label to pod ([906947e](https://github.com/devsy-org/devsy-provider-kubernetes/commit/906947e72a138eb8b73b5a97eab1de385d6b341a))
* add run option UID label to pod ([cfa7694](https://github.com/devsy-org/devsy-provider-kubernetes/commit/cfa76942c8f7818bafdadeac34a423aaac38b50c))
* add STRICT_SECURITY option to control the devpod container and init container security contexts ([26ba5cd](https://github.com/devsy-org/devsy-provider-kubernetes/commit/26ba5cd6b2056230019560ef2b7da5178ac4192a))
* add STRICT_SECURITY option to control the devpod container and init container security contexts ([89e5475](https://github.com/devsy-org/devsy-provider-kubernetes/commit/89e547578ad6aad1646b880417c15acec0dad78d))
* add WORKSPACE_VOLUME_MOUNT option. Allows configuration of the ([38d10a6](https://github.com/devsy-org/devsy-provider-kubernetes/commit/38d10a6068ed45887ec11976099f904d70cb398c))
* allow restart policy to be set by users ([3848817](https://github.com/devsy-org/devsy-provider-kubernetes/commit/3848817079c2bd853dfea73798fd71093427f149))
* allow users to customize the architecture detection pod with the ARCH_DETECTION_POD_MANIFEST_TEMPLATE option ([270063d](https://github.com/devsy-org/devsy-provider-kubernetes/commit/270063d7378e76df172c60eef095e7cb094cbbb5))
* allow users to override the workspace image; add warnings of potential impact ([a5994fc](https://github.com/devsy-org/devsy-provider-kubernetes/commit/a5994fcacb4f104e56560742b3ea718142e6b3c7))
* allow users to override the workspace image; add warnings of potential impact ([146fb0f](https://github.com/devsy-org/devsy-provider-kubernetes/commit/146fb0ff23e4679ed90d972e280e368877ab3f9a))
* implement driver log interface ([4365b0a](https://github.com/devsy-org/devsy-provider-kubernetes/commit/4365b0a6f1ab38189a5b8522e28a1ca9539b4c43))
* initial setup ([#1](https://github.com/devsy-org/devsy-provider-kubernetes/issues/1)) ([b8bf5ba](https://github.com/devsy-org/devsy-provider-kubernetes/commit/b8bf5ba03fa417b414a2068916a51dacbaeab072))
* inline pod manifest template ([039996e](https://github.com/devsy-org/devsy-provider-kubernetes/commit/039996e95167fa48392215234ee8147da92095b4))
* modernize tooling to match AWS/gcloud providers ([#5](https://github.com/devsy-org/devsy-provider-kubernetes/issues/5)) ([f7eced8](https://github.com/devsy-org/devsy-provider-kubernetes/commit/f7eced888e9bf933538c93b47baeaac8ab6bc727))
* reprovision pod when provider configuration changed ([ae41781](https://github.com/devsy-org/devsy-provider-kubernetes/commit/ae41781a5e56d16a5bd78fca7f20d402467ecb6d))
* use PodAffinity to ensure that Workspace pod is scheduled on the same node where Architecture detection was ([9686a40](https://github.com/devsy-org/devsy-provider-kubernetes/commit/9686a40e25ada7e070cef29f911acc933a3cef59))
* use PodAffinity to ensure that Workspace pod is scheduled on the same node where Architecture detection was ([d357ede](https://github.com/devsy-org/devsy-provider-kubernetes/commit/d357ede4e8c1d59d457e9f55a08f887aac6a9cdd))


### Bug Fixes

* add affinity to existing, when present ([6013261](https://github.com/devsy-org/devsy-provider-kubernetes/commit/6013261caa7666274ccf5b7dfb322cf27808626c))
* command problems ([fe6756c](https://github.com/devsy-org/devsy-provider-kubernetes/commit/fe6756c5dca81638e2510158ca07e91e6cc5a439))
* **deps:** update kubernetes monorepo to v0.35.4 ([#22](https://github.com/devsy-org/devsy-provider-kubernetes/issues/22)) ([42a6a1a](https://github.com/devsy-org/devsy-provider-kubernetes/commit/42a6a1a9006df577f950d3ad9963d6e32bcb611e))
* **deps:** update module github.com/docker/cli to v29.4.0+incompatible ([#24](https://github.com/devsy-org/devsy-provider-kubernetes/issues/24)) ([5671ad0](https://github.com/devsy-org/devsy-provider-kubernetes/commit/5671ad06e59032b2972e029f5652abb5840cb769))
* **deps:** update module github.com/skevetter/devpod to v0.22.1 ([#26](https://github.com/devsy-org/devsy-provider-kubernetes/issues/26)) ([f91afc7](https://github.com/devsy-org/devsy-provider-kubernetes/commit/f91afc7d14c289398c67383f1a9e8258317d8a39))
* do not use rm flag ([53ce6a7](https://github.com/devsy-org/devsy-provider-kubernetes/commit/53ce6a71755ac96900292e3a0a69a92a0ce32c22))
* escape provider binary path ([bc31097](https://github.com/devsy-org/devsy-provider-kubernetes/commit/bc310973057cd9c906d6bdda200fdaaaa15798eb))
* escape provider binary path ([b0397fa](https://github.com/devsy-org/devsy-provider-kubernetes/commit/b0397fa528d187c1826bf9b7223b9dcd1863134b))
* FindDevContainer tries to lookup the pod and returns first result ([6dfa6fe](https://github.com/devsy-org/devsy-provider-kubernetes/commit/6dfa6fee2e307a5e65d82fa3a6f13e9caa6de09b))
* import k8s client authentication package ([#2](https://github.com/devsy-org/devsy-provider-kubernetes/issues/2)) ([5fd56a0](https://github.com/devsy-org/devsy-provider-kubernetes/commit/5fd56a05e23b65252cd79958fbd1201f112db45b))
* increase waitPodRunning timeout 5 -&gt; 10 minutes ([51c0912](https://github.com/devsy-org/devsy-provider-kubernetes/commit/51c09123388008cc73f5f103bd001b2210e4c469))
* increase waitPodRunning timeout 5 -&gt; 10 minutes ([f1c0f55](https://github.com/devsy-org/devsy-provider-kubernetes/commit/f1c0f555f8760e65c3606981a256e57c9a9e686a))
* linting and formatting ([cb9da70](https://github.com/devsy-org/devsy-provider-kubernetes/commit/cb9da70074f03deec0afb59b5514086373de2431))
* merge init container definition of pod manifest template with our init container ([ae4c9d5](https://github.com/devsy-org/devsy-provider-kubernetes/commit/ae4c9d5e1b141f9e6713fa20c9fab13e87d6f03f))
* only apply architecture discovery based pod affinity if NODE_SELECTOR is not defined ([ba596ac](https://github.com/devsy-org/devsy-provider-kubernetes/commit/ba596ac9c2ba9a81de835402ae24d860ff89c66f))
* read arch detection pod template ([586af27](https://github.com/devsy-org/devsy-provider-kubernetes/commit/586af27d7c6f73120d8c79c547c7bca7b6f6c562))
* refactor runCommand to use cmdIO struct ([#9](https://github.com/devsy-org/devsy-provider-kubernetes/issues/9)) ([fbfd2c1](https://github.com/devsy-org/devsy-provider-kubernetes/commit/fbfd2c1f574481b4e2a6d29e00a9575aee1d0e7b))
* rename workspace label, make it detect properly ([7b05f2b](https://github.com/devsy-org/devsy-provider-kubernetes/commit/7b05f2b4a3c4bf1637c7de0fd80375c982b3ba13))
* resolve all golangci-lint violations ([#7](https://github.com/devsy-org/devsy-provider-kubernetes/issues/7)) ([3deee89](https://github.com/devsy-org/devsy-provider-kubernetes/commit/3deee89d60b61c0ebf2fe530f0e733bc89b8a242))
* resolve all remaining lint violations ([#11](https://github.com/devsy-org/devsy-provider-kubernetes/issues/11)) ([c92527f](https://github.com/devsy-org/devsy-provider-kubernetes/commit/c92527f384a85675f784116ce277c9341c385d07))
* resolve resources based on pod template and resource option ([ff2a394](https://github.com/devsy-org/devsy-provider-kubernetes/commit/ff2a394169f76be3552aa8b12a14e8ba00485a41))
* resolve resources based on pod template and resource option ([3b7062f](https://github.com/devsy-org/devsy-provider-kubernetes/commit/3b7062f2297edf361f91128af1db459820547911))
* resolve yamllint errors in tests.yaml and defect_report.yaml ([#18](https://github.com/devsy-org/devsy-provider-kubernetes/issues/18)) ([d21b328](https://github.com/devsy-org/devsy-provider-kubernetes/commit/d21b3280e92021f145cdbe8d84b6351e6d5f8481))
* strip leading / in mount path ([a73c8ab](https://github.com/devsy-org/devsy-provider-kubernetes/commit/a73c8ab53622eed091033694a9e9e46643f3aa5c))
* unmarshal inline (arch detection) pod manifest template ([b83a4c3](https://github.com/devsy-org/devsy-provider-kubernetes/commit/b83a4c39272511284a8c23806ad60518940c4d8f))
* workspace labels ([ed6e3f7](https://github.com/devsy-org/devsy-provider-kubernetes/commit/ed6e3f74c2704cb32d8343f3e18f511932804182))

## 1.0.0 (2026-04-19)

Initial release as Devsy Kubernetes Provider.
