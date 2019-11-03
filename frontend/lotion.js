const lotion = require('lotion')
const {ValidateSPV, utils} = require('@summa-tx/bitcoin-spv-js')

const app = lotion({
  initialState: {
    tokens: 0,
    proofs: []
  }
})
function generateTokens (index) {
  console.log({index})
  const decimals = parseInt(Math.random() * 10) * index
  const randomNumOfTokens = parseInt(Math.random() * 1000 * decimals)
  state.tokens += randomNumOfTokens
}
function transactionHandler (state, transaction) {
  const proof = transaction.proof
  const txid = proof.tx_id
  console.log(proof.tx_id)
  if (state.proofs.length > 0) {
    if (!state.proofs.find(p => p === txid)) {
      generateTokens(proof.index)
      state.proofs.push(txid)
    } else {
      throw new Error('Proof already exists')
    }
  } else  {
    state.proofs.push(txid)
  }
  // if (ValidateSPV.validateProof(proof)) {
  //   state.proof_count++
  // }
  if (proof) {
    state.tokens = state.tokens + moreTokens
    state.proof_count++
  }
}

app.use(transactionHandler)

app.start().then(appInfo => console.log(appInfo.GCI))
