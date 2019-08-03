export default {
    methods: {
        stateString(state) {
            if (state === "UP") return "Запущен";
            else if (state === "REMOVING") return "Удяляется";
            else if (state === "UNKNOWN") return "Не отвечает";
            else return "Неизвестно";
        },
        stateIcon(state) {
            if (state === "UP") return "check_circle";
            else if (state === "REMOVING") return "delete";
            else if (state === "UNKNOWN") return "error";
            else return "warning";
        },
        stateColor(state) {
            if (state === "UP") return "green";
            else if (state === "REMOVING") return "red";
            else if (state === "UNKNOWN") return "red";
            else return "red";
        }
    }
}