<template>
  <div class="flex-grow">
    <div v-if="!ready" class="v--loader-wrapper">
      <div v-if="!error" class="flex-center flex-column">
        <v-progress-circular :width="4" :size="64" indeterminate></v-progress-circular>
        <p class="headline font-weight-thin mt-4">{{status}}</p>
      </div>
      <div v-else class="flex-center flex-column">
        <v-icon style="font-size: 96px" color="red">error</v-icon>
        <p class="headline red--text font-weight-light mt-4">{{errorMessage}}</p>
      </div>
    </div>
    <v-slide-x-transition leave-absolute>
      <slot v-if="ready && !error"></slot>
    </v-slide-x-transition>
  </div>
</template>

<script>
export default {
  props: {
    ready: {
      type: Boolean,
      default: false
    },
    error: {
      type: Boolean,
      default: false
    },
    errorMessage: {
      type: String,
      default: "Произошла ошибка"
    },
    status: {
      type: String,
      default: "Загрузка"
    }
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
</style>
