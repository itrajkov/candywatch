<script lang="ts">
    import { afterUpdate } from "svelte";
    import { sendChatMessage, startWebsocket } from "../lib/socket";
    import { chatMessages } from "../lib/state";
    import { getCookieValue } from "../lib/util";
    import ChatMessage from "./ChatMessage.svelte";
    let ws = startWebsocket();
    function sendMessage() {
        let msg = {
            userId: getCookieValue("session_id") || "",
            content: "hello world!",
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
        <!-- TODO: Move this button to an input field component-->
        <button id="sendBtn" class="button" on:click={sendMessage}>SEND</button>
    </div>
</main>

<style>
    #chat {
        gap: 10px;
        display: flex;
        padding-bottom: 30px;
        flex-direction: column;
        background-color: #f74040;
        min-width: 80vh;
        min-height: 40vh;
        max-height: 40vh;
        border-radius: 25px;
        overflow-y: scroll;
        scrollbar-width: none;
    }

    .button {
        text-decoration: none;
        padding: 10px;
        color: white;
        display: flex;
        justify-content: center;
        align-items: center;
        margin: 10px;
        min-width: 30px;
        min-height: 40px;
        border-radius: 25px;
        border: 0;
        background-color: #f74040;
        font-size: 1rem;
        float: right;
    }

    .button:hover {
        background-color: #ff0022;
    }

    .button:active {
        scale: 0.95;
        scrollbar-width: none;
    }
</style>
