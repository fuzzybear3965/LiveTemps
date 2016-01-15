// *** WEBSOCKETS STUFF BELOW *** //
var ws WebSocket;

window.onload = upgrade;

function upgrade() {
   ws = new WebSocket("ws://localhost")
};

ws.onmessage = function (event) {
   console.log(event.data)
}

// *** PLOTLY STUFF BELOW *** //
plot = document.getElementById("Plotter");

Plotly.plot( 
      plot, 
      [ { x: [1, 2, 3, 4, 5], y: [1, 2, 4, 8, 16] }], 
      { margin: { t : 0 } } 
);


