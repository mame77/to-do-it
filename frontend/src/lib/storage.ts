import { Game, Schedule, FixedEvent, NotificationRecord, UserPoints, BonusPenaltyRecord, NotificationSettings } from '../types';

const STORAGE_KEYS = {
  GAMES: 'gameScheduler_games',
  SCHEDULES: 'gameScheduler_schedules',
  FIXED_EVENTS: 'gameScheduler_fixedEvents',
  NOTIFICATIONS: 'gameScheduler_notifications',
  POINTS: 'gameScheduler_points',
  BONUS_PENALTY: 'gameScheduler_bonusPenalty',
  NOTIFICATION_SETTINGS: 'gameScheduler_notificationSettings',
};

export const storage = {
  // Games
  getGames: (): Game[] => {
    const data = localStorage.getItem(STORAGE_KEYS.GAMES);
    return data ? JSON.parse(data) : [];
  },
  saveGames: (games: Game[]) => {
    localStorage.setItem(STORAGE_KEYS.GAMES, JSON.stringify(games));
  },

  // Schedules
  getSchedules: (): Schedule[] => {
    const data = localStorage.getItem(STORAGE_KEYS.SCHEDULES);
    return data ? JSON.parse(data) : [];
  },
  saveSchedules: (schedules: Schedule[]) => {
    localStorage.setItem(STORAGE_KEYS.SCHEDULES, JSON.stringify(schedules));
  },

  // Fixed Events
  getFixedEvents: (): FixedEvent[] => {
    const data = localStorage.getItem(STORAGE_KEYS.FIXED_EVENTS);
    return data ? JSON.parse(data) : [];
  },
  saveFixedEvents: (events: FixedEvent[]) => {
    localStorage.setItem(STORAGE_KEYS.FIXED_EVENTS, JSON.stringify(events));
  },

  // Notifications
  getNotifications: (): NotificationRecord[] => {
    const data = localStorage.getItem(STORAGE_KEYS.NOTIFICATIONS);
    return data ? JSON.parse(data) : [];
  },
  saveNotifications: (notifications: NotificationRecord[]) => {
    localStorage.setItem(STORAGE_KEYS.NOTIFICATIONS, JSON.stringify(notifications));
  },

  // Points
  getPoints: (): UserPoints => {
    const data = localStorage.getItem(STORAGE_KEYS.POINTS);
    return data ? JSON.parse(data) : { total: 0, streak: 0, lastPlayedDate: null };
  },
  savePoints: (points: UserPoints) => {
    localStorage.setItem(STORAGE_KEYS.POINTS, JSON.stringify(points));
  },

  // Bonus/Penalty Records
  getBonusPenaltyRecords: (): BonusPenaltyRecord[] => {
    const data = localStorage.getItem(STORAGE_KEYS.BONUS_PENALTY);
    return data ? JSON.parse(data) : [];
  },
  saveBonusPenaltyRecords: (records: BonusPenaltyRecord[]) => {
    localStorage.setItem(STORAGE_KEYS.BONUS_PENALTY, JSON.stringify(records));
  },

  // Notification Settings
  getNotificationSettings: (): NotificationSettings => {
    const data = localStorage.getItem(STORAGE_KEYS.NOTIFICATION_SETTINGS);
    return data ? JSON.parse(data) : { enabled: true, minutesBefore: 15, soundEnabled: true };
  },
  saveNotificationSettings: (settings: NotificationSettings) => {
    localStorage.setItem(STORAGE_KEYS.NOTIFICATION_SETTINGS, JSON.stringify(settings));
  },
};
