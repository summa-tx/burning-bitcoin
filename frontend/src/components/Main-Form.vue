<template>
  <section class="">
    {{ proofs }}
    <v-layout column justify-center align-center>
      <v-container>
        <h2>Instructions:</h2>
        <p>
          <ol>
            <li>Send Bitcoin to the address below to burn Bitcoin.</li>
            <li>Collect Proof.</li>
            <li>Submit Proof.</li>
          </ol>
        </p>
        <v-layout column class="input-section">
          <v-flex>
            <h3>Burn Address:</h3>
          </v-flex>
          <Click-To-Copy
            input-value="0x924852n4h52o2030283-5902dd85-9028j304823oim4iounyoyojiu9yu5928t"
            :is-read-only="true"
            :input="true"
            :crop="true"
            input-margin="ma-1"
            item-id="burnAddress"
          ></Click-To-Copy>
        </v-layout>
        <v-layout column class="input-section">
          <v-form v-model="validTxid">
            <v-layout column>
              <v-flex>
                <h3>Collect Proof</h3>
              </v-flex>
              <v-flex>
                <v-text-field
                  v-model="txid"
                  label="Enter your TXID"
                  :rules="rules.txid"
                  required
                ></v-text-field>
              </v-flex>
              <v-flex>
                <v-text-field
                  v-model="numOfHeaders"
                  label="Select number of headers"
                  type="number"
                  max="100"
                  :rules="rules.headers"
                  required
                ></v-text-field>
              </v-flex>
              <v-layout justify-center >
                <v-btn
                  :disabled="!validTxid"
                  @click="handleCollectProof"
                >Go</v-btn>
              </v-layout>
            </v-layout>
          </v-form>
        </v-layout>
        <v-layout column class="input-section">
          <v-form v-model="validProof">
            <v-layout column>
              <v-flex>
                <h3>Submit Proof</h3>
              </v-flex>
              <v-flex>
                {{ }}
              </v-flex>
              <v-container>
                <v-row dense>
                  <v-col sm="2" class="pt-5">Version</v-col>
                  <v-col sm="10">
                    <v-text-field
                      readonly
                      class="proof-item"
                      :placeholder="proof.version"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row dense>
                  <v-col sm="2" class="pt-5">VIN</v-col>
                  <v-col sm="10">
                    <v-textarea
                      readonly
                      class="proof-item"
                      :placeholder="proof.vin"
                    ></v-textarea>
                  </v-col>
                </v-row>
                <v-row dense>
                  <v-col sm="2" class="pt-5">VOUT</v-col>
                  <v-col sm="10">
                    <v-textarea
                      readonly
                      class="proof-item"
                      :placeholder="proof.vout"
                    ></v-textarea>
                  </v-col>
                </v-row>
                <v-row dense>
                  <v-col sm="2" class="pt-5">Locktime</v-col>
                  <v-col sm="10">
                    <v-text-field
                      disabled
                      class="proof-item"
                      :placeholder="proof.locktime"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row dense>
                  <v-col sm="2" class="pt-5">TXID</v-col>
                  <v-col sm="10">
                    <v-text-field
                      disabled
                      class="proof-item"
                      :placeholder="proof.tx_id"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row dense>
                  <v-col sm="2" class="pt-5">TXID LE</v-col>
                  <v-col sm="10">
                    <v-text-field
                      disabled
                      class="proof-item"
                      :placeholder="proof.tx_id_le"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row dense>
                  <v-col sm="2" class="pt-5">Index</v-col>
                  <v-col sm="10">
                    <v-text-field
                      disabled
                      class="proof-item"
                      type="number"
                      :placeholder="proof.index"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row dense>
                  <v-col sm="2" class="pt-5">Intermediate Nodes</v-col>
                  <v-col sm="10">
                    <v-textarea
                      disabled
                      class="proof-item"
                      :placeholder="proof.intermediate_nodes"
                    ></v-textarea>
                  </v-col>
                </v-row>
                <v-row dense>
                  <v-col><h4>Confirming Header</h4></v-col>
                  <v-col></v-col>
                </v-row>
                <v-row dense>
                  <v-col sm="2" class="pt-5">Raw</v-col>
                  <v-col sm="10">
                    <v-text-field
                      disabled
                      class="proof-item"
                      :placeholder="proof.confirming_header.raw"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row dense>
                  <v-col sm="2" class="pt-5">Hash</v-col>
                  <v-col sm="10">
                    <v-text-field
                      disabled
                      class="proof-item"
                      :placeholder="proof.confirming_header.hash"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row dense>
                  <v-col sm="2" class="pt-5">Hash LE</v-col>
                  <v-col sm="10">
                    <v-text-field
                      disabled
                      class="proof-item"
                      :placeholder="proof.confirming_header.hash_le"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row dense>
                  <v-col sm="2" class="pt-5">Height</v-col>
                  <v-col sm="10">
                    <v-text-field
                      disabled
                      class="proof-item"
                      :placeholder="proof.confirming_header.height"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row dense>
                  <v-col sm="2" class="pt-5">PrevHash</v-col>
                  <v-col sm="10">
                    <v-text-field
                      disabled
                      class="proof-item"
                      :placeholder="proof.confirming_header.prevhash"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row dense>
                  <v-col sm="2" class="pt-5">Merkle Root</v-col>
                  <v-col sm="10">
                    <v-text-field
                      disabled
                      class="proof-item"
                      :placeholder="proof.confirming_header.merkle_root"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row dense>
                  <v-col sm="2" class="pt-5">Merkle Root LE</v-col>
                  <v-col sm="10">
                    <v-text-field
                      disabled
                      class="proof-item"
                      :placeholder="proof.confirming_header.merkle_root_le"
                    ></v-text-field>
                  </v-col>
                </v-row>
              </v-container>
              <v-layout justify-center>
                <v-btn
                  :disabled="!validProof"
                  @click="handleSubmitProof"
                >Submit Proof</v-btn>
              </v-layout>
            </v-layout>
          </v-form>
        </v-layout>
      </v-container>
    </v-layout>
  </section>
