'use client';

import { useOptimistic, useTransition } from 'react';

import { ServerTaskResponse } from '@/app/api/kupolog';
import { Spinner } from '@/components/spinner';
import { Checkbox } from '@/components/ui/checkbox';
import { cn } from '@/lib/utils';

export function TaskCard({
  task,
  updateTask,
}: {
  task: ServerTaskResponse;
  updateTask: (taskID: string) => void;
}) {
  const [optimisticTask, setOptimisticTask] = useOptimistic<
    ServerTaskResponse,
    boolean
  >(task, (state, completed) => ({ ...state, completed }));

  const [isPending, startTransition] = useTransition();

  return (
    <div
      className={cn(
        'bg-card text-card-foreground cursor-pointer rounded-xl border shadow',
        { 'bg-muted border-muted shadow-none': optimisticTask.completed },
      )}
    >
      <div className="flex items-center justify-between space-x-4 px-4">
        <div className="flex flex-row items-center gap-4 p-6">
          <Checkbox
            id={`optimisticTask-${optimisticTask.taskID?.toString()}`}
            checked={optimisticTask.completed}
            onCheckedChange={() => {
              startTransition(() => {
                if (!optimisticTask.taskID) {
                  return;
                }

                setOptimisticTask(!optimisticTask.completed);
                updateTask(optimisticTask.taskID);
              });
            }}
          />
          <label
            htmlFor={`optimisticTask-${optimisticTask.taskID?.toString()}`}
            className={cn(
              'cursor-pointer select-none leading-none tracking-tight',
              {
                'text-muted-foreground line-through': optimisticTask.completed,
              },
            )}
          >
            {task.title}
          </label>
        </div>
        <Spinner isLoading={isPending} className="size-4" />
      </div>
    </div>
  );
}
