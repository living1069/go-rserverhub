<template>
  <v-layout fill-height>
    <v-layout column fill-height>
      <p class="display-1">Конфигурации</p>
      <div>
        <v-btn color="primary" outline @click="add" class="ml-0 mb-4">Добавить конфигурацию</v-btn>
        <v-spacer/>
      </div>
      <v-data-table :headers="headers" :loading="busy" :items="configurations" hide-actions>
        <v-progress-linear v-slot:progress color="primary" indeterminate></v-progress-linear>
        <template slot="items" slot-scope="props">
          <td class="pointer" @click="edit(props.item.id)">{{ props.item.id }}</td>
          <td class="pointer" @click="edit(props.item.id)">{{ props.item.name }}</td>
          <td class="pointer" @click="edit(props.item.id)">{{ props.item.executable }}</td>
          <td class="pointer" @click="edit(props.item.id)">{{ props.item.directory }}</td>
          <td class="pointer" @click="edit(props.item.id)">{{ props.item.port }}</td>
          <td class="pointer" @click="edit(props.item.id)">{{ props.item.map }}</td>
        </template>
      </v-data-table>
    </v-layout>
    <configuration-dialog public/>
  </v-layout>
</template>

<script>
import ConfigurationDialog from "@/components/configuration/ConfigurationDialog";
export default {
  components: { ConfigurationDialog },
  data: () => ({
    busy: false,
    configurations: [],
    headers: [
      {
        text: "ID",
        align: "left",
        value: "id"
      },
      { text: "Название", value: "name" },
      { text: "Испольняемый файл", value: "executable" },
      { text: "Рабочая папка", value: "directory" },
      { text: "Порт", value: "port" },
      { text: "Карта", value: "map" }
    ]
  }),
  mounted() {
    this.get();
    this.$events.$on("CONFIGURATION_UPDATE", this.get.bind(this));
    this.$events.$on("CONFIGURATION_CREATE", this.get.bind(this));
    this.$events.$on("CONFIGURATION_DELETE", this.get.bind(this));
  },
  methods: {
    edit(id) {
      this.$events.$emit("SHOW_CONFIGURATION_DIALOG", id);
    },
    add() {
      this.$events.$emit("SHOW_CONFIGURATION_DIALOG");
    },
    async get() {
      try {
        this.busy = true;
        const response = await this.$api.configurations.list();
        this.configurations = response.data;
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
