<script lang="ts">
    import { room } from "../lib/state";
    import { joinRoom } from "../lib/api";
    import { onMount } from "svelte";
    import { pop } from "svelte-spa-router";
    export let params = { id: null };

    onMount(async () => {
        if (!$room.id) {
            // TODO: Validate roomId here instead in api.ts
            console.log("Joining room..");
            try {
                let r = await joinRoom(params.id || "");
                console.log(r);
                room.set(r);
            } catch(err) {
                console.error("Couldn't join room:", err)
                pop();
            }
        }
    });
</script>

<main>
    <h1 id="title">Room: {params.id || $room.id}</h1>
</main>
