<template>
  <div>
    <div class="layui-col-md3 layuimini-qiuck-module">
      <a href="#" onClick="javascript:triggerNavMenuClick('ALL', 'images')" title="镜像管理" data-icon="fa fa-window-maximize">
        <i class="fa fa-window-maximize"></i>
        <cite>镜像管理</cite>
      </a>
    </div>
  </div>
</template>

<script>

export default {
  name: 'QuickEntry',
  props: ['title']

}
</script>

<style lang="scss" scoped>


</style>
