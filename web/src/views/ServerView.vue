<template>
  <v-layout fill-height column>
    <request-loader :ready="ready" :error="error">
      <v-layout fill-height column>
        <div class="flex-center mb-4">
          <span class="display-1">Сервер #{{id}}</span>
          <v-spacer/>
          <server-actions :server="server"/>
        </div>
        <v-tabs v-model="active">
          <v-tab style="min-width: 160px" :key="index" v-for="(text, index) in tabs">{{text}}</v-tab>
        </v-tabs>
        <v-layout column>
          <server-info v-show="active==0" :server="server" class="mt-4"/>
          <server-console v-show="active==1" :server="server"/>          
        </v-layout>
      </v-layout>
    </request-loader>
  </v-layout>
</template>

<script>
import ServerInfo from "@/components/server/ServerInfo";
import ServerConsole from "@/components/server/ServerConsole";
import ServerActions from "@/components/server/ServerActions";
import RequestLoader from "@/components/shared/RequestLoader";

export default {
  components: {
    RequestLoader,
    ServerInfo,
    ServerConsole,    
    ServerActions
  },
  props: ["id"],
  data: () => ({
    active: 0,
    serverAddDialogVisible: false,
    tabs: ["Информация", "Консоль"],
    ready: false,
    server: null,
    error: false
  }),
  watch: {
    active: function(val) {
      this.$events.$emit("server-view-tab-change", val);
    }
  },
  created() {
    this.get();
    const self = this;

    this.$events.$on("SERVER_UPDATE", server => {
      if (self.server.id == server.id) self.server = server;
    });

    this.$events.$on("SERVER_DELETE", id => {
      if (self.server.id == id) self.$router.push("/servers");
    });
  },
  methods: {
    async get() {
      try {
        this.ready = false;
        const result = await this.$api.servers.get(this.id);
        this.server = result.data;
      } catch (e) {
        this.$events.$emit("ERROR", e);
      } finally {
        this.ready = true;
      }
    }
  }
};
</script>

<style>
.v--fill-height,
.v--fill-height > .v-window,
.v--fill-height > .v-window > .v-window__container,
.v--fill-height > .v-window > .v-window__container > .v-window-item {
  display: flex;
  flex-direction: column;
  flex: 1 1 auto;
}
</style>
