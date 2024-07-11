import { type ChatMessage } from "./interfaces";
import { chatMessages } from './state';
import { get } from 'svelte/store'

export function handleChatMessage(message: ChatMessage){
    chatMessages.update(currentMessages => [
        ...currentMessages,
        message
    ]);
    console.log("Current messages:", get(chatMessages));
}
