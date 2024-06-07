import { revalidatePath } from 'next/cache';

export type Task = {
  id: number;
  name: string;
  completed: boolean;
  type: 'weekly' | 'daily';
  subtasks?: Subtask[];
};

type Subtask = {
  id: number;
  name: string;
  completed: boolean;
};

const tasks: Task[] = [
  {
    id: 1,
    name: 'Weekly Hunt Marks',
    completed: false,
    type: 'weekly',
    subtasks: [
      {
        id: 1,
        name: 'ARR Hunt Marks',
        completed: false,
      },
      {
        id: 2,
        name: 'Heavenward Hunt Marks',
        completed: false,
      },
      {
        id: 3,
        name: 'Stormblood Hunt Marks',
        completed: false,
      },
      {
        id: 4,
        name: 'Shadowbringers Hunt Mark',
        completed: false,
      },
      {
        id: 5,
        name: 'Endwalker Hunt Mark',
        completed: false,
      },
    ],
  },
  {
    id: 2,
    name: 'Challenge Log',
    completed: true,
    type: 'weekly',
  },
  {
    id: 3,
    name: 'Custom Delivery',
    completed: false,
    type: 'weekly',
  },
  {
    id: 4,
    name: 'Wondrous Tails',
    completed: false,
    type: 'weekly',
  },
  {
    id: 5,
    name: 'Retainer ventures',
    completed: false,
    type: 'daily',
  },
  {
    id: 6,
    name: 'Leve Allowance',
    completed: false,
    type: 'daily',
  },
  {
    id: 7,
    name: 'Fashion Report',
    completed: false,
    type: 'weekly',
  },
];

export const getWeeklyTasks = () =>
  tasks.filter((task) => task.type === 'weekly');

export const getDailyTasks = () =>
  tasks.filter((task) => task.type === 'daily');

export const toggleTask = (id: number) => {
  const taskIdx = tasks.findIndex((task) => task.id === id);
  if (tasks[taskIdx]) {
    tasks[taskIdx].completed = !tasks[taskIdx].completed;
  }

  revalidatePath('/dailies');
};

export const toggleSubtask = (taskId: number, subtaskId: number) => {
  const taskIdx = tasks.findIndex((task) => task.id === taskId);

  if (!tasks[taskIdx]) {
    return;
  }

  const subtasks = tasks[taskIdx].subtasks;

  if (!subtasks) {
    return;
  }

  const subtaskIdx = subtasks.findIndex((subtask) => subtask.id === subtaskId);

  if (!subtasks[subtaskIdx]) {
    return;
  }

  subtasks[subtaskIdx].completed = !subtasks[subtaskIdx].completed;

  revalidatePath('/dailies');
};
