import { type ChatMessage } from "./interfaces"
import { handleChatMessage } from "./socket_handlers"
import { getCookieValue } from "./util";

localStorage.debug = '*'

export function startWebsocket(): WebSocket {
    let ws: WebSocket | null = new WebSocket('ws://localhost:8080')
    let userId = getCookieValue("session_id")
    ws.binaryType = "arraybuffer"

    ws.onmessage = function (e) {
        console.log("Server: ",)
        var enc = new TextDecoder();
        let content = enc.decode(e.data);
        let msg = { userId, content } as ChatMessage
        handleChatMessage(msg)
    }


    ws.onopen = function () {
        console.log("Connected to socket!")
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
        ws.send(message.content);
    } else {
        console.log("Socket is closed.")
    }
}

export function isOpen(ws: WebSocket) { return ws.readyState === ws.OPEN }
