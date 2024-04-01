<script>
  import { Tabs, TabItem } from 'flowbite-svelte'
  import '../components/Config.svelte'
  import Config from '../components/Config.svelte';

  export let visible = true
</script>

{#if visible}
<div>

  <Tabs style="underline">
    <TabItem open title="Browse Mods">
      Loading...
    </TabItem>
    <TabItem title="Installed Mods">
      You need to specify a mods directory!
    </TabItem>
    <TabItem title="Settings">
      <Config />
    </TabItem>
  </Tabs>
</div>
{/if}


<style>
  div {
    padding: 1rem;
  }
</style>