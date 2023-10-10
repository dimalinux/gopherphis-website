<script>
  import { base } from "$app/paths";
  import Paragraph from "./Paragraph.svelte";
  import SectionHeader from "./SectionHeader.svelte";
  import URL from "./URL.svelte";

  const apiDocs =
    "https://pkg.go.dev/github.com/dimalinux/gopherphis#section-directories";
  const gopherphisSources = "https://github.com/dimalinux/gopherphis";
  const websiteSources = gopherphisSources + "-website";
</script>

<div class="mx-auto max-w-prose m-4">
  <SectionHeader>Project</SectionHeader>
  <Paragraph>
    Gopherphis provides golang libraries for Monero with a focus on Seraphis
    support. By the time Seraphis is released, the project will, at minimum,
    have a full CLI wallet.
  </Paragraph>

  <SectionHeader>Website</SectionHeader>
  <Paragraph>
    This website provides a Seraphis <URL href="{base}/keys"
      >key testing page</URL
    >, inspired by the
    <URL href="https://xmr.llcoins.net/addresstests.html"
      >Cryptonote key tester</URL
    >, to help other projects when transitioning their code to Seraphis.
  </Paragraph>

  <SectionHeader>Source Code</SectionHeader>
  <Paragraph>
    Go libraries and future wallet:<br />
    <URL href={gopherphisSources} />
  </Paragraph>
  <Paragraph>
    API docs:<br />
    <URL href={apiDocs} />
  </Paragraph>
  <Paragraph>
    Website, which generates web assembly from the primary repository for the
    keys page:<br />
    <URL href={websiteSources} />
  </Paragraph>
</div>
