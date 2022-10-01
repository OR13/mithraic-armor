const mithras = require('../index')

;(async () => {
  const message = `🌈 I wear the Armor of Mithras and the Light.`;

  // console.log(message);
  const headers = {
    p: {
      alg: 'ES256',
      kid: '11',
      // need support for this...
      cty: 'application/jose',
    },
    // u: {},
  };
  const signer = {
    key: {
      d: Buffer.from(
          '6c1382765aec5358f117733d281c1c7bdc39884d04a45a1e6c67c858bc206c19',
          'hex',
      ),
    },
  };
  const signed = await mithras.sign.create(headers, message, signer);

  const verifier = {
    key: {
      x: Buffer.from(
          '143329cce7868e416927599cf65a34f3ce2ffda55a7eca69ed8919a394d42f0f',
          'hex',
      ),
      y: Buffer.from(
          '60f7f1a780d8a783bfb7a2dd6b2796e8128dbbcef9d3d168db9529971a36e7b9',
          'hex',
      ),
    },
  };

  const verified = await mithras.sign.verify(signed, verifier);
  console.log('Verified message: ' + verified.toString('utf8'));
})();
