import type { NextAuthConfig } from 'next-auth';
 
export const authConfig = {
  pages: {
    signIn: '/admin/login',
  },
  callbacks: {
    authorized({ auth, request: { nextUrl } }) {
      const isLoggedIn = !!auth?.user;
      const isOnLogin = nextUrl.pathname.startsWith('/login');
      if (isOnLogin) {
        if (isLoggedIn) {
          return Response.redirect(new URL('/admin', nextUrl));
        }
      } else {
        if (isLoggedIn) return true;
        return Response.redirect(new URL('/admin/login', nextUrl));
      }
      return true;
    },
  },
  providers: [], // Add providers with an empty array for now
} satisfies NextAuthConfig;