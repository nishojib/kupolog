'use client';

import { useTransition } from 'react';

import { ServerTaskResponse } from '@/app/api/kupolog';
import { Checkbox } from '@/components/ui/checkbox';
import { cn } from '@/lib/utils';

export function TaskCard({ task }: { task: ServerTaskResponse }) {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [_, startTransition] = useTransition();

  return (
    <div
      className={cn(
        'bg-card text-card-foreground cursor-pointer rounded-xl border shadow',
      )}
    >
      <div className="flex items-center justify-between space-x-4 px-4">
        <div className="flex flex-row items-center gap-4 p-6">
          <Checkbox
            id={`task-${task.taskID?.toString()}`}
            checked={false}
            onCheckedChange={() => {}}
          />
          <label
            htmlFor={`task-${task.taskID?.toString()}`}
            className={cn(
              'cursor-pointer select-none leading-none tracking-tight',
            )}
          >
            {task.title}
          </label>
        </div>
      </div>
    </div>
  );
}
