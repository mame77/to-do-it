import { useGameData } from "./hooks/useGameData";
import { usePointsSystem } from "./hooks/usePointsSystem";
import { useNotificationSystem } from "./hooks/useNotificationSystem";
import { useBonusAnimation } from "./hooks/useBonusAnimation";
import { GameListAdvanced } from "./components/GameListAdvanced";
import { GameFormAdvanced } from "./components/GameFormAdvanced";
import { ScheduleCalendar } from "./components/ScheduleCalendar";
import { CalendarView } from "./components/CalendarView";
import { FixedEventManager } from "./components/FixedEventManager";
import { PointsDisplay } from "./components/PointsDisplay";
import { NotificationPanel } from "./components/NotificationPanel";
import { BonusPenaltyHistory } from "./components/BonusPenaltyHistory";
import { MotivationGauge } from "./components/MotivationGauge";
import { NotificationSettings } from "./components/NotificationSettings";
import { BonusAnimation } from "./components/BonusAnimation";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "./components/ui/tabs";
import { Button } from "./components/ui/button";
import {
  Gamepad2,
  Calendar,
  Clock,
  Settings,
  Award,
} from "lucide-react";
import { Toaster } from "./components/ui/sonner";

function App() {
  // カスタムフックで状態管理を分離
  const {
    games,
    schedules,
    fixedEvents,
    handleAddGame,
    handleDeleteGame,
    handleUpdateStatus,
    handleUpdateGenre,
    handleAddFixedEvent,
    handleDeleteFixedEvent,
    handleGenerateSchedules,
    handleMoveSchedule,
    handleCompleteSchedule: completeSchedule,
    handleSkipSchedule: skipSchedule,
  } = useGameData();

  const {
    points,
    bonusPenaltyRecords,
    awardCompletionBonus,
    applySkipPenalty,
  } = usePointsSystem();

  const {
    notifications,
    notificationSettings,
    handleMarkAsRead,
    handleMarkAllAsRead,
    handleUpdateNotificationSettings,
  } = useNotificationSystem(schedules, games);

  const { showBonusAnimation, bonusAnimationData, showAnimation, hideAnimation } =
    useBonusAnimation();

  // スケジュール完了時の処理（ポイント付与とアニメーション）
  const handleCompleteSchedule = (scheduleId: string) => {
    const schedule = schedules.find(s => s.id === scheduleId);
    if (!schedule) return;

    const game = games.find(g => g.id === schedule.gameId);
    completeSchedule(scheduleId);

    const result = awardCompletionBonus(game?.title);
    showAnimation(result.type, result.points, "プレイスケジュールを完了しました");
  };

  // スケジュールスキップ時の処理（ペナルティとアニメーション）
  const handleSkipSchedule = (scheduleId: string) => {
    const schedule = schedules.find(s => s.id === scheduleId);
    if (!schedule) return;

    const game = games.find(g => g.id === schedule.gameId);
    skipSchedule(scheduleId);

    const result = applySkipPenalty(game?.title);
    showAnimation(result.type, result.points, "スケジュールをスキップしました");
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <Toaster />
      <BonusAnimation
        show={showBonusAnimation}
        type={bonusAnimationData?.type || "bonus"}
        points={bonusAnimationData?.points || 0}
        message={bonusAnimationData?.message || ""}
        onClose={hideAnimation}
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
