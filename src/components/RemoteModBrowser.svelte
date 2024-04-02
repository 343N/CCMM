<script>
  import ModInfoCard from './ModInfoCard.svelte'
  export let ccModPath
  export let apiKey
  export let apiUserId

  const CC_GAME_ID = 508

  $: modList = []

  let getMods =  async () => {
    let res = await fetch(`https://u-${apiUserId}.modapi.io/v1/${CC_GAME_ID}/mods?api_key${apiKey}`)
    let json = await res.json()
    modList = json.data
  }

  getMods()

</script>
<div>
  <h1>Mod List</h1>
  {#each modList as modInfo}
  <ModInfoCard data={modInfo}/>  
  {/each}

</div>
  
<style>

</style>