<template>
  <div>
    <v-layout row align-center>
      <v-btn
        :loading="busy.auto"
        :disabled="anyBusy"
        @click="autorestart = !autorestart"
        small
        flat
        fab
        :color="autorestart ? 'green' : 'grey'"
        class="my-0 mr-0"
      >
        <v-icon medium big>{{autorestart ? 'toggle_on' : 'toggle_off'}}</v-icon>
      </v-btn>

      <v-btn
        :loading="busy.update"
        :disabled="updateDialogVisible"
        @click="doUpdate"
        small
        flat
        fab
        color="primary"
        class="my-0 mr-0"
      >
        <v-icon medium big>refresh</v-icon>
      </v-btn>

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
    <server-dialog ref="dialog" />
    <v-dialog v-model="updateDialogVisible" persistent hide-overlay width="600">
      <v-card :color="updateStatus ? 'blue darken-1' : 'red darken-1'" dark>
        <v-card-title class="headline pb-0">Обновление ресурсов</v-card-title>
        <v-card-text class="pb-0">
          <v-progress-linear v-if="busy.update" indeterminate color="white" class="mb-0"></v-progress-linear>
          <v-sheet
            v-if="!busy.update"
            class="pa-2"
           :color="updateStatus ? 'blue darken-4' : 'red darken-4'"
            style="max-height:400px; overflow-y:auto"
          >
            <p v-for="(p, i) in updateResult" class="mb-0" :key="i" v-text="p" />
          </v-sheet>
        </v-card-text>
        <v-card-actions>
          <v-layout justify-end>
            <v-btn flat @click="updateDialogVisible = false">Закрыть</v-btn>
          </v-layout>
        </v-card-actions>
      </v-card>
    </v-dialog>
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
      auto: false,
      update: false
    },
    serverAddDialogVisible: false,
    updateDialogVisible: false,
    updateStatus: true,
    updateResult: []
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
        this.updateResult = false;
        this.up;
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

    async doUpdate() {
      try {
        this.busy.update = true;
        this.updateStatus = true;
        this.updateDialogVisible = true;
        const response = await this.$api.servers.update(this.server.id);
        this.updateResult = response.data.buffer;
        this.updateStatus = response.data.exitCode == 0;
      } catch (e) {
        this.updateDialogVisible = false;
        this.$events.$emit("ERROR", e);
      } finally {
        this.busy.update = false;
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
