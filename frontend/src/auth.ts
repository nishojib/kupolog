import 'next-auth/jwt';

import NextAuth, { type NextAuthConfig, User } from 'next-auth';
import Discord from 'next-auth/providers/discord';
import Google from 'next-auth/providers/google';

import { kupologApi } from '@/app/api';

declare module 'next-auth' {
  interface Session {
    accessToken?: string;
    error?: 'RefreshAccessTokenError';
  }
}

declare module 'next-auth/jwt' {
  interface JWT {
    access_token: string;
    expires_at: number;
    refresh_token: string;
    user?: User;
    error?: 'RefreshAccessTokenError';
  }
}

const authOptions = {
  pages: { signIn: '/auth/login' },
  callbacks: {
    async jwt({ token, account }) {
      if (account) {
        // First login, save the `access_token`, `refresh_token`, and other
        // details into the JWT

        const { data } = await kupologApi.auth.loginCreate({
          provider: account?.provider,
          provider_account_id: account?.providerAccountId,
          access_token: account?.access_token,
          expires_at: account?.expires_at,
        });

        const userProfile: User = {
          id: data.user?.userID,
          name: data.user?.name,
          email: data.user?.email,
          image: data?.user?.image,
        };

        // TODO: add expires_at to the response in the API
        return {
          access_token: data.token?.access_token ?? token.access_token,
          expires_at: account.expires_at ?? token.expires_at,
          refresh_token: data.token?.refresh_token ?? token.refresh_token,
          user: userProfile,
        };
      } else if (Date.now() < token.expires_at * 1000) {
        return token;
      } else {
        if (!token.refresh_token) {
          throw new Error('Missing refresh token');
        }

        try {
          const { data } = await kupologApi.auth.refreshCreate({
            headers: { Authorization: `Bearer ${token.refresh_token}` },
          });

          return {
            ...token,
            access_token: data.access_token ?? token.access_token,
            expires_at: Date.now() + 1e8,
            refresh_token: token.refresh_token,
          };
        } catch (error) {
          console.error('Error refreshing access token', error);
          // The error property can be used client-side to handle the refresh token error
          return { ...token, error: 'RefreshAccessTokenError' as const };
        }
      }
    },
    async session({ session, token }) {
      if (token.user) {
        session.user = {
          id: token.user.id || '',
          name: token.user.name,
          email: token.user.email || '',
          image: token.user.image,
          emailVerified: null,
        };

        session.accessToken = token.access_token;
      }

      return session;
    },
  },
  secret: process.env.AUTH_SECRET,
  providers: [Google({}), Discord({})],
  trustHost: true,
} satisfies NextAuthConfig;

export const { handlers, signIn, signOut, auth } = NextAuth(authOptions);
