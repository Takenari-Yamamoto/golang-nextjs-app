import React, { useEffect } from "react";
import { useRouter } from "next/router";
import { useTask } from "@/hooks/api/useTask";
import { Button } from "@mui/material";
import WithAuth from "@/middleware/withAuth";

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
    <>
      <div>
        <div>タイトル：{task.title}</div>
        <div>内容：{task.content}</div>
        <div>作成日：{task.created_at}</div>
        <div>作成者：{task.created_by}</div>
      </div>
      <Button
        variant="contained"
        color="success"
        onClick={() => router.push(`/task/${id}/edit`)}
      >
        編集
      </Button>
      <Button variant="contained" color="error">
        削除
      </Button>
    </>
  );
};

export default WithAuth(TaskDetail);
