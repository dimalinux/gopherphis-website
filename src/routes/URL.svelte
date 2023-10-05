<script>
  import { A } from "flowbite-svelte";
  export let href = "";
</script>

<A class="text-blue-500" {href}><slot>{href}</slot></A>
