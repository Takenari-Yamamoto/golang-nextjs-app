import { AxiosInstance } from "axios";

export type Task = {
  content: string;
  created_at: string;
  created_by: string;
  id: string;
  title: string;
};

interface TaskRepository {
  getTasks(): Promise<any>;
  //   getTask(id: number): Promise<Task>;
  //   createTask(task: Task): Promise<Task>;
  //   updateTask(task: Task): Promise<Task>;
  //   deleteTask(id: number): Promise<void>;
}

export const taskRepository = (client: AxiosInstance): TaskRepository => {
  return {
    getTasks: async (): Promise<Task> => await client.get("/tasks"),
  };
};
