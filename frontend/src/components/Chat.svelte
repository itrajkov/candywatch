<script lang="ts">
    import ChatMessage from "./ChatMessage.svelte";
    let messages = [
        { userId: "ivche", content: "hello world!" },
        { userId: "ivche", content: "hello world!" },
    ];
    function sendMessage() {
        messages = [...messages, { userId: "ivche", content: "hello world!" }];
    }
</script>

<main>
    <div id="box">
        <div id="chat">
            <ul id="chat-messages">
                {#each messages as message}
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
</style>
