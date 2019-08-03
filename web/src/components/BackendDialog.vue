<template>
  <v-dialog persistent width="500" v-model="visible">
    <v-card>
      <v-card-title class="headline">Укажите адрес сервера</v-card-title>
      <v-divider/>
      <v-card-text class="px-4 pb-0">
        <v-text-field v-model="url" label="Адрес" placeholder="localhost:8080"></v-text-field>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn flat @click="save">Сохранить</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  data: () => ({
    url: null,
    visible: false
  }),
  created() {
      const self = this;
      this.$events.$on('SHOW_BACKEND_DIALOG', () => {
          self.url = self.$store.getters.endpoint;
          self.visible = true;
      })
  },
  methods: {
    save() {
      const self = this;
      setTimeout(() => {
        self.$store.dispatch("SET_ENDPOINT", self.url);
        location.reload();
      }, 500);
    }
  }
};
</script>

<style>
</style>
