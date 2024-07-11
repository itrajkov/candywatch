<script lang="ts">
    import Router from "svelte-spa-router";
    import Home from "./routes/Home.svelte";
    import Room from "./routes/Room.svelte";
    import { startWebsocket } from "./lib/socket";
    import { ws } from "./lib/state";
    import { onMount } from "svelte";

    onMount(() => {
        ws.set(startWebsocket());
    });

    const routes = {
        // Exact path
        "/": Home,

        // Using named parameters, with last being optional
        "/rooms/": Room,

        // Using named parameters, with last being optional
        "/rooms/:id": Room,

        // Catch-all
        // This is optional, but if present it must be the last
        // "*": NotFound,
    };
</script>

<main>
    <Router {routes} />
</main>
