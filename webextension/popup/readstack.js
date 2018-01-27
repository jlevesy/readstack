function handleClientError(error) {
    // TODO
    console.log(error);
}

function handleServerError(response) {
    // TODO
    console.log(response);
}

function renderItem(item) {
    const entry = document.createElement('li');

    entry.appendChild(document.createTextNode(item.name));
    entry.addEventListener('click', (e) => {
        e.preventDefault();
        browser.tabs.create(
            {
                url: item.url,
                active: false
            }
        );
    });

    return entry;
}

function renderItems(items) {
    const list = document.createElement('ul');

    items.forEach((item) => {
        list.appendChild(renderItem(item));
    });

    return list;
}

function clear(node) {
    while (node.hasChildNodes()) {
        node.removeChild(node.lastChild);
    }
}

const container = document.getElementById('container');

fetch('http://localhost:8080/api/v1/item').then((response) => {
    clear(container)

    if (!response.ok) {
        handleError(response);
        return
    }

    response.json().then((data) => {
        container.appendChild(renderItems(data.items));
    });
}).catch((e) => {
    clear(container);
    handleClientError(e);
});
