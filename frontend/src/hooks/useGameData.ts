import { useState, useEffect } from "react";
import { Game, Schedule, FixedEvent, PlayStatus, GameGenre } from "../types";
import { storage } from "../lib/storage";
import { generateSchedules } from "../lib/scheduleGenerator";
import { toast } from "sonner@2.0.3";

export function useGameData() {
  const [games, setGames] = useState<Game[]>([]);
  const [schedules, setSchedules] = useState<Schedule[]>([]);
  const [fixedEvents, setFixedEvents] = useState<FixedEvent[]>([]);

  // データの読み込み
  useEffect(() => {
    setGames(storage.getGames());
    setSchedules(storage.getSchedules());
    setFixedEvents(storage.getFixedEvents());
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

  const handleAddGame = (title: string, genre?: GameGenre) => {
    const newGame: Game = {
      id: `game_${Date.now()}`,
      title,
      genre,
      status: "unstarted",
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
    toast.success("ステータスを更新しました");
  };

  const handleUpdateGenre = (id: string, genre: GameGenre) => {
    setGames(games.map(g => (g.id === id ? { ...g, genre } : g)));
    toast.success("ジャンルを更新しました");
  };

  const handleAddFixedEvent = (event: Omit<FixedEvent, "id">) => {
    const newEvent: FixedEvent = {
      ...event,
      id: `event_${Date.now()}`,
    };
    setFixedEvents([...fixedEvents, newEvent]);
    toast.success("固定予定を追加しました");
  };

  const handleDeleteFixedEvent = (id: string) => {
    setFixedEvents(fixedEvents.filter(e => e.id !== id));
    toast.success("固定予定を削除しました");
  };

  const handleGenerateSchedules = () => {
    const newSchedules = generateSchedules(games, fixedEvents);
    setSchedules(newSchedules);
    toast.success(`${newSchedules.length}件のスケジュールを生成しました`);
  };

  const handleMoveSchedule = (
    scheduleId: string,
    newDate: string,
    newStartTime: string
  ) => {
    setSchedules(
      schedules.map(s =>
        s.id === scheduleId
          ? { ...s, date: newDate, startTime: newStartTime }
          : s
      )
    );
    toast.success("スケジュールを移動しました");
  };

  const handleCompleteSchedule = (scheduleId: string) => {
    setSchedules(
      schedules.map(s => (s.id === scheduleId ? { ...s, completed: true } : s))
    );
  };

  const handleSkipSchedule = (scheduleId: string) => {
    setSchedules(
      schedules.map(s => (s.id === scheduleId ? { ...s, skipped: true } : s))
    );
  };

  return {
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
    handleCompleteSchedule,
    handleSkipSchedule,
  };
}
