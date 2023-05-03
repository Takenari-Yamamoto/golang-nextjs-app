import { useEffect } from "react";
import { useTask } from "@/hooks/api/useTask";
import { TaskCard } from "@/components/task/card/TaskCard";
import Head from "next/head";
import WithAuth from "@/middleware/withAuth";
import style from "@/styles/pages/task-list.module.scss";

const Home: React.FC = () => {
  const { tasks, fetchAllTasks, loading, error } = useTask();

  useEffect(() => {
    fetchAllTasks();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  if (loading) {
    <div>Loading</div>;
  }

  if (error) {
    return <div>Error</div>;
  }

  if (!tasks) {
    return <div>No tasks</div>;
  }

  return (
    <>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <div className={style.root}>
          {tasks.map((task, i) => {
            return <TaskCard className={style.card} key={i} {...task} />;
          })}
        </div>
      </main>
    </>
  );
};

export default WithAuth(Home);
