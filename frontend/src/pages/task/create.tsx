import React, { useState } from "react";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import WithAuth from "@/middleware/withAuth";
import style from "@/styles/pages/task-create.module.scss";
import { useTask } from "@/hooks/api/useTask";
import { handleLink } from "@/libs/router";

const CreateTask = () => {
  const { createTask, loading } = useTask();

  const [value, setValue] = useState({
    title: "",
    content: "",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    console.log(e.target.name);
    setValue({
      ...value,
      [e.target.name]: e.target.value,
    });
  };

  const handleClick = async () => {
    const res = await createTask({
      title: value.title,
      content: value.content,
    });
    if (res === "success") {
      setValue({
        title: "",
        content: "",
      });
      handleLink("/");
    }
  };

  return (
    <div className={style.root}>
      <h2>タスクを新規作成</h2>
      <TextField
        label="タイトル"
        name="title"
        variant="outlined"
        value={value.title}
        onChange={handleChange}
      />
      <TextField
        label="コンテンツ"
        name="content"
        variant="outlined"
        value={value.content}
        onChange={handleChange}
      />
      <Button variant="contained" disabled={loading} onClick={handleClick}>
        作成
      </Button>
    </div>
  );
};

export default WithAuth(CreateTask);
