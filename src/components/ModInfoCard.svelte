<script>
  import { Button, Card, Heading, Img, P } from "flowbite-svelte";
 
 export let data = null;
 export let clickable = null;
 export let remote = false;

 export let downloadButtonHandler = () => { return null }

</script>

<Card class="animate clickable sm:p-0 w-48 h-256 mb-4 ml-2 mr-2">
  <img src={data.logo.original} class="object-cover h-48 w-48 rounded-lg rounded-b-none"/>
  <div class="align-items-center">
    <P class="pl-4 pr-4 pt-4 text-m font-semibold w-full w-auto">{data.name}</P>
    <button data-mod-id={data.id} on:click={downloadButtonHandler}>Download</button>
  </div>
    <!-- <P class="pl-4 pr-4 pb-4 text-sm">{data.summary.substring(0, 50) + (data.summary.length > 50 ? "..." : "")}</P> -->
</Card>

<style>
  .animate {
    transition: all 0.2s;
  }
  :global(.clickable) { 
    transition: all 0.2s;
  }

  :global(.clickable) > img{ 
    transition: all 0.2s;
    cursor: pointer;
  }

  :global(.clickable):hover {
    scale: 1;
    box-shadow:
      0 0 10px 0px #ffffff08;  /* inner white */
      /* 0 0 10px 6px #f0f, middle magenta */
      /* 0 0 14px 9px #0ff; outer cyan */
  }

  button {
    background-color: rgba(0,0,0,0.2);
    width: fit-content;
    text-align: center;
    margin: 1rem;

  }

  div {
    display: flex;
    height: 8rem;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;
  }
  
</style>  