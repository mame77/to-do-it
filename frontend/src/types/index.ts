export type PlayStatus = 'unstarted' | 'playing' | 'completed';

export type GameGenre = 'RPG' | 'アクション' | 'アドベンチャー' | 'シミュレーション' | 'パズル' | 'スポーツ' | 'その他';

export interface Game {
  id: string;
  title: string;
  genre?: GameGenre;
  status: PlayStatus;
  addedAt: string;
}

export interface Schedule {
  id: string;
  gameId: string;
  date: string;
  startTime: string;
  endTime: string;
  completed: boolean;
  skipped: boolean;
}

export interface FixedEvent {
  id: string;
  title: string;
  dayOfWeek: number[]; // 0-6 (Sunday-Saturday)
  startTime: string;
  endTime: string;
  isRecurring: boolean;
  specificDate?: string;
}

export interface NotificationRecord {
  id: string;
  scheduleId: string;
  gameTitle: string;
  scheduledTime: string;
  read: boolean;
  createdAt: string;
}

export interface UserPoints {
  total: number;
  streak: number;
  lastPlayedDate: string | null;
}

export interface BonusPenaltyRecord {
  id: string;
  type: 'bonus' | 'penalty';
  points: number;
  reason: string;
  gameTitle?: string;
  createdAt: string;
}

export interface NotificationSettings {
  enabled: boolean;
  minutesBefore: number; // 何分前に通知するか
  soundEnabled: boolean;
}
