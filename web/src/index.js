var body = document.body;

var formElt = document.createElement("form");
var nameInputElt = document.createElement("input");
var urlInputElt = document.createElement("input");
var envoyerElt = document.createElement("input");

envoyerElt.type = "submit";

formElt.appendChild(nameInputElt);
formElt.appendChild(urlInputElt);
formElt.appendChild(envoyerElt);

body.appendChild(formElt);

formElt.addEventListener("submit", function (e) {
    //e.preventDefault();
    var newItem = new FormData();
    ajaxPost("http://localhost:8081/api/v1/item", {"name":"superpierre", "url":"https://supersite.com"}, function (reponse) {console.log(newItem)}, true);
});

ajaxGet("http://localhost:8081/api/v1/item", function (reponse) {
    var items = JSON.parse(reponse);
    items.items.forEach(function (item) {
        var itemsElt = document.createElement("div");
        var lienElt = document.createElement("a");
        lienElt.textContent = item.name;
        lienElt.href = item.url;
        itemsElt.appendChild(lienElt);
        body.appendChild(itemsElt);
    });
});


function ajaxGet(url, callback) {
    var req = new XMLHttpRequest();
    req.open("GET", url);
    req.addEventListener("load", function () {
        if (req.status >= 200 && req.status < 400) {
            callback(req.responseText);
        } else {
            console.error(req.status + " " + req.statusText + " " + url);
        }
    });
    req.addEventListener("error", function () {
        console.error("Erreur réseau avec l'URL " + url);
    });
    req.send(null);
}

function ajaxPost(url, data, callback, isJson) {
    var req = new XMLHttpRequest();
    req.open("POST", url);
    req.addEventListener("load", function () {
        if (req.status >= 200 && req.status < 400) {
            callback(req.responseText);
        } else {
            console.error(req.status + " " + req.statusText + " " + url);
        }
    });
    req.addEventListener("error", function () {
        console.error("Erreur réseau avec l'URL " + url);
    });
    if (isJson) {
        req.setRequestHeader("Content-Type", "application/json");
        data = JSON.stringify(data);
    }
    req.send(data);
}