<script>
  import { Card, Label, Textarea, Button } from "flowbite-svelte";
  import {} from "./wasm_exec.js";
  import { submitPhrase } from "./gopherphis.js";
  import MnemonicSelector from "./MnemonicSelector.svelte";

  function clickHandler() {
    submitPhrase();
  }
</script>

<div class="flex justify-center">
  <Card class="w-full max-w-md">
    <Label class="space-y-2">
      <MnemonicSelector/>
      <Textarea
        id="phrase"
        rows="10"
        cols="30"
        placeholder="Enter a mnemonic phrase"
      />
    </Label>
    <Button type="button" class="w-full" on:click={clickHandler}>Submit</Button>
  </Card>
  <Card>
    <Label class="space-y-2">
      <span>Private Spend Key</span>
      <Textarea id="key" rows="10" cols="30" />
    </Label>
  </Card>
</div>
