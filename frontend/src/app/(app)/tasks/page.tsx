import { revalidatePath } from 'next/cache';

import { getTasks, toggleComplete } from '@/actions/tasks';
import { auth } from '@/auth';

import { Tasks } from './tasks';

export default async function Page() {
  const tasks = await getTasks();
  const session = await auth();

  return (
    <div className="container">
      <Tasks
        tasks={tasks}
        completeWeeklyTask={async (taskID) => {
          'use server';

          if (!session?.user?.id) {
            return;
          }

          await toggleComplete(taskID, 'weekly');
          revalidatePath('/tasks');
        }}
        completeDailyTask={async (taskID) => {
          'use server';

          if (!session?.user?.id) {
            return;
          }

          await toggleComplete(taskID, 'daily');
          revalidatePath('/tasks');
        }}
      />
    </div>
  );
}
