<script>
import { Line } from "vue-chartjs";

export default {
  extends: Line,
  props: ["dataset"],
  data: () => ({
    enable: false,
    options: {
      responsive: true,
      maintainAspectRatio: false,
      aspectRatio: 10,
      legend: {
        display: false
      },
      tooltips: {
        enabled: false
      },
      elements: {
        point: {
          radius: 0
        },
        line: {
          borderColor: "rgba(0,0,0,0)"
        }
      },
      scales: {
        xAxes: [
          {
            type: "time",
            time: {
              unit: "second",
              stepSize: 25
            },
            bounds: "ticks"
          }
        ],
        yAxes: [
          {
            ticks: {
              stepSize: 20
            }
          }
        ]
      }
    }
  }),
  created() {
    const self = this;
    setInterval(() => (self.enable ? self.$data._chart.update() : null), 1000);
    this.$events.$on("server-view-tab-change", i => {
      if (i == 2)
        setTimeout(() => self.renderChart(self.dataset, self.options), 100);
      self.enable = i == 2;
    });
  }
};
</script>
