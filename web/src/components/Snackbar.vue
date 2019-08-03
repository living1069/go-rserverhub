<template>
  <v-snackbar v-model="snackbar" bottom right :timeout="5000">
    {{ message }}
    <v-btn color="white" flat @click="snackbar = false">Закрыть</v-btn>
  </v-snackbar>
</template>

<script>
export default {
  data: () => ({
    snackbar: false,
    message: ""
  }),
  created() {
    const self = this;
    this.$events.$on("ERROR", error => {
      console.error(error);
      let message = self._.get(error, "response.data.message");
      if (!message) message = "Ошибка запроса";
      self.message = message;
      self.snackbar = true;
    });

    this.$events.$on("MESSAGE", message => {
      self.message = message;
      self.snackbar = true;
    });
  }
};
</script>

<style>
.v-snack__content {
  /* height: 100% !important;*/
}
</style>
