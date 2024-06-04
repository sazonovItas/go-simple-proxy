<script setup>
import NavBar from "@/components/layouts/NavBar.vue";
</script>

<template>
  <div class="v-account-box">
    <NavBar />
    <div class="v-account-info-box"></div>
    <div class="v-account-info">
      <h2>Stats</h2>
      <div class="v-account-request-info">
        <p>Requests created</p>
        <p>Upload</p>
        <p>Download</p>
      </div>
      <div class="v-account-request" style="background-color: #262626">
        <p>{{ totalStats.cnt }}</p>
        <p>{{ prettyBytes(totalStats.upload) }}</p>
        <p>{{ prettyBytes(totalStats.download) }}</p>
      </div>
    </div>
    <div class="v-account-request-box">
      <h2>Proxy Requests</h2>
      <h3 v-if="requests.length === 0">Not found</h3>
      <div v-else class="v-account-requests">
        <div class="v-account-request-info">
          <p>Host</p>
          <p>Remote IP</p>
          <p>Upload</p>
          <p>Download</p>
        </div>
        <div v-for="(request, index) in requests">
          <div v-if="index % 2 === 1" class="v-account-request" style="background-color: #262626">
            <p>{{ request.host }}</p>
            <p>{{ request.remote_ip }}</p>
            <p>{{ prettyBytes(request.upload) }}</p>
            <p>{{ prettyBytes(request.download) }}</p>
          </div>
          <div v-else class="v-account-request" style="background-color: #343434">
            <p>{{ request.host }}</p>
            <p>{{ request.remote_ip }}</p>
            <p>{{ prettyBytes(request.upload) }}</p>
            <p>{{ prettyBytes(request.download) }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent, ref } from "vue";
import UserService from "@/services/User";
import { prettyBytes } from "@/lib";

export default defineComponent({
  components: {
    NavBar,
  },
  data() {
    const requests = ref([]);

    UserService.Requests()
      .then((resp) => {
        requests.value = resp.data.requests;
      })
      .catch((error) => {
        console.log(error);
      });

    return {
      requests: requests,
    };
  },
  methods: {
    AccountRequests() { },
  },
  computed: {
    totalStats() {
      return {
        upload: this.requests.reduce((sum, req) => sum + req.upload, 0),
        download: this.requests.reduce((sum, req) => sum + req.download, 0),
        cnt: this.requests.length,
      };
    },
  },
});
</script>

<style scoped>
.v-account-box {
  width: 100%;
}

.v-account-info {
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
  width: 125px;
}

.v-account-request-info p {
  width: 100px;
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
