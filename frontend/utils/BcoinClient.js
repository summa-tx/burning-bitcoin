/*!
 * BcoinClient.js - bcoin based proof fetcher
 * Copyright (c) 2019, Mark Tyneway (Apache-2.0 License).
 * https://github.com/summa-tx/agoric-bitcoin-spv
 */

'use strict';

const {NodeClient} = require('bcoin/lib/client');
const {Block, ChainEntry} = require('bcoin');
const assert = require('bsert');
const {merkle} = require('bcrypto');
const hash256 = require('bcrypto/lib/hash256');
const consensus = require('bcoin/lib/protocol/consensus');
const BN = require('bcrypto/lib/bn.js');

/**
 * BcoinClient extends the bcoin NodeClient
 * and adds the methods necessary to create
 * the proofs used in this library.
 */
class BcoinClient extends NodeClient {
  /**
   * bcoinclient constructor.
   * @param {object} options
   */
  constructor(options) {
    super(options);
  }

  /**
   * Get the block height of a transaction by txid.
   * Note: requires bcoin tx-indexer enabled.
   * @param {String} txid
   * @returns {Number}
   */
  async getHeightByTX(txid) {
    const tx = await super.getTX(txid);
    if (!tx)
      return null;

    return tx.height;
  }

  /**
   * Get the transaction hex by txid.
   * Note: requires bcoin tx-indexer enabled.
   * @param {String} txid
   * @returns {String}
   */
  async getTX(txid) {
    const tx = await super.getTX(txid);
    return tx.hex;
  }

  /**
   * Validate the merkle tree of a block and then compute
   * a merkle proof of inclusion for the txid.
   * @param {String} txid
   * @param {Number} height
   * @returns {[][]String, Number} - a merkle proof and the index of the leaf
   */
  async getMerkleProof(txid, height) {
    const block = await super.execute('getblockbyheight', [height]);

    let index = -1;
    const txs = [];
    for (const [i, tx] of Object.entries(block.tx)) {
      if (tx === txid)
        index = i >>> 0; // cast to uint from string
      txs.push(Buffer.from(tx, 'hex').reverse());
    }

    assert(index >= 0, 'Transaction not in block.');

    const [root] = merkle.createRoot(hash256, txs.slice());
    assert.bufferEqual(Buffer.from(block.merkleroot, 'hex').reverse(), root);

    const branch = merkle.createBranch(hash256, index, txs.slice());

    const proof = [];
    for (const hash of branch)
      proof.push(hash.toString('hex'));

    return [proof, index];
  }

  /**
   * Get a valid chain of headers starting at a height
   * that have greater than or equal to an amount of work.
   * @param {Number} height
   * @param {String} nwork - hex number
   * @returns {[]String} - a list of hex headers
   */
  async getHeaderChain(height, nwork) {
    const headers = [];
    let accumulated = new BN(0);

    const target = new BN(nwork, 16, 'be');

    if (target.eq(new BN(0)))
      throw new Error('nwork is too small.');

    while (accumulated.lte(target)) {
      const json = await super.getBlock(height);
      const block = Block.fromJSON(json);
      const header = block.toHeaders();

      const valid = consensus.verifyPOW(block.hash(), block.bits);
      assert(valid, 'Invalid Proof of Work.');

      headers.push(header.toRaw().toString('hex'));

      const entry = ChainEntry.fromBlock(block);
      const proof = entry.getProof();

      accumulated = accumulated.add(proof);
      height += 1;
    }

    return headers;
  }
}

module.exports = BcoinClient;

