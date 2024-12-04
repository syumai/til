// with.mjs
var obj = { a: 1, b: 2 };

with (obj) {
  console.log(a); // 1
  console.log(b); // 2
  b = 3;
}

console.log(obj.b); // 3
export const a = 1;
