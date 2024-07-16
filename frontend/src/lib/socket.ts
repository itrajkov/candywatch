import { type ChatMessage } from "./interfaces"
import { handleChatMessage } from "./socket_handlers"

localStorage.debug = '*'

export function startWebsocket(): WebSocket {
    let ws: WebSocket | null = new WebSocket(import.meta.env.VITE_WS_URL)
    ws.binaryType = "arraybuffer"

    ws.onmessage = function (e) {
        var enc = new TextDecoder();
        let str = enc.decode(e.data);
        let [userId, content] = str.split(":")
        let msg = { userId, content } as ChatMessage
        console.log("Server: ", msg)
        handleChatMessage(msg)
    }

    ws.onopen = function () {
        console.log("Connected to socket!", ws?.readyState)
    }

    ws.onclose = function () {
        console.log("Disonnected from socket!")
        ws = null
        setTimeout(startWebsocket, 1000)
    }

    return ws
}


export function sendChatMessage(ws: WebSocket, message: ChatMessage) {
    if (ws == null) {
        return
    }
    if (isOpen(ws)) {
        ws.send(`${message.userId}:${message.content}`);
    } else {
        console.log("Socket is closed.")
    }
}

export function isOpen(ws: WebSocket) { return ws.readyState === ws.OPEN }
