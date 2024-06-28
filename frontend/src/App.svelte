<script lang="ts">
 import { socket } from "./lib/socket"
</script>

<main>
  <h1 id="title">candywatch</h1>
  <div id="menu">
    <button class="button">CREATE ROOM</button>
    <button class="button">JOIN ROOM</button>
  </div>
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

 #menu {
   gap: 10px;
   display: flex;
   padding: 30px;
   flex-direction: column;
   justify-content: space-around;
   background-color: red;
   min-width: 30vh;
   min-height: 15vh;
   border-radius: 25px;
 }

 .button {
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
