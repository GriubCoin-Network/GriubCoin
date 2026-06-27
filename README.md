<p align="center">
  <img src="./griubcoin.png" width="300">
  <br/>
  <b>GriubCoin (GRBC)</b>
</p>

<div align="center">

[![version](https://img.shields.io/github/tag/GriubCoin-Network/griubcoin.png)](https://github.com/GriubCoin-Network/GriubCoin/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/GriubCoin-Network/GriubCoin)](https://goreportcard.com/report/github.com/GriubCoin-Network/GriubCoin)
[![GitHub](https://img.shields.io/github/license/GriubCoin-Network/GriubCoin.svg)](https://github.com/GriubCoin-Network/GriubCoin/blob/master/LICENSE.md)
[![Twitter Follow](https://img.shields.io/twitter/follow/griubcoin.svg?label=Follow\&style=social)](https://twitter.com/griubcoin)

</div>

<div align="center">

### [Telegram](https://t.me/griubcoin)

</div>

---

GriubCoin (GRBC) is a Layer 1 blockchain network built as a fork of the Kava codebase, extended and modified for the GriubCoin ecosystem.

The project focuses on simplicity, validator participation, and community-driven development.

---

## 🚀 Mainnet / Testnet

Current recommended version: **v1.0.0**

> The `main` branch may contain unstable or experimental changes and should not be used in production.

---

## Installation (Source Build)

```bash
git clone https://github.com/GriubCoin-Network/GriubCoin.git
cd GriubCoin
make build
```

---

## 🧪 Running a Node

### Initialize node

```bash
grbc init mynode
```

### Start node

```bash
grbc start
```

---

## Docker (Recommended)

Prebuilt image:
https://github.com/orgs/GriubCoin-Network/packages/container/package/grbc-node

Validators MUST use a fixed version tag (do NOT use `latest`)

```bash
docker pull ghcr.io/griubcoin-network/grbc-node:1.0.0
```

Run node:

```bash
docker run -d \
  --name grbc-node \
  -p 26656:26656 \
  -p 26657:26657 \
  -v $HOME/.grbc:/root/.grbc \
  ghcr.io/griubcoin-network/grbc-node:1.0.0 \
  grbc start
```

## CLI Reference

This section provides basic command-line interface (CLI) usage for GriubCoin (GRBC) node.

---

## General
```bash
Check version:
grbc version

Check node status:
grbc status

Create a new wallet:
grbc keys add <key-name>

Save your mnemonic securely. It cannot be recovered !!!

List wallets:
grbc keys list

Show wallet address:
grbc keys show <key-name> -a

Initialize node:
grbc init "moniker-name" --chain-id griubcoin_44211-1

Add genesis account (testnet / local setup):
grbc add-genesis-account <address> 1000000ugrbc

Validate genesis file:
grbc validate-genesis

Query Blockchain:
Check account balance
grbc query bank balances <address>

List validators:
grbc query staking validators

Transactions:
Send tokens
grbc tx bank send <from> <to> <amount>ugrbc --chain-id griubcoin_44211-1

Start full node:
grbc start

Stop node (Docker):
docker stop grbc-node

Restart node (Docker):
docker restart grbc-node

## Network Info

* Chain ID: `griubcoin_44211-1`
* Binary: `grbc`
* RPC: `http://127.0.0.1:26657`
* P2P: `26656`

---

## Support

For technical issues or validator onboarding:

* GitHub: https://github.com/GriubCoin-Network/GriubCoin
* Telegram: https://t.me/griubcoin

---

## License

Licensed under Apache 2.0
© GriubCoin Network contributors
