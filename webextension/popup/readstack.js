function renderError(message) {
    const errElement = document.createElement('p');

    errElement.appendChild(
        document.createTextNode(message)
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

const defaultHost = 'http://localhost:8080/';
const container = document.getElementById('container');

browser.storage.local.get('readstack-options').then((res) => {
    let host = defaultHost;

    if (res['readstack-options']) {
        host = res['readstack-options']['host'] || defaultHost;
    }

    fetch(`${host}api/v1/item`).then((response) => {
        clear(container)
        response.json().then((data) => {
            if (!response.ok) {
                container.appendChild(
                    renderError(data['message'])
                );
                return
            }

            container.appendChild(renderItems(data.items));
        }).catch((e) => {
            clear(container);
            container.appendChild(
                renderError(`Failed to decode JSON response ${e}`)
            );
        });
    }).catch((e) => {
        clear(container);
        container.appendChild(
            renderError(`Failed to perform request ${e}`)
        );
    });
});
