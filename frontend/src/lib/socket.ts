localStorage.debug = '*';
export const socket = new WebSocket("ws://localhost:8080") // TODO: Extract into env var

socket.binaryType = "arraybuffer";

// Events
// socket.addEventListener("open", (_) => {});

socket.addEventListener("message", (event) => {
    console.log("Server: ", event.data)
});

function isOpen(ws: WebSocket) { return ws.readyState === ws.OPEN }

// setInterval(() => {
//     let message_buffer = new Uint8Array(8)
//     message_buffer[0] = 1
//     if (isOpen(socket)) {
//         socket.send(message_buffer);
//     } else {
//         console.log("Socket is closed.")
//     }
// }, 1000)
