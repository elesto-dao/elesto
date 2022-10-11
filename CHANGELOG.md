<a name="unreleased"></a>
## [Unreleased]


<a name="v4.0.0-rc1"></a>
## [v4.0.0-rc1] - 2022-10-11
### Chore
- rename module for version v4
- upgrade cosmos-sdk and tendermint to latest patch version
- update changelog for v3.0.0-rc2
- rename module version to v3

### Docs
- add verification diagram for pvc
- amend ADR [#8](https://github.com/elesto-dao/elesto/issues/8), introduce epoch concept

### Feat
- remove public credentials listing in genesis
- add support for CosmosADR036 based proofs ([#274](https://github.com/elesto-dao/elesto/issues/274))
- remove IsPublic field from credential definition
- add fixed inflation distribution ([#259](https://github.com/elesto-dao/elesto/issues/259))
- remove team rewards and community tax from mint
- **credential:** add support allowed list for credential definitions ([#265](https://github.com/elesto-dao/elesto/issues/265))

### Fix
- use pascal case for fields names in credential proof
- fields names in Proof are not compliant to standard
- rollback implementation to epoch base inflation
- **credential:** audit changes credential ([#250](https://github.com/elesto-dao/elesto/issues/250)) ([#269](https://github.com/elesto-dao/elesto/issues/269))
- **docs:** minor edits for about


<a name="v3.0.0-rc2"></a>
## [v3.0.0-rc2] - 2022-09-14
### Chore
- update changelog for v3.0.0-rc2
- rename module version to v3
- remove link-aries command from did cli
- updating the seeds
- update changelog for v3.0.0-rc1
- **cosmoscmd:** remove cosmoscmd from node
- **update:** updgrade to v0.45.7 & fix tests

### Docs
- amend ADR [#8](https://github.com/elesto-dao/elesto/issues/8), introduce epoch concept

### Feat
- remove team rewards and community tax from mint
- add stricter validation of did document
- **did:** inhibiting the creation vanity DIDs ([#220](https://github.com/elesto-dao/elesto/issues/220))
- **docs:** use template for code of conduct

### Fix
- rollback implementation to epoch base inflation
- conflict between /credentials/{id} and /credentials/definitions
- credential definition id should not be DID
- **did:** simulation modifications
- **docs:** remove CoC from contributing because now isolated in CoC file
- **docs:** update ADR description


<a name="v3.0.0-rc1"></a>
## [v3.0.0-rc1] - 2022-08-22
### Chore
- update changelog for v3.0.0-rc1
- prepare codebase for Oak audit
- grammatical improvements
- **credential:** goimport -local
- **credential:** replace ioutils with os
- **credential:** sort imports

### Docs
- add Mint and Inflation ADR
- update docs about network upgrade
- **ADR:** add verifiable credential ADR 006 ([#179](https://github.com/elesto-dao/elesto/issues/179))

### Feat
- add stale github action
- use ghcr.io container image for kubernetes devnet
- add github action release for containers
- add k3d local devnet scripts
- add dockerfile
- **docs:** create mint module readme
- **mint:** chain reset-resistant mint schedule ([#180](https://github.com/elesto-dao/elesto/issues/180))

### Fix
- expose tendermint rpc endpoint on 0.0.0.0 for liveness probe
- link aries command generates invalid service ids ([#195](https://github.com/elesto-dao/elesto/issues/195))
- linter error
- **simulations:** allow simulations to pass even if Interchain Accounts ([#202](https://github.com/elesto-dao/elesto/issues/202))

### Test
- **credential:** add test for update-revocation-list
- **credential:** add test for create-revocation-list
- **credential:** add test for issue-public-credential
- **credential:** add test for publish-credential-definition


<a name="v2.0.0-rc2"></a>
## [v2.0.0-rc2] - 2022-07-04
### Chore
- update changelog for v2.0.0-rc2
- update changelog for v2.0.0-rc1
- **makefile:** build only the node binary for the build target

### Docs
- add how-to for network upgrades

### Feat
- **devnet:** adding script to spin up local devnet ([#160](https://github.com/elesto-dao/elesto/issues/160))

### Fix
- disable CGO for node builds
- **mint:** fix pruning panic, add developer and community funding ([#161](https://github.com/elesto-dao/elesto/issues/161))
- **proto:** fix proto go package for v2

### Refactor
- credential definition ids are not required to be DID


<a name="v2.0.0-rc1"></a>
## [v2.0.0-rc1] - 2022-06-21
### Chore
- update changelog for v2.0.0-rc1
- go mod tidy
- update golang-ci config ([#102](https://github.com/elesto-dao/elesto/issues/102))
- **app:** ignore unused param in simulation call
- **credential:** comment non-working test, will be fixed late
- **go.mod:** require go 1.18

### Docs
- fix mkdocs installation instruction
- add validator documentation ([#87](https://github.com/elesto-dao/elesto/issues/87))

### Feat
- add Interchain Account host module ([#108](https://github.com/elesto-dao/elesto/issues/108))
- prepare v2 release series
- add query for credential definition listing ([#138](https://github.com/elesto-dao/elesto/issues/138))
- implement custom inflation ([#104](https://github.com/elesto-dao/elesto/issues/104))
- add credential support ([#94](https://github.com/elesto-dao/elesto/issues/94))
- **docs:** add ADR for ibc-enabled dids ([#88](https://github.com/elesto-dao/elesto/issues/88))
- **docs:** initial draft of adr-004-light-client-resolver ([#86](https://github.com/elesto-dao/elesto/issues/86))
- **mint:** in-place store migrations ([#147](https://github.com/elesto-dao/elesto/issues/147))

### Fix
- update some settings for the simulation
- **Makefile:** run 4 jobs, 50 blocks for test-sim-after-import target
- **app:** add missing x/credential begin/end block, initgenesis
- **app:** ignore "no validator commission" errors
- **app:** parse validator key with library func
- **app:** register credential storekey and reinstate CLI test for credentials
- **did:** don't abort simulation on verification method call failure

### Test
- increase test coverage for model critical section (>80%)
- increase test coverage (toward 80%)
- increase test coverage for critical sections  ([#120](https://github.com/elesto-dao/elesto/issues/120))


<a name="v1.0.0-rc2"></a>
## [v1.0.0-rc2] - 2022-04-01
### Chore
- update changelog for v1.0.0-rc2
- update the aries command example
- update the aries command example
- **integration:** change vars to be idomatic go
- **metadata:** remove did metadata
- **upgrade:** update sdk version to 0.45.x
- **upgrade:** update ibc-go version to 3.0.0
- **vm:** validate vm when creating
- **vue:** remove frontend code

### Feat
- add configuration for automatic changelog generation ([#81](https://github.com/elesto-dao/elesto/issues/81))
- remove net prefix from did doc id
- adding update did seed
- **audit:** improve core logic of did module, marshaling, msg_server, did struct
- **docs:** adding comments to elaborate on did module functionality
- **genesis:** adding import and export of genesis state
- **migration:** set up migration script
- **sims:** add update did doc simulation message
- **sims:** added simulations for determinism, import/export, benchmark
- **test:** add tests to keeper.go and did.go in keeper package


<a name="v1.0.0-rc1"></a>
## v1.0.0-rc1 - 2022-02-23
### Chore
- remove update did state transition
- remove update did state transition
- pre-v1 cleanup  ([#30](https://github.com/elesto-dao/elesto/issues/30))
- remove testutil folder
- remove testutil folder
- update tests after refactoring
- more ci stuff
- fix test and linting settings
- **proto:** update package in proto files [#23](https://github.com/elesto-dao/elesto/issues/23)
- **proto:** update package in proto files

### Ci
- fix codecov settings

### Docs
- amend edits for the did adr ([#40](https://github.com/elesto-dao/elesto/issues/40))
- port documentation from cosmos-cash ([#34](https://github.com/elesto-dao/elesto/issues/34))

### Feat
- add support for publicKeyJwt verification method
- add quality github workflow
- add quality github workflow
- add did module
- add license
- **did:** update module structure
- **proto:** set version of did proto files to v1
- **proto:** set version of did proto files to v1
- **simulation:** set up simulation tests with create did message ([#13](https://github.com/elesto-dao/elesto/issues/13))
- **simulation:** add simulation for add/delete controller msg
- **simulation:** add simulation for add/delete controller msg
- **simulation:** add unit tests to the simulation framework
- **simulation:** add simulations ([#24](https://github.com/elesto-dao/elesto/issues/24))

### Fix
- check on uri and did method string ([#22](https://github.com/elesto-dao/elesto/issues/22))
- check on uri and did method string

### Test
- increase test coverage


[Unreleased]: https://github.com/elesto-dao/elesto/compare/v4.0.0-rc1...HEAD
[v4.0.0-rc1]: https://github.com/elesto-dao/elesto/compare/v3.0.0-rc2...v4.0.0-rc1
[v3.0.0-rc2]: https://github.com/elesto-dao/elesto/compare/v3.0.0-rc1...v3.0.0-rc2
[v3.0.0-rc1]: https://github.com/elesto-dao/elesto/compare/v2.0.0-rc2...v3.0.0-rc1
[v2.0.0-rc2]: https://github.com/elesto-dao/elesto/compare/v2.0.0-rc1...v2.0.0-rc2
[v2.0.0-rc1]: https://github.com/elesto-dao/elesto/compare/v1.0.0-rc2...v2.0.0-rc1
[v1.0.0-rc2]: https://github.com/elesto-dao/elesto/compare/v1.0.0-rc1...v1.0.0-rc2
