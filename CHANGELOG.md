# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [8.1.0] - 2025-09-17

### Changed

- Go: Update dependencies.
- Update k8s modules to v0.34.1 (#390)
- Update module github.com/spf13/viper to v1.21.0 (#389)
- Update module sigs.k8s.io/controller-runtime to v0.22.1 (#388)
- Update module k8s.io/apiextensions-apiserver to v0.33.2 (#372)
- Update dependency go to v1.24.3 (#362)
- Bump golang.org/x/net from 0.36.0 to 0.38.0 (#355)
- Update module github.com/google/go-cmp to v0.7.0 (#343)
- Update module github.com/giantswarm/micrologger to v1.1.2 (#330)

## [8.0.0] - 2024-10-22

### Changed

- Upgrade controller-runtime to v0.19.0
- Update go to v1.22.6
- Upgrade k8s modules to v0.31.1

## [7.2.0] - 2023-11-09

### Changed

- Upgrade go to 1.20

## [7.1.0] - 2023-10-24

### Changed

- Upgrade controller-runtime to v0.16.3

## [7.0.1] - 2021-12-21

### Fixed

- Fix k8sclient/fake missing `CRDClient()` method.

## [7.0.0] - 2021-12-17

### Changed

- Upgrade github.com/giantswarm/backoff v0.2.0 to v1.0.0
- Upgrade github.com/giantswarm/microerror v0.3.0 to v0.4.0
- Upgrade github.com/giantswarm/micrologger v0.5.0 to v0.6.0

## [6.1.0] - 2021-12-17

### Added

- Add back CRDClient removed in v6.0.0.

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



[Unreleased]: https://github.com/giantswarm/k8sclient/compare/v8.1.0...HEAD
[8.1.0]: https://github.com/giantswarm/k8sclient/compare/v8.0.0...v8.1.0
[8.0.0]: https://github.com/giantswarm/k8sclient/compare/v7.2.0...v8.0.0
[7.2.0]: https://github.com/giantswarm/k8sclient/compare/v7.1.0...v7.2.0
[7.1.0]: https://github.com/giantswarm/k8sclient/compare/v7.0.1...v7.1.0
[7.0.1]: https://github.com/giantswarm/k8sclient/compare/v7.0.0...v7.0.1
[7.0.0]: https://github.com/giantswarm/k8sclient/compare/v6.1.0...v7.0.0
[6.1.0]: https://github.com/giantswarm/k8sclient/compare/v6.0.0...v6.1.0
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
