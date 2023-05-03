import { taskRepository } from "@/api/taskRepository";
import { apiClient } from "@/libs/api-client";
import { useState } from "react";

export const useTask = () => {
  const [tasks, setTasks] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const taskRepo = taskRepository(apiClient);

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

  return {
    tasks,
    fetchAllTasks,
    loading,
    error,
  };
};
