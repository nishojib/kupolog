import { Api } from '@/app/api/kupolog';
import { auth } from '@/auth';

export const kupologApi = new Api({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
  headers: { 'Content-Type': 'application/json' },
});

kupologApi.instance.interceptors.request.use(
  async (config) => {
    if (config.url?.includes('/tasks')) {
      const session = await auth();

      if (session?.accessToken) {
        config.headers.Authorization = `Bearer ${session.accessToken}`;
      }
    }

    return config;
  },
  (error) => {
    return Promise.reject(error);
  },
);
