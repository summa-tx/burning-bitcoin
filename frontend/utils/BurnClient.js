const {Client} = require('bcurl');
const assert = require('assert');

class BurnClient extends Client {
  constructor(options) {
    super(options);

    this.on('error', console.log);
  }

  validated(txid) {
    assert(typeof txid === 'string');
    debugger;
    return this.get(`/custom/burn/validated/${txid}`)
  }

  getAddressesByBlockHeight(height) {
    assert((height >>> 0) === height);
    return;
  }
}

(async () => {
  const b = new BurnClient({
    port: 6060
  });

  const res = await b.validated('e77b13569e689afe29272392f6c6c24d86133fc9b901128ce52668df2792c02c')
  console.log(res);

})().catch(e => {
  console.log(e);
});

