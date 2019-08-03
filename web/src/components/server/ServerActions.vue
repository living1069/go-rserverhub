<template>
  <div>
    <v-layout row align-center>
      <v-switch
        v-model="autorestart"
        :disabled="anyBusy"
        :loading="busy.auto"
        class="flex"
        label="Автоматический перезапуск"
      ></v-switch>

      <v-btn
        :disabled="!canStart"
        :loading="busy.start"
        @click="doStart"
        small
        flat
        fab
        color="green"
        class="my-0 mr-0"
      >
        <v-icon medium big>play_arrow</v-icon>
      </v-btn>
      <v-btn
        :disabled="!canStop"
        :loading="busy.stop"
        @click="doStop"
        small
        flat
        fab
        color="red"
        class="my-0 mr-0"
      >
        <v-icon medium>stop</v-icon>
      </v-btn>
      <v-btn small flat fab color="orange" class="my-0 mr-0" @click="edit">
        <v-icon medium>edit</v-icon>
      </v-btn>
      <v-btn
        small
        flat
        fab
        color="red"
        class="my-0 mr-0"
        :disabled="anyBusy"
        :loading="busy.delete"
        @click="doDelete"
      >
        <v-icon medium>delete</v-icon>
      </v-btn>
    </v-layout>
    <server-dialog ref="dialog"/>
  </div>
</template>

<script>
import ServerDialog from "@/components/server/ServerDialog";
import { Promise } from "q";
import { setTimeout } from "timers";

export default {
  components: { ServerDialog },
  props: ["server"],
  data: () => ({
    busy: {
      start: false,
      stop: false,
      delete: false,
      auto: false
    },
    serverAddDialogVisible: false
  }),
  computed: {
    canStart() {
      return !this.server.session && !this.anyBusy;
    },

    canStop() {
      return this.server.session && !this.anyBusy;
    },

    anyBusy() {
      return (
        this.busy.start || this.busy.stop || this.busy.delete || this.busy.auto
      );
    },

    autorestart: {
      get() {
        return this.server.autorestart;
      },
      set(value) {
        this.doAuto(value);
        //this.server.autorestart = value;
      }
    }
  },

  methods: {
    async doStart() {
      try {
        this.busy.start = true;
        await this.timeout(500);
        await this.$api.servers.start(this.server.id);
        this.$events.$emit("MESSAGE", "Сервер запускается");        
      } catch (e) {
        this.$events.$emit("ERROR", e);
      } finally {
        this.busy.start = false;
      }
    },

    async doStop() {
      try {
        this.busy.stop = true;
        await this.timeout(500);
        await this.$api.servers.stop(this.server.id);
        this.$events.$emit("MESSAGE", "Сервер остановлен");
        this.server.session = null;
      } catch (e) {
        this.$events.$emit("ERROR", e);
      } finally {
        this.busy.stop = false;
      }
    },

    async doAuto(value) {
      try {
        this.busy.auto = true;
        this.server.autorestart = value;
        await this.timeout(500);      
        await this.$api.servers.save(this.server);
        this.$events.$emit("MESSAGE", "Сервер обновлен");
      } catch (e) {
        this.$events.$emit("ERROR", e);
      } finally {
        this.busy.auto = false;
      }
    },

    edit() {
      this.$refs.dialog.open(this.server.id, this.server.host);
    },

    async doDelete() {
      try {
        this.busy.delete = true;
        await this.$api.servers.delete(this.server.id);
      } catch (e) {
        this.$events.$emit("ERROR", e);
      } finally {
        this.busy.delete = false;
      }
    },

    async timeout(val) {
      return new Promise(resolve => setTimeout(() => resolve(), val));
    }
  }
};
</script>

<style>
</style>
