<template>
  <div>
    <router-view v-if="$route.meta.keepAlive" v-slot="{ Component }">
      <transition mode="out-in" name="el-fade-in-linear">
        <keep-alive>
          <component :is="Component"/>
        </keep-alive>
      </transition>
    </router-view>
    <router-view v-if="!$route.meta.keepAlive" v-slot="{ Component }">
      <transition mode="out-in" name="el-fade-in-linear">
        <component :is="Component"/>
      </transition>
    </router-view>
  </div>
</template>

<script>
import { getInfo, getVersion } from "@/api/docker";

export default {
  name: 'Docker',
  data() {
    return {}
  },
  created() {
    this.version()

  },

  methods: {
    async version() {
      const d1 = await getInfo()
      console.log(d1)
      const d2 = await getVersion()
      console.log(d2)
    },
  }
}
</script>
