import axiosApi from "../http";

export default class AuthService {
  static async Register(email, login, password) {
    return axiosApi.post("/api/v1/users/register", {
      user: {
        email: email,
        login: login,
        password: password,
      },
    });
  }

  static async Login(login, password) {
    return axiosApi
      .post("/api/v1/users/login", {
        user: {
          login: login,
          password: password,
        },
      })
      .then((resp) => {
        localStorage.setItem("access_token", resp.data.token);
        localStorage.setItem("is_auth", true);
        return resp;
      });
  }
}
