import { UTCDate } from '@date-fns/utc';
import { addDays, differenceInSeconds, nextTuesday, set } from 'date-fns';

export function getNextDailyResetTimeUTC() {
  const now = new UTCDate();
  let resetTimeUTC = set(now, {
    hours: 15,
    minutes: 0,
    seconds: 0,
    milliseconds: 0,
  }); // 3 PM UTC

  // If the reset time is in the past for today, set it for tomorrow
  if (now >= resetTimeUTC) {
    resetTimeUTC = addDays(resetTimeUTC, 1);
  }

  return resetTimeUTC;
}

export function getNextWeeklyResetTimeUTC() {
  const now = new UTCDate();
  let resetTimeUTC = set(nextTuesday(now), {
    hours: 8,
    minutes: 0,
    seconds: 0,
    milliseconds: 0,
  }); // 8 AM UTC

  // If the reset time is in the past for today, set it for next week
  if (now >= resetTimeUTC) {
    resetTimeUTC = addDays(resetTimeUTC, 7);
  }

  return resetTimeUTC;
}

export function calculateTimeLeft(resetTime: Date) {
  const now = new UTCDate();
  const totalSeconds = differenceInSeconds(resetTime, now);

  const days = Math.floor(totalSeconds / (24 * 3600));
  const hours = Math.floor((totalSeconds % (24 * 3600)) / 3600);
  const minutes = Math.floor((totalSeconds % 3600) / 60);
  const seconds = totalSeconds % 60;

  return { days, hours, minutes, seconds };
}

export function formatTimeLeft(
  timeLeft: {
    days: number;
    hours: number;
    minutes: number;
    seconds: number;
  },
  showDays: boolean,
) {
  const { days, hours, minutes, seconds } = timeLeft;

  const hoursText = hours > 0 ? hours.toString().padStart(2, '0') : '';
  const minutesText = minutes > 0 ? minutes.toString().padStart(2, '0') : '';
  const secondsText = seconds > 0 ? seconds.toString().padStart(2, '0') : '';

  // Format the output
  if (days > 0 && showDays) {
    return `${days}d ${hoursText}h`;
  } else if (hours > 0) {
    return `${hoursText}h ${minutesText}m`;
  } else if (minutes > 0) {
    return `${minutesText}m ${secondsText}s`;
  } else if (seconds > 0) {
    return `${secondsText}s`;
  } else {
    return '0s';
  }
}
