<script lang="ts">
    function sendMessage() {
        console.log("Sending message.");
    }
</script>

<main>
    <div id="box">
        <div id="chat"></div>
        <button id="sendBtn" class="button" on:click={sendMessage}>SEND</button>
    </div>
</main>

<style>
    #sendBtn {
        margin: 10px;
        float: right;
    }

    #chat {
        gap: 10px;
        display: flex;
        padding: 30px;
        flex-direction: column;
        justify-content: space-around;
        background-color: #f74040;
        min-width: 80vh;
        min-height: 40vh;
        border-radius: 25px;
    }

    .button {
        text-decoration: none;
        padding: 10px;
        color: white;
        display: flex;
        justify-content: center;
        align-items: center;
        min-width: 30px;
        min-height: 40px;
        border-radius: 25px;
        border: 0;
        background-color: #f74040;
        font-size: 1rem;
    }

    .button:hover {
        background-color: grey;
    }

    .button:active {
        scale: 0.95;
    }
</style>
