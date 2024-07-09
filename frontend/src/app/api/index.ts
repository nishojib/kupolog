import { Api } from '@/app/api/kupolog';

export const kupologApi = new Api({ baseURL: process.env.NEXT_PUBLIC_API_URL });
