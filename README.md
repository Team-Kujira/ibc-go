<div align="center">
  <h1>ibc-go</h1>
</div>

![banner](docs/ibc-go-image.png)

<div align="center">
  <a href="https://github.com/cosmos/ibc-go/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/cosmos/ibc-go.svg" />
  </a>
  <a href="https://github.com/cosmos/ibc-go/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/cosmos/ibc-go.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/cosmos/ibc-go?tab=doc">
    <img alt="GoDoc" src="https://godoc.org/github.com/cosmos/ibc-go?status.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/cosmos/ibc-go">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/cosmos/ibc-go" />
  </a>
  <a href="https://codecov.io/gh/cosmos/ibc-go">
    <img alt="Code Coverage" src="https://codecov.io/gh/cosmos/ibc-go/branch/main/graph/badge.svg" />
  </a>
</div>
<div align="center">
  <a href="https://github.com/cosmos/ibc-go">
    <img alt="Lines Of Code" src="https://tokei.rs/b1/github/cosmos/ibc-go" />
  </a>
  <a href="https://discord.gg/AzefAFd">
    <img alt="Discord" src="https://img.shields.io/discord/669268347736686612.svg" />
  </a>
  <a href="https://sourcegraph.com/github.com/cosmos/ibc-go?badge">
    <img alt="Imported by" src="https://sourcegraph.com/github.com/cosmos/ibc-go/-/badge.svg" />
  </a>
    <img alt="Lint Status" src="https://github.com/cosmos/cosmos-sdk/workflows/Lint/badge.svg" />
</div>

The Inter-Blockchain Communication protocol (IBC) allows blockchains to talk to each other. IBC handles transport across different sovereign blockchains. This end-to-end, connection-oriented, stateful protocol provides reliable, ordered, and authenticated communication between heterogeneous blockchains. This IBC implementation in Golang is built as a Cosmos SDK module.

## Contents

1. **[Core IBC Implementation](https://github.com/cosmos/ibc-go/tree/main/modules/core)**

    1.1 [ICS 02 Client](https://github.com/cosmos/ibc-go/tree/main/modules/core/02-client)

    1.2 [ICS 03 Connection](https://github.com/cosmos/ibc-go/tree/main/modules/core/03-connection)

    1.3 [ICS 04 Channel](https://github.com/cosmos/ibc-go/tree/main/modules/core/04-channel)

    1.4 [ICS 05 Port](https://github.com/cosmos/ibc-go/tree/main/modules/core/05-port)

    1.5 [ICS 23 Commitment](https://github.com/cosmos/ibc-go/tree/main/modules/core/23-commitment/types)

    1.6 [ICS 24 Host](https://github.com/cosmos/ibc-go/tree/main/modules/core/24-host)

2. **Applications**

    2.1 [ICS 20 Fungible Token Transfers](https://github.com/cosmos/ibc-go/tree/main/modules/apps/transfer)

    2.2 [ICS 27 Interchain Accounts](https://github.com/cosmos/ibc-go/tree/main/modules/apps/27-interchain-accounts)

3. **Light Clients**

    3.1 [ICS 07 Tendermint](https://github.com/cosmos/ibc-go/tree/main/modules/light-clients/07-tendermint)

    3.2 [ICS 06 Solo Machine](https://github.com/cosmos/ibc-go/tree/main/modules/light-clients/06-solomachine)

## Roadmap

For an overview of upcoming changes to ibc-go take a look at the [roadmap](./docs/roadmap/roadmap.md).

## Support

We have active, helpful communities on Discord and Telegram.

For questions and support please use the `new-devs-support` channel in the [Cosmos Network Discord server](https://discord.com/channels/669268347736686612/699634178173698108) or join the [IBC Gang Discord server](https://discord.gg/RdpdkaXKpZ). The issue list of this repo is exclusively for bug reports and feature requests.

To receive announcements of new releases or other technical updates, please join the [IBC is expansive Telegram group](https://t.me/ibc_is_expansive).

## Resources

- [IBC Website](https://ibcprotocol.org/)
- [IBC Specification](https://github.com/cosmos/ibc)
- [Documentation](https://ibc.cosmos.network/main/ibc/overview.html)
