// exit-from-strict-1.mjs
globalThis.obj = {
  message: "Hello, World!",
};

// これは動かない
eval(`
  with(obj) {
    console.log(message);
  }
`);
