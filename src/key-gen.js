const did = require('@or13/did-jwk')

;(async () => {
  const k = await did.generateKeyPair('ES512');
  console.log(k);
})();
