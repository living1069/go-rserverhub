export default {
  methods: {
      stateString(session) {
          if (!session) return 'Не запущен';
          switch (session.status) {
        case "LOADING":
          return "Запускается";
        case "UP":
            return "Запущен";
        case "DOWN":
          return "Остановлен";
        case "ERROR":
            return "Ошибка";
      }
    },
      stateIcon(session) {
          if (!session) return 'stop';
          switch (session.status) {
        case "LOADING":
          return "sync";
        case "UP":
          return "check_circle";
        case "DOWN":
          return "stop";
        case "ERROR":
            return "error";
      }
    },
      stateColor(session) {
          if (!session) return 'blue-grey';
          switch (session.status) {
        case "LOADING":
          return "orange";
        case "UP":
            return "green";
        case "DOWN":
          return "blue-grey";
        case "ERROR":
            return "error";
      }
    }
  }
}