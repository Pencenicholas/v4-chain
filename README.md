[grpc]
# Enable grpc. The Cosmos gRPC service is used by various daemon processes,
# and must be enabled in order for the protocol to operate:
enable = true
# Non-standard gRPC ports are not supported at this time. Please run on port 9090, which is the default
# port specified in the config file.
# Note: grpc can be also be configured via start flags. Be careful not to change the default settings
# with either of the following flags: `--grpc.enable`, `--grpc.address`.
address = "0.0.0.0:9090""><img src="https://dydx.exchange/icon.svg?" width="256" /></p>

<h1 align="center">dYdX Chain</h1>

<div align="center">
  <a href='https://github.com/dydxprotocol/v4-chain/blob/main/LICENSE'>
    <img src='https://img.shields.io/badge/License-AGPL_v3-blue.svg' alt='License' />
  </a>
</div>

The dYdX v4 software (the ”dYdX Chain”) is a sovereign blockchain software built using Cosmos SDK and CometBFT. It powers a robust decentralized perpetual futures exchange with a high-performance orderbook and matching engine for a feature-rich, self-custodial perpetual trading experience on any market.

This repository contains the source code for the Cosmos SDK application responsible for running the chain and the associated indexer services. The indexer services are a read-only layer that indexes on-chain and off-chain data from a node and serves it to users in a more performant and user-friendly way.

# Getting Started

