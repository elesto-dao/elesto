
[![Go Reference](https://pkg.go.dev/badge/github.com/elesto-dao/elesto.svg)](https://pkg.go.dev/github.com/elesto-dao/elesto)
[![build](https://github.com/elesto-dao/elesto/actions/workflows/quality.yaml/badge.svg?branch=main)](https://github.com/elesto-dao/elesto/actions/workflows/quality.yaml)
[![codecov](https://codecov.io/gh/elesto-dao/elesto/branch/main/graph/badge.svg?token=NLT5ZWM460)](https://codecov.io/gh/elesto-dao/elesto)
[![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/elesto-dao/elesto)](https://libraries.io/go/github.com%2Felesto-dao%2Felesto)
[![DeepSource](https://deepsource.io/gh/elesto-dao/elesto.svg/?label=active+issues&show_trend=true&token=BRR7kVLyskz5-N1etTDRay5J)](https://deepsource.io/gh/elesto-dao/elesto/?ref=repository-badge)

# Elesto

## Get started

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.com).

### Web Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.com/elesto-dao/elesto@latest! | sudo bash
```
`elesto-dao/elesto` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Starport](https://starport.com)
- [Tutorials](https://docs.starport.com/guide)
- [Starport docs](https://docs.starport.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/H6wGTY8sxw)
