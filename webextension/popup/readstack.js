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

fetch('http://localhost:8080/api/v1/item').then((response) => {
    if (!response.ok) {
        handleError(response);
        return
    }

    response.json().then((data) => {
        document.body.appendChild(renderItems(data.items));
    });
}).catch(handleClientError);
