import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";

import { Toaster } from "@/components/ui/toaster"
import MainNav from "@/components/main-nav"
import SignOut from "@/components/sign-out"

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Reborn MMORPG Admin Panel",
  description: "Edit game configs and objects",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <div className="w-full bg-slate-100">
          <div className="container mx-auto px-2">
            <MainNav />
            <SignOut />
          </div>
        </div>
        <div className="container mx-auto px-2 mt-2">
            {children}
        </div>
        <Toaster />
      </body>
    </html>
  );
}
