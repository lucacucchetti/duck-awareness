<script>
  import {
  Content,Header,SkipToContent
  } from "carbon-components-svelte";
  import "carbon-components-svelte/css/white.css";
  import NoteList from "./lib/NoteList.svelte";

</script>

<Header company="Duck" platformName="Awareness">
  <svelte:fragment slot="skip-to-content">
    <SkipToContent />
  </svelte:fragment>
</Header>

<Content>
  <NoteList />
</Content>
