import React from "react";
import { Task } from "@/api/taskRepository";
import style from "./task-card.module.scss";
import { handleLink } from "@/libs/router";

type Props = Task & {
  className?: string;
};

export const TaskCard: React.FC<Props> = ({
  id,
  title,
  content,
  created_at,
  created_by,
  className = "",
}) => {
  //　本当は親で持たせたい
  const moveToDetail = () => {
    handleLink(`/task/${id}`);
  };

  return (
    <div className={`${style.root} ${className}`} onClick={moveToDetail}>
      <p>ID: {id}</p>
      <p>タイトル：{title}</p>
      <p>内容：{content}</p>
      <p>作成日：{created_at}</p>
      <p>作成者：{created_by}</p>
    </div>
  );
};
