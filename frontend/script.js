
   
function createPlayer() {
  var button = document.getElementById("nameButton");

  var inputName = document.getElementById("playerName").value;
  console.log(inputName)
  var body = {
    name: inputName,
  };
  console.log(body)

  const reqInfo = {
    method: 'POST',
    // mode: "cors", // TODO:  cors should be allowed. Security reasons.
    origin: "http://localhost:3333",
    headers: {
      "content-type": "application/json",
    },
    body: JSON.stringify(body),
  };

  fetch("http://localhost:3333/player/create", reqInfo)
    .then((response) => response.json())
    .then((data) => {
      console.log(data)
      button.style.backgroundColor="green"
    })
    .catch((err) => {
      console.log(err)
      button.style.backgroundColor="red"
    });
}

function PlayerList() {
  fetch("http://localhost:3333/playerlist")
    .then((response) => response.json())
    .then((data) => renderTable(data["list"]))
    .catch((err) => console.log(err));
}

function renderTable(data) {
  document.querySelectorAll("tbody tr").forEach(function(e){e.remove()})
  let pl = document.getElementById("playerlist");

  data.forEach(d => {
    let tr = document.createElement("tr");
    let td = document.createElement("td");
    td.innerText = d.name
    tr.appendChild(td)
    pl.appendChild(tr)
  })

}


// const ws = new WebSocket("ws://localhost:3333/ws/start")
// // Show a connected message when the WebSocket is opened.
// ws.onopen = function(event) {
//   obj = {
//     name: "Lucas",
//     gid: "lucas",
//     pid: "lucas",
//     jumpPosition: {
//       Row: 1,
//       Column: 1,
//     },
//   }
//   j = JSON.stringify(obj)
//   ws.send(j);
// };

// ws.onmessage = function(event) {
//   console.log(event);
// };