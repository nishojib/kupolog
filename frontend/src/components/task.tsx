'use client';

import { CaretSortIcon } from '@radix-ui/react-icons';
import Image from 'next/image';
import { useState, useTransition } from 'react';

import { type Task } from '@/actions/dailies';
import { Button } from '@/components/ui/button';
import { Checkbox } from '@/components/ui/checkbox';
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from '@/components/ui/collapsible';
import { cn } from '@/lib/utils';

type TaskProps = {
  task: Task;
  updateTask: (id: number) => Promise<void>;
  updateSubtask: (taskId: number, subtaskId: number) => Promise<void>;
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
            id={`task-${task.id.toString()}`}
            checked={task.completed}
            onCheckedChange={() => {
              startTransition(() => {
                updateTask(task.id);
              });
            }}
          />
          <label
            htmlFor={`task-${task.id.toString()}`}
            className={cn(
              'cursor-pointer select-none leading-none tracking-tight',
              {
                'line-through': task.completed,
              },
            )}
          >
            {task.name}
          </label>
        </div>
        {task.content && (
          <CollapsibleTrigger asChild>
            <Button variant="ghost" size="sm">
              <CaretSortIcon className="size-4" />
              <span className="sr-only">Toggle</span>
            </Button>
          </CollapsibleTrigger>
        )}
      </div>
      {Array.isArray(task.content) && task.content?.length > 0 && (
        <CollapsibleContent className="space-y-2 border-t py-4 pl-16">
          {task.content?.map((subtask) => (
            <div key={subtask.id} className="flex flex-row items-center gap-4">
              <Checkbox
                className="rounded-full"
                id={`subtask-${subtask.id.toString()}`}
                checked={subtask.completed}
                onCheckedChange={() => {
                  startTransition(() => {
                    updateSubtask(task.id, subtask.id);
                  });
                }}
              />
              <label
                htmlFor={`subtask-${subtask.id.toString()}`}
                className={cn(
                  'cursor-pointer select-none leading-none tracking-tight',
                  {
                    'line-through': subtask.completed,
                  },
                )}
              >
                {subtask.name}
              </label>
            </div>
          ))}
        </CollapsibleContent>
      )}
      {typeof task.content === 'string' && (
        <CollapsibleContent className="space-y-2 border-t p-4">
          <Image src={task.content} alt={task.name} width={800} height={300} />
        </CollapsibleContent>
      )}
      {typeof task.content === 'object' && 'key' in task.content && (
        <CollapsibleContent className="space-y-2 border-t py-4 pl-16">
          <div>Some content</div>
        </CollapsibleContent>
      )}
    </Collapsible>
  );
}
