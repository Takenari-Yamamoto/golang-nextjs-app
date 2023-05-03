import React, { FC } from "react";
import { useAuth } from "@/contexts/AuthContext";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import router from "next/router";
import login from "@/pages/login";
import { authRepository } from "@/api/authRepository";

export const Header: FC = () => {
  const { user } = useAuth();

  const handleLink = (path: string) => {
    router.push(path);
  };

  const authButton = (() => {
    if (user) {
      return {
        text: "ログアウト",
        func: () => authRepository.logout(),
      };
    }
    return {
      text: "ログイン",
      func: () => handleLink("/login"),
    };
  })();

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <Typography
            variant="h6"
            component="div"
            onClick={() => handleLink("/")}
          >
            Todo App
          </Typography>
          <Box sx={{ display: "flex", marginLeft: "auto" }}>
            <Button
              sx={{ mx: 1, color: "white" }}
              onClick={() => handleLink("/task/create")}
            >
              新規作成
            </Button>
            <Button
              sx={{ mx: 1, color: "white" }}
              onClick={() => handleLink("/user/list")}
            >
              ユーザー
            </Button>
          </Box>
          <Box sx={{ flexGrow: 1 }} />

          <Button color="inherit" onClick={authButton.func}>
            {authButton.text}
          </Button>
        </Toolbar>
      </AppBar>
    </Box>
  );
};
