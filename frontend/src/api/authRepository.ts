import { auth } from "@/config/firebase";
import {
  UserCredential,
  createUserWithEmailAndPassword,
  signInWithEmailAndPassword,
  signOut,
} from "firebase/auth";

interface AuthRepository {
  register: (email: string, password: string) => Promise<UserCredential>;
  login: (email: string, password: string) => Promise<UserCredential>;
  logout: () => Promise<void>;
}

export const authRepository: AuthRepository = {
  register: async (email: string, password: string) =>
    await createUserWithEmailAndPassword(auth, email, password),
  login: async (email: string, password: string) =>
    await signInWithEmailAndPassword(auth, email, password),
  logout: async () => await signOut(auth),
};
