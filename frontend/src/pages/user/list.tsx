import WithAuth from "@/middleware/withAuth";
import React from "react";

const UserList = () => {
  return <div>UserList</div>;
};

export default WithAuth(UserList);
