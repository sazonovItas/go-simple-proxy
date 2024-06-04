import { createRouter, createWebHistory } from "vue-router";

import LoginView from "@/views/LoginView.vue";
import RegisterView from "@/views/RegisterView.vue";
import ProxyView from "@/views/ProxyView.vue";
import UserView from "@/views/UserView.vue";

const authGuard = (to, from, next) => {
  if (
    !["login", "register"].includes(to.name) &&
    !localStorage.getItem("is_auth")
  ) {
    next({ name: "login" });
  } else if (
    ["login", "register"].includes(to.name) &&
    localStorage.getItem("is_auth")
  ) {
    next({ name: from.name });
  } else next();
};

const routes = [
  {
    path: "/register",
    name: "register",
    component: RegisterView,
  },
  {
    path: "/login",
    name: "login",
    component: LoginView,
  },
  {
    path: "/:path(.*)*",
    name: "proxy",
    component: ProxyView,
  },
  {
    path: "/account",
    name: "account",
    component: UserView,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes: routes,
});

router.beforeEach(authGuard);

export default router;
