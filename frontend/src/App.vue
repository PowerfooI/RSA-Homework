<template>
  <div id="app">
    <RsaPanel></RsaPanel>
  </div>
</template>

<script>
import RsaPanel from "@/RsaPanel";

export default {
  name: 'App',
  components: {
    RsaPanel,
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
