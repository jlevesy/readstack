function renderClientError(error) {
    const errElement = document.createElement('p')

    errElement.appendChild(
        document.createTextNode(
            `Failed to communicate with the server : ${error}`
        )
    );

    return errElement;
}

function renderServerError(error) {
    const errElement = document.createElement('p')

    errElement.appendChild(
        document.createTextNode(
            `Server returned an error : ${error.title}`
        )
    );

    return errElement;
}

function renderItem(item) {
    const entry = document.createElement('li');

    entry.appendChild(document.createTextNode(item.name));
    entry.addEventListener('click', (e) => {
        e.preventDefault();
        browser.tabs.create({ url: item.url, active: false});
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
    response.json().then((data) => {
        if (!response.ok) {
            container.appendChild(renderServerError(data));
            return
        }

        container.appendChild(renderItems(data.items));
    });
}).catch((e) => {
    clear(container);
    container.appendChild(renderClientError(e));
});
