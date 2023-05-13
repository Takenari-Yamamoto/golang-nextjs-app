import axios from "axios";
import { getToken } from "./firebase";

export const apiClient = axios.create({
  baseURL: "http://localhost:8000/v1",
  // baseURL: "https://golang-nextjs-deploy-umeslybd3q-an.a.run.app/v1",
});

apiClient.interceptors.request.use(async (config) => {
  const token = await getToken();
  if (token) {
    config.headers["Authorization"] = `Bearer ${token}`;
  }
  return config;
});
