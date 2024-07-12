<script lang="ts">
    import Menu from "../components/Menu.svelte";
    import {startWebsocket} from "../lib/socket";
    startWebsocket();
</script>

<main>
    <h1 id="title">Candywatch</h1>
    <Menu />
</main>

<style>
    main {
        display: flex;
        flex-direction: column;
        text-align: center;
    }

    #title {
        user-select: none;
        -webkit-tap-highlight-color: transparent;
        -webkit-user-drag: none;
        margin: 10;
        font-family: "Pacifico", cursive;
        font-size: 90px;
        color: #f74040;
        text-shadow: 5px 5px #ffffff;
        text-align: center;
    }
</style>
