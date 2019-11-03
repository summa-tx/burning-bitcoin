const express = require('express')
const app = express()
const port = 3000
const bodyParser = require('body-parser')
const cors = require('cors')
require('dotenv').config()

const BcoinClient = require('./bcoinstuff/BcoinClient')
const apiKey = process.env.BCOIN_API_KEY
const host = process.env.BCOIN_HOST
const bcoinPort = 8888

const axios = require('axios')
let { connect } = require('lotion')

const client = new BcoinClient({
  apiKey, host, port: bcoinPort
})

app.use(cors())
app.use(bodyParser.json())
app.use(bodyParser.urlencoded({ extended: true }))

app.get('/', (req, res) => res.send('Hello World!'))

app.post('/getProof', (req, res) => {
  client.getProof(req.body.txid).then((proof) => {
    res.send({proof})
  })
    .catch((e) => {
      console.log('bcoin getProof error', e)
    })
})

app.post('/submitProof', async (req, res) => {
  let GCI = '764b1598b02227cd0fa3c0cc08ae3618a498690b1fc817e4238e978378e3a216'
  let { state, send } = await connect(GCI)
  let myBalance = await state.proof_count
  console.log({ myBalance })

  const proof = req.body.proof
  try {
    await send({ proof })
    let { proofCount, tokens, proofs } = await state
    res.send({ proofCount, tokens, proofs })
  } catch (e) {
    res.send(e)
  }

})

app.listen(port, () => console.log(`Example app listening on port ${port}!`))
