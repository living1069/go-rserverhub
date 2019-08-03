<template>
  <v-dialog v-model="visible" max-width="600">
    <v-card>
      <v-card-title class="headline">{{data.id ? 'Редактирование сервера' : 'Добавление сервера'}}</v-card-title>
      <v-divider/>
      <v-expand-transition>
        <v-progress-linear v-show="busy" class="pa-0 ma-0" :indeterminate="true"></v-progress-linear>
      </v-expand-transition>
      <v-card-text class="px-4">
        <v-layout column>
          <v-text-field
            v-if="data.host"
            readonly
            v-model="data.host.name"
            prepend-icon="lock"
            label="Хост"
            disabled
            required
          ></v-text-field>
          <v-select
            prepend-icon="list"
            v-model="data.configuration"
            :items="configurations"
            return-object
            item-text="name"
            item-value="id"
            label="Конфигурация сервера"
          ></v-select>
        </v-layout>
        <configuration-fields ref="fields" :busy="busy" :data="data.configuration"/>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn
          :disabled="busy"
          :loading="busy"
          flat
          @click="save"
        >{{this.data.id ? 'Сохранить' : 'Добавить'}}</v-btn>
        <v-btn flat @click="close">Закрыть</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import ConfigurationFields from "@/components/configuration/ConfigurationFields.vue";

export default {
  components: { ConfigurationFields },
  data: () => ({
    data: null,
    busy: false,
    host: null,
    visible: false,
    configurations: null,
    oldConfiguration: { id: null }
  }),
  created() {
    this.reset();
  },
  methods: {
    async loadConfigurations() {
      try {
        this.busy = true;
        const result = await this.$api.configurations.list();
        this.configurations = [{ name: "Не выбран", id: null }].concat(
          result.data
        );
      } catch (e) {
        this.$events.$emit(
          "MESSAGE",
          "Не удалось загрузить список конфигураций"
        );
      } finally {
        this.busy = false;
      }
    },

    async load(id) {
      try {
        this.busy = true;
        const result = await this.$api.servers.get(id);
        this.data = result.data;
        this.oldConfiguration = this.data.configuration;
        if (this.oldConfiguration == null) this.oldConfiguration = { id: null };
      } catch (e) {
        this.$events.$emit("ERROR", e);
      } finally {
        this.busy = false;
      }
    },

    async save() {
      if (!this.$refs.fields.validate()) return;
      try {
        this.busy = true;
        if (this.host) this.data.host = this.host;        
        if (this.data.id) await this.$api.servers.save(this.data);
        else await this.$api.servers.add(this.data);
        this.close();
      } catch (e) {
        this.$events.$emit("ERROR", e);
      } finally {
        this.busy = false;
      }
    },

    close() {
      this.visible = false;
    },

    reset() {
      this.data = {
        id: null,
        host: null,
        configuration: {
          id: null,
          name: null,
          executable: null,
          directory: null,
          fsgame: null,
          arguments: null,
          port: null,
          map: null
        }
      };
    },

    open(id, host) {
      this.reset();
      this.$refs.fields.reset();
      this.loadConfigurations();
      this.host = host;
      if (id) this.load(id);
      this.visible = true;
    }
  }
};
</script>

<style>
</style>
