import harden from '@agoric/harden';

export default harden((terms, _inviteMaker) => {
  const {contracts} = terms;
  return harden({
    getCommandHandler() {
      return harden({
        processInbound(obj, home) {
          switch (obj.type) {
            case 'burning-bitcoin-2Message': {
              return harden({type: 'burning-bitcoin-2Response', orig: obj, contracts});
            }
          }
          return undefined;
        },
      });
    }
  });
});
