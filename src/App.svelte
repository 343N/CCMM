<script>
  import Intro from './pages/Intro.svelte'
  import Browser from './pages/Browser.svelte'
  import './styles.css'

  $: browserShown = false
  let introShown

  let introFinished = () =>{
    browserShown = true
    introShown = false
  }

</script>

  <Intro onfinish={introFinished} bind:visible={introShown}/>
  <Browser visible={browserShown} />


<style>
  .logo.vite:hover {
    filter: drop-shadow(0 0 2em #747bff);
  }

  .logo.svelte:hover {
    filter: drop-shadow(0 0 2em #ff3e00);
  }

  :global(body, #app){
    width: 100%;
    height: 100%;
  }

</style>
