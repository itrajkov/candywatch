<script lang="ts">
    import { type ChatMessage } from "../lib/interfaces";
    function sendMessage() {
        console.log("Sending message.");
    }
    export let message: ChatMessage;
</script>

<main>
    <li id="chat-message">{message.userId}: {message.content}</li>
</main>

<style>
 #chat-message{
     font-size: 1.7rem;
     list-style: none;
     color: white;
 }
</style>
