<template>
  <v-layout>
    <v-flex xs12 :class="paddings">
      <slot/>
    </v-flex>
    <v-divider v-if="!right" vertical/>
  </v-layout>
</template>

<script>
export default {
  props: {
    left: {
      type: Boolean
    },
    right: {
      type: Boolean
    }
  },
  computed: {
    paddings() {
      if (this.left) return "pr-4";
      else if (this.right) return "pl-4";
      else return "px-4";
    }
  }
};
</script>
