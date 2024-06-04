import axiosApi from "../http";

export default class UserService {
  static async Account() {
    return axiosApi.get("/api/v1/user/account");
  }

  static async Requests() {
    return axiosApi.get("/api/v1/user/request");
  }
}
