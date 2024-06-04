<template>
  <div class="v-login-box">
    <h1 class="v-login-header">Login to Your Account</h1>
    <div class="v-input-box">
      <p>
        <svg height="18px" width="18px" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg"
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
        Login:
      </p>
      <input class="v-input" type="text" v-model="login" />
    </div>
    <div class="v-input-box">
      <p>
        <svg xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="18px" height="18px" viewBox="0 0 50 50"
          style="fill: #ffffff">
          <path
            d="M 25 3 C 18.363281 3 13 8.363281 13 15 L 13 20 L 9 20 C 7.300781 20 6 21.300781 6 23 L 6 47 C 6 48.699219 7.300781 50 9 50 L 41 50 C 42.699219 50 44 48.699219 44 47 L 44 23 C 44 21.300781 42.699219 20 41 20 L 37 20 L 37 15 C 37 8.363281 31.636719 3 25 3 Z M 25 5 C 30.566406 5 35 9.433594 35 15 L 35 20 L 15 20 L 15 15 C 15 9.433594 19.433594 5 25 5 Z M 25 30 C 26.699219 30 28 31.300781 28 33 C 28 33.898438 27.601563 34.6875 27 35.1875 L 27 38 C 27 39.101563 26.101563 40 25 40 C 23.898438 40 23 39.101563 23 38 L 23 35.1875 C 22.398438 34.6875 22 33.898438 22 33 C 22 31.300781 23.300781 30 25 30 Z">
          </path>
        </svg>
        Password:
      </p>
      <input class="v-input" type="password" v-model="password" />
      <p class="v-error">{{ error }}</p>
    </div>
    <button class="v-btn" @click="SignIn">Login</button>
    <p class="v-help">
      Don't have account? <a href="/register">Register here</a>
    </p>
  </div>
</template>

<script>
import { defineComponent, ref } from "vue";
import { validateLogin, validatePassword } from "@/lib";
import AuthService from "@/services/Auth";

export default defineComponent({
  data() {
    const login = ref("");
    const password = ref("");
    const error = ref("");

    return {
      login: login,
      password: password,
      error: error,
    };
  },
  methods: {
    SignIn() {
      this.error = this.validateForm();
      if (this.error !== "") {
        return;
      }

      AuthService.Login(this.login, this.password)
        .then(() => {
          this.$router.push({ name: "proxy" });
        })
        .catch((error) => {
          console.log(error);
          this.error = error.response?.data?.message;
        });
    },
    validateForm() {
      return validateLogin(this.login) || validatePassword(this.password);
    },
  },
  computed() { },
});
</script>

<style scoped>
.v-error {
  color: #aa2222;

  margin: 5px 0 0 0;
  font-size: 18px;
}

.v-help {
  padding: 0 0 20px 0;
}

.v-login-header {
  padding: 40px 0 20px 0;
  font-size: 28px;
}

.v-login-box {
  max-width: 450px;
  min-width: 430px;

  display: flex;
  flex-direction: column;
  align-self: center;
  background-color: #343434;
  border-radius: 20px;
}

.v-input-box {
  min-width: 300px;
  align-self: center;

  padding: 5px;
}

.v-input-box p {
  margin: 5px 0 5px 0;
  font-size: 20px;
  text-align: left;
}

.v-input {
  min-width: 300px;
  height: 30px;

  padding: 2px 10px 2px 10px;
  font-size: 18px;
}

.v-btn {
  min-width: 75px;
  margin: 10px 10px 10px 10px;

  display: flex;
  justify-content: center;
  align-self: center;
  font-size: 18px;

  background-color: #643cff;
}

.v-btn:hover {
  background-color: #444cff;
}
</style>
