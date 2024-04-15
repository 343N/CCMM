<script>
  import { Heading } from "flowbite-svelte"
  
  export let timer = 2000
  export let introDelay = 500
  export let onfinish = () => {}
  export let visible = true

  let logoPath = "logo-256.png"
  let audioSrc = "intro.mp3"
  let audioVol = 0.0

  

  let onContainerLoad = (event) => {
    let eAudio = event.target
    let container = eAudio.parentElement
    setTimeout(()=>{
      console.log(eAudio)
      console.log("Loaded!")
      eAudio.volume = audioVol
      eAudio.play()
      eAudio.addEventListener("play", ()=> {
        container.style.opacity = 1
        setTimeout(onfinish, timer)
      })
    }, introDelay)

}

</script>

{#if visible}
<div class="container">
  <div class="contents">
    <img src={logoPath} />
    <h1>CCMM</h1>
    <h2>Cortex Command Mod Manager</h2>
  </div>
  <audio on:loadeddata={onContainerLoad} src={audioSrc} volume={audioVol}>
</div>
{/if}

<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@100..900&display=swap');

div.container {
font-family: "Inter", sans-serif;
display: flex;
flex-direction:  column;
justify-content: center;
height: 90%;
width: 100%;
  opacity: 0;
}



img {
  width: 8rem;
  margin-bottom: 2rem;
}

div.contents {
  align-items: center;
  display: flex;
  /* heightg  */
  flex-direction: column;
}

h1, h2 {
  font-weight: 200;
  display: inline-block;
}
h1 { 

  font-size: 3rem;
  line-height: 3rem;
}
h2 {
  font-size: 2rem;
  line-height: 2rem;
}

</style>
