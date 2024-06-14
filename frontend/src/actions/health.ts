'use server';

import { Api } from '@/app/api/api';

export async function getHealth() {
  const api = new Api({ baseURL: process.env.NEXT_PUBLIC_API_URL });
  const data = await api.health.healthList();
  return data.data;
}
