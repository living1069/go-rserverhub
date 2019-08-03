<template>
  <div>
    <v-layout>
      <v-flex xs3>
        <flex-divider left>
          <div class="flex-center">
            <div>
              <p class="headline mb-2">Состояние</p>
              <p
                :class="`subheading ma-0 ${stateColor(server.session)}--text`"
              >{{stateString(server.session)}}</p>
            </div>
            <v-spacer/>
            <v-icon x-large :color="stateColor(server.session)">{{stateIcon(server.session)}}</v-icon>
          </div>
        </flex-divider>
      </v-flex>
      <v-flex xs3>
        <flex-divider>
          <p class="headline mb-2">Адрес</p>
          <span class="subheading">{{hostAddress}}</span>
        </flex-divider>
      </v-flex>
      <v-flex xs3>
        <flex-divider>
          <p class="headline mb-2">Порт</p>
          <span class="subheading">{{server.configuration.port}}</span>
        </flex-divider>
      </v-flex>
      <v-flex xs3>
        <flex-divider right>
          <div class="flex-center">
            <div>
              <p class="headline mb-2">Хост</p>
              <p class="subheading ma-0">{{hostName}}</p>
            </div>
            <v-spacer/>
            <v-btn
              :href="hostLink"
              @click.prevent="$router.push(hostLink)"
              flat
              fab
              big
              class="ma-0"
            >
              <v-icon medium>arrow_forward_ios</v-icon>
            </v-btn>
          </div>
        </flex-divider>
      </v-flex>
    </v-layout>
    <v-divider class="my-4"/>
    <v-layout>
      <v-flex xs3>
        <flex-divider left>
          <p class="headline mb-0">Конфигурация сервера</p>
          <server-configuration :configuration="server.configuration"/>
        </flex-divider>
      </v-flex>
      <v-flex xs9>
        <flex-divider right>
          <v-layout align-center justify-space-between>
            <p class="headline mb-0">История</p>
            <v-btn
              small
              flat
              color="red"
              class="my-0 mr-0"
              :disabled="isClearingSessions"
              :loading="isClearingSessions"
              @click="clearSessions"
            >Очистить историю</v-btn>
          </v-layout>
          <server-sessions-list :server="server"/>
        </flex-divider>
      </v-flex>
    </v-layout>
  </div>
</template>

<script>
import FlexDivider from "@/components/shared/FlexDivider";
import ServerConfiguration from "@/components/server/ServerConfiguration";
import ServerSessionsList from "@/components/server/ServerSessionsList";
import ServerMixin from "@/mixins/server";

export default {
  components: {
    FlexDivider,
    ServerConfiguration,
    ServerSessionsList
  },
  data: () => ({
    isClearingSessions: false
  }),
  props: ["server"],
  mixins: [ServerMixin],
  computed: {
    hostLink() {
      return `/host/${this.server.host.id}`;
    },
    hostName() {
      return this.server.host.name;
    },
    hostAddress() {
      return this.server.host.address;
    }
  },
  methods: {
    async clearSessions() {
      try {
        this.isClearingSessions = true;
        await this.$api.servers.sessions.clear(this.server.id);
        this.$events.$emit("SERVER_SESSIONS_UPDATE");
      } catch (e) {
        this.$events.$emit("error", e);
      } finally {
        this.isClearingSessions = false;
      }
    }
  }
};
</script>

<style>
</style>
