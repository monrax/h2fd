# h2fd - HTTP2 Frame Decoder

Do you have a bunch of HTTP2 frame bytes? Are you looking for a way to decode them easily?

I got you covered.

## What is this?

```ascii
[ hex byte array ] -->  [ this app ] --> [ http2 frame details ]
```

## How to use this?

You can use this app in 3 different ways

Go here:

- [Live website!](https://monrax.github.io/h2fd/) <-- easiest (no need to install anything)

Or, you can run the app locally:

- [CLI app](#cli-app) <-- CLI
- [Node app](#node-app) <-- browser-based GUI

These three alternatives do not depend on each other. You can pick and use the one you like the most.

## CLI app

### Requirements

- Go 1.24+

### Steps

Clone this repo

```sh
https://github.com/monrax/h2fd.git
```

Run the app

```sh
go run .
```

Enter your bytes

```sh
# Enter raw bytes: 00 00 04 00 00 00 01 02 03 be be ca fe
```

See the frame details

```sh
#Raw bytes read: [00 00 04 00 00 00 01 02 03 be be ca fe]

#Frame at index: 0
#Length: 4
#Type: DATA (0)
#Flag: 00000000 (0x0)
#        PADDED: false
#        END_STREAM: false
#R bit: 0b0, Stream ID: 66051 [00 01 02 03]
#Data: [be be ca fe]
```

## Node app

### Requirements

- Go 1.24+
- Node.js v20+

### Steps

Clone this repo

```sh
https://github.com/monrax/h2fd.git
```

Compile to wasm

```sh
GOARCH=wasm GOOS=js go build -o app/main.wasm
```

Copy `wasm_exec.js`

```sh
cp $(go env GOROOT)/lib/wasm/wasm_exec.js ./app/wasm_exec.js
```

Install express.js

```sh
npm install express
```

Start the node server

```sh
node app.js
```

Go to https://localhost:8000/ and use the app!
