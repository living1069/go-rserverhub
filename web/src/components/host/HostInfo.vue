<template>
  <div>
    <v-layout>
      <v-flex xs3>
        <flex-divider left>
          <div class="flex-center">
            <div>
              <p class="headline mb-2">Состояние</p>
              <p :class="`subheading ma-0 ${stateColor(status)}--text`">{{stateString(status)}}</p>
            </div>
            <v-spacer/>
            <v-icon x-large :color="stateColor(status)">{{stateIcon(status)}}</v-icon>
          </div>
        </flex-divider>
      </v-flex>
      <v-flex xs3>
        <flex-divider>
          <p class="headline mb-2">Имя</p>
          <span class="subheading">{{host.name}}</span>
        </flex-divider>
      </v-flex>
      <v-flex xs3>
        <flex-divider>
          <p class="headline mb-2">Адрес</p>
          <span class="subheading">{{host.address}}</span>
        </flex-divider>
      </v-flex>
      <v-flex xs3>
        <flex-divider right>
          <p class="headline mb-2">Имя машины</p>
          <span class="subheading">Неизвестно</span>
        </flex-divider>
      </v-flex>
    </v-layout>
    <v-divider class="my-4"/>
    <v-layout>
      <v-flex xs3 class="text-xs-center">
        <flex-divider left>
          <p class="headline">Состояние машины</p>
          <div>
            <v-progress-circular
              class="mb-3"
              :rotate="90"
              :size="180"
              :width="12"
              :value="cpu"
              :color="cpuColor"
            >
              <p class="headline ma-0">{{Math.round(cpu)}}%</p>
            </v-progress-circular>
            <p class="subheading">Загрузка CPU</p>
          </div>
          <div>
            <v-progress-circular
              class="mb-3"
              :rotate="90"
              :size="180"
              :width="12"
              :value="(memoryUsed > 0 && memoryTotal) ? (memoryUsed / memoryTotal) * 100 : 0"
              :color="memoryColor"
            >
              <p class="title mb-2">{{Math.round(memoryUsed)}} MB</p>
              <p class="subheading ma-0">{{Math.round(memoryTotal)}} MB</p>
            </v-progress-circular>
            <p class="subheading">Использование памяти</p>
          </div>
        </flex-divider>
      </v-flex>
      <v-flex xs9>
        <flex-divider right>
          <div class="flex-center">
            <p class="headline ma-0" style="float: left">Сервера</p>
            <v-spacer/>
            <v-btn flat color="green" class="my-0" @click="addServer">Добавить сервер</v-btn>
          </div>
          <server-list :host="host"/>
          <server-dialog ref="serverDialog" :host="host" :public="false"/>
        </flex-divider>
      </v-flex>
    </v-layout>
  </div>
</template>

<script>
import FlexDivider from "@/components/shared/FlexDivider";
import ServerList from "@/components/server/ServerList";
import ServerDialog from "@/components/server/ServerDialog";
import HostMixin from "@/mixins/host";

export default {
  components: {
    FlexDivider,
    ServerList,
    ServerDialog
  },
  mixins: [HostMixin],
  props: ["host"],
  data: () => ({
    serverAddDialogVisible: false,
    stats: null
  }),
  computed: {
    status() {
      return this.host.status;
    },
    memoryUsed() {
      if (!this.stats) return 0;
      return this.stats.memoryTotal - this.stats.memoryAvailable;
    },
    memoryTotal() {
      return this.stats ? this.stats.memoryTotal : 0;
    },
    cpu() {
      return this.stats ? this.stats.cpu : 0;
    },
    cpuColor() {
      return this.cpu > 90 ? "red" : this.cpu > 75 ? "orange" : "green";
    },
    memoryColor() {
      const value = (this.memoryUsed / this.memoryTotal) * 100;
      return value > 90 ? "red" : value > 75 ? "orange" : "green";
    }
  },
  created() {
    this.$events.$on("HOST_STATS", this.onStats.bind(this));
  },
  methods: {
    onStats(stats) {
      if (stats.host != this.host.id) return;
      this.stats = stats;
    },
    addServer() {
      this.$refs.serverDialog.open(null, this.host);
    }
  }
};
</script>