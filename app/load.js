const go = new Go();

const resp = fetch("main.wasm", {mode: "no-cors"});

WebAssembly.instantiateStreaming(resp, go.importObject).then((result) => {
    go.run(result.instance);
});