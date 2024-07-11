<script lang="ts">
    import { createRoom, joinRoom } from "../lib/api";
    import { room } from "../lib/state";
    import { push } from "svelte-spa-router";

    async function handleCreate() {
        try {
            let r = await createRoom();
            r = await joinRoom(r.id);
            console.log(r)
            room.set(r)
            push(`#/rooms/${r.id}`);
        } catch (error) {
            console.error("Failed to create and join room:", error);
        }
    }

    async function handleJoin() {
        console.error("Not implemented.");
    }
</script>

<main>
    <div id="menu">
        <button class="button" on:click={handleCreate}>Create Room</button>
        <!-- Not implemented -->
        <button class="button" on:click={handleJoin}>Join Room</button>
    </div>
</main>

<style>
    #menu {
        gap: 10px;
        display: flex;
        padding: 30px;
        flex-direction: column;
        justify-content: space-around;
        min-width: 30vh;
        min-height: 15vh;
        border-radius: 25px;
    }

    .button {
        text-decoration: none;
        display: flex; justify-content: center;
        align-items: center;
        min-width: 40vh;
        min-height: 60px;
        border-radius: 25px;
        border: 0;
        background-color: #f74040;
        color: white;
        font-size: 2rem;
        font-family: 'Pacifico';
    }

    .button:hover {
        background-color: #ff0022;
    }

    .button:active {
        scale: 0.95;
    }
</style>
