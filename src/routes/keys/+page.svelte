<!-- svelte:head>
    <script src="./wasm_exec.js" defer></script>
    <script src="./gopherphis.js" defer></script>
</svelte:head -->

<script>
    import {} from './wasm_exec.js';
    import {} from './gopherphis.js';

    function clickHandler() {
        window.submitPhrase();
    }
</script>

<div>
    <textarea id="phrase" rows="10" cols="30" placeholder="Enter a mnemonic phrase"></textarea>
    <button type="button" on:click="{clickHandler}">Submit</button>
    <textarea id="key" rows="10" cols="30"></textarea>
</div>
