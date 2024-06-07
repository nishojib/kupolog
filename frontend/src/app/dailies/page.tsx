import {
  getDailyTasks,
  getWeeklyTasks,
  toggleSubtask,
  toggleTask,
} from '@/actions/dailies';
import { TaskCard } from '@/components/task';

export default async function Page() {
  const weeklyTasks = getWeeklyTasks();
  const dailyTasks = getDailyTasks();

  return (
    <div className="container">
      <div className="my-20 grid grid-cols-1 gap-4 md:grid-cols-2">
        <div className="space-y-6">
          <div className="flex items-end justify-between border-b pb-2 ">
            <h3 className="text-2xl">Weeklies</h3>
          </div>
          <ul className="space-y-4">
            {weeklyTasks.map((task) => (
              <TaskCard
                key={task.id}
                task={task}
                updateTask={async (id) => {
                  'use server';
                  toggleTask(id);
                }}
                updateSubtask={async (taskId, subtaskId) => {
                  'use server';
                  toggleSubtask(taskId, subtaskId);
                }}
              />
            ))}
          </ul>
        </div>
        <div className="space-y-6">
          <div className="flex items-end justify-between border-b pb-2">
            <h3 className="text-2xl">Dailies</h3>
          </div>
          <ul className="space-y-4">
            {dailyTasks.map((task) => (
              <TaskCard
                key={task.id}
                task={task}
                updateTask={async () => {
                  'use server';
                  toggleTask(task.id);
                }}
                updateSubtask={async (taskId, subtaskId) => {
                  'use server';
                  toggleSubtask(taskId, subtaskId);
                }}
              />
            ))}
          </ul>
        </div>
      </div>
    </div>
  );
}
