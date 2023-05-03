import { useTask } from "@/hooks/api/useTask";
import WithAuth from "@/middleware/withAuth";
import { useRouter } from "next/router";
import React, { useEffect } from "react";

export const TaskDetail = () => {
  const { fetchTaskById, task, loading, error } = useTask();
  const router = useRouter();
  const { id } = router.query;

  useEffect(() => {
    console.log(id);
    if (id && typeof id === "string") {
      fetchTaskById(id);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  if (loading) {
    return <div>Loading</div>;
  }

  if (error) {
    return <div>Error</div>;
  }

  if (!task) {
    return null;
  }

  return (
    <div>
      <div>{task.title}</div>
      <div>{task.content}</div>
      <div>{task.created_at}</div>
      <div>{task.created_by}</div>
    </div>
  );
};

export default WithAuth(TaskDetail);
