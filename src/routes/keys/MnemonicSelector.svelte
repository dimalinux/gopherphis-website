<script>
  import { Button, Dropdown, DropdownItem } from "flowbite-svelte";
  import { ChevronDownSolid } from "flowbite-svelte-icons";

  import { page } from "$app/stores";
  $: activeUrl = $page.url.pathname;
  const activeClass =
    "text-green-700 dark:text-green-300 hover:text-green-900 dark:hover:text-green-500";

  const cryptonoteMnemonic = "cryptonote_mnemonic";
  const polyseedMnemonic = "polyseed_mnemonic";

  const descriptions = new Map([
    [cryptonoteMnemonic, "Cryptonote Mnemonic Phrase"],
    [polyseedMnemonic, "Polyseed Mnemonic Phrase"],
  ]);
  let selectedMnemonic = "";
  selectMnemonicType(cryptonoteMnemonic);

  /**
   * @param {string} mType - mnemonic type value
   */
  function selectMnemonicType(mType) {
    selectedMnemonic = mType;
    window.selectedMnemonic = mType;
  }

  /**
   * @return {boolean} whether the selected mnemonic is cryptonote
   */
  export function isCryptonoteMnemonic() {
    return selectedMnemonic === cryptonoteMnemonic;
  }

  isCryptonoteMnemonic();
</script>

<Button>
  {descriptions.get(selectedMnemonic)}
  <ChevronDownSolid class="ml-2 text-white dark:text-white" />
</Button>
<Dropdown {activeUrl} {activeClass}>
  {#each [...descriptions] as [key, description]}
    <DropdownItem on:click={() => selectMnemonicType(key)}>
      {description}
    </DropdownItem>
  {/each}
</Dropdown>
