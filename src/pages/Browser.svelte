<script>


  import { Tabs, TabItem } from 'flowbite-svelte'
  import { AppStore } from '../AppStore.js'
  import '../components/Config.svelte'
  import Config from '../components/Config.svelte';
  import RemoteModBrowser from '../components/RemoteModBrowser.svelte';

  export let visible = true

  $: has_config = false;

  let apiKey
  let apiUserId
  let ccModPath

  let checkHasConfig = async () => {
    apiKey = await AppStore.get("apiKey")
    apiUserId = await AppStore.get("apiUserId")
    ccModPath = await AppStore.get("ccModPath")
    console.log(`Checking config! '${apiKey}', '${apiUserId}', '${ccModPath}'`)
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
      <RemoteModBrowser {apiKey} {apiUserId} {ccModPath}/>
      {/if}

    </TabItem>
    <TabItem title="Installed Mods">
      {#if !has_config}
      You need to fill out your settings!
      {:else}
      TODO: add a browser lel
      <!-- <RemoteModBrowser/> -->
      {/if}

    </TabItem>
    <TabItem title="Settings">
      <Config on:save={checkHasConfig}/>
    </TabItem>
  </Tabs>
</div>
{/if}


<style>
  div {
    padding: 1rem;
  }

</style>