[Architecture Overview](https://dydx.exchange/blog/v4-technical-architecture-overview)

[Indexer Deep Dive](https://dydx.exchange/blog/v4-deep-dive-indexer)

[Front End Deep Dive](https://dydx.exchange/blog/v4-deep-dive-front-end)

# Resources and Documentation

[Official Documentation](https://docs.dydx.exchange/)

[dYdX Academy](https://dydx.exchange/crypto-learning#)

[dYdX Blog](https://dydx.exchange/blog#)

# Clients

[v4-clients](https://github.com/dydxprotocol/v4-clients)

# Third-party Clients

[C++ Client](https://github.com/asnefedovv/dydx-v4-client-cpp)

[Python Client](https://github.com/kaloureyes3/v4-clients/tree/main/v4-client-py)

By clicking the above links to third-party clients, you will leave the dYdX Trading Inc. (“dYdX”) GitHub repository and join repositories made available by third parties, which are independent from and unaffiliated with dYdX. dYdX is not responsible for any action taken or content on third-party repositories.

# Directory Structure

`audits` — Audit reports live here.

`docs` — Home for documentation pertaining to the entire repo.

`indexer` — Contains source code for indexer services. See its [README](https://github.com/dydxprotocol/v4-chain/blob/main/indexer/README.md) for developer documentation.

`proto` — Contains all protocol buffers used by protocol and indexer.

`protocol` — Contains source code for the Cosmos SDK app that runs the protocol. See its [README](https://github.com/dydxprotocol/v4-chain/blob/main/protocol/README.md) for developer documentation.

`v4-proto-js` — Scripts for publishing proto package to [npm](https://www.npmjs.com/package/@dydxprotocol/v4-proto).

`v4-proto-py` — Scripts for publishing proto package to [PyPI](https://pypi.org/project/v4-proto/).

# Contributing

If you encounter a bug, please file an [issue](https://github.com/dydxprotocol/v4-chain/issues) to let us know. Alternatively, please feel free to contribute a bug fix directly. If you are planning to contribute changes that involve significant design choices, please open an issue for discussion instead.

# License and Terms

The dYdX Chain is licensed under AGPLv3 and subject to the [v4 Terms of Use](https://dydx.exchange/v4-terms). The full license can be found [here](https://github.com/dydxprotocol/v4-chain/blob/main/LICENSE). dYdX does not deploy or run v4 software for public use, or operate or control any dYdX Chain infrastructure. dYdX products and services are not available to persons or entities who reside in, are located in, are incorporated in, or have registered offices in the United States or Canada, or Restricted Persons (as defined in the dYdX [Terms of Use](https://dydx.exchange/terms)).
# This is a TOML document.

title = "TOML Example"

[owner]
name = "Tom Preston-Werner"
dob = 1979-05-27T07:32:00-08:00 # First class dates

[database]
server = "192.168.1.1"
ports = [ 8000, 8001, 8002 ]
connection_max = 5000
enabled = true

[servers]

  # Indentation (tabs and/or spaces) is allowed but not required
  [servers.alpha]
  ip = "10.0.0.1"
  dc = "eqdc10"

  [servers.beta]
  ip = "10.0.0.2"
  dc = "eqdc10"

[clients]
data = [ ["gamma", "delta"], [1, 2] ]

# Line breaks are OK when inside arrays
hosts = [
  "alpha",
  "omega"
]# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

###############################################################################
###                           Base Configuration                            ###
###############################################################################

# The minimum gas prices a validator is willing to accept for processing a
# transaction. A transaction's fees must meet the minimum of any denomination
# specified in this config (e.g. 0.25token1,0.0001token2).
minimum-gas-prices = "0stake"

# default: the last 362880 states are kept, pruning at 10 block intervals
# nothing: all historic states will be saved, nothing will be deleted (i.e. archiving node)
# everything: 2 latest states will be kept; pruning at 10 block intervals.
# custom: allow pruning options to be manually specified through 'pruning-keep-recent', and 'pruning-interval'
pruning = "default"

# These are applied if and only if the pruning strategy is custom.
pruning-keep-recent = "0"
pruning-interval = "0"

# HaltHeight contains a non-zero block height at which a node will gracefully
# halt and shutdown that can be used to assist upgrades and testing.
#
# Note: Commitment of state will be attempted on the corresponding block.
halt-height = 0

# HaltTime contains a non-zero minimum block time (in Unix seconds) at which
# a node will gracefully halt and shutdown that can be used to assist upgrades
# and testing.
#
# Note: Commitment of state will be attempted on the corresponding block.
halt-time = 0

# MinRetainBlocks defines the minimum block height offset from the current
# block being committed, such that all blocks past this offset are pruned
# from Tendermint. It is used as part of the process of determining the
# ResponseCommit.RetainHeight value during ABCI Commit. A value of 0 indicates
# that no blocks should be pruned.
#
# This configuration value is only responsible for pruning Tendermint blocks.
# It has no bearing on application state pruning which is determined by the
# "pruning-*" configurations.
#
# Note: Tendermint block pruning is dependent on this parameter in conjunction
# with the unbonding (safety threshold) period, state pruning and state sync
# snapshot parameters to determine the correct minimum value of
# ResponseCommit.RetainHeight.
min-retain-blocks = 0

# InterBlockCache enables inter-block caching.
inter-block-cache = true

# IndexEvents defines the set of events in the form {eventType}.{attributeKey},
# which informs Tendermint what to index. If empty, all events will be indexed.
#
# Example:
# ["message.sender", "message.recipient"]
index-events = []

# IavlCacheSize set the size of the iavl tree cache (in number of nodes).
iavl-cache-size = 781250

# IAVLDisableFastNode enables or disables the fast node feature of IAVL. 
# Default is false.
iavl-disable-fastnode = false

# IAVLLazyLoading enable/disable the lazy loading of iavl store.
# Default is false.
iavl-lazy-loading = false

# AppDBBackend defines the database backend type to use for the application and snapshots DBs.
# An empty string indicates that a fallback will be used.
# The fallback is the db_backend value set in Tendermint's config.toml.
app-db-backend = ""

###############################################################################
###                         Telemetry Configuration                         ###
###############################################################################

[telemetry]

# Prefixed with keys to separate services.
service-name = ""

# Enabled enables the application telemetry functionality. When enabled,
# an in-memory sink is also enabled by default. Operators may also enabled
# other sinks such as Prometheus.
enabled = false

# Enable prefixing gauge values with hostname.
enable-hostname = false

# Enable adding hostname to labels.
enable-hostname-label = false

# Enable adding service to labels.
enable-service-label = false

# PrometheusRetentionTime, when positive, enables a Prometheus metrics sink.
prometheus-retention-time = 0

# GlobalLabels defines a global set of name/value label tuples applied to all
# metrics emitted using the wrapper functions defined in telemetry package.
#
# Example:
# [["chain_id", "cosmoshub-1"]]
global-labels = []

###############################################################################
###                           API Configuration                             ###
###############################################################################

[api]

# Enable defines if the API server should be enabled.
enable = false

# Swagger defines if swagger documentation should automatically be registered.
swagger = false

# Address defines the API server to listen on.
address = "tcp://localhost:1317"

# MaxOpenConnections defines the number of maximum open connections.
max-open-connections = 1000

# RPCReadTimeout defines the Tendermint RPC read timeout (in seconds).
rpc-read-timeout = 10

# RPCWriteTimeout defines the Tendermint RPC write timeout (in seconds).
rpc-write-timeout = 0

# RPCMaxBodyBytes defines the Tendermint maximum request body (in bytes).
rpc-max-body-bytes = 1000000

# EnableUnsafeCORS defines if CORS should be enabled (unsafe - use it at your own risk).
enabled-unsafe-cors = false

###############################################################################
###                           Rosetta Configuration                         ###
###############################################################################

[rosetta]

# Enable defines if the Rosetta API server should be enabled.
enable = false

# Address defines the Rosetta API server to listen on.
address = ":8080"

# Network defines the name of the blockchain that will be returned by Rosetta.
blockchain = "app"

# Network defines the name of the network that will be returned by Rosetta.
network = "network"

# Retries defines the number of retries when connecting to the node before failing.
retries = 3

# Offline defines if Rosetta server should run in offline mode.
offline = false

# EnableDefaultSuggestedFee defines if the server should suggest fee by default.
# If 'construction/medata' is called without gas limit and gas price,
# suggested fee based on gas-to-suggest and denom-to-suggest will be given.
enable-fee-suggestion = false

# GasToSuggest defines gas limit when calculating the fee
gas-to-suggest = 200000

# DenomToSuggest defines the default denom for fee suggestion.
# Price must be in minimum-gas-prices.
denom-to-suggest = "uatom"

###############################################################################
###                           gRPC Configuration                            ###
###############################################################################

[grpc]

# Enable defines if the gRPC server should be enabled.
enable = true

# Address defines the gRPC server address to bind to.
address = "localhost:9090"

# MaxRecvMsgSize defines the max message size in bytes the server can receive.
# The default value is 10MB.
max-recv-msg-size = "10485760"

# MaxSendMsgSize defines the max message size in bytes the server can send.
# The default value is math.MaxInt32.
max-send-msg-size = "2147483647"

###############################################################################
###                        gRPC Web Configuration                           ###
###############################################################################

[grpc-web]

# GRPCWebEnable defines if the gRPC-web should be enabled.
# NOTE: gRPC must also be enabled, otherwise, this configuration is a no-op.
enable = true

# Address defines the gRPC-web server address to bind to.
address = "localhost:9091"

# EnableUnsafeCORS defines if CORS should be enabled (unsafe - use it at your own risk).
enable-unsafe-cors = false

###############################################################################
###                        State Sync Configuration                         ###
###############################################################################

# State sync snapshots allow other nodes to rapidly join the network without replaying historical
# blocks, instead downloading and applying a snapshot of the application state at a given height.
[state-sync]

# snapshot-interval specifies the block interval at which local state sync snapshots are
# taken (0 to disable).
snapshot-interval = 0

# snapshot-keep-recent specifies the number of recent snapshots to keep and serve (0 to keep all).
snapshot-keep-recent = 2

###############################################################################
###                         Store / State Streaming                         ###
###############################################################################

[store]
streamers = []

[streamers]
[streamers.file]
keys = ["*"]
write_dir = ""
prefix = ""

# output-metadata specifies if output the metadata file which includes the abci request/responses 
# during processing the block.
output-metadata = "true"

# stop-node-on-error specifies if propagate the file streamer errors to consensus state machine.
stop-node-on-error = "true"

# fsync specifies if call fsync after writing the files.
fsync = "false"

###############################################################################
###                         Mempool                                         ###
###############################################################################

[mempool]
# Setting max-txs to 0 will allow for a unbounded amount of transactions in the mempool.
# Setting max_txs to negative 1 (-1) will disable transactions from being inserted into the mempool.
# Setting max_txs to a positive number (> 0) will limit the number of transactions in the mempool, by the specified amount.
#
# Note, this configuration only applies to SDK built-in app-side mempool
# implementations.
max-txs = 5000### Gas Prices ###
minimum-gas-prices = "0.025ibc/8E27BA2D5493AF5636760E354E46004562C46AB7EC0CC4C1CA14E9E20E2545B5,12500000000$NATIVE_TOKEN_DENOM"

### Pruning ###
pruning = "custom"
# Small numbers >= "2" for validator nodes.
# Larger numbers could be used for full-nodes if they are used for historical queries.
pruning-keep-recent = "7"
# Any prime number between "13" and "97", inclusive.
pruning-interval = "17"dydxprotocold start --p2p.seeds="..." --bridge-daemon-eth-rpc-endpoint="<eth rpc endpoint>" --non-validating-full-node=true# For the deployment by DYDX token holders
# CHAIN_ID=dydx-mainnet-1

# For testnet
CHAIN_ID=dydx-testnet-4