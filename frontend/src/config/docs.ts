import { MainNavItem } from '@/types/nav';

interface DocsConfig {
  mainNav: MainNavItem[];
}

export const docsConfig: DocsConfig = {
  mainNav: [
    {
      title: 'Profile',
      href: '/profile',
    },
    {
      title: 'Dailies',
      href: '/dailies',
    },
  ],
};
