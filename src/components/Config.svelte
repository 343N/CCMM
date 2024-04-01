<script> 
  import { Input, Label, A, Helper, Button} from 'flowbite-svelte'
  import {open as openShell} from '@tauri-apps/api/shell'
  import {open as openDialog} from '@tauri-apps/api/dialog' 
  import { AppStore } from "../AppStore"
  import { createEventDispatcher } from 'svelte';
  
  export let apiKey = ""
  export let apiUserId = ""
  export let ccModPath = ""
 
  export let visible = true

  export let save = ()=>{}

  const dispatch = createEventDispatcher()

  let openInBrowser = async (e) => {
    await openShell(e.target.dataset.externalLink)
    e.preventDefault()
  }

  $: modsPickerLabel = "No folder chosen."

  let setModsFolder = async (e) => {
    e.preventDefault()
    let selected = (await openDialog({directory: true}))
    
    // @ts-ignore
    modsPickerLabel = selected || "No folder chosen."
    console.log(selected)
  }

  let saveConfig = async (e) => {
    await AppStore.set("apiKey", apiKey)
    await AppStore.set("apiUserId", apiUserId)
    await AppStore.set("ccModPath", ccModPath)

    dispatch("save",
      {
        "apiKey": apiKey,
        "apiUserId": apiUserId,
        "ccModPath": ccModPath
      })
  }

</script>

{#if visible}

<div id="configContainer">

  <div id="optionsContainer">
  <p>In order to use CCMM, you need to grab your API details for mod.io</p>

  <A href="#" data-external-link='https://mod.io/me/access' on:click={openInBrowser}>Click here to go to the Mod.io API details page</A>
  <div id="apiContainer">
    <div>
      <Label for="api_key">Mod.io API Key
      </Label>
      <Input type="password" id="api_key" bind:apiKey required />
    </div>
    <div>
      <Label for="api_userid">API User ID
      </Label>  
      <Input type="text" id="api_userid" bind:apiUserId required />
    </div>
  </div>
  <div>
    <Label for="api_key">Mod Directory
      <Helper>Example: C:\Games\Cortex Command\Mods</Helper>
    </Label>
    <Button id="configModsFolderPicker" webkitdirectory on:click={setModsFolder}>Choose Folder</Button>
    <Label>Folder: {modsPickerLabel}</Label>
  </div>
</div>
  <div id="saveContainer">
    <Button id="save" on:click={saveConfig}>Save</Button>
  </div>

</div>

{/if}
<style>

#configContainer {
  display: flex;
  flex-direction: column;
  justify-content: space-between; 
  height: 100%;
  width: 100%;
}

#configContainer > #optionsContainer {
  /* margin-left: 4rem; */
  /* margin-right: 4rem; */
  display: flex;
  flex-direction: column;
  padding: 2rem;
}

#apiContainer {
  display: flex;
  flex-direction: row;
  width: 100%;
}


#apiContainer > div {
  flex-grow: 1;
  /* padding-left: 1rem; */
}

#apiContainer > div:first-child {
  margin-right:0.5rem;
}
#apiContainer > div:last-child {
  margin-left:0.5rem;
}


#saveContainer {
  width: 100%;
  display: flex;
  justify-content: center;
}

#configContainer > div {
  margin-bottom: 1rem;
  width: 100%;
  /* color: rgba(240,240,240,255) */
}


#optionsContainer > div {
  margin-bottom: 1rem;
}
</style>