<script>
  import "../app.postcss";
  import { base } from "$app/paths";
  import { Navbar, NavBrand, NavLi, NavUl, DarkMode } from "flowbite-svelte";
</script>

<Navbar>
  <NavBrand>Gopherphis</NavBrand>
  <NavUl>
    <NavLi href="{base}/">Home</NavLi>
    <NavLi href="{base}/api-docs">API Docs</NavLi>
    <NavLi href="{base}/keys">Key Playground</NavLi>
    <NavLi href="https://github.com/dimalinux/gopherphis">GitHub</NavLi>
    <DarkMode class="text-sm" />
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
