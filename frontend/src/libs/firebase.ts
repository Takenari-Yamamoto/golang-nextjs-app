import { getAuth, getIdToken } from "firebase/auth";

export const getToken = async (): Promise<string | null> => {
  const auth = getAuth();
  if (auth.currentUser) {
    const token = await getIdToken(auth.currentUser);
    return token;
  }
  return null;
};
