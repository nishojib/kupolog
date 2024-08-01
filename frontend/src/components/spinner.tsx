import { cn } from '@/lib/utils';

export function Spinner({
  isLoading,
  className,
}: {
  isLoading: boolean;
  className?: string;
}) {
  if (!isLoading) {
    return undefined;
  }

  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
      className={cn('text-primary animate-spin', className)}
    >
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
  );
}
