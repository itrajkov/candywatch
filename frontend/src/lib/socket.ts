localStorage.debug = '*';


export function startWebsocket(): WebSocket {
    let ws: WebSocket | null = new WebSocket('ws://localhost:8080')
    ws.binaryType = "arraybuffer";

    ws.onmessage = function(e){
        console.log("Server: ", e.data)
    }


    ws.onopen = function(){
        console.log("Connected to socket!")
    }

    ws.onclose = function(){
        console.log("Disonnected from socket!")
        ws = null
        setTimeout(startWebsocket, 1000)
    }


    setInterval(() => {
        let message_buffer = new Uint8Array(8)
        message_buffer[0] = 1
        if (ws == null) {
            return
        }
        if (isOpen(ws)) {
            ws.send(message_buffer);
        } else {
            ws
            console.log("Socket is closed.")
        }
    }, 1000)

    return ws
}

function isOpen(ws: WebSocket) { return ws.readyState === ws.OPEN }
