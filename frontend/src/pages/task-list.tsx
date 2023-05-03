import { useTask } from "@/hooks/api/useTask";
import WithAuth from "@/middleware/withAuth";
import { useEffect } from "react";

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

  return (
    <div>
      {tasks.map((task, i) => {
        return <div key={i}>oo</div>;
      })}
    </div>
  );
};

export default WithAuth(TaskList);
