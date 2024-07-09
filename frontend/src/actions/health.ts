'use server';

import { kupologApi } from '@/app/api';

export async function getHealth() {
  const { data } = await kupologApi.health.healthList();
  return data;
}
