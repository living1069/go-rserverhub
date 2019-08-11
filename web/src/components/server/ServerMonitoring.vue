<template>
  <v-layout column>
    <v-layout class="px-4 pb-2 pt-4" style="max-height: 300px">
      <v-flex xs6 class="pr-2">
        <iframe
          :class="{border: !this.$store.getters.dark}"
          :src="`http://10.8.0.1:9400/d-solo/Or3PhOZZz/roh?orgId=1&var-server=${server.id}&theme=${theme}&panelId=2`"
          frameborder="0"
        ></iframe>
      </v-flex>
      <v-flex class="pl-2">
        <iframe
          :class="{border: !this.$store.getters.dark}"
          :src="`http://10.8.0.1:9400/d-solo/Or3PhOZZz/roh?orgId=1&var-server=${server.id}&theme=${theme}&panelId=4`"
          frameborder="0"
        ></iframe>
      </v-flex>
    </v-layout>
    <v-layout class="px-4 py-2" style="max-height: 300px">
      <v-flex xs6 class="pr-2">
        <iframe
          :class="{border: !this.$store.getters.dark}"
          :src="`http://10.8.0.1:9400/d-solo/Or3PhOZZz/roh?orgId=1&var-server=${server.id}&theme=${theme}&panelId=14`"
          frameborder="0"
        ></iframe>
      </v-flex>
      <v-flex class="pl-2">
        <iframe
          :class="{border: !this.$store.getters.dark}"
          :src="`http://10.8.0.1:9400/d-solo/Or3PhOZZz/roh?orgId=1&var-server=${server.id}&theme=${theme}&panelId=12`"
          frameborder="0"
        ></iframe>
      </v-flex>
    </v-layout>
  </v-layout>
</template>

<script>
export default {
  props: ["server"],
  computed: {
    theme() {
      return this.$store.getters.dark ? "dark" : "light";
    }
  }
};
</script>

<style>
.border {
  border: 1px solid lightgray;
  padding: 1px;
}

iframe {
  height: 100%;
  width: 100%;
}
</style>
