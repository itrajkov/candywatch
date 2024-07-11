<script lang="ts">
    import { createRoom, joinRoom } from "../lib/api";
    import { get } from "svelte/store";
    import { room } from "../lib/state";
    import { push, pop, replace } from "svelte-spa-router";
    import { onMount } from "svelte";

    async function createAndJoin() {
        try {
            let r = await createRoom();
            r = await joinRoom(r.id);
            room.set(r);
            push(`#/rooms/${$room.id}`)
        } catch (error) {
            console.error("Failed to create and join room:", error);
        }
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
