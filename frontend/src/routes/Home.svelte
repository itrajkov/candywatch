<script lang="ts">
    //import { socket } from "./lib/socket";
    import Menu from '../components/Menu.svelte'
</script>

<main>
    <h1 id="title">candywatch</h1>
    <Menu/>
</main>

<style>
    main {
        display: flex;
        flex-direction: column;
        text-align: center;
    }

    #title {
        font-size: 3rem;
    }

</style>
