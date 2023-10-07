# CoCage
> This name is a placeholder and will be changed in the future. Currently, it's a meme ref and answer
> to [the question](https://x.com/nickwh8te/status/1682779788350566402?s=20).

CoCage is Cosmos SDK Data Availability/Data Publication Module. This enables on-chain verification of 
published data by an external network such as Celestia. Validators continually run a side-car process(light client) 
that samples the published blocks on the external network. They then employ vote extensions to attest to data 
commitments (merkle roots) over the data published by the external network.

## Motivation

CoCage is a foundational component for **trust minimized proof-based IBC bridging** (as opposed to committee based). In order for Cosmos to be able to bridge to other chains 
and Rollups in trust minimized way, it needs to be able to verify the data published by those chains. The goal of CoCage
is to provide a reusable module that enables data publication verification for a variety of publication networks. 

The original design was inspired by a collaboration between [Neutron](https://neutron.org/) and [Celestia](httsp://celestia.org). CoCage builds on top of Neutron's design and simplifies
it substantially.

## Design

### Vote Extensions

On every new block, validators check if their Light Node sampled any new DP network heights and if so they submit a vote
extension with the latest DP height sampled.

### Prepare Proposal

Proposer checks for vote extensions from previous height. If so, they find the latest
submitted DP height with +2/3 votes, get the data commitments for the heights sampled from their DA Light Node and 
propose them as a TX.

### Process Proposal

ProcessProposal then verifies the data commitments from against the DP network and if they are valid, it updates 
the state of Keeper.

### Keeper

Keeper stores the mapping between DP heights and respective data commitments, so that future modules and IBC client
can verify proofs against the commitments.

## Future Work

* Abstract the DP network interface, so that we can support multiple DP networks.
* Add support for IBC client to verify proofs against the DP network commitments.

## Supported Networks

### Celestia

The first external network supported by CoCage is [Celestia](https://celestia.network/). Celestia is a first modular
DA/DP (Data Availability/Publication) network that provides a decentralized data availability layer for blockchains.

We use Rollkit's [Celestia OpenRPC API](https://github.com/rollkit/celestia-openrpc/) to query and sample data from the
Celestia's Light Node. We cannot depend on celestia-node's API directly, and we cannot embed LN in Validator process,
due to conflicting SDK versions/forks.
