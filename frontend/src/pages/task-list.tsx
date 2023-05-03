import { useEffect } from "react";
import { TaskCard } from "@/components/task/card/TaskCard";
import { useTask } from "@/hooks/api/useTask";
import WithAuth from "@/middleware/withAuth";
import style from "@/styles/pages/task-list.module.scss";

const TaskList: React.FC = () => {
  const { tasks, fetchAllTasks, loading, error } = useTask();

  useEffect(() => {
    fetchAllTasks();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  if (loading) {
    console.log("Fetching tasks...");
    <div>Loading</div>;
  }

  if (error) {
    console.error("ERROR", error);
    return <div>Error</div>;
  }

  if (!tasks) {
    return <div>No tasks</div>;
  }

  console.log("TASKS", tasks);

  return (
    <div className={style.root}>
      {tasks.map((task, i) => {
        return <TaskCard className={style.card} key={i} {...task} />;
      })}
    </div>
  );
};

export default WithAuth(TaskList);
