<template>
  <div class="layui-col-space10">
    <div class="layuimini-qiuck-module">
      <a href="#" :title="title">
        <i class="fa" :class="className"></i>
        <span>{{ title }}</span>
      </a>
    </div>
  </div>
</template>

<script>

export default {
  name: 'QuickEntry',
  props: ['title', 'className']

}
</script>

<style lang="scss" scoped>
.layui-col-space10{
  margin-left: 5px;
  margin-right: 5px;
}
.layuimini-qiuck-module a {
  color: #333;
  text-decoration: none;
}

.layuimini-qiuck-module {
  text-align: center;
  margin-top: 9px;
  margin-bottom: 9px;
}

.layuimini-qiuck-module a i {
  display: inline-block;
  width: 100%;
  height: 90px !important;
  line-height: 80px !important;
  text-align: center;
  border-radius: 2px;
  font-size: 40px !important;
  background-color: #F8F8F8;
  color: #333;
  transition: all .3s;
  -webkit-transition: all .3s;
}

.layuimini-qiuck-module a span {
  position: relative;
  top: 8px;
  display: block;
  color: #666;
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;
  font-size: 14px;
}

</style>
