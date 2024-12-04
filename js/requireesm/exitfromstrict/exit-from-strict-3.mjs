// exit-from-strict-3.mjs
globalThis.obj = {
  message: "Hello, World!",
};

/*
function f() {
  // ここはModuleなので構文エラーになる
  with(obj) {
    console.log(message);
  }
}
f();
*/

// これは動く
new Function(`
  console.log(import("hoge.js"));
  with(obj) {
    console.log(message);
  }
`)();
