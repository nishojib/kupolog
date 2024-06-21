import { revalidatePath } from 'next/cache';

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

const tasks: Task[] = [
  {
    id: 1,
    name: 'Weekly repeatable quests',
    completed: false,
    type: 'weekly',
  },
  {
    id: 2,
    name: 'Weekly Hunt Marks',
    completed: false,
    type: 'weekly',
    content: [
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
    id: 3,
    name: 'Squadron priority mission',
    completed: false,
    type: 'weekly',
  },
  {
    id: 4,
    name: 'Challenge Logs',
    completed: false,
    type: 'weekly',
    content: [
      {
        id: 1,
        name: 'Battles',
        completed: false,
      },
      {
        id: 2,
        name: 'PvP',
        completed: false,
      },
      {
        id: 3,
        name: 'FATE',
        completed: false,
      },
      {
        id: 4,
        name: 'Levequests',
        completed: false,
      },
      {
        id: 5,
        name: 'Crafting & Gathering',
        completed: false,
      },
      {
        id: 6,
        name: 'Treasure Hunt',
        completed: false,
      },
      {
        id: 7,
        name: 'Tribal Quests',
        completed: false,
      },
      {
        id: 8,
        name: 'Grand Company',
        completed: false,
      },
      {
        id: 9,
        name: 'Retainer Ventures',
        completed: false,
      },
      {
        id: 10,
        name: 'Gold Saucer',
        completed: false,
      },
      {
        id: 11,
        name: 'Other (The Forbidden Land, Eureka)',
        completed: false,
      },
      {
        id: 12,
        name: 'Other (Island Sanctuary)',
        completed: false,
      },
      {
        id: 13,
        name: 'Overall Completion',
        completed: false,
      },
    ],
  },
  {
    id: 5,
    name: 'Custom Delivery',
    completed: false,
    type: 'weekly',
    content: [
      {
        id: 1,
        name: 'Zhloe Aliapoh',
        completed: false,
      },
      {
        id: 2,
        name: "M'naago",
        completed: false,
      },
      {
        id: 3,
        name: 'Kurenai',
        completed: false,
      },
      {
        id: 4,
        name: 'Adkiragh',
        completed: false,
      },
      {
        id: 5,
        name: 'Kai-Shirr',
        completed: false,
      },
      {
        id: 6,
        name: 'Ehil Tou',
        completed: false,
      },
      {
        id: 7,
        name: 'Charlemald',
        completed: false,
      },
      {
        id: 8,
        name: 'Ameliance',
        completed: false,
      },
      {
        id: 9,
        name: 'Anden',
        completed: false,
      },
      {
        id: 10,
        name: 'Margrat',
        completed: false,
      },
    ],
  },
  {
    id: 6,
    name: 'Wondrous Tails',
    completed: false,
    type: 'weekly',
    content: {},
  },
  {
    id: 7,
    name: 'Normal raid lockout',
    completed: false,
    type: 'weekly',
    content: [],
  },
  {
    id: 8,
    name: 'Savage raid lockout',
    completed: false,
    type: 'weekly',
    content: [],
  },
  {
    id: 9,
    name: 'Alliance raid lockout',
    completed: false,
    type: 'weekly',
    content: [],
  },
  {
    id: 10,
    name: 'Faux Hollows',
    completed: false,
    type: 'weekly',
  },
  {
    id: 11,
    name: 'Tomestone cap',
    completed: false,
    type: 'weekly',
    content: [],
  },
  {
    id: 12,
    name: 'Doman Enclave donation',
    completed: false,
    type: 'weekly',
  },
  {
    id: 13,
    name: 'Masked Carnivale targets',
    completed: false,
    type: 'weekly',
    content: [],
  },
  {
    id: 14,
    name: 'Blue Mage Log',
    completed: false,
    type: 'weekly',
    content: [],
  },
  {
    id: 15,
    name: 'Gold Saucer tournaments',
    completed: false,
    type: 'weekly',
    content: [],
  },
  {
    id: 16,
    name: 'Fashion Report',
    completed: false,
    type: 'weekly',
    content: 'https://i.imgur.com/BOigGZO.png',
  },
  {
    id: 17,
    name: 'Jumbo Cactpot',
    completed: false,
    type: 'weekly',
  },
  {
    id: 18,
    name: 'Daily repeatable quests',
    completed: false,
    type: 'daily',
    content: [],
  },
  {
    id: 19,
    name: 'Daily Hunt marks',
    completed: false,
    type: 'daily',
    content: [],
  },
  {
    id: 20,
    name: 'Squadron missions',
    completed: false,
    type: 'daily',
  },
  {
    id: 21,
    name: 'Grand Company turn-ins',
    completed: false,
    type: 'daily',
  },
  {
    id: 22,
    name: 'Free Company voyages',
    completed: false,
    type: 'daily',
  },
  {
    id: 23,
    name: 'Retainer ventures',
    completed: false,
    type: 'daily',
    content: [],
  },
  {
    id: 24,
    name: 'Housing gardening',
    completed: false,
    type: 'daily',
  },
  {
    id: 25,
    name: 'Tribal quests',
    completed: false,
    type: 'daily',
    content: [],
  },
  {
    id: 26,
    name: 'Treasure map allowance',
    completed: false,
    type: 'daily',
  },
  {
    id: 27,
    name: 'Leve allowance',
    completed: false,
    type: 'daily',
  },
  {
    id: 28,
    name: 'Duty roulettes',
    completed: false,
    type: 'daily',
    content: [],
  },
  {
    id: 29,
    name: 'Mini Cactpot',
    completed: false,
    type: 'daily',
    content: '',
  },
  {
    id: 30,
    name: 'Island Sanctuary dailies',
    completed: false,
    type: 'daily',
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

  const subtasks = tasks[taskIdx].content;

  if (!subtasks) {
    return;
  }

  if (!Array.isArray(subtasks)) {
    return;
  }

  const subtaskIdx = subtasks.findIndex((subtask) => subtask.id === subtaskId);

  if (!subtasks[subtaskIdx]) {
    return;
  }

  subtasks[subtaskIdx].completed = !subtasks[subtaskIdx].completed;

  revalidatePath('/dailies');
};
