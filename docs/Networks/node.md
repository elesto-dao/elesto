---
title: Install Elesto node
---


## Minimum Requirements
The minimum recommended specs for running the Elesto node (`elestod`) is as follows:

- 4-core (2 physical core), x86_64 architecture processor
- 16 GB RAM (or equivalent swap file set up)
- 200 GB of storage space


## Manual installation 

!!! hint
    We'll be using Ubuntu linux during on this guide, if you have chosen a different operating system, you will need to modify your commands to suit your operating system.

### Install pre-requisites

```shell
# update the local package list and install any available upgrades
sudo apt update && sudo apt upgrade -y

# install toolchain and ensure accurate time synchronization
sudo apt install make build-essential git jq ufw curl snapd -y
```

### Install Golang

```shell
# use snap to install the latest stable version of go
sudo snap install go --classic
```

Update your execution path to be able to launch the go binaries:

```shell
 echo 'PATH="$HOME/go/bin:$PATH"' >> ~/.profile && source ~/.profile
```

### Fetch the code from GitHub

```shell
cd ~
git clone https://github.com/elesto-dao/elesto
cd elesto
git checkout v1.0.0-rc2
```

### Build and install the Elesto binary

```shell
make install 
```

the binary will be installed in `~/go/bin/elestod`

### Enable the host firewall [OPTIONAL]

```shell
## allow ssh connection to the server
sudo ufw allow ssh

## allow port to submit transactions
sudo ufw allow 26657/tcp

## start the firewall
sudo ufw enable
```


## Network configuration 

To configure the node to join the **Testnet** network follow [this link](./testnet.md).
