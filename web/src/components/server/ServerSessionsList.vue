<template>
  <div>
    <v-data-table
      :headers="headers"
      :items="sessions"
      hide-actions
      :pagination.sync="pagination"
      :loading="busy"
    >
      <v-progress-linear v-slot:progress color="primary" indeterminate></v-progress-linear>
      <template v-slot:items="props">
        <td>{{ time(props.item.dateStart) }}</td>
        <td>{{ time(props.item.dateStop) }}</td>
        <td>
          <v-icon :color="stateColor(props.item)">{{stateIcon(props.item)}}</v-icon>
          <span :class="`${stateColor(props.item)}--text ml-2`">{{stateString(props.item)}}</span>
        </td>
        <td>
          <v-btn
            :disabled="props.item.log == null"
            color="blue-grey"
            class="white--text"
            :href="log(props.item.log)"
            small
            flat
            fab
          >
            <v-icon dark>cloud_download</v-icon>
          </v-btn>
        </td>
        <td>
          <v-btn
            :disabled="props.item.report == null"
            color="blue-grey"
            class="white--text"
            :href="report(props.item.report)"
            small
            flat
            fab
          >
            <v-icon dark>cloud_download</v-icon>
          </v-btn>
        </td>
      </template>
    </v-data-table>
    <div class="text-xs-center pt-2">
      <v-pagination v-model="pagination.page" :length="pages" total-visible="10"></v-pagination>
    </div>
  </div>
</template>

<script>
import ServerMixin from "@/mixins/server";
import moment from "moment";

export default {
  props: ["server"],
  mixins: [ServerMixin],
  data: () => ({
    headers: [
      { text: "Запущен", value: "dateStart", sortable: false },
      { text: "Остановлен", value: "dateStop", sortable: false },
      { text: "Статус", value: "status", sortable: false },
      { text: "Лог-файл", value: "status", sortable: false },
      { text: "Отчет об ошибке", value: "status", sortable: false }
    ],
    pagination: {
      descending: true,
      page: 0,
      rowsPerPage: 7,
      totalItems: 0
    },
    sessions: [],
    busy: false
  }),
  computed: {
    pages() {
      if (
        this.pagination.rowsPerPage == null ||
        this.pagination.totalItems == null
      )
        return 0;

      return Math.ceil(
        this.pagination.totalItems / this.pagination.rowsPerPage
      );
    }
  },
  mounted() {
    this.get();
    const self = this;
    this.$events.$on("SERVER_UPDATE", server => {
      if (server.id == self.server.id) self.get();
    });

    this.$events.$on("SERVER_SESSIONS_UPDATE", this.get.bind(this));
  },
  methods: {
    async get() {
      try {
        this.busy = true;
        const result = await this.$api.servers.sessions.list(this.server.id);
        this.sessions = result.data;
        this.pagination.totalItems = this.sessions.length;
      } catch (e) {
        this.$events.$emit("ERROR", e);
      } finally {
        this.busy = false;
      }
    },
    time(unix) {
      return moment.unix(unix).format("YYYY.MM.DD HH:mm:ss");
    },
    log(log) {
      if (!log) return "#";
      return `http://${this.$store.getters.endpoint}/logs/${log.filename}`;
    },
    report(report) {
      if (!report) return "#";
      return `http://${this.$store.getters.endpoint}/reports/${report.filename}`;
    }
  }
};
</script>
