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

```bash
docker pull dmrdev2/grbc-node:latest
```

Run node:

```bash
docker run -d \
  --name grbc-node \
  -p 26656:26656 \
  -p 26657:26657 \
  -v $HOME/.grbc:/root/.grbc \
  dmrdev2/grbc-node:latest \
  grbc start
```

---

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
