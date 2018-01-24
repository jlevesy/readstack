var body = document.body;

var formElt = document.createElement("form");
formElt.enctype = "application/json";

var nameInputElt = document.createElement("input");
nameInputElt.name = "name";

var urlInputElt = document.createElement("input");
urlInputElt.name = "url";

var envoyerElt = document.createElement("input");
envoyerElt.type = "submit";

formElt.appendChild(nameInputElt);
formElt.appendChild(urlInputElt);
formElt.appendChild(envoyerElt);

body.appendChild(formElt);

fetch('http://localhost:8081/api/v1/item').then(function(response) {
    if(response.ok) {
    response.json().then(function(objectJson) {
        objectJson.items.forEach(function (item) {
            var itemsElt = document.createElement("div");
            var lienElt = document.createElement("a");
            lienElt.textContent = item.name;
            lienElt.href = item.url;
            itemsElt.appendChild(lienElt);
            body.appendChild(itemsElt);
        });
    });
    } else {
      console.log('Mauvaise réponse du réseau');
    }
  })
  .catch(function(error) {
    console.log('Il y a eu un problème avec l\'opération fetch: ' + error.message);
  });

formElt.addEventListener("submit", function (e) {
    var name = formElt.name.value;
    var url = formElt.url.value;
    var newItem = {"name" : name, "url": url};

    fetch("http://localhost:8081/api/v1/item", {
        method: "POST",
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(newItem)
    })
    .then( response =>  response.json() );
});