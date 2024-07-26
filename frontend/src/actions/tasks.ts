import { kupologApi } from '@/app/api';

export async function getTasks(kind?: string) {
  const { data } = await kupologApi.tasks.sharedList({ kind: kind ?? '' });
  return data;
}

export async function toggleComplete(taskID: string, kind: string) {
  return await kupologApi.tasks.sharedUpdate(taskID, {
    hasCompleted: true,
    hasHidden: false,
    kind,
  });
}
