<script>
  import { Tabs, TabItem } from 'flowbite-svelte'
  import { AppStore } from '../AppStore.js'
  import '../components/Config.svelte'
  import Config from '../components/Config.svelte';
  import RemoteModBrowser from '../components/RemoteModBrowser.svelte';

  export let visible = true

  $: has_config = false;

  let checkHasConfig = async () => {
    let apiKey = await AppStore.get("apiKey")
    let apiUserId = await AppStore.get("apiUserId")
    let ccModPath = await AppStore.get("ccModPath")
    has_config = apiKey && apiUserId && ccModPath

    
  }
  checkHasConfig()

</script>

{#if visible}
<div>

  <Tabs style="underline">
    <TabItem open title="Browse Mods">
      {#if !has_config}
      You need to fill out your settings!
      {:else}
      <RemoteModBrowser/>
      {/if}

    </TabItem>
    <TabItem title="Installed Mods">
      {#if !has_config}
      You need to fill out your settings!
      {:else}
      <RemoteModBrowser/>
      {/if}

    </TabItem>
    <TabItem title="Settings">
      <Config on:save={checkHasConfig}/>
    </TabItem>
  </Tabs>
</div>
{/if}


<style>
  /* .tabFixes {
    /* align-items: center; */
    /* flex-wrap: nowrap !important; */
  /* } */
  div {
    padding: 1rem;
  }

</style>