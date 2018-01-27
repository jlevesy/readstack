document.addEventListener('DOMContentLoaded', () => {
  browser.storage.local.get('readstackOptions').then((res) => {
    document.getElementById('readstack-host').value = res['host'] || 'http://localhost:8080';
  });
});

document.getElementById("readstack-options").addEventListener("submit", (e) => {
    e.preventDefault();

    value = document.getElementById('readstack-host').value

    try {
        url = new URL(value);
    } catch(_) {
        return;
    }

    readstackOptions = {
        host: url.href
    };

    browser.storage.local.set({'readstack-options': readstackOptions}).catch((err) => console.log(err));
});
