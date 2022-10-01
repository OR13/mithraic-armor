// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

'use strict';

const path = require('path');

globalThis.require = require;
globalThis.fs = require('fs');
globalThis.TextEncoder = require('util').TextEncoder;
globalThis.TextDecoder = require('util').TextDecoder;

globalThis.performance = {
  now() {
    const [sec, nsec] = process.hrtime();
    return sec * 1000 + nsec / 1000000;
  },
};

const crypto = require('crypto');
globalThis.crypto = {
  getRandomValues(b) {
    crypto.randomFillSync(b);
  },
};

require('../bin/wasm_exec');

const go = new Go();
go.argv = ['hello'];
go.env = Object.assign({TMPDIR: require('os').tmpdir()}, process.env);
go.exit = process.exit;

const run = () => {
  WebAssembly.instantiate(
      fs.readFileSync(path.resolve(__dirname, '../bin/main.wasm')),
      go.importObject,
  )
      .then((result) => {
        process.on('exit', (code) => {
        // Node.js exits if no event handler is pending
          if (code === 0 && !go.exited) {
          // deadlock, make Go print error and stack traces
            go._pendingEvent = {id: 0};
            go._resume();
          }
        });
        return go.run(result.instance);
      })
      .catch((err) => {
        console.error(err);
        process.exit(1);
      });
}

;(async () => {
  console.log('main...');
  run();
})();
