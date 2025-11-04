import { useState, useEffect } from 'react';
import { Game, Schedule, FixedEvent, NotificationRecord, UserPoints, PlayStatus, GameGenre, BonusPenaltyRecord, NotificationSettings as NotificationSettingsType } from './types';
import { storage } from './lib/storage';
import { generateSchedules } from './lib/scheduleGenerator';
import { GameListAdvanced } from './components/GameListAdvanced';
import { GameFormAdvanced } from './components/GameFormAdvanced';
import { ScheduleCalendar } from './components/ScheduleCalendar';
import { CalendarView } from './components/CalendarView';
import { FixedEventManager } from './components/FixedEventManager';
import { PointsDisplay } from './components/PointsDisplay';
import { NotificationPanel } from './components/NotificationPanel';
import { BonusPenaltyHistory } from './components/BonusPenaltyHistory';
import { MotivationGauge } from './components/MotivationGauge';
import { NotificationSettings } from './components/NotificationSettings';
import { BonusAnimation } from './components/BonusAnimation';
import { Tabs, TabsContent, TabsList, TabsTrigger } from './components/ui/tabs';
import { Button } from './components/ui/button';
import { Gamepad2, Calendar, Clock, Bell, Settings, Award, Gauge } from 'lucide-react';
import { toast } from 'sonner@2.0.3';
import { Toaster } from './components/ui/sonner';

