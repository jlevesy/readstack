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
const itemApiPath = 'api/v1/item';
const container = document.getElementById('container');

function loadItems(host) {
  fetch(`${host}${itemApiPath}`).then((response) => {
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
}

function saveItem(host, item) {
  console.log(JSON.stringify(item));
  fetch(`${host}${itemApiPath}`, {
    method: 'POST',
    body: JSON.stringify(item),
    headers: new Headers({
      'Content-Type': 'application/json'
    })
  }).then(() => {
    lists = document.getElementsByTagName('ul');

    // We're supposed to have exactly one ul
    // brittle, but does the job.
    if (lists.length !== 1) {
      return;
    }

    lists[0].appendChild(renderItem(item));
  }).catch((e) => console.log(e));
}

browser.storage.local.get('readstack-options').then((res) => {
  let host = defaultHost;

  if (res['readstack-options']) {
    host = res['readstack-options']['host'] || defaultHost;
  }

  document.getElementById('options').addEventListener('click', (e) => {
    e.preventDefault();

    browser.runtime.openOptionsPage();
  });

  document.getElementById('add-current').addEventListener('click', (e) => {
    e.preventDefault();

    browser.tabs.query({currentWindow: true, active: true}).then((tabs) => {
      tabs.forEach((tab) => saveItem(host, {name: tab.title, url: tab.url}));
    });
  });

  loadItems(host);
});