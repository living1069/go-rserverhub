<template>
  <v-dialog v-model="visible" max-width="600">
    <v-card>
      <v-card-title
        class="headline"
      >{{data.id ? 'Редактирование конфигурации' : 'Добавление конфигурации'}}</v-card-title>
      <v-divider/>
      <v-expand-transition>
        <v-progress-linear v-show="busy" class="pa-0 ma-0" :indeterminate="true"></v-progress-linear>
      </v-expand-transition>
      <v-card-text class="px-4">
        <configuration-fields ref="fields" :busy="busy" :data="data"/>
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
  props: {
    public: {
      type: Boolean,
      default: false
    }
  },
  data: () => ({
    busy: false,
    data: null,
    visible: false
  }),
  created() {
    const self = this;
    this.reset();
    this.$events.$on("SHOW_CONFIGURATION_DIALOG", id => {
      self.$refs.fields.reset();
      self.reset();
      if (id) self.get(id);
      self.visible = true;
    });
  },
  methods: {
    async get(id) {
      try {
        this.busy = true;
        const response = await this.$api.configurations.get(id);
        this.data = response.data;
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
        this.data.public = this.public;
        const response = await this.$api.configurations.save(this.data);
        this.data = response.data;
        this.$events.$emit("MESSAGE", "Конфигурация обновлена");
        this.close();
      } catch (e) {
        this.$events.$emit("ERROR", e);
      } finally {
        this.busy = false;
      }
    },

    reset() {
      this.data = {
        id: null,
        name: null,
        executable: null,
        directory: null,
        fsgame: null,
        arguments: null,
        port: null,
        map: null,
        public: false
      };
    },

    close() {
      this.visible = false;
    }
  }
};
</script>

<style>
.data-fields {
  display: flex;
  flex-direction: column;
}

.data-field {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
</style>