function App() {
  const [games, setGames] = useState<Game[]>([]);
  const [schedules, setSchedules] = useState<Schedule[]>([]);
  const [fixedEvents, setFixedEvents] = useState<FixedEvent[]>([]);
  const [notifications, setNotifications] = useState<NotificationRecord[]>([]);
  const [points, setPoints] = useState<UserPoints>({ total: 0, streak: 0, lastPlayedDate: null });
  const [bonusPenaltyRecords, setBonusPenaltyRecords] = useState<BonusPenaltyRecord[]>([]);
  const [notificationSettings, setNotificationSettings] = useState<NotificationSettingsType>({
    enabled: true,
    minutesBefore: 15,
    soundEnabled: true,
  });
  const [showBonusAnimation, setShowBonusAnimation] = useState(false);
  const [bonusAnimationData, setBonusAnimationData] = useState<{
    type: 'bonus' | 'penalty';
    points: number;
    message: string;
  } | null>(null);

  // データの読み込み
  useEffect(() => {
    setGames(storage.getGames());
    setSchedules(storage.getSchedules());
    setFixedEvents(storage.getFixedEvents());
    setNotifications(storage.getNotifications());
    setPoints(storage.getPoints());
    setBonusPenaltyRecords(storage.getBonusPenaltyRecords());
    setNotificationSettings(storage.getNotificationSettings());

    // 通知権限のリクエスト
    if ('Notification' in window && Notification.permission === 'default') {
      Notification.requestPermission();
    }
  }, []);

  // データの保存
  useEffect(() => {
    storage.saveGames(games);
  }, [games]);

  useEffect(() => {
    storage.saveSchedules(schedules);
  }, [schedules]);

  useEffect(() => {
    storage.saveFixedEvents(fixedEvents);
  }, [fixedEvents]);

  useEffect(() => {
    storage.saveNotifications(notifications);
  }, [notifications]);

  useEffect(() => {
    storage.savePoints(points);
  }, [points]);

  useEffect(() => {
    storage.saveBonusPenaltyRecords(bonusPenaltyRecords);
  }, [bonusPenaltyRecords]);

  useEffect(() => {
    storage.saveNotificationSettings(notificationSettings);
  }, [notificationSettings]);

  // 通知のチェック（1分ごと）
  useEffect(() => {
    if (!notificationSettings.enabled) return;

    const checkNotifications = () => {
      const now = new Date();
      const upcomingSchedules = schedules.filter(schedule => {
        if (schedule.completed || schedule.skipped) return false;
        
        const scheduleTime = new Date(`${schedule.date}T${schedule.startTime}`);
        const timeDiff = scheduleTime.getTime() - now.getTime();
        const minutesDiff = timeDiff / (1000 * 60);
        
        // 設定した時間前に通知
        return minutesDiff > 0 && minutesDiff <= notificationSettings.minutesBefore;
      });

      upcomingSchedules.forEach(schedule => {
        // 既に通知済みかチェック
        const alreadyNotified = notifications.some(
          n => n.scheduleId === schedule.id
        );

        if (!alreadyNotified) {
          const game = games.find(g => g.id === schedule.gameId);
          if (game) {
            const notification: NotificationRecord = {
              id: `notif_${Date.now()}_${Math.random()}`,
              scheduleId: schedule.id,
              gameTitle: game.title,
              scheduledTime: `${schedule.date}T${schedule.startTime}`,
              read: false,
              createdAt: new Date().toISOString(),
            };

            setNotifications(prev => [notification, ...prev]);

            // ブラウザ通知
            if ('Notification' in window && Notification.permission === 'granted') {
              new Notification('ゲームプレイの時間です！', {
                body: `${game.title} - ${schedule.startTime}から`,
                icon: '/favicon.ico',
              });
            }

            toast.info(`${game.title}のプレイ時間が近づいています`, {
              description: `${schedule.startTime}から開始予定です`,
            });
          }
        }
      });
    };

    checkNotifications();
    const interval = setInterval(checkNotifications, 60000); // 1分ごと

    return () => clearInterval(interval);
  }, [schedules, games, notifications, notificationSettings]);

  const handleAddGame = (title: string, genre?: GameGenre) => {
    const newGame: Game = {
      id: `game_${Date.now()}`,
      title,
      genre,
      status: 'unstarted',
      addedAt: new Date().toISOString(),
    };
    setGames([...games, newGame]);
    toast.success(`${title}を追加しました`);
  };

  const handleDeleteGame = (id: string) => {
    const game = games.find(g => g.id === id);
    setGames(games.filter(g => g.id !== id));
    setSchedules(schedules.filter(s => s.gameId !== id));
    toast.success(`${game?.title}を削除しました`);
  };

  const handleUpdateStatus = (id: string, status: PlayStatus) => {
    setGames(games.map(g => (g.id === id ? { ...g, status } : g)));
    toast.success('ステータスを更新しました');
  };

  const handleUpdateGenre = (id: string, genre: GameGenre) => {
    setGames(games.map(g => (g.id === id ? { ...g, genre } : g)));
    toast.success('ジャンルを更新しました');
  };

  const handleAddFixedEvent = (event: Omit<FixedEvent, 'id'>) => {
    const newEvent: FixedEvent = {
      ...event,
      id: `event_${Date.now()}`,
    };
    setFixedEvents([...fixedEvents, newEvent]);
    toast.success('固定予定を追加しました');
  };

  const handleDeleteFixedEvent = (id: string) => {
    setFixedEvents(fixedEvents.filter(e => e.id !== id));
    toast.success('固定予定を削除しました');
  };

  const handleGenerateSchedules = () => {
    const newSchedules = generateSchedules(games, fixedEvents);
    setSchedules(newSchedules);
    toast.success(`${newSchedules.length}件のスケジュールを生成しました`);
  };

  const addBonusPenaltyRecord = (
    type: 'bonus' | 'penalty',
    points: number,
    reason: string,
    gameTitle?: string
  ) => {
    const record: BonusPenaltyRecord = {
      id: `bp_${Date.now()}`,
      type,
      points,
      reason,
      gameTitle,
      createdAt: new Date().toISOString(),
    };
    setBonusPenaltyRecords(prev => [record, ...prev]);

    // アニメーション表示
    setBonusAnimationData({ type, points, message: reason });
    setShowBonusAnimation(true);
  };

  const handleCompleteSchedule = (scheduleId: string) => {
    const schedule = schedules.find(s => s.id === scheduleId);
    if (!schedule) return;

    const game = games.find(g => g.id === schedule.gameId);
    
    setSchedules(
      schedules.map(s =>
        s.id === scheduleId ? { ...s, completed: true } : s
      )
    );

    // ポイント加算
    const today = new Date().toISOString().split('T')[0];
    const lastPlayed = points.lastPlayedDate;
    const isConsecutive = lastPlayed === new Date(Date.now() - 86400000).toISOString().split('T')[0];

    setPoints({
      total: points.total + 10,
      streak: isConsecutive || !lastPlayed ? points.streak + 1 : 1,
      lastPlayedDate: today,
    });

    addBonusPenaltyRecord(
      'bonus',
      10,
      'プレイスケジュールを完了しました',
      game?.title
    );

    toast.success('+10ポイント獲得！', {
      description: 'プレイを完了しました',
    });
  };

  const handleSkipSchedule = (scheduleId: string) => {
    const schedule = schedules.find(s => s.id === scheduleId);
    if (!schedule) return;

    const game = games.find(g => g.id === schedule.gameId);

    setSchedules(
      schedules.map(s =>
        s.id === scheduleId ? { ...s, skipped: true } : s
      )
    );

    // ポイント減点
    setPoints({
      ...points,
      total: Math.max(0, points.total - 5),
      streak: 0,
    });

    addBonusPenaltyRecord(
      'penalty',
      -5,
      'スケジュールをスキップしました',
      game?.title
    );

    toast.error('-5ポイント', {
      description: 'スケジュールをスキップしました',
    });
  };

  const handleMoveSchedule = (scheduleId: string, newDate: string, newStartTime: string) => {
    setSchedules(
      schedules.map(s =>
        s.id === scheduleId
          ? { ...s, date: newDate, startTime: newStartTime }
          : s
      )
    );
    toast.success('スケジュールを移動しました');
  };

  const handleMarkAsRead = (id: string) => {
    setNotifications(
      notifications.map(n => (n.id === id ? { ...n, read: true } : n))
    );
  };

  const handleMarkAllAsRead = () => {
    setNotifications(notifications.map(n => ({ ...n, read: true })));
    toast.success('すべての通知を既読にしました');
  };

  const handleUpdateNotificationSettings = (settings: NotificationSettingsType) => {
    setNotificationSettings(settings);
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <Toaster />
      <BonusAnimation
        show={showBonusAnimation}
        type={bonusAnimationData?.type || 'bonus'}
        points={bonusAnimationData?.points || 0}
        message={bonusAnimationData?.message || ''}
        onClose={() => setShowBonusAnimation(false)}
      />
      
      <header className="bg-white shadow-sm border-b">
        <div className="max-w-7xl mx-auto px-4 py-4 sm:px-6 lg:px-8">
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-3">
              <Gamepad2 className="h-8 w-8 text-blue-600" />
              <h1 className="text-2xl text-gray-900">ゲームスケジューラー</h1>
            </div>
            <Button onClick={handleGenerateSchedules} size="lg">
              <Calendar className="h-5 w-5 mr-2" />
              スケジュール生成
            </Button>
          </div>
        </div>
      </header>

      <main className="max-w-7xl mx-auto px-4 py-8 sm:px-6 lg:px-8">
        <div className="mb-8 grid grid-cols-1 lg:grid-cols-3 gap-6">
          <div className="lg:col-span-2">
            <PointsDisplay points={points} />
          </div>
          <div>
            <MotivationGauge points={points} />
          </div>
        </div>

        <Tabs defaultValue="schedule" className="space-y-6">
          <TabsList className="grid w-full grid-cols-6">
            <TabsTrigger value="schedule" className="flex items-center gap-2">
              <Calendar className="h-4 w-4" />
              <span className="hidden sm:inline">スケジュール</span>
            </TabsTrigger>
            <TabsTrigger value="calendar" className="flex items-center gap-2">
              <Calendar className="h-4 w-4" />
              <span className="hidden sm:inline">カレンダー</span>
            </TabsTrigger>
            <TabsTrigger value="games" className="flex items-center gap-2">
              <Gamepad2 className="h-4 w-4" />
              <span className="hidden sm:inline">ゲーム</span>
            </TabsTrigger>
            <TabsTrigger value="events" className="flex items-center gap-2">
              <Clock className="h-4 w-4" />
              <span className="hidden sm:inline">固定予定</span>
            </TabsTrigger>
            <TabsTrigger value="history" className="flex items-center gap-2">
              <Award className="h-4 w-4" />
              <span className="hidden sm:inline">履歴</span>
            </TabsTrigger>
            <TabsTrigger value="settings" className="flex items-center gap-2">
              <Settings className="h-4 w-4" />
              <span className="hidden sm:inline">設定</span>
            </TabsTrigger>
          </TabsList>

          <TabsContent value="schedule" className="space-y-6">
            <ScheduleCalendar
              schedules={schedules}
              games={games}
              onCompleteSchedule={handleCompleteSchedule}
              onSkipSchedule={handleSkipSchedule}
            />
          </TabsContent>

          <TabsContent value="calendar" className="space-y-6">
            <CalendarView
              schedules={schedules}
              games={games}
              onCompleteSchedule={handleCompleteSchedule}
              onSkipSchedule={handleSkipSchedule}
              onMoveSchedule={handleMoveSchedule}
            />
          </TabsContent>

          <TabsContent value="games" className="space-y-6">
            <GameFormAdvanced onAddGame={handleAddGame} />
            <GameListAdvanced
              games={games}
              onDeleteGame={handleDeleteGame}
              onUpdateStatus={handleUpdateStatus}
              onUpdateGenre={handleUpdateGenre}
            />
          </TabsContent>

          <TabsContent value="events" className="space-y-6">
            <FixedEventManager
              events={fixedEvents}
              onAddEvent={handleAddFixedEvent}
              onDeleteEvent={handleDeleteFixedEvent}
            />
          </TabsContent>

          <TabsContent value="history" className="space-y-6">
            <BonusPenaltyHistory records={bonusPenaltyRecords} />
          </TabsContent>

          <TabsContent value="settings" className="space-y-6">
            <NotificationSettings
              settings={notificationSettings}
              onUpdateSettings={handleUpdateNotificationSettings}
            />
            <NotificationPanel
              notifications={notifications}
              onMarkAsRead={handleMarkAsRead}
              onMarkAllAsRead={handleMarkAllAsRead}
            />
          </TabsContent>
        </Tabs>
      </main>
    </div>
  );
}

export default App;
