/*!
 * BcoinClient.js - bcoin based proof fetcher
 * Copyright (c) 2019, Mark Tyneway (Apache-2.0 License).
 * https://github.com/summa-tx/agoric-bitcoin-spv
 */

'use strict';

const {NodeClient} = require('bcoin/lib/client');
const {TX, Block, ChainEntry, Headers} = require('bcoin');
const assert = require('bsert');
const {merkle} = require('bcrypto');
const hash256 = require('bcrypto/lib/hash256');
const consensus = require('bcoin/lib/protocol/consensus');
const BN = require('bcrypto/lib/bn.js');
const {utils} = require('@summa-tx/bitcoin-spv-js');

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
   * @params {String} txid - big endian
   * @params {Number} count
   */

  async getProof(txid) {
    assert(typeof txid === 'string');

    const tx = await super.getTX(txid);

    if (!tx)
      throw new Error('Cannot find transaction');

    if (tx.height < 0)
      throw new Error('Transaction not confirmed');

    const json = await super.getBlockHeader(tx.height);

    if (!json)
      throw new Error('Cannot find header');

    const header = await this.getHeader(tx.height);
    const txinfo = parseBcoinTx(tx.hex);

    let [nodes, index] = await this.getMerkleProof(txid, tx.height);
    nodes = Buffer.from(nodes.join(), 'hex');

    return {
      version: Number(utils.bytesToUint(utils.reverseEndianness(txinfo.version))),
      vin: txinfo.vin,
      vout: txinfo.vout,
      locktime: Number(utils.bytesToUint(Buffer.from(txinfo.locktime).reverse())),
      tx_id: new Uint8Array(Buffer.from(txid, 'hex')),
      tx_id_le: new Uint8Array(Buffer.from(reverse(txid), 'hex')),
      index: index,
      confirming_header: header,
      intermediate_nodes: new Uint8Array(nodes)
    }
  }

  async getHeader(height) {
    const json = await super.getBlockHeader(height);
    const header = Headers.fromJSON(json);

    return {
      raw: new Uint8Array(header.toRaw().slice(0, 80)),
      hash: new Uint8Array(header.hash()),
      hash_le: new Uint8Array(header.hash().reverse()),
      height: height,
      prevhash: new Uint8Array(header.prevBlock),
      merkle_root: new Uint8Array(header.merkleRoot),
      merkle_root_le: new Uint8Array(header.merkleRoot.reverse()),
    }
  }

  async getHeaders(height, count) {
    const headers = [];

    for (let i = 0; i < count; i++) {
      const header = await this.getHeader(height + i);
      headers.push(header);
    }

    return headers;
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

  async getHeaderChainByCount(height, count) {
    assert((height >>> 0) === height);
    assert((count >>> 0) === count);

    const headers = [];

    for (let i = 0; i < count; i++) {
      const json = await super.getBlockHeader(height);

      if (!json)
        throw new Error('Cannot find header');

      const header = Headers.fromJSON(json);
      const hex = header.toRaw().toString('hex');
      headers.push(hex);
    }

    return new Uint8Array(headers.join(''));
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

// helper functions
function parseBcoinTx(hex) {
  if (typeof hex !== 'string') {
    throw new Error('Must pass string')
  }

  const tx = TX.fromRaw(hex, 'hex');

  if (tx.inputs.length > 253 || tx.outputs.length > 253) {
    throw RangeError('too many ins/outs');
  }

  const raw = tx.toRaw();
  // version and witness flag if any
  const baseOffset = raw[4] === 0 ? 6 : 4;

  const vinBytes = 1 + tx.inputs.reduce((a, b) => a + b.getSize(), 0);
  const voutBytes = 1 + tx.outputs.reduce((a, b) => a + b.getSize(), 0);

  return {
    version: new Uint8Array(raw.subarray(0, 4)),
    vin: new Uint8Array(raw.subarray(baseOffset, baseOffset + vinBytes)),
    vout: new Uint8Array(raw.subarray(baseOffset + vinBytes, baseOffset + vinBytes + voutBytes)),
    locktime: new Uint8Array(raw.subarray(-4))
  };
}

/**
 * Reverse the endianess of a hex string
 */

function reverse(str) {
  return Buffer.from(str, 'hex').reverse().toString('hex');
}
