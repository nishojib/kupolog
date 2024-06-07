'use client';

import { UTCDate } from '@date-fns/utc';
import React, { useState } from 'react';
import { useInterval } from 'usehooks-ts';

import {
  calculateTimeLeft,
  formatTimeLeft,
  getNextDailyResetTimeUTC,
  getNextWeeklyResetTimeUTC,
} from '@/lib/time';

export function WeeklyTimer() {
  return <Timer getResetTimeUTC={getNextWeeklyResetTimeUTC} showDays={true} />;
}

export function DailyTimer() {
  return <Timer getResetTimeUTC={getNextDailyResetTimeUTC} showDays={false} />;
}

type TimerProps = {
  getResetTimeUTC: () => UTCDate;
  showDays: boolean;
};

function Timer(props: TimerProps) {
  const { getResetTimeUTC, showDays } = props;

  const [timeLeft, setTimeLeft] = useState(
    calculateTimeLeft(getResetTimeUTC()),
  );

  useInterval(() => {
    setTimeLeft(calculateTimeLeft(getResetTimeUTC()));
  }, 1000);

  return <span>{formatTimeLeft(timeLeft, showDays)}</span>;
}
