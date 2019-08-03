<template>
  <v-dialog v-model="visible" width="500">
    <v-card>
      <v-card-title class="headline">Добавление агента</v-card-title>
      <v-divider/>
      <v-expand-transition>
        <v-progress-linear v-show="busy" class="pa-0 ma-0" :indeterminate="true"></v-progress-linear>
      </v-expand-transition>
      <v-card-text class="pb-0">
        <v-text-field
          v-model="form.name"
          prepend-icon="title"
          label="Имя агента"
          :disabled="busy"
          required
        ></v-text-field>
        <v-text-field
          v-model="form.address"
          prepend-icon="dns"
          label="Публичный адрес"
          :disabled="busy"
          required
        ></v-text-field>
        <div v-if="form.id">
          <v-text-field
            readonly
            v-model="form.id"
            prepend-icon="lock"
            label="Идентификатор"
            disabled
            required
          ></v-text-field>
          <div class="ml-1 mb-2">Команда для запуска агента:</div>
          <v-sheet color="ml-1 blue-grey darken-4">
            <div
              class="pa-2 light-blue--text text--lighten-4"
              style="word-break: break-word;"
            >{{command}}</div>
          </v-sheet>
        </div>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn :disabled="busy" :loading="busy" flat @click="add">Сохранить</v-btn>
        <v-btn flat @click="close">Закрыть</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  data: () => ({
    busy: false,
    visible: false,
    form: {
      id: null,
      name: null,
      address: null
    }
  }),
  computed: {
    command() {
      return `RServerAgent.exe 
      --backend=${this.$store.getters.endpoint}
      --id=${this.form.id}`;
    }
  },
  created() {
    const self = this;
    this.$events.$on("SHOW_HOST_DIALOG", self.show);
  },
  methods: {
    async add() {
      try {
        this.busy = true;
        let response;

        if (this.form.id) response = await this.$api.hosts.save(this.form);
        else response = await this.$api.hosts.add(this.form);

        this.form = response.data;
        this.$events.$emit("MESSAGE", "Машина успешно обновлена");
      } catch (e) {
        this.$events.$emit("ERROR", e);
      } finally {
        this.busy = false;
      }
    },

    close() {
      this.visible = false;
    },

    async get(id) {
      try {
        this.busy = true;
        const response = await this.$api.hosts.get(id);
        this.form = response.data;
      } catch (e) {
        this.$events.$emit("error", e);
      } finally {
        this.busy = false;
      }
    },

    show(id) {
      this.visible = true;
      if (id) {
        this.get(id);
      } else {
        this.form = {
          id: null,
          name: null,
          address: null
        };
      }
    }
  }
};
</script>

<style>
</style>
