<a name="unreleased"></a>
## [Unreleased]


<a name="v2.0.0-rc1"></a>
## [v2.0.0-rc1] - 2022-06-15
### Chore
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
## [v1.0.0-rc1] - 2022-02-23
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


<a name="latest"></a>
## latest - 2022-01-28

[Unreleased]: https://github.com/elesto-dao/elesto/compare/v2.0.0-rc1...HEAD
[v2.0.0-rc1]: https://github.com/elesto-dao/elesto/compare/v1.0.0-rc2...v2.0.0-rc1
[v1.0.0-rc2]: https://github.com/elesto-dao/elesto/compare/v1.0.0-rc1...v1.0.0-rc2
[v1.0.0-rc1]: https://github.com/elesto-dao/elesto/compare/latest...v1.0.0-rc1
