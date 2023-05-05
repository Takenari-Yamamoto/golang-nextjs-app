import { Header } from "@/components/header/Header";
import { AuthProvider } from "@/contexts/AuthContext";
import "@/styles/globals.css";
import type { AppProps } from "next/app";

export default function App({ Component, pageProps }: AppProps) {
  return (
    <AuthProvider>
      <Header />
      <div style={{ padding: "48px" }}>
        <Component {...pageProps} />
      </div>
    </AuthProvider>
  );
}
