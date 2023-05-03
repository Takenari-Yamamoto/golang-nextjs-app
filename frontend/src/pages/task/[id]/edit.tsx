import WithAuth from "@/middleware/withAuth";
import { useRouter } from "next/router";

export const TaskEdit = () => {
  const router = useRouter();
  const { id } = router.query;
  return <div>TaskEdit</div>;
};

export default WithAuth(TaskEdit);
