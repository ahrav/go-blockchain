/*
Package node is the implementation of a blockchain DB.

**READING**

- Books
[Build a blockchain from scratch](https://web3.coach/) - Lukas Lukae

- Video
[Is It Possible To Get The Same 24 BIP39 Seed Words?](https://www.youtube.com/watch?v=hjRntYh0ot8) - John Chow

- Articles
[Transparent Logs for Skeptical Clients](https://research.swtch.com/tlog) - Russ Cox

- Ethereum
[Ethereum Mining](https://ethereum.org/en/developers/docs/consensus-mechanisms/pow/mining/) - Ethereum Website
[Ethereum Blocks](https://ethereum.org/en/developers/docs/blocks/) - Ethereum Website

**NOTES**

Block time refers to the time it takes to mine a new block. In Ethereum, the
average block time is between 12 to 14 seconds and is evaluated after each block.
The expected block time is set as a constant at the protocol level and is used to
protect the network's security when the miners add more computational power. The
average block time gets compared with the expected block time, and if the average
block time is higher, then the difficulty is decreased in the block header. If
the average block time is smaller, then the difficulty in the block header will
be increased.

After a user initiates a transaction, it is propagated peer-to-peer in each node’s
mempool. Each transaction has a fee attached to it. The fee signals the desire to
purchase blockspace, which allows the transaction to be processed and included in a block.

Every moment there are numerous proposed blocks existing in this “Schrödinger’s
state” between unconfirmed and confirmed, competing to find the first hash output
that satisfies the difficulty target.

Nodes also run a series of validity checks on these transactions. These checks
include ensuring that the funds are still available, the output is not exceeding
the input, the signature is valid, etc

Transactions can be dropped from the pool.

When two nodes want to communicate, they send each other some cryptographic data
(public keys and such) to make sure all of the subsequent data transfer is encrypted.

In Proof of Work, the protocol sets out conditions for what makes a block valid.
It might say, for instance, only a block whose hash begins with 00 will be valid.
The only way for the miner to create one that matches that combination is to
brute-force inputs. They can tweak a parameter in their data to produce a different
outcome for every guess until they get the right hash.
*/
package node