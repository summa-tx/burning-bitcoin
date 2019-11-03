let lotion = require('lotion')
const {ValidateSPV} = require('@summa-tx/bitcoin-spv-js')

let app = lotion({
  initialState: {
    proof_count: 0
  }
})

function transactionHandler (state, transaction) {
  const proof = JSON.parse(transaction.proof)
  if (ValidateSPV.validateProof(proof)) {
    state.proof_count++
  }
}

app.use(transactionHandler)

app.start().then(appInfo => console.log(appInfo.GCI))
