import { kupologApi } from '@/app/api';

export type Task = {
  id: number;
  name: string;
  completed: boolean;
  type: 'weekly' | 'daily';
  content?: Subtask[] | Record<string, string> | string;
};

type Subtask = {
  id: number;
  name: string;
  completed: boolean;
};

export async function getWeeklies() {
  const { data } = await kupologApi.dailies.weeklyList();
  return data?.weeklies;
}

export async function getDailies() {
  const { data } = await kupologApi.dailies.dailyList();
  return data?.dailies;
}

export async function toggleTask(taskID: string) {
  const { data } = await kupologApi.dailies.tasksUpdate(taskID);
  return data;
}

export async function toggleSubtask(subtaskID: string) {
  const { data } = await kupologApi.dailies.subtasksUpdate(subtaskID);
  return data;
}
