import axiosApi from "../http";

export default class ProxyService {
  static async ProxyInfo() {
    return axiosApi.get("/api/v1/proxy");
  }
}
