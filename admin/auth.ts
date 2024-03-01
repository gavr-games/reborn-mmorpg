import NextAuth from 'next-auth';
import Credentials from 'next-auth/providers/credentials';
import bcrypt from 'bcryptjs';
import prisma from './lib/prisma';
import { z } from 'zod';
import type { User } from '@prisma/client';
import { authConfig } from './auth.config';

async function getUser(email: string): Promise<User | null> {
  try {
    const user = await prisma.user.findUnique({
      where: {
        email: email,
      },
    })
    return user
  } catch (error) {
    console.error('Failed to fetch user:', error);
    throw new Error('Failed to fetch user.');
  }
}

export const { auth, signIn, signOut } = NextAuth({
  ...authConfig,
  providers: [
    Credentials({
      credentials: { 
        email: { label: "Username", type: "text", placeholder: "johndoe@domain.com" },
        password: { label: "Password", type: "password", placeholder: "verysecurepassword" },
      },
      async authorize(credentials) {
        const parsedCredentials = z
          .object({ email: z.string().email(), password: z.string().min(6) })
          .safeParse(credentials);

        if (parsedCredentials.success) {
          const { email, password } = parsedCredentials.data;

          const user = await getUser(email);
          if (!user) return null;

          const passwordsMatch = await bcrypt.compare(password, user.password);
          if (passwordsMatch) return user as any;
        }

        console.log('Invalid credentials');
        return null;
      },
    }),
  ],
});