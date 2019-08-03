<template>
  <v-app :dark="$store.getters.dark">
    <v-layout fill-height>
      <app v-if="$store.getters.endpoint"/>
      <backend-dialog v-if="!$store.getters.endpoint"/>
    </v-layout>
    <snackbar/>
  </v-app>
</template>

<script>
import Snackbar from "@/components/Snackbar";
import App from "@/components/App";
import BackendDialog from "@/components/BackendDialog";

export default {
  components: {
    App,
    Snackbar,
    BackendDialog
  },
  mounted() {
    if (this.$store.getters.endpoint == null)
      this.$events.$emit("SHOW_BACKEND_DIALOG");
  }
};
</script>

<style>
.v--loader-wrapper {
  display: flex;
  height: 100%;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}

.v--loader-content {
  display: flex;
  justify-content: center;
  align-items: center;
}

.h-100 {
  height: 100%;
}
</style>

