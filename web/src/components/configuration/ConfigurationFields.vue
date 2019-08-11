<template>
  <v-layout column>
    <v-text-field
      label="Название конфигурации"
      :disabled="busy"
      :error-messages="validation.name"
      v-model="data.name"
      prepend-icon="title"
    ></v-text-field>
    <v-textarea
      label="Путь к исполняемому файлу"
      v-model="data.executable"
      :error-messages="validation.executable"
      :disabled="busy"
      prepend-icon="extension"
      rows="1"
      auto-grow
    ></v-textarea>
    <v-textarea
      label="Путь к рабочей директории"
      v-model="data.directory"
      :error-messages="validation.directory"
      :disabled="busy"
      prepend-icon="folder"
      rows="1"
      auto-grow
    ></v-textarea>
    <v-textarea
      label="Путь к обновляемым ресурсам игры"
      v-model="data.resources"
      :error-messages="validation.resources"
      :disabled="busy"
      prepend-icon="folder"
      rows="1"
      auto-grow
    ></v-textarea>
    <v-textarea
      label="Путь к файлу fsgame.ltx"
      :disabled="busy"
      v-model="data.fsgame"
      :error-messages="validation.fsgame"
      prepend-icon="gesture"
      rows="1"
      auto-grow
    ></v-textarea>
    <v-textarea
      label="Аргументы запуска"
      :disabled="busy"
      v-model="data.arguments"
      :error-messages="validation.arguments"
      prepend-icon="code"
      :rows="2"
      auto-grow
    ></v-textarea>
    <v-text-field
      label="Порт сервера"
      :disabled="busy"
      :error-messages="validation.port"
      v-model="data.port"
      type="number"
      prepend-icon="dns"
    ></v-text-field>
    <v-text-field
      label="Карта"
      :disabled="busy"
      :error-messages="validation.map"
      v-model="data.map"
      prepend-icon="layers"
    ></v-text-field>
  </v-layout>
</template>

<script>
import { required, numeric } from "vuelidate/lib/validators";

export default {
  props: ["data", "busy"],
  validations: {
    data: {
      name: { required },
      executable: { required },
      directory: { required },
      fsgame: { required },
      arguments: { required },
      resources: { required },
      port: { required, numeric },
      map: { required }
    }
  },
  computed: {
    validation() {
      const ret = {};
      if (this.$v.data.name.$dirty && !this.$v.data.name.required)
        ret.name = "Поле не может быть пустым";

      if (this.$v.data.executable.$dirty && !this.$v.data.executable.required)
        ret.executable = "Поле не может быть пустым";

      if (this.$v.data.directory.$dirty && !this.$v.data.directory.required)
        ret.directory = "Поле не может быть пустым";

      if (this.$v.data.resources.$dirty && !this.$v.data.resources.required)
        ret.directory = "Поле не может быть пустым";

      if (this.$v.data.name.$dirty && !this.$v.data.fsgame.required)
        ret.fsgame = "Поле не может быть пустым";

      if (this.$v.data.arguments.$dirty && !this.$v.data.arguments.required)
        ret.arguments = "Поле не может быть пустым";

      if (this.$v.data.port.$dirty && !this.$v.data.port.required)
        ret.port = "Поле не может быть пустым";

      if (this.$v.data.port.$dirty && !this.$v.data.port.numeric)
        ret.port = "Поле должно содержать только цифры";

      if (this.$v.data.name.$dirty && !this.$v.data.map.required)
        ret.map = "Поле не может быть пустым";

      return ret;
    }
  },
  methods: {
    validate() {
      this.$v.$touch();
      return !this.$v.$invalid;
    },
    reset() {
      this.$v.$reset();
    }
  }
};
</script>

<style>
</style>
