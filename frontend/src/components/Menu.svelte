<script lang="ts">
    import { createRoom, joinRoom } from "../lib/api";
    import { startWebsocket } from "../lib/socket";
    let socket = startWebsocket();
    export async function createAndJoin() {
        let room = await createRoom();
        room = await joinRoom(room.id);
        console.log(room)
    }
</script>

<main>
    <div id="menu">
        <button class="button" on:click={createAndJoin}>CREATE ROOM</button>
        <a class="button" href="/room/someid">JOIN ROOM</a>
        <!-- Not implemented -->
    </div>
</main>

<style>
    #menu {
        gap: 10px;
        display: flex;
        padding: 30px;
        flex-direction: column;
        justify-content: space-around;
        background-color: indigo;
        min-width: 30vh;
        min-height: 15vh;
        border-radius: 25px;
    }

    .button {
        text-decoration: none;
        color: black;
        display: flex;
        justify-content: center;
        align-items: center;
        min-width: 30px;
        min-height: 40px;
        border-radius: 25px;
        border: 0;
        background-color: white;
        font-size: 1.3rem;
    }

    .button:hover {
        background-color: grey;
    }

    .button:active {
        scale: 0.95;
    }
</style>
