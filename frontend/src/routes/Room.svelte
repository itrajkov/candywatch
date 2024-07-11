<script lang="ts">
    import { room } from "../lib/state";
    import { onMount } from "svelte";
    onMount(() => {
        if (!$room.id) {
            // TODO: Join room and add it to state
            console.log("No room!");
        }
    });
</script>

<main>
    <h1 id="title">Room: {$room.id}</h1>
</main>
