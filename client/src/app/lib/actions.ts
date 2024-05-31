'use server';

import { z } from 'zod';

import axios from './axios';

const HealthSchema = z.object({
  status: z.literal('available'),
  system_info: z.object({
    environment: z.enum(['development', 'production']),
    version: z.string(),
  }),
});

export async function getHealth() {
  const data = await axios.get('/health');

  const health = HealthSchema.safeParse(data.data);

  if (!health.success) {
    throw new Error(health.error.message);
  }

  return health.data;
}
