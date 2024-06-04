<template>
  <div class="v-navbar">
    <div class="v-navigation">
      <h1 class="v-logo">Proxy Manager</h1>
      <div class="v-links">
        <a class="v-link" href="/proxy">Proxies</a>
      </div>
    </div>
    <div class="v-user-box">
      <div class="v-user">
        <a href="/account">
          <div class="v-user-login">
            <svg height=" 26px" width="26px" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg"
              xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 60.671 60.671" xml:space="preserve" fill="#ffffff"
              stroke="#ffffff">
              <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
              <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
              <g id="SVGRepo_iconCarrier">
                <g>
                  <g>
                    <ellipse style="fill: #ffffff" cx="30.336" cy="12.097" rx="11.997" ry="12.097"></ellipse>
                    <path style="fill: #ffffff"
                      d="M35.64,30.079H25.031c-7.021,0-12.714,5.739-12.714,12.821v17.771h36.037V42.9 C48.354,35.818,42.661,30.079,35.64,30.079z">
                    </path>
                  </g>
                </g>
              </g>
            </svg>
            {{ login }}
          </div>
        </a>
        <button class="v-btn" @click="LogOut">Log out</button>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent, ref } from "vue";
import UserService from "@/services/User";

export default defineComponent({
  data() {
    const login = ref("");

    UserService.Account()
      .then((resp) => {
        console.log(resp);
        login.value = resp.data.user.login;
      })
      .catch((error) => {
        console.log(error);
      });

    return {
      login: login,
    };
  },
  mounted() { },
  methods: {
    async UserData() { },
    Account() {
      this.$router.push({ name: "account" });
    },
    Login() {
      this.$router.push({ name: "login" });
    },
    Register() {
      this.$router.push({ name: "register" });
    },
    LogOut() {
      localStorage.removeItem("is_auth");
      localStorage.removeItem("access_token");
      this.$router.push({ name: "login" });
    },
  },
  computed() { },
});
</script>

<style scoped>
.v-navbar {
  width: 100%;
  padding: 10px;

  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #343434;
  box-shadow: 10px 8px 8px #282828;
  border-radius: 20px;
}

.v-navigation {
  margin: 10px 0 0 0;
  display: flex;
  justify-content: space-evenly;
}

.v-logo {
  margin: 5px 0 0 0;

  font-size: 36px;
}

.v-links {
  margin: 10px 30px 30px 10px;

  display: flex;
  justify-content: space-evenly;

  font-size: 24px;
}

.v-link {
  margin: 0 0px 0 30px;
}

.v-user-box {}

.v-user {
  display: flex;
}

.v-user-login {
  margin: 10px 20px 10px 20px;

  font-size: 22px;
}

.v-btn {
  min-width: 75px;
  margin: 10px 10px 10px 10px;

  display: flex;
  justify-content: center;
  align-self: center;
  font-size: 14px;

  background-color: #343434;
  border-color: #643cff;
}

.v-btn:hover {
  background-color: #444cff;
}
</style>
