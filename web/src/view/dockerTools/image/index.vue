<template>
  <div>

    <div v-show="scene == 0">
      <ImageTable/>
    </div>

    <div v-show="scene == 1">
      <ImageDetail/>
    </div>
  </div>
</template>

<script>
import ImageTable from './ImageTable'
import ImageDetail from './ImageDetail'

export default {
  name: 'Image',
  components: {
    ImageTable,
    ImageDetail
  },
  data(){
    return{
      scene:1
    }
  }
}
</script>

<style scoped lang="scss">

</style>
