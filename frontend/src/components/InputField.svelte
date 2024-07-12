<script lang="ts">
    export let onButtonClick: (text: string) => void;

    let inputValue = "";

    const handleClick = () => {
        if (onButtonClick && typeof onButtonClick === "function") {
            onButtonClick(inputValue);
            inputValue = "";
        }
    };

    const handleKeyPress = (event: KeyboardEvent) => {
        if (event.key === "Enter") {
            onButtonClick(inputValue);
            inputValue = "";
        }
    };
</script>

<main>
    <div id="box">
        <input
            type="text"
            on:keypress={handleKeyPress}
            name="msg"
            bind:value={inputValue}
        />
        <button id="sendBtn" class="button" on:click={handleClick}>SEND</button>
    </div>
</main>

<style>
    #box {
        margin: 5px;
        display: flex;
        align-items: center;
        border: 1px solid #ccc;
        border-radius: 5px;
        padding: 5px;
        background-color: #fff;
    }

    input {
        flex: 1;
        padding: 5px;
        border: none;
        outline: none;
        border-radius: 5px;
        font-size: 1.5rem;
        font-family: "Pacifico", cursive;
        color: #f74040;
    }

    .button {
        text-decoration: none;
        padding: 10px;
        color: white;
        display: flex;
        justify-content: center;
        align-items: center;
        margin: 10px;
        min-width: 30px;
        min-height: 40px;
        border-radius: 25px;
        border: 0;
        background-color: #f74040;
        font-size: 1rem;
        float: right;
    }

    .button:hover {
        background-color: #ff0022;
    }

    .button:active {
        scale: 0.95;
        scrollbar-width: none;
    }
</style>
