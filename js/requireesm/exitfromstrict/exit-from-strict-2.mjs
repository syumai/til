// exit-from-strict-2.mjs
globalThis.obj = {
  message: "Hello, World!",
};

// これは動く
(0, eval)(`
  with(obj) {
    console.log(message);
  }
`);
