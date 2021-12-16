# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [Unreleased]

### Changed

- Upgrade github.com/giantswarm/backoff v0.2.0 to v1.0.0
- Upgrade github.com/giantswarm/microerror v0.3.0 to v0.4.0
- Upgrade github.com/giantswarm/micrologger v0.5.0 to v0.6.0

## [6.0.0] - 2021-10-29

### Changed

- Update Kubernetes dependencies to v1.20.12.
- Update controller-runtime to v0.8.3.
- Update architect-orb to v4.6.0.

### Removed

- Remove CRD client (CRDClient) and typed client (G8sClient) to avoid dependency on apiextensions.

## [5.12.0] - 2021-08-05

### Changed

- Update apiextensions to v3.30 to introduce `v1alpha3` GS AWS CR's.

## [5.11.0] - 2021-02-18

### Added

- Handle timeout when discovering dynamic client and return as custom error.

## [5.10.0] - 2021-02-08

### Added

- Fake clientset for testing (port from v4)

## [5.0.0] - 2020-10-27

### Changed

- Update apiextensions to v3 and replace CAPI with Giant Swarm fork.
- Prepare module v5.

### Fixed

- Improved error message in `k8sclient.NewClients` function

## [4.0.0] - 2020-08-10

### Changed

- Updated Kubernetes dependencies to v1.18.5

## [3.1.2] - 2020-07-16

### Changed

- Updated apiextensions to v0.4.15

## [3.1.1] - 2020-07-06

### Added

- Add release workflows.

### Changed

- Update `go.mod` file.



## [3.1.0] 2020-05-16

### Changed

- Simplify test interface.
- Add `k8sclienttest.NewEmpty`.



## [3.0.0] 2020-04-14

### Changed

- Graduate CRD Client to `apiextensions/v1`.
- Prepare module v3.



## [2.0.0] 2020-04-14

### Changed

- Aligned import path with upcoming major version `v2.0.0`.



## [1.0.1] 2020-04-14

### Changed

- Aligned import path with major version.



## [1.0.0] 2020-04-14

### Changed

- Prepare project structure for `v1.0.0` by having all go code in `pkg/k8sclient/`.
- Use architect orb `v0.8.9`.



## [0.2.0] 2020-03-20

### Changed

- Switch from dep to Go modules.
- Use architect orb.



## [0.1.0] 2020-03-18

### Added

- First release.



[Unreleased]: https://github.com/giantswarm/k8sclient/compare/v6.0.0...HEAD
[6.0.0]: https://github.com/giantswarm/k8sclient/compare/v5.12.0...v6.0.0
[5.12.0]: https://github.com/giantswarm/k8sclient/compare/v5.11.0...v5.12.0
[5.11.0]: https://github.com/giantswarm/k8sclient/compare/v5.10.0...v5.11.0
[5.10.0]: https://github.com/giantswarm/k8sclient/compare/v5.0.0...v5.10.0
[5.0.0]: https://github.com/giantswarm/k8sclient/compare/v4.0.0...v5.0.0
[4.0.0]: https://github.com/giantswarm/k8sclient/compare/v3.1.2...v4.0.0
[3.1.2]: https://github.com/giantswarm/k8sclient/compare/v3.1.1...v3.1.2
[3.1.1]: https://github.com/giantswarm/k8sclient/compare/v3.1.0...v3.1.1
[3.1.0]: https://github.com/giantswarm/k8sclient/compare/v3.0.0...v3.1.0
[3.0.0]: https://github.com/giantswarm/k8sclient/compare/v2.0.0...v3.0.0
[2.0.0]: https://github.com/giantswarm/k8sclient/compare/v1.0.1...v2.0.0
[1.0.1]: https://github.com/giantswarm/k8sclient/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/giantswarm/k8sclient/compare/v0.2.0...v1.0.0
[0.2.0]: https://github.com/giantswarm/k8sclient/compare/v0.1.0...v0.2.0

[0.1.0]: https://github.com/giantswarm/k8sclient/releases/tag/v0.1.0
