localStorage.debug = '*';
export const socket = new WebSocket("ws://localhost:8080")

socket.binaryType = "arraybuffer";

// Events
socket.addEventListener("open", (event) => {
    let message_buffer = new Uint16Array(8)
    message_buffer[0] = 10
    socket.send(message_buffer);
});

socket.addEventListener("message", (event) => {
    console.log("Message: ", event.data)
});
