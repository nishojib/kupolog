'use client';

import { TaskCard } from '@/app/(app)/tasks/task';
import { DailyTimer, WeeklyTimer } from '@/app/(app)/tasks/timer';
import { ServerSharedTaskResponse } from '@/app/api/kupolog';

import { useSSEEvents } from './use-sse-events';

export function Tasks({
  tasks,
  completeWeeklyTask,
  completeDailyTask,
}: {
  tasks: ServerSharedTaskResponse;
  completeWeeklyTask: (taskID: string) => void;
  completeDailyTask: (taskID: string) => void;
}) {
  useSSEEvents(process.env.NEXT_PUBLIC_API_URL + '/events?stream=messages');

  return (
    <div className="my-20 grid grid-cols-1 gap-4 md:grid-cols-2">
      <div className="space-y-6">
        <div className="flex items-end justify-between border-b pb-2 ">
          <h3 className="text-2xl">Weeklies</h3>
          <p className="text-muted-foreground text-sm">
            <WeeklyTimer />
          </p>
        </div>
        <ul className="space-y-4">
          {tasks?.weeklies?.map((task) => (
            <TaskCard
              key={task.taskID}
              task={task}
              updateTask={completeWeeklyTask}
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
          {tasks?.dailies?.map((task) => (
            <TaskCard
              key={task.taskID}
              task={task}
              updateTask={completeDailyTask}
            />
          ))}
        </ul>
      </div>
    </div>
  );
}
