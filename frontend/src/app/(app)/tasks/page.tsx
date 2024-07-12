import { getTasks } from '@/actions/tasks';
import { TaskCard } from '@/components/task';
import { DailyTimer, WeeklyTimer } from '@/components/timer';

export default async function Page() {
  const tasks = await getTasks();

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
            {tasks?.weeklies?.map((task) => (
              <TaskCard key={task.taskID} task={task} />
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
              <TaskCard key={task.taskID} task={task} />
            ))}
          </ul>
        </div>
      </div>
    </div>
  );
}
