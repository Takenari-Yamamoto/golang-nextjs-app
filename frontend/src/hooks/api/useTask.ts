import { CreateTaskParams, Task, taskRepository } from "@/api/taskRepository";
import { apiClient } from "@/libs/api-client";
import { useState } from "react";

export const useTask = () => {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [task, setTask] = useState<Task>();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const taskRepo = taskRepository(apiClient);

  const createTask = async (task: CreateTaskParams) => {
    setLoading(true);
    try {
      await taskRepo.createTask(task);
      return "success";
    } catch (e) {
      alert(e);
      setError(e as string);
    }
    setLoading(false);
  };

  const fetchAllTasks = async () => {
    setLoading(true);
    try {
      const res = await taskRepo.getTasks();

      setTasks(res.data);
    } catch (e) {
      setError(e as string);
    }
    setLoading(false);
  };

  // 詳細を取得する
  const fetchTaskById = async (id: string) => {
    setLoading(true);
    try {
      const res = await taskRepo.getTaskById(id);
      setTask(res.data);
    } catch (e) {
      setError(e as string);
    }
    setLoading(false);
  };

  // 削除
  const deleteTask = async (id: string) => {
    setLoading(true);
    try {
      await taskRepo.deleteTask(id);
      return "success";
    } catch (e) {
      alert(e);
      setError(e as string);
    }
    setLoading(false);
  };

  return {
    loading,
    error,
    tasks,
    task,
    fetchAllTasks,
    fetchTaskById,
    createTask,
    deleteTask,
  };
};
