<script>
  import "../app.postcss";
  import {
    DarkMode,
    Img,
    Navbar,
    NavBrand,
    NavHamburger,
    NavLi,
    NavUl,
  } from "flowbite-svelte";
  import { base } from "$app/paths";
  import { page } from "$app/stores";

  $: activeUrl = $page.url.pathname;
  const activeClass =
    "text-white bg-green-700 md:bg-transparent md:text-green-700 md:dark:text-white dark:bg-green-600 md:dark:bg-transparent p-0";
  const nonActiveClass =
    "text-gray-700 hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-green-700 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent p-0";
</script>

<Navbar>
  <NavBrand class="text-2xl font-bold dark:text-yellow-400">
    <Img
      width="30"
      src="{base}/seraphim-gopher.jpeg"
      alt="Seraphim Gopher"
      title="Seraphim Gopher"
    />
    <span class="pl-2">Gopherphis</span>
  </NavBrand>
  <NavHamburger />
  <NavUl {activeUrl} {activeClass}>
    <NavLi href="{base}/">Home</NavLi>
    <NavLi href="{base}/keys">Key Playground</NavLi>
    <DarkMode class="text-sm p-0" />
  </NavUl>
</Navbar>
<slot />

<style>
  :global(:root) {
    color-scheme: dark light;
  }

  :global(html, body) {
    height: 100%;
    width: 100%;
  }
</style>
