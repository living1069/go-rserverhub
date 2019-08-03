<template>
  <v-data-table :headers="headers" :loading="busy" hide-actions :items="data">
    <v-progress-linear v-slot:progress color="primary" indeterminate></v-progress-linear>
    <template slot="items" slot-scope="props">
      <td class="pointer" @click="doOpenServerView(props.item.id)">{{ props.item.id }}</td>
      <td class="pointer" @click="open(props.item.id)">{{ props.item.host.name }}</td>
      <td class="pointer" @click="open(props.item.id)">{{ props.item.configuration.port }}</td>
      <td class="pointer" @click="open(props.item.id)">{{ props.item.host.address }}</td>
      <td
        class="pointer"
        @click="open(props.item.id)"
      >{{ props.item.stats ? props.item.stats.players : 0 }}</td>
      <td
        class="pointer"
        @click="open(props.item.id)"
      >{{ props.item.stats ? props.item.stats.fps : 0 }}</td>
      <td class="pointer flex-center" @click="open(props.item.id)">
        <v-icon :color="stateColor(props.item.session)">{{stateIcon(props.item.session)}}</v-icon>
        <span
          :class="`${stateColor(props.item.session)}--text ml-2`"
        >{{stateString(props.item.session)}}</span>
      </td>
    </template>
  </v-data-table>
</template>

<script>
import ServerMixin from "@/mixins/server";

export default {
  props: ["host"],
  data: () => ({
    headers: [
      {
        text: "ID",
        align: "left",
        value: "id"
      },
      { text: "Адрес", value: "address" },
      { text: "Порт", value: "configuration.port" },
      { text: "Хост", value: "host" },
      { text: "Игроки", value: "players" },
      { text: "FPS", value: "stats.fps" },
      { text: "Состояние", value: "status" }
    ],
    busy: false,
    data: []
  }),
  mixins: [ServerMixin],
  mounted() {
    this.get();
    this.$events.$on("SERVER_CREATE", this.get.bind(this));
    this.$events.$on("SERVER_UPDATE", this.get.bind(this));
    this.$events.$on("SERVER_DELETE", this.get.bind(this));
  },
  methods: {
    open(id) {
      this.$router.push(`/server/${id}`);
    },
    async get() {
      try {
        this.busy = true;
        let result;        
        if (this.host) result = await this.$api.hosts.servers(this.host.id);
        else result = await this.$api.servers.list();        
        this.data = result.data;
      } catch (e) {
        this.$events.$emit("ERROR", e);
      } finally {
        this.busy = false;
      }
    }
  }
};
</script>