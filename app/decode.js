const handleClickButton = async () => {
    const results = document.getElementById("result");
    while (results.firstChild) {
        results.firstChild.remove();
    }
    
    const inHex = Uint8Array.fromHex(document.getElementById('inputarea').value.replace(/\s/g, ''));
    const nCols = 16;
    const nRows = Math.floor(inHex.length / nCols) + 1;
    const frames = await wasmGetFrames(inHex);
    countUp();

    function makeRow (data, isHeader, headerCount) {
        const row = document.createElement("tr");
        const hLen = isHeader ? headerCount : data.length;
        let col
        for (let i = 0; i < hLen; i++) {
            if (isHeader) {
                row.setAttribute("id", "hrow");
                col = document.createElement("th");
                col.textContent = i.toString();
            } else {
                col = document.createElement("td");
                col.textContent = ("0" + data[i].toString(16)).slice(-2);;
            }
            row.appendChild(col);
        }
        return row
    }

    function makeHeaders() {
        return makeRow(null, true, nCols);
    }

    for (const idx in frames) {
        if (Object.prototype.hasOwnProperty.call(frames, idx)) {
            const frame = frames[idx];

            const table = document.createElement("table");
            table.appendChild(makeHeaders());

            for (let i = 0; i < nRows * nCols; i+=nCols) {
                console.log(i);
                table.appendChild(makeRow(inHex.slice(i, i + nCols)));
            }

            results.appendChild(table);
            results.appendChild(document.createElement("br"));
            const node = document.createTextNode(JSON.stringify(frame));
            results.appendChild(node);
        }
    }
}