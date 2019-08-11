<template>
  <div :class="`flex-column flex-grow`">
    <div v-if="!ready" class="console-loader-wrapper flex-grow">
      <v-progress-circular :width="4" :size="64" indeterminate></v-progress-circular>
      <p class="headline font-weight-thin mt-4">Подключение к серверу</p>
    </div>
    <div v-show="ready" id="log-container" class="flex-column flex-grow pa-2">
      <div id="log-content" class="log-content" v-show="logs.length > 0">
        <p :class="`mb-0 body-1 ${logColor(log)}`" v-for="(log, id) in logs" :key="id">{{log}}</p>
      </div>
      <v-layout fill-height justify-center align-center v-show="logs.length == 0">
        <v-icon large>watch_later</v-icon>
      </v-layout>

      <div class="log-command">
        <div class="mt-2">
          <v-divider />
          <v-text-field
            class="console-input pa-0"
            v-model="command"
            :disabled="commandBusy"
            :loading="commandBusy"
            prepend-icon="arrow_right"
            flat
            label="Консоль"
            single-line
            @keyup="send"
          ></v-text-field>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ["server"],
  data: () => ({
    command: null,
    commandBusy: false,
    ready: false,
    logs: [],
    ws: null
  }),
  mounted() {
    const container = window.document.getElementById("log-content");
    container.addEventListener("DOMNodeInserted", () => {
      container.scrollTop = container.scrollHeight;
    });
    this.fetch();
  },
  methods: {
    logColor(log) {
      if (
        log.startsWith("!") ||
        log.startsWith("FATAL ERROR") ||
        log.startsWith("[error]")
      )
        return "red--text text--lighten-1";
      else if (log.startsWith("*")) return "grey--text text--darken--2";
      else if (log.startsWith("-")) return "green--text text--darken--2";
      else if (log.startsWith("@")) return "indigo--text text--accent-2";
      else return "";
    },
    async send(event) {
      if (!(event.keyCode == 13 && this.command)) return;

      try {
        this.commandBusy = true;
        await this.timeout(250);
        await this.$api.servers.command(this.server.id, this.command);
        this.command = null;
      } catch (e) {
        this.$events.$emit("ERROR", e);
      } finally {
        this.commandBusy = false;
      }
    },

    log(e) {      
      const logs = this.logs.concat(JSON.parse(e.data));
      while (logs.length > 150) logs.shift();
      this.logs = logs;
    },

    async fetch() {
      try {
        this.ready = false;
        const result = await this.$api.servers.logs(this.server.id);
        this.logs = result.data;
        let ws = new WebSocket(
          `ws://${this.$store.getters.endpoint}/ws/queue/logs/${this.server.id}`
        );
        ws.onmessage = this.log.bind(this);
        this.ws = ws;
      } catch (e) {
        this.$events.$emit("ERROR", e);
      } finally {
        this.ready = true;
      }
    },

    async timeout(val) {
      return new Promise(resolve => setTimeout(() => resolve(), val));
    }
  },
  beforeDestroy() {
    this.ws.close();
  }
};
</script>

<style>
.console-input {
  margin-left: -8px !important;
}

.console-input > .v-input__prepend-outer {
  margin-right: 0px !important;
}
.console-input > .v-input__control > .v-input__slot {
  padding-left: 0px !important;
  margin-bottom: 0px !important;
}

.console-input > .v-input__control > .v-text-field__details {
  display: none;
}

.console-loader-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}

.log-content {
  flex: 1 1 auto;
  position: relative; /* need this to position inner content */
  overflow-y: auto;
}

.log-command {
  flex: 0 0 auto;
}

#log-container {
  height: 1px;
}
</style>
