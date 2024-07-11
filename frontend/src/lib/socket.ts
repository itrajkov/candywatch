localStorage.debug = '*';
export const socket = new WebSocket(import.meta.env.VITE_WS_URL)

socket.binaryType = "arraybuffer";

socket.addEventListener("open", (_) => {console.log("Connected to websocket.")});

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
