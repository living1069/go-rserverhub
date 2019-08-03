<template>
  <request-loader :ready="ready" class="h-100">
    <div v-if="host">      
      <host-alerts :status="host.status"/>
      <host-actions :host="host"/>
      <host-info :host="host" class="my-4"/>
    </div>
  </request-loader>
</template>

<script>
import HostInfo from "@/components/host/HostInfo";
import HostAlerts from "@/components/host/HostAlerts";
import HostActions from "@/components/host/HostActions";

export default {
  components: {
    HostInfo,    
    HostAlerts,
    HostActions
  },
  props: ["id"],
  data: () => ({    
    ready: false,
    host: null
  }),
  mounted() {
    this.get();
    const self = this;
    this.$events.$on("HOST_UPDATE", this.update.bind(this));
    this.$events.$on("HOST_DELETE", id => {
      if (self.host.id == id) self.$router.push("/hosts");
    });
  },
  methods: {
    async get() {
      try {
        this.ready = false;
        const result = await this.$api.hosts.get(this.id);
        this.host = result.data;
      } catch (e) {
        this.$events.$emit("ERROR", e);
      } finally {
        this.ready = true;
      }
    },

    update(host) {
      if (this.host.id == host.id) this.host = host;
    }
  }
};
</script>

<style>
</style>
