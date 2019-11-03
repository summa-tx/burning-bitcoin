// SIMULATED ON AGORIC CHAIN

import { makeMint } from '@agoric/ertp/core/mint'

// Mint
const burntBitsMint = makeMint('burntBits')

// blocks with Addresses you pay
// Array of Objects. Objects[blockNumber: ['addresses',]]
const upline = []
// Addresses that pay you
const downline = []

/**
 * Distributes the total amount of burntBits
 */
function distribute (address, amount, upline) {
  // Create a new purse named after the address
  const totalBitsToDistribute = burntBitsMint.mint(amount, address)

  // For every address in the upline, create a payment
  upline.forEach((block, index) => {
    block.forEach((address) => {
      totalBitsToDistribute.
    })
  })
}
