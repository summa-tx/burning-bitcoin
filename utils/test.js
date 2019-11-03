const {utils} = require('@summa-tx/bitcoin-spv-js');
// const { TX, Block, ChainEntry, Headers } = require('bcoin');
const { merkle } = require('bcrypto');
const TX = require('bcoin/lib/primitives/tx');
const Block = require('bcoin/lib/primitives/block')
const ChainEntry = require('bcoin/lib/blockchain/chainentry');
const Headers = require('bcoin/lib/primitives/headers');
console.log({utils})

module.exports = {
  thing () {
    return {
      hex: utils.deserializeHex,
      entry: TX,
      Block, ChainEntry, Headers, merkle
    }
  }
}
