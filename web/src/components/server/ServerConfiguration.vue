<template>
  <v-list two-line>
    <v-list-tile v-for="(item, i) in list" :key="i" avatar class="v--remove-padding">
      <v-list-tile-avatar>
        <v-icon :class="[item.iconClass]">{{ item.icon }}</v-icon>
      </v-list-tile-avatar>
      <v-list-tile-content>
        <v-list-tile-title>{{ item.title }}</v-list-tile-title>
        <v-list-tile-sub-title>{{ item.value }}</v-list-tile-sub-title>
      </v-list-tile-content>
    </v-list-tile>
  </v-list>
</template>

<script>
export default {
  props: ["configuration"],
  computed: {
    list() {
      return [
        {
          icon: "extension",
          iconClass: "green lighten-1 white--text",
          title: "Путь к исполняемому файлу",
          value: this.configuration.executable
        },
        {
          icon: "folder",
          iconClass: "purple lighten-1 white--text",
          title: "Путь к рабочей папке файлу",
          value: this.configuration.directory
        },
        {
          icon: "gesture",
          iconClass: "indigo lighten-1 white--text",
          title: "Путь к файлу fsgame.ltx",
          value: this.configuration.fsgame
        },
        {
          icon: "code",
          iconClass: "blue lighten-1 white--text",
          title: "Параметры запуска",
          value: this.configuration.arguments
        },
        {
          icon: "dns",
          iconClass: "cyan lighten-1 white--text",
          title: "Порт",
          value: this.configuration.port
        },
        {
          icon: "layers",
          iconClass: "lime lighten-1 white--text",
          title: "Карта",
          value: this.configuration.map
        }        
      ];
    }
  }
};
</script>

<style>
.v--remove-padding > .v-list__tile {
  padding: 0px !important;
}
</style>
