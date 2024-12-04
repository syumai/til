// CommonJS modules側 (index.js)
const {
  currentCount,
  increment,
  decrement,
} = require("./counter.mjs");
console.log(currentCount());
console.log(increment());
console.log(increment());
console.log(decrement());
