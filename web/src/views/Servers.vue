<template>
  <div>
    <p class="display-1">Сервера</p>
    <v-btn color="primary" outline class="ml-0 mb-4">Добавить сервер</v-btn>
    <server-list/>
  </div>
</template>

<script>
import ServerList from "@/components/server/ServerList";

export default {
  components: {
    ServerList
  }
};
</script>