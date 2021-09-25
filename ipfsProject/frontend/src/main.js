import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import Echo from 'laravel-echo';
import { inject } from 'vue'

Vue.use(VueAxios, axios)

new Vue({
    render: h => h(App)
}).$mount('#app');

axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*';
var postform1 = new Vue({
  el:'#postform',
  data: {
    key:'',
    data:'',
    encodedData:'',
    postStatus:''
  },
  methods: {
    sendForm: function() {
      var link='http://localhost:1323/add'
      axios.post(link, {
        key:this.key,
        data:this.data,
      })
      .then(response => {
        if (response.data.status == "200") {
          this.encodedData = response.data.cid,
          this.postStatus = "Success! Status:" + response.data.status
        }
        console.log(response.data)
      })
      .catch(e => {
        console.log(e.response)
        console.log("error")
      })
    }
  }
})
var getform1 = new Vue({
  el:'#getform',
  data: {
    key:'',
    data:'',
    decodedData:''
  },
  methods: {
    getForm: function() {
      var link='http://localhost:1323/get/' + this.data + '/?key=' + this.key
      axios.get(link)
      .then(response => {
        this.decodedData = response.data
        console.log(response.data)
      })
      .catch(e => {
        console.log(e.response)
        console.log("error")
      })
    }
  }
})
