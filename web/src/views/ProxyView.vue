<script setup>
import NavBar from "@/components/layouts/NavBar.vue";
</script>

<template>
  <div class="v-proxy-box">
    <NavBar />
    <div class="v-account-request-box">
      <h2>Proxies</h2>
      <h3 v-if="proxies.length === 0">Not found</h3>
      <div v-else class="v-account-requests">
        <div class="v-account-request-info">
          <p>Address</p>
          <p>Status</p>
          <p>Started</p>
        </div>
        <div v-for="(proxy, index) in proxies">
          <div v-if="index % 2 === 1" class="v-account-request" style="background-color: #262626">
            <p>{{ proxy.address }}</p>
            <p>{{ proxy.status }}</p>
            <p>{{ new Date(proxy.started_at).toUTCString() }}</p>
          </div>
          <div v-else class="v-account-request" style="background-color: #343434">
            <p>{{ proxy.address }}</p>
            <p>{{ proxy.status }}</p>
            <p>{{ new Date(proxy.started_at).toUTCString() }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent, ref } from "vue";
import ProxyService from "@/services/Proxy";

export default defineComponent({
  components: {
    NavBar,
  },
  data() {
    const proxies = ref([]);

    ProxyService.ProxyInfo()
      .then((resp) => {
        proxies.value = resp.data.proxies;
      })
      .catch((error) => {
        console.log(error);
      });
    return {
      proxies: proxies,
    };
  },
  methods: {},
  computed() { },
});
</script>

<style scoped>
.v-proxy-box {
  width: 100%;
}

.v-account-requests {
  background-color: #343434;
  border-radius: 20px;
}

.v-account-request-info {
  margin: 20px 0 20px 0;

  width: 100%;

  display: flex;
  justify-content: space-around;

  font-weight: 600;
}

.v-account-request {
  padding: 8px 0 8px 0;
  height: 75px;
  width: 100%;

  display: flex;
  justify-content: space-around;
  align-items: center;
  background-color: #303030;
  word-break: break-all;

  border-radius: 20px;
}

.v-account-request p {
  width: 150px;
}

.v-account-request-info p {
  width: 150px;
}

.v-account-request-box {
  margin: 20px 0 0 0;
  width: 100%;

  padding: 10px;
  display: flex;
  flex-direction: column;
  justify-content: center;

  background-color: #343434;
  box-shadow: 10px 8px 8px #282828;
  border-radius: 20px;
}
</style>
