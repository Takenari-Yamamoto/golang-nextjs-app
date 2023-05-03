import React, { useEffect, ComponentType } from "react";
import { useRouter } from "next/router";
import { useAuth } from "../contexts/AuthContext";

// ここで、Propsというジェネリック型を定義しています
type Props<P = {}> = P & {
  Component: ComponentType<P>;
};

const WithoutAuth = <P extends Record<string, unknown>>(
  Component: ComponentType<P>
) => {
  const NotAuthComponent: React.FC<Props<P>> = (props) => {
    const { user, loading } = useAuth();
    const router = useRouter();

    useEffect(() => {
      if (!loading && user) {
        router.replace("/");
      }
      // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [user, loading]);

    return <>{!user && <Component {...(props as P)} />}</>;
  };

  return NotAuthComponent;
};

export default WithoutAuth;
