import React from "react";
import { Task } from "@/api/taskRepository";
import style from "./task-card.module.scss";

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
  return (
    <div className={`${style.root} ${className}`}>
      <p>{id}</p>
      <p>{title}</p>
      <p>{content}</p>
      <p>{created_at}</p>
      <p>{created_by}</p>
    </div>
  );
};
