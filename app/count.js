const BASE_URL = "https://api.counterapi.dev/v1/monrax/h2fd";


async function updateCount(up) {
    const resp = await fetch(BASE_URL + (up ? "/up" : ""));
    if (resp.ok) {
        resp.json().then(j => document.getElementById("count").textContent = "Frames decoded: " + j.count);
    }
    return resp.ok;
}

async function countUp() {
    return updateCount(true);
}

updateCount();