</template>

<script>
import axios from 'axios'
import { makeMint } from '@agoric/ertp/core/mint'

export default {
  name: 'MainForm',

  components: {
    ClickToCopy: () => import(/* webpackChunkName: 'Click to Copy' */ './Click-To-Copy')
  },

  async mounted () {
    const burntBitsMint = makeMint('burntBits')
    this.burntBitsMint = burntBitsMint
    const assay = burntBitsMint.getAssay()
    this.assay = assay
    this.tokens = assay.makeEmptyPurse()
    this.myTokens = assay.makeEmptyPurse()
    this.everyoneElseTokens = assay.makeEmtpyPurse()
  },

  computed: {
    headerRange: function() {
      let range = []
      for (let i = 1; i <= 100; i++) {
        range.push(i)
      }
      return range
    }
  },

  data: () => {
    return {
      txid: 'eeab514286b0968957b875faa9486945e76bec9d6d1fa4762351f2cfe6a2aa8d',
      numOfHeaders: null,
      proof: {
        version: null,
        vin: null,
        vout: null,
        locktime: null,
        tx_id: null,
        tx_id_le: null,
        index: null,
        intermediate_nodes: null,
        confirming_header: {
          raw: null,
          hash: null,
          hash_le: null,
          height: null,
          prevhash: null,
          merkle_root: null,
          merkle_root_le: null,
        }
      },
      validTxid: true,
      validProof: false,
      rules: {
        txid: [v => !!v || 'TXID is required'],
        headers: [v => v < 101 || 'Max headers is 100']
      },
      rawProof: undefined,
      burntBitsMint: null,
      totalBits: null,
      proofs: [],
      tokens: 0, // counter
      myTokens: null, // purse
      assay: null,
      everyoneElseTokens: null // purse
    }
  },

  methods: {
    handleCollectProof () {
      console.log('collect proof')
      axios
        .post('http://localhost:3000/getProof', { txid: this.txid })
        .then((res) => {
          console.log({res})
          this.proof.version = res.data.proof.pretty.version
          this.proof.vin = res.data.proof.pretty.vin
          this.proof.vout = res.data.proof.pretty.vout
          this.proof.locktime = res.data.proof.pretty.locktime
          this.proof.tx_id = res.data.proof.pretty.tx_id
          this.proof.tx_id_le = res.data.proof.pretty.tx_id_le
          this.proof.index = res.data.proof.pretty.index
          this.proof.confirming_header.raw = res.data.proof.pretty.confirming_header.raw
          this.proof.confirming_header.hash = res.data.proof.pretty.confirming_header.hash
          this.proof.confirming_header.hash_le = res.data.proof.pretty.confirming_header.hash_le
          this.proof.confirming_header.height = res.data.proof.pretty.confirming_header.height
          this.proof.confirming_header.prevhash = res.data.proof.pretty.confirming_header.prevhash
          this.proof.confirming_header.merkle_root = res.data.proof.pretty.confirming_header.merkle_root
          this.proof.confirming_header.merkle_root_le = res.data.proof.pretty.confirming_header.merkle_root_le
          this.proof.intermediate_nodes = res.data.proof.pretty.intermediate_nodes
          console.log(this.proof.version)
          this.rawProof = res.data.proof.raw
          this.validProof = true
        })
        .catch((err) => {
          console.log({err})
          this.validProof = false
        })
    },

    handleSubmitProof () {
      console.log('submit proof')
      axios.post('http://localhost:3000/submitProof', { proof: this.proof})
        .then((res) => {
          console.log({res})
          this.proofs = res.data.proofs
          this.makeBurntBits()
        })
    },

    makeBurntBits () {
      // Randomly generate a number of tokens
      const decimals = parseInt(Math.random() * 10) * 2
      const randomNumOfTokens = parseInt(Math.random() * 1000 * decimals)

      // Mint random number of tokens
      const purse = this.burntBitsMint.mint(randomNumOfTokens)
      this.tokens += randomNumOfTokens
      // Give some to me
      const payment = purse.withdraw(0.88 * randomNumOfTokens)

      this.myTokens.depositAll(payment)
    }
  }
}
</script>

<style lang='stylus' scoped>
.input-section
  padding 20px 30px
  margin 20px 0
  border 2px solid #e3e3e3
  border-radius 10px

</style>
