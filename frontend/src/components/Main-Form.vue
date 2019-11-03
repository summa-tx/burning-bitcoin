<template>
  <section class="">
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
                {{ proof }}
              </v-flex>
              <v-layout justify-center>
                <v-btn
                  :disabled="validProof"
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
import BcoinClient from '../../../utils/BcoinClient'
const apiKey = process.env.BCOIN_API_KEY
const host = process.env.BCOIN_HOST
const port = process.env.BCOIN_PORT

const client = new BcoinClient({
  apiKey, host, port
})

export default {
  name: 'MainForm',

  components: {
    ClickToCopy: () => import(/* webpackChunkName: 'Click to Copy' */ './Click-To-Copy')
  },

  mounted () {
    console.log(test.thing())
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
      proof: null,
      validTxid: true,
      validProof: false,
      rules: {
        txid: [v => !!v || 'TXID is required'],
        headers: [v => v < 101 || 'Max headers is 100']
      }
    }
  },

  methods: {
    handleCollectProof () {
      console.log('collect proof')
      // client.getProof(this.txid).then((proof) => {
      //   this.proof = proof
      //   console.log({ proof })
      //   this.validProof = true
      // })
      // .catch((e) => {
      //   console.log('bcoin getProof error', e)
      // })
    },

    handleSubmitProof: () => {
      console.log('submit proof')
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
