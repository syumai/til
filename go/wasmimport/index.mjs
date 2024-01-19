import "./wasm_exec.js";
import fs from "node:fs";

const go = new Go();
const result = await WebAssembly.instantiate(fs.readFileSync("./wasmimport"), {
  ...go.importObject,
  math: {
    add(a, b) {
      return a + b;
    },
  },
});
await go.run(result.instance);
