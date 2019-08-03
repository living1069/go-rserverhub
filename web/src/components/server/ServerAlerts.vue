<template>
  <div>
    <v-alert
      v-if="status === 'UNKNOWN'"
      :value="true"
      class="ma-0 mb-4"
      color="error"
      icon="warning"
      outline
    >Сервер не отвечает. Управление и мониторинг сервера невозможен.</v-alert>
    <v-alert
      v-if="status === 'REMOVING'"
      :value="true"
      class="ma-0 mb-4"
      color="error"
      icon="warning"
      outline
    >Сервер находится в состоянии удаления. Управление и мониторинг сервера невозможен.</v-alert>
  </div>
</template>

<script>
export default {
    props: ["status"]
};
</script>

<style>
</style>
