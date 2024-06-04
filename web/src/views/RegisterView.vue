<template>
  <div class="v-login-box">
    <h1 class="v-login-header">Registration Form</h1>
    <div class="v-input-box">
      <p>
        <svg margin="2px 0 0 0" version="1.0" id="Layer_1" xmlns="http://www.w3.org/2000/svg"
          xmlns:xlink="http://www.w3.org/1999/xlink" width="18px" height="18px" viewBox="0 0 64 64"
          enable-background="new 0 0 64 64" xml:space="preserve" fill="#ffffff">
          <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
          <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
          <g id="SVGRepo_iconCarrier">
            <g>
              <path fill="#ffffff"
                d="M32,34.934L63.617,10.34C62.984,8.965,61.613,8,60,8H4c-1.613,0-2.984,0.965-3.617,2.34L32,34.934z">
              </path>
              <path fill="#ffffff"
                d="M34.457,43.156C33.734,43.719,32.867,44,32,44s-1.734-0.281-2.457-0.844L0,20.18V52c0,2.211,1.789,4,4,4h56 c2.211,0,4-1.789,4-4V20.18L34.457,43.156z">
              </path>
            </g>
          </g>
        </svg>
        Email:
      </p>
      <input class="v-input" type="text" v-model="email" />
    </div>
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
    </div>
    <div class="v-input-box">
      <p>
        <svg xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="18px" height="18px" viewBox="0 0 50 50"
          style="fill: #ffffff">
          <path
            d="M 25 3 C 18.363281 3 13 8.363281 13 15 L 13 20 L 9 20 C 7.300781 20 6 21.300781 6 23 L 6 47 C 6 48.699219 7.300781 50 9 50 L 41 50 C 42.699219 50 44 48.699219 44 47 L 44 23 C 44 21.300781 42.699219 20 41 20 L 37 20 L 37 15 C 37 8.363281 31.636719 3 25 3 Z M 25 5 C 30.566406 5 35 9.433594 35 15 L 35 20 L 15 20 L 15 15 C 15 9.433594 19.433594 5 25 5 Z M 25 30 C 26.699219 30 28 31.300781 28 33 C 28 33.898438 27.601563 34.6875 27 35.1875 L 27 38 C 27 39.101563 26.101563 40 25 40 C 23.898438 40 23 39.101563 23 38 L 23 35.1875 C 22.398438 34.6875 22 33.898438 22 33 C 22 31.300781 23.300781 30 25 30 Z">
          </path>
        </svg>
        Confirm password:
      </p>
      <input class="v-input" type="password" v-model="confirmPassword" />
      <p class="v-error">{{ error }}</p>
    </div>
    <p />
    <button class="v-btn" @click="SignUp">Register</button>
    <p class="v-help">Already have account? <a href="/login">Login here</a></p>
  </div>
</template>

<script>
import { defineComponent, ref } from "vue";
import { validateEmail, validateLogin, validatePassword } from "@/lib";
import AuthService from "@/services/Auth";

export default defineComponent({
  data() {
    const email = ref("");
    const login = ref("");
    const password = ref("");
    const confirmPassword = ref("");
    const error = ref("");

    return {
      email: email,
      login: login,
      password: password,
      confirmPassword: confirmPassword,
      error: error,
    };
  },
  methods: {
    SignUp() {
      this.error = this.validateForm();
      if (this.error !== "") {
        return;
      }

      AuthService.Register(this.email, this.login, this.password)
        .then((resp) => {
          console.log(resp.data);
          this.$router.push({ name: "login" });
        })
        .catch((error) => {
          console.log(error);
          this.error = error.response?.data?.message;
        });
    },
    validateForm() {
      return (
        validateEmail(this.email) ||
        validateLogin(this.login) ||
        validatePassword(this.password) ||
        (this.password !== this.confirmPassword
          ? "passwords are different"
          : "")
      );
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

.v-login-header {
  padding: 40px 0 20px 0;
  font-size: 28px;
}

.v-login-box {
  max-width: 550px;
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

.v-help {
  padding: 0 0 20px 0;
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
