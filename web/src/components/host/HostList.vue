<template>
  <v-layout fill-height class="fill-width">
    <v-layout column fill-height>
      <p class="display-1">Машины</p>
      <div>
        <v-btn outline color="primary" @click="add" class="ml-0 mb-4">Добавить машину</v-btn>
        <v-spacer/>
      </div>
      <v-data-table :headers="headers" :loading="busy" :items="hosts" hide-actions>
        <v-progress-linear v-slot:progress color="primary" indeterminate></v-progress-linear>
        <template slot="items" slot-scope="props">
          <td class="pointer" @click="edit(props.item.id)">{{ props.item.id }}</td>
          <td class="pointer" @click="edit(props.item.id)">{{ props.item.name }}</td>
          <td class="pointer" @click="edit(props.item.id)">{{ props.item.address }}</td>
          <td class="pointer flex-center" @click="edit(props.item.id)">
            <v-icon :color="stateColor(props.item.status)">{{stateIcon(props.item.status)}}</v-icon>
            <span
              :class="`${stateColor(props.item.status)}--text ml-2`"
            >{{stateString(props.item.status)}}</span>
          </td>
        </template>
      </v-data-table>
    </v-layout>
    <host-dialog/>
  </v-layout>
</template>

<script>
import host from "@/mixins/host";
import HostDialog from "@/components/host/HostDialog";

export default {
  mixins: [host],
  components: { HostDialog },
  data: () => ({
    busy: false,
    hosts: [],
    headers: [
      {
        text: "ID",
        align: "left",
        value: "id"
      },
      { text: "Название", value: "name" },
      { text: "Адрес", value: "address" },
      { text: "Состояние", value: "state" }
    ]
  }),
  mounted() {
    this.get();
    this.$events.$on("HOST_CREATE", this.get.bind(this));
    this.$events.$on("HOST_UPDATE", this.get.bind(this));
    this.$events.$on("HOST_DELETE", this.get.bind(this));
  },
  methods: {
    edit(id) {
      this.$router.push(`/host/${id}`);
    },
    add() {
      this.$events.$emit("SHOW_HOST_DIALOG");
    },
    async get() {
      try {
        this.busy = true;
        const response = await this.$api.hosts.list();
        this.hosts = response.data;
      } catch (e) {
        this.$events.$emit("ERROR", e);
      } finally {
        this.busy = false;
      }
    }
  }
};
</script>

<style>
</style>
