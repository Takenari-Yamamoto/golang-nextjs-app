import WithAuth from "@/middleware/withAuth";

const TaskList: React.FC = () => {
  return <div>TaskList</div>;
};

export default WithAuth(TaskList);
