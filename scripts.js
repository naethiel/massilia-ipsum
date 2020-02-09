async function handleSubmit(e) {
    e.preventDefault();

    await getSentences();
}

async function getSentences() {
    const content = await generate();

    const container = document.querySelector("#output");

    container.innerHTML = "";

    content.data.forEach(paragraph => {
        let p = document.createElement("p");
        p.innerHTML = paragraph;

        container.appendChild(p);
    });
}

async function generate() {
    const FD = new FormData(document.querySelector("#form"));

    try {
        return await fetch(
            `${window.location.origin}/api/generate.php?length=${FD.get(
                "length"
            )}&size=${FD.get("size")}`
        ).then(r => r.json());
    } catch (err) {
        console.error(err);
    }
}

window.addEventListener("load", async () => await getSentences());