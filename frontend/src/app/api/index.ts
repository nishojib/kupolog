import { Api } from '@/app/api/kupolog';
import { auth } from '@/auth';

export const kupologApi = new Api({ baseURL: process.env.NEXT_PUBLIC_API_URL });

kupologApi.instance.interceptors.request.use(
  async (config) => {
    const session = await auth();

    if (session?.accessToken) {
      config.headers.Authorization = `Bearer ${session.accessToken}`;
    }

    return config;
  },
  (error) => {
    return Promise.reject(error);
  },
);
