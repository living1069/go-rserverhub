<template>
  <div class="flex-center">
    <span class="display-1 mb-2">Хост #{{host.id}}</span>
    <v-spacer/>
    <v-btn small flat fab color="orange" class="my-0" @click="edit">
      <v-icon style="font-size: 24px">edit</v-icon>
    </v-btn>
    <v-btn
      small
      flat
      fab
      color="red"
      class="my-0 mr-0"
      @click="doDelete"
      :disabled="deleteBusy"
      :loading="deleteBusy"
    >
      <v-icon style="font-size: 24px">delete</v-icon>
    </v-btn>
    <host-dialog/>
  </div>
</template>

<script>
import HostDialog from "@/components/host/HostDialog";
export default {
  components: { HostDialog },
  props: ["host"],
  data: () => ({
    deleteBusy: false
  }),
  methods: {
    edit() {
      this.$events.$emit("SHOW_HOST_DIALOG", this.host.id);
    },
    async doDelete() {
      const message = {
        title: "Потверждение",
        message:
          "Вы уверены, что хотите удалить хост? Все связанные ресурсы будут удалены.",
        action: "MODAL_ACTION_YESNO"
      };
      try {
        await this.$store.dispatch("MODAL_SHOW", message);
        try {
          this.deleteBusy = true;
          await this.$api.hosts.delete(this.host.id);
        } catch (e) {
          this.$events.$emit("ERROR", e);
        } finally {
          this.deleteBusy = false;
        }
      } catch (e) {
        //
      }
    }
  }
};
</script>

<style>
</style>
