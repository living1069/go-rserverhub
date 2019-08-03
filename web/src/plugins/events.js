export default {
    install: function (Vue) {    
      Object.defineProperty(Vue.prototype, '$events', { value: new Vue() });
    },
  };  