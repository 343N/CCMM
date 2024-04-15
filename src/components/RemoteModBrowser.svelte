<script>
  import ModInfoCard from './ModInfoCard.svelte'
  import BrowserCache from './BrowserCache';
  import { Heading } from "flowbite-svelte"
  export let ccModPath
  export let apiKey
  export let apiUserId

  const CC_GAME_ID = 508


  $: modList = BrowserCache.RemoteBrowserCache;
  $: isLoading = true
  $: errorRetrievingData = null

  let getMods = async () => {
    if (BrowserCache.RemoteBrowserCache.length > 0){
      modList = BrowserCache.RemoteBrowserCache
      console.log(BrowserCache.RemoteBrowserCache)
      console.log("loaded mods from cache poggies")
    } else { 
      let res = await fetch(`https://u-${apiUserId}.modapi.io/v1/games/${CC_GAME_ID}/mods?api_key=${apiKey}`)
      console.log(res)
      if (res.status >= 400) {
        errorRetrievingData = res.statusText
        errorRetrievingData = "\n" + await res.text()
        return
      }
      let json = await res.json()
      console.log(json)
      modList = json.data
      BrowserCache.RemoteBrowserCache = modList
      console.log(modList)
    }


    isLoading = false; 
    console.log(isLoading)
  }

  getMods() 



</script>
<div>
  <Heading tag="h4" class="ml-2">Mod List</Heading>
  <p style="display: {isLoading ? "block" : "none"}">Retrieving Data</p>
  <div class="cardListContainer"> 
    
    {#if errorRetrievingData}
      {errorRetrievingData}   
    {:else}
       {#each modList as modInfo}
       <ModInfoCard data={modInfo} remote={true}/>  
       {/each}
    {/if}
       
  </div>

</div>
  
<style>
  .cardListContainer {
    display: flex;
    flex-wrap: wrap;
    flex-grow: 1;
    justify-content: space-between;
    /* padding-right:  */
  }

</style>