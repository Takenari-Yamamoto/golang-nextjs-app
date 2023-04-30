import React, { useEffect, ComponentType } from "react";
import { useRouter } from "next/router";
import { useAuth } from "../contexts/AuthContext";

type Props<P = {}> = P & {
  Component: ComponentType<P>;
};

const withAuth = <P extends Record<string, unknown>>(
  Component: ComponentType<P>
) => {
  const AuthComponent: React.FC<Props<P>> = (props) => {
    const { user, loading } = useAuth();
    const router = useRouter();

    useEffect(() => {
      if (!loading && !user) {
        router.replace("/login");
      }
      // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [user, loading]);

    return <>{user && <Component {...(props as P)} />}</>;
  };

  return AuthComponent;
};

export default withAuth;
