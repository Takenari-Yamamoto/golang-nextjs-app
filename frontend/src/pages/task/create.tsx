import React, { useState } from "react";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import WithAuth from "@/middleware/withAuth";
import style from "@/styles/pages/task-create.module.scss";

const CreateTask = () => {
  const [value, setValue] = useState("");
  const handleChange = (event: any) => {
    setValue(event.target.value);
  };

  return (
    <div className={style.root}>
      <h2>タスクを新規作成</h2>
      <TextField
        label="テキスト入力"
        variant="outlined"
        value={value}
        onChange={handleChange}
      />
      <TextField
        label="テキスト入力"
        variant="outlined"
        value={value}
        onChange={handleChange}
      />
      {/* ボタンを配置 */}
      <Button variant="contained">作成</Button>
    </div>
  );
};

export default WithAuth(CreateTask);
