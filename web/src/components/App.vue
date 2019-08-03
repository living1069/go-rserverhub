<template>
  <v-layout fill-height>
    <v-navigation-drawer app clipped fixed permanent>
      <v-list>
        <v-list-tile v-for="item in items" :key="item.title" :to="item.path">
          <v-list-tile-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>{{ item.title }}</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar app fixed clipped-left :height="72">
      <v-toolbar-title>
        <span class="font-weight-medium">ROH</span>
        <span class="font-weight-light">SERVERHUB</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn fab small flat @click="dark = !dark">
        <v-icon medium>invert_colors</v-icon>
      </v-btn>
    </v-toolbar>
    <v-content class="h-100">
      <v-card class="pa-4 h-100">
        <v-slide-x-transition leave-absolute>
          <router-view class="fill-width"></router-view>
        </v-slide-x-transition>
      </v-card>
    </v-content>
  </v-layout>
</template>

<script>
export default {
  data: () => ({
    items: [
      { title: "Машины", icon: "cloud", path: "/hosts" },
      { title: "Сервера", icon: "storage", path: "/servers" },
      { title: "Конфигурации", icon: "settings", path: "/configurations" }
    ],
    ws: null
  }),
  computed: {
    dark: {
      get() {
        return this.$store.getters.dark;
      },
      set(value) {
        this.$store.dispatch("SET_DARK", value);
      }
    }
  },
  mounted() {
    const self = this;
    let ws = new WebSocket(
      `ws://${this.$store.getters.endpoint}/ws/queue/info`
    );
    ws.onerror = msg => self.$events.$emit("error", msg);
    ws.onmessage = msg => {
      let data = JSON.parse(msg.data);
      self.$events.$emit(data.type, data.payload);
    };
    this.ws = ws;
  }
};
</script>

<style>
</style>
