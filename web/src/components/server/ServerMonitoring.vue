<template>
  <div>
    <v-layout>
      <v-flex xs3>
        <flex-divider left>
          <p class="headline mb-2 text-truncate">Игроки</p>
          <p class="subheading mb-0 text-truncate">{{players}}</p>
        </flex-divider>
      </v-flex>
      <v-flex xs3>
        <flex-divider>
          <p class="headline mb-2 text-truncate">Время работы</p>
          <p class="subheading mb-0 text-truncate">{{uptime}}</p>
        </flex-divider>
      </v-flex>
      <v-flex xs3>
        <flex-divider>
          <p class="headline mb-2 text-truncate">Игровое время</p>
          <p class="subheading mb-0 text-truncate">{{gametime}}</p>
        </flex-divider>
      </v-flex>
    </v-layout>
    <v-divider class="my-4"/>
    <v-layout>
      <v-flex xs6>
        <flex-divider>
          <v-card flat>
            <p class="title font-weight-light mb-0 ml-1">FPS</p>
            <server-monitoring-chart
              :dataset="fps"
              :height="200"
              class="mt-3"
            />
          </v-card>
        </flex-divider>
      </v-flex>
      <v-flex xs6>
        <flex-divider>
          <v-card flat>
            <p class="title font-weight-light mb-0 ml-1">Update (ms)</p>
            <server-monitoring-chart
              :dataset="engine"
              :height="200"
              class="mt-3"
            />
          </v-card>
        </flex-divider>
      </v-flex>
    </v-layout>
  </div>
</template>

<script>
import ServerMonitoringChart from "./ServerMonitoringChart";
import FlexDivider from "@/components/shared/FlexDivider";
import moment from "moment";
import momentDurationFormatSetup from "moment-duration-format";

momentDurationFormatSetup(moment);
moment.locale("ru");

var self = null;

export default {
  props: ["id"],
  components: {
    ServerMonitoringChart,
    FlexDivider
  },
  computed: {
    uptime() {
      if (this.stats)
        return moment
          .duration(this.stats.uptime, "milliseconds")
          .format("HH:mm:ss");
      else return "00:00:00";
    },
    gametime() {
      if (this.stats)
        return moment(this.stats.gametime).format("YYYY.MM.DD HH:mm:ss");
      else return "00:00:00";
    },
    players() {
      if (this.stats) return this.stats.players;
      else return 0;
    },
    stats: () => self.$store.state.server.stats,
    fps: () => self.$store.state.server.dataset.fps,
    engine: () => self.$store.state.server.dataset.engine
  },
  watch: {
    id: function(value, old) {
      console.log(value, old);
    }
  },
  created() {
    self = this;
  }
};
</script>

<style>
</style>
