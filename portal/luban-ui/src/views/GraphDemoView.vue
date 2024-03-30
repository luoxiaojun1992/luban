<script>
export default {
  name: 'GraphDemoView',
  data() {
    return {
      goAstNodeList: ''
    }
  },
  mounted() {
    const go = new Go()
    WebAssembly.instantiateStreaming(fetch('/luban.wasm'), go.importObject)
      .then((result) => {
        go.run(result.instance)

        fetch('/graph_demo.json')
          .then((response) => response.text())
          .then((data) => {
            this.goAstNodeList = JSON.parse(LubanProgramParse(data))
          })
          .catch((err) => {
            console.error('Failed to fetch json example: ' + err)
          })
      })
      .catch((err) => {
        console.error('Failed to instantiate wasm: ' + err)
      })
  }
}
</script>

<template>
  {{ goAstNodeList }}
</template>
