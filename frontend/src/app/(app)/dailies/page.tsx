import { revalidatePath } from 'next/cache';

import {
  getDailies,
  getWeeklies,
  toggleSubtask,
  toggleTask,
} from '@/actions/dailies';
import { TaskCard } from '@/components/task';
import { DailyTimer, WeeklyTimer } from '@/components/timer';

export default async function Page() {
  const [weeklyTasks, dailyTasks] = await Promise.all([
    getWeeklies(),
    getDailies(),
  ]);

  return (
    <div className="container">
      <div className="my-20 grid grid-cols-1 gap-4 md:grid-cols-2">
        <div className="space-y-6">
          <div className="flex items-end justify-between border-b pb-2 ">
            <h3 className="text-2xl">Weeklies</h3>
            <p className="text-muted-foreground text-sm">
              <WeeklyTimer />
            </p>
          </div>
          <ul className="space-y-4">
            {weeklyTasks?.map((task) => (
              <TaskCard
                key={task.taskID}
                task={task}
                updateTask={async (taskID) => {
                  'use server';
                  await toggleTask(taskID);
                  revalidatePath('/');
                }}
                updateSubtask={async (subtaskID) => {
                  'use server';
                  await toggleSubtask(subtaskID);
                  revalidatePath('/');
                }}
              />
            ))}
          </ul>
        </div>
        <div className="space-y-6">
          <div className="flex items-end justify-between border-b pb-2">
            <h3 className="text-2xl">Dailies</h3>
            <p className="text-muted-foreground text-sm">
              <DailyTimer />
            </p>
          </div>
          <ul className="space-y-4">
            {dailyTasks?.map((task) => (
              <TaskCard
                key={task.taskID}
                task={task}
                updateTask={async (taskID) => {
                  'use server';
                  await toggleTask(taskID);
                  revalidatePath('/');
                }}
                updateSubtask={async (subtaskID) => {
                  'use server';
                  await toggleSubtask(subtaskID);
                  revalidatePath('/');
                }}
              />
            ))}
          </ul>
        </div>
      </div>
    </div>
  );
}
