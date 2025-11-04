import { Game, Schedule, FixedEvent } from '../types';

// 時間の文字列（HH:MM）を分に変換
const timeToMinutes = (time: string): number => {
  const [hours, minutes] = time.split(':').map(Number);
  return hours * 60 + minutes;
};

// 分を時間の文字列（HH:MM）に変換
const minutesToTime = (minutes: number): string => {
  const hours = Math.floor(minutes / 60);
  const mins = minutes % 60;
  return `${String(hours).padStart(2, '0')}:${String(mins).padStart(2, '0')}`;
};

// 日付と時刻から完全な日時文字列を作成
const createDateTime = (date: Date, time: string): string => {
  const dateStr = date.toISOString().split('T')[0];
  return `${dateStr}T${time}:00`;
};

// 特定の日に固定イベントがある時間帯を取得
const getBlockedTimes = (date: Date, fixedEvents: FixedEvent[]): { start: number; end: number }[] => {
  const dayOfWeek = date.getDay();
  const dateStr = date.toISOString().split('T')[0];
  const blocked: { start: number; end: number }[] = [];

  fixedEvents.forEach(event => {
    const isRelevant = event.isRecurring 
      ? event.dayOfWeek.includes(dayOfWeek)
      : event.specificDate === dateStr;

    if (isRelevant) {
      blocked.push({
        start: timeToMinutes(event.startTime),
        end: timeToMinutes(event.endTime),
      });
    }
  });

  return blocked.sort((a, b) => a.start - b.start);
};

// 利用可能な時間帯を見つける
const findAvailableSlots = (
  date: Date,
  fixedEvents: FixedEvent[],
  minDuration: number = 60 // 最小プレイ時間（分）
): { start: number; end: number }[] => {
  const blockedTimes = getBlockedTimes(date, fixedEvents);
  const slots: { start: number; end: number }[] = [];

  // プレイ可能な時間帯を設定（例：10:00-23:00）
  const dayStart = 10 * 60; // 10:00
  const dayEnd = 23 * 60; // 23:00

  let currentStart = dayStart;

  blockedTimes.forEach(blocked => {
    if (currentStart < blocked.start) {
      const duration = blocked.start - currentStart;
      if (duration >= minDuration) {
        slots.push({ start: currentStart, end: blocked.start });
      }
    }
    currentStart = Math.max(currentStart, blocked.end);
  });

  // 最後のブロック後の時間
  if (currentStart < dayEnd) {
    const duration = dayEnd - currentStart;
    if (duration >= minDuration) {
      slots.push({ start: currentStart, end: dayEnd });
    }
  }

  return slots;
};

// ゲームのプレイスケジュールを自動生成
export const generateSchedules = (
  games: Game[],
  fixedEvents: FixedEvent[],
  startDate: Date = new Date(),
  daysToSchedule: number = 14
): Schedule[] => {
  const schedules: Schedule[] = [];
  
  // 未開始またはプレイ中のゲームのみを対象
  const gamesToSchedule = games.filter(
    game => game.status === 'unstarted' || game.status === 'playing'
  );

  if (gamesToSchedule.length === 0) return [];

  let gameIndex = 0;
  const sessionDuration = 90; // 1セッション90分

  for (let dayOffset = 0; dayOffset < daysToSchedule; dayOffset++) {
    const currentDate = new Date(startDate);
    currentDate.setDate(startDate.getDate() + dayOffset);

    // 土日はより多くの時間を確保、平日は1-2セッション
    const targetSessions = currentDate.getDay() === 0 || currentDate.getDay() === 6 ? 3 : 1;
    
    const availableSlots = findAvailableSlots(currentDate, fixedEvents, sessionDuration);

    let sessionsScheduled = 0;
    for (const slot of availableSlots) {
      if (sessionsScheduled >= targetSessions) break;

      const slotDuration = slot.end - slot.start;
      const possibleSessions = Math.floor(slotDuration / sessionDuration);

      for (let i = 0; i < possibleSessions && sessionsScheduled < targetSessions; i++) {
        const startMinutes = slot.start + (i * sessionDuration);
        const endMinutes = startMinutes + sessionDuration;

        const game = gamesToSchedule[gameIndex % gamesToSchedule.length];

        schedules.push({
          id: `schedule_${Date.now()}_${schedules.length}`,
          gameId: game.id,
          date: currentDate.toISOString().split('T')[0],
          startTime: minutesToTime(startMinutes),
          endTime: minutesToTime(endMinutes),
          completed: false,
          skipped: false,
        });

        gameIndex++;
        sessionsScheduled++;
      }
    }
  }

  return schedules;
};
