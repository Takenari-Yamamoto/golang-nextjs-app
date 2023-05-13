import axios from "axios";
import { getToken } from "./firebase";

export const apiClient = axios.create({
  baseURL: `${process.env.NEXT_PUBLIC_API_ENDPOINT}`,
});

apiClient.interceptors.request.use(async (config) => {
  const token = await getToken();
  if (token) {
    config.headers["Authorization"] = `Bearer ${token}`;
  }
  return config;
});
