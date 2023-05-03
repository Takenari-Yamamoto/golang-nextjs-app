import { auth } from "@/config/firebase";
import {
  createUserWithEmailAndPassword,
  signInWithEmailAndPassword,
} from "firebase/auth";

interface AuthRepository {
  register: (email: string, password: string) => Promise<any>;
  login: (email: string, password: string) => Promise<any>;
}

export const authRepository: AuthRepository = {
  register: async (email: string, password: string): Promise<any> =>
    await createUserWithEmailAndPassword(auth, email, password),
  login: async (email: string, password: string): Promise<any> =>
    await signInWithEmailAndPassword(auth, email, password),
};
