<template>
  <div>
    <v-alert
      v-if="status === 'UNKNOWN'"
      :value="true"
      class="mb-4"
      color="error"
      icon="warning"
      outline
    >Хост не отвечает. Управление серверами невозможно.</v-alert>

    <v-alert
      v-if="status === 'REMOVING'"
      :value="true"
      class="mb-4"
      color="error"
      icon="warning"
      outline
    >Хост удаляется. Управление серверами невозможно.</v-alert>
  </div>
</template>

<script>
export default {
  props: ["status"]
};
</script>

<style>
</style>
