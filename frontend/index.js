const express = require('express')
const app = express()
const port = 3000
const bodyParser = require('body-parser')
const cors = require('cors')

const BcoinClient = require('./bcoinstuff/BcoinClient')
const apiKey = "the voice of the bell at gion shrine"
const host = "157.245.167.26"
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
    console.log({ raw: proof.pretty })
    res.send({proof})
  })
    .catch((e) => {
      console.log('bcoin getProof error', e)
    })
})

app.post('/submitProof', async (req, res) => {
  let GCI = '07296574bddca30975ffabb5103079a6755667dead4c9448a1f5793ebdd5c4d2'
  let { state, send } = await connect(GCI)
  let myBalance = await state.proof_count
  console.log({ myBalance })

  const proof = JSON.stringify(req.body.proof)
  console.log(await send({ proof }))
})

app.listen(port, () => console.log(`Example app listening on port ${port}!`))
