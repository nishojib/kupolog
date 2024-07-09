'use client';

import { CollapsibleTriggerProps } from '@radix-ui/react-collapsible';
import { CaretSortIcon } from '@radix-ui/react-icons';
import { useState, useTransition } from 'react';

import { ModelsTask } from '@/app/api/kupolog';
import { Button } from '@/components/ui/button';
import { Checkbox } from '@/components/ui/checkbox';
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from '@/components/ui/collapsible';
import { cn } from '@/lib/utils';

type TaskProps = {
  task: ModelsTask;
  updateTask: (taskID: string) => Promise<void>;
  updateSubtask: (subtaskID: string) => Promise<void>;
};

export function TaskCard(props: TaskProps) {
  const { task, updateTask, updateSubtask } = props;

  const [isOpen, setIsOpen] = useState(false);
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [_, startTransition] = useTransition();

  return (
    <Collapsible
      open={isOpen}
      onOpenChange={setIsOpen}
      className={cn(
        'bg-card text-card-foreground cursor-pointer rounded-xl border shadow',
        {
          'bg-background shadow-none text-muted-foreground': task.completed,
        },
      )}
    >
      <div className="flex items-center justify-between space-x-4 px-4">
        <div className="flex flex-row items-center gap-4 p-6">
          <Checkbox
            id={`task-${task.taskID?.toString()}`}
            checked={task.completed}
            onCheckedChange={() => {
              startTransition(() => {
                if (!task.taskID) {
                  return;
                }

                updateTask(task.taskID);
              });
            }}
          />
          <label
            htmlFor={`task-${task.taskID?.toString()}`}
            className={cn(
              'cursor-pointer select-none leading-none tracking-tight',
              {
                'line-through': task.completed,
              },
            )}
          >
            {task.title}
          </label>
        </div>

        <DailyContentTrigger
          show={
            (task.contentType === 'subtask' &&
              task.subtasks?.length !== undefined &&
              task.subtasks?.length > 0) ||
            task.contentType === 'image' ||
            task.contentType === 'embed'
          }
        >
          <Button variant="ghost" size="sm">
            <CaretSortIcon className="size-4" />
            <span className="sr-only">Toggle</span>
          </Button>
        </DailyContentTrigger>
      </div>
      {task.contentType === 'subtask' && (
        <CollapsibleContent className="space-y-2 border-t py-4 pl-16">
          {task.subtasks?.map((subtask) => (
            <div
              key={subtask.subtaskID}
              className="flex flex-row items-center gap-4"
            >
              <Checkbox
                className="rounded-full"
                id={`subtask-${subtask.subtaskID?.toString()}`}
                checked={subtask.completed}
                onCheckedChange={() => {
                  startTransition(() => {
                    if (!subtask.subtaskID) {
                      return;
                    }

                    updateSubtask(subtask.subtaskID);
                  });
                }}
              />
              <label
                htmlFor={`subtask-${subtask.subtaskID?.toString()}`}
                className={cn(
                  'cursor-pointer select-none leading-none tracking-tight',
                  {
                    'line-through': subtask.completed,
                  },
                )}
              >
                {subtask.title}
              </label>
            </div>
          ))}
        </CollapsibleContent>
      )}
      {task.contentType === 'image' && (
        <CollapsibleContent className="space-y-2 border-t p-4">
          <div>Get the image here</div>
          {/* <Image src={} alt={task.title ?? ""} width={800} height={300} /> */}
        </CollapsibleContent>
      )}
      {task.contentType === 'embed' && (
        <CollapsibleContent className="space-y-2 border-t py-4 pl-16">
          <div>Get the embed here</div>
        </CollapsibleContent>
      )}
    </Collapsible>
  );
}

function DailyContentTrigger(
  props: CollapsibleTriggerProps & { show: boolean },
) {
  const { show, ...rest } = props;

  if (!show) {
    return <></>;
  }

  return (
    <CollapsibleTrigger asChild {...rest}>
      <Button variant="ghost" size="sm">
        <CaretSortIcon className="size-4" />
        <span className="sr-only">Toggle</span>
      </Button>
    </CollapsibleTrigger>
  );
}
