<script lang="ts">
    import { afterUpdate } from "svelte";
    import { sendChatMessage, startWebsocket } from "../lib/socket";
    import { chatMessages } from "../lib/state";
    import { getCookieValue } from "../lib/util";
    import ChatMessage from "./ChatMessage.svelte";
    import InputField from "./InputField.svelte";
    let ws = startWebsocket();
    function sendMessage(text: string) {
        let msg = {
            userId: getCookieValue("session_id") || "",
            content: text,
        };
        sendChatMessage(ws, msg);
        chatMessages.update((currentMessages) => [...currentMessages, msg]);
    }

    let chat: HTMLDivElement;
    afterUpdate(() => {
        chat.scrollTop = chat.scrollHeight;
    });
</script>

<main>
    <div id="box">
        <div id="chat" bind:this={chat}>
            <ul id="chat-messages">
                {#each $chatMessages as message}
                    <ChatMessage {message} />
                {/each}
            </ul>
        </div>
        <InputField onButtonClick={sendMessage}/>
    </div>
</main>

<style>
    #chat {
        margin: 10px;
        gap: 10px;
        display: flex;
        padding-bottom: 30px;
        flex-direction: column;
        background-color: #f74040;
        min-height: 70vh;
        max-height: 70vh;
        border-radius: 25px;
        overflow-y: scroll;
        scrollbar-width: none;
    }

    #chat-messages {
        margin: 0;
        padding: 40px;
    }
</style>
