import { useState, useEffect } from "react";
import {
  NotificationRecord,
  NotificationSettings as NotificationSettingsType,
  Schedule,
  Game,
} from "../types";
import { storage } from "../lib/storage";
import { toast } from "sonner@2.0.3";

export function useNotificationSystem(schedules: Schedule[], games: Game[]) {
  const [notifications, setNotifications] = useState<NotificationRecord[]>([]);
  const [notificationSettings, setNotificationSettings] =
    useState<NotificationSettingsType>({
      enabled: true,
      minutesBefore: 15,
      soundEnabled: true,
    });

  // データの読み込み
  useEffect(() => {
    setNotifications(storage.getNotifications());
    setNotificationSettings(storage.getNotificationSettings());

    // 通知権限のリクエスト
    if ("Notification" in window && Notification.permission === "default") {
      Notification.requestPermission();
    }
  }, []);

  // データの保存
  useEffect(() => {
    storage.saveNotifications(notifications);
  }, [notifications]);

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

        const scheduleTime = new Date(
          `${schedule.date}T${schedule.startTime}`
        );
        const timeDiff = scheduleTime.getTime() - now.getTime();
        const minutesDiff = timeDiff / (1000 * 60);

        // 設定した時間前に通知
        return (
          minutesDiff > 0 && minutesDiff <= notificationSettings.minutesBefore
        );
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
            if (
              "Notification" in window &&
              Notification.permission === "granted"
            ) {
              new Notification("ゲームプレイの時間です！", {
                body: `${game.title} - ${schedule.startTime}から`,
                icon: "/favicon.ico",
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

  const handleMarkAsRead = (id: string) => {
    setNotifications(
      notifications.map(n => (n.id === id ? { ...n, read: true } : n))
    );
  };

  const handleMarkAllAsRead = () => {
    setNotifications(notifications.map(n => ({ ...n, read: true })));
    toast.success("すべての通知を既読にしました");
  };

  const handleUpdateNotificationSettings = (
    settings: NotificationSettingsType
  ) => {
    setNotificationSettings(settings);
  };

  return {
    notifications,
    notificationSettings,
    handleMarkAsRead,
    handleMarkAllAsRead,
    handleUpdateNotificationSettings,
  };
}
