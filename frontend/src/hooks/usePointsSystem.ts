import { useState, useEffect } from "react";
import { UserPoints, BonusPenaltyRecord } from "../types";
import { storage } from "../lib/storage";
import { toast } from "sonner@2.0.3";

export function usePointsSystem() {
  const [points, setPoints] = useState<UserPoints>({
    total: 0,
    streak: 0,
    lastPlayedDate: null,
  });
  const [bonusPenaltyRecords, setBonusPenaltyRecords] = useState<
    BonusPenaltyRecord[]
  >([]);

  // データの読み込み
  useEffect(() => {
    setPoints(storage.getPoints());
    setBonusPenaltyRecords(storage.getBonusPenaltyRecords());
  }, []);

  // データの保存
  useEffect(() => {
    storage.savePoints(points);
  }, [points]);

  useEffect(() => {
    storage.saveBonusPenaltyRecords(bonusPenaltyRecords);
  }, [bonusPenaltyRecords]);

  const addBonusPenaltyRecord = (
    type: "bonus" | "penalty",
    pointsValue: number,
    reason: string,
    gameTitle?: string
  ) => {
    const record: BonusPenaltyRecord = {
      id: `bp_${Date.now()}`,
      type,
      points: pointsValue,
      reason,
      gameTitle,
      createdAt: new Date().toISOString(),
    };
    setBonusPenaltyRecords(prev => [record, ...prev]);
    return record;
  };

  const awardCompletionBonus = (gameTitle?: string) => {
    const today = new Date().toISOString().split("T")[0];
    const lastPlayed = points.lastPlayedDate;
    const isConsecutive =
      lastPlayed ===
      new Date(Date.now() - 86400000).toISOString().split("T")[0];

    setPoints({
      total: points.total + 10,
      streak: isConsecutive || !lastPlayed ? points.streak + 1 : 1,
      lastPlayedDate: today,
    });

    addBonusPenaltyRecord(
      "bonus",
      10,
      "プレイスケジュールを完了しました",
      gameTitle
    );

    toast.success("+10ポイント獲得！", {
      description: "プレイを完了しました",
    });

    return { points: 10, type: "bonus" as const };
  };

  const applySkipPenalty = (gameTitle?: string) => {
    setPoints({
      ...points,
      total: Math.max(0, points.total - 5),
      streak: 0,
    });

    addBonusPenaltyRecord(
      "penalty",
      -5,
      "スケジュールをスキップしました",
      gameTitle
    );

    toast.error("-5ポイント", {
      description: "スケジュールをスキップしました",
    });

    return { points: -5, type: "penalty" as const };
  };

  return {
    points,
    bonusPenaltyRecords,
    addBonusPenaltyRecord,
    awardCompletionBonus,
    applySkipPenalty,
  };
}
