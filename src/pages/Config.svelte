<script> 
  import { Input, Label, Fileupload, A} from 'flowbite-svelte'
  let apiKey = ""
  let apiUserId = ""
  let ccModPath = ""
</script>

<div id="configContainer">

  <A href="https://mod.io/me/access" >API details page</A>
  <div>
    <Label for="api_key">Mod.io API Key
      <Input type="password" id="api_key" bind:apiKey required />
    </Label>
  </div>
  <div>
    <Label for="api_userid">API User ID
      
      <Input type="text" id="api_userid" bind:apiUserId required />
    </Label>
  </div>
  <div>
    <Label for="api_key">Mod Directory
      
      <Fileupload bind:ccModPath />
    </Label>
  </div>
</div>

<style>
#configContainer {
  /* margin-left: 4rem; */
  /* margin-right: 4rem; */
  padding: 4rem;
}

#configContainer > div {
  margin-bottom: 4rem;
  /* color: rgba(240,240,240,255) */
}
</style>