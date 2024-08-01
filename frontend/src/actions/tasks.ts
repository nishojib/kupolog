'use server';

import { kupologApi } from '@/app/api';
import { auth } from '@/auth';

export async function getTasks(kind?: string) {
  const session = await auth();

  if (!session?.user) {
    throw new Error('You must be signed in to perform this action');
  }

  const { data } = await kupologApi.tasks.sharedList({ kind: kind ?? '' });
  return data;
}

export async function toggleComplete(taskID: string, kind: string) {
  const session = await auth();

  if (!session?.user) {
    throw new Error('You must be signed in to perform this action');
  }

  return await kupologApi.tasks.sharedUpdate(taskID, {
    hasCompleted: true,
    hasHidden: false,
    kind,
  });
}
