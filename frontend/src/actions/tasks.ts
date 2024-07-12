import { kupologApi } from '@/app/api';

export async function getTasks(kind?: string) {
  const { data } = await kupologApi.tasks.sharedList({ kind: kind ?? '' });
  return data;
}
