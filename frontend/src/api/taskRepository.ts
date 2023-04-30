import { apiClient } from "@/libs/api-client";
import { AxiosInstance } from "axios";

interface TaskRepository {
  getTasks(): Promise<any>;
  //   getTask(id: number): Promise<Task>;
  //   createTask(task: Task): Promise<Task>;
  //   updateTask(task: Task): Promise<Task>;
  //   deleteTask(id: number): Promise<void>;
}

export const taskRepository = (client: AxiosInstance): TaskRepository => {
  return {
    getTasks: async (): Promise<any> => await client.get("/tasks"),
  };
};
