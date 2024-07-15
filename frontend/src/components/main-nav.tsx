'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';

import { Icons } from '@/components/icons';
import { siteConfig } from '@/config/site';
import { cn } from '@/lib/utils';

export function MainNav() {
  const pathname = usePathname();

  return (
    <div className="mr-4 hidden md:flex">
      <Link href="/" className="mr-6 flex items-center gap-2">
        <Icons.logo className="text-primary size-6 fill-current" />
        <span className="text-primary hidden font-bold sm:inline-block">
          {siteConfig.name}
        </span>
      </Link>
      <nav className="flex items-center gap-4 text-sm lg:gap-6">
        <Link
          href="/profile"
          className={cn(
            'hover:text-foreground/80 transition-colors',
            pathname === '/profile' ? 'text-foreground' : 'text-foreground/60',
          )}
        >
          Profile
        </Link>
        <Link
          href="/tasks"
          className={cn(
            'hover:text-foreground/80 transition-colors',
            pathname?.startsWith('/tasks')
              ? 'text-foreground'
              : 'text-foreground/60',
          )}
        >
          Tasks
        </Link>
      </nav>
    </div>
  );
}
