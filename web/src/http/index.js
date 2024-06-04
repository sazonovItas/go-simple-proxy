import axios from "axios";
import router from "@/router";

const token = "access_token";
const isAuth = "is_auth";

const axiosApi = axios.create({
  withCredentials: true,
  baseURL: import.meta.env.VITE_BACKEND_URL,
});

axiosApi.interceptors.request.use((config) => {
  config.headers.Authorization = `Bearer ${localStorage.getItem(token)}`;
  return config;
});

axiosApi.interceptors.response.use(
  (config) => {
    return config;
  },
  async (error) => {
    if (error.response.status === 401) {
      localStorage.removeItem(token);
      localStorage.removeItem(isAuth);
      router.push({ name: "login" });
    }

    throw error;
  },
);

export default axiosApi;
