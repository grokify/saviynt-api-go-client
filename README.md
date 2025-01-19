# Saviynt API Go Client

[![Build Status][build-status-svg]][build-status-link]
[![Lint Status][lint-status-svg]][lint-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

 [build-status-svg]: https://github.com/grokify/saviynt-api-go-client/workflows/test/badge.svg
 [build-status-link]: https://github.com/grokify/saviynt-api-go-client/actions/workflows/test.yaml
 [lint-status-svg]: https://github.com/grokify/saviynt-api-go-client/workflows/lint/badge.svg
 [lint-status-link]: https://github.com/grokify/saviynt-api-go-client/actions/workflows/lint.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/saviynt-api-go-client
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/saviynt-api-go-client
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/saviynt-api-go-client
 [docs-godoc-link]: https://pkg.go.dev/github.com/grokify/saviynt-api-go-client
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/saviynt-api-go-client/blob/master/LICENSE

## Overview

This is an auto-generated client using OpenAPI Generator `7.10.0`.

## Supported Saviynt APIs

| No. | Tag | Name | Endpoint | In Spec | Test: cURL | Test: OpenAPI Generator | Test: Automated |
| - | - | - | - | - | - | - | - |
| 1 | Delegated Administration | Get Delegate User List | `GET /ECM/api/v5/getDelegateUserList` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 2 | Delegated Administration | Create Delegate | `POST /ECM/api/v5/createDelegate` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 3 | Delegated Administration | Fetch Delegates List | `POST /ECM/api/v5/fetchDelegatesList` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 4 | Delegated Administration | Edit Delegate | `POST /ECM/api/v5/editDelegate` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 5 | Delegated Administration | Delete Delegate | `POST /ECM/api/v5/deleteDelegate` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 6 | Job Control | Create and Update Trigger | `POST /ECM/api/v5/createUpdateTrigger` | :white_check_mark: | :x: | :x: | :x: |
| 7 | Job Control | Check Job Status | `POST /ECM/api/v5/checkJobStatus` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 8 | Job Control | Delete Trigger | `POST /ECM/api/v5/deleteTrigger` | :white_check_mark: | :x: | :x: | :x: |
| 9 | Job Control | Run Job Trigger | `POST /ECM/api/v5/runJobTrigger` | :white_check_mark: | :x: | :x: | :x: |
| 10 | Job Control | Fetch Job Metadata | `POST /ECM/api/v5/fetchJobMetadata` | :white_check_mark: | :x: | :x: | :x: |
| 11 | Job Control | Create Triggers | `POST /ECM/api/v5/createTriggers` | :white_check_mark: | :x: | :x: | :x: |
| 12 | Job Control | Resume Pause Jobs | `POST /ECM/api/v5/resumePauseJobs` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 13 | mTLS Authentication | Upload KeyStore | `POST /ECM/api/v5/uploadKeyStore` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 14 | mTLS Authentication | Get KeyStore Details | `POST /ECM/api/v5/getKeyStoreCertificateDetails` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 15 | mTLS Authentication | Delete KeyStore | `POST /ECM/api/v5/deleteKeyStoreAlias/{alias}` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 16 | SAV Roles | Get All SAV Roles | `POST /ECMv6/api/userms/savroles` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 17 | SAV Roles | Get SAV Role Users | `POST /ECMv6/api/userms/savroles/{roleName/users` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 18 | Transport | Export Package | `POST /ECM/api/v5/exportTransportPackage` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 19 | Transport | Import Package | `POST /ECM/api/v5/importTransportPackage` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |
| 20 | Transport | Transport Status | `GET /ECM/api/v5/transportPackageStatus` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x: |