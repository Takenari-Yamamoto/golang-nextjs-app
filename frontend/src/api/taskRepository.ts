import { AxiosInstance, AxiosResponse } from "axios";

export type Task = {
  id: string;
  title: string;
  content: string;
  created_by: string;
  created_at: string;
};

export type CreateTaskParams = {
  title: string;
  content: string;
};

interface TaskRepository {
  getTasks(): Promise<any>;
  getTaskById(id: string): Promise<AxiosResponse<Task>>;
  createTask(task: CreateTaskParams): Promise<Task>;
  updateTask(task: Task): Promise<Task>;
  deleteTask(id: string): Promise<void>;
}

export const taskRepository = (client: AxiosInstance): TaskRepository => {
  return {
    getTasks: (): Promise<Task> => client.get("/tasks"),
    getTaskById: (id: string): Promise<AxiosResponse<Task>> =>
      client.get(`/tasks/${id}`),
    createTask: (task: CreateTaskParams): Promise<Task> =>
      client.post("/task", task),
    updateTask: (task: Task): Promise<Task> =>
      client.patch(`/tasks/${task.id}`, task),
    deleteTask: (id: string): Promise<void> => client.delete(`/tasks/${id}`),
  };
};
