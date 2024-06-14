import NextAuth, { type NextAuthConfig } from 'next-auth';
import Discord from 'next-auth/providers/discord';
import Google from 'next-auth/providers/google';

const authOptions = {
  pages: {
    signIn: '/auth/login',
  },
  callbacks: {
    async signIn({ account }) {
      console.log({ account });
      return true;
    },
  },
  secret: process.env.AUTH_SECRET,
  providers: [Google({}), Discord({})],
} satisfies NextAuthConfig;

export const { handlers, signIn, signOut, auth } = NextAuth(authOptions);
