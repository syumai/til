// index.mjs
import { default as cjs } from "./cjs.cjs";
import * as esm from "./esm.mjs";
console.log(cjs.message);
console.log(esm.message);
