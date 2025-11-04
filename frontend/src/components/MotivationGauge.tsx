import { UserPoints } from '../types';
import { Card, CardContent, CardHeader, CardTitle } from './ui/card';
import { Progress } from './ui/progress';
import { Flame, Target, Zap } from 'lucide-react';

interface MotivationGaugeProps {
  points: UserPoints;
}

export function MotivationGauge({ points }: MotivationGaugeProps) {
  // モチベーションレベルを計算（0-100）
  const calculateMotivation = (): number => {
    const pointsFactor = Math.min(points.total / 2, 50); // 最大50ポイント分
    const streakFactor = Math.min(points.streak * 5, 50); // 最大50ポイント分
    return Math.min(Math.round(pointsFactor + streakFactor), 100);
  };

  const motivation = calculateMotivation();

  const getMotivationLevel = (value: number): { label: string; color: string } => {
    if (value >= 80) return { label: '絶好調！', color: 'text-green-600' };
    if (value >= 60) return { label: '好調', color: 'text-blue-600' };
    if (value >= 40) return { label: '普通', color: 'text-yellow-600' };
    if (value >= 20) return { label: 'やや低め', color: 'text-orange-600' };
    return { label: '要注意', color: 'text-red-600' };
  };

  const level = getMotivationLevel(motivation);

  const getProgressColor = (value: number): string => {
    if (value >= 80) return 'bg-green-500';
    if (value >= 60) return 'bg-blue-500';
    if (value >= 40) return 'bg-yellow-500';
    if (value >= 20) return 'bg-orange-500';
    return 'bg-red-500';
  };

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center gap-2">
          <Zap className="h-5 w-5 text-yellow-500" />
          モチベーションゲージ
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="space-y-2">
          <div className="flex items-center justify-between">
            <span className="text-sm">現在のレベル</span>
            <span className={`${level.color}`}>{level.label}</span>
          </div>
          
          <div className="relative">
            <Progress value={motivation} className="h-4" />
            <div
              className={`absolute inset-0 h-4 rounded-full ${getProgressColor(motivation)} transition-all`}
              style={{ width: `${motivation}%` }}
            />
          </div>
          
          <div className="text-center text-2xl">{motivation}%</div>
        </div>

        <div className="grid grid-cols-2 gap-4 pt-4 border-t">
          <div className="text-center">
            <div className="flex items-center justify-center gap-1 mb-1">
              <Target className="h-4 w-4 text-blue-500" />
              <span className="text-sm text-gray-600">ポイント</span>
            </div>
            <div className="text-xl">{points.total}</div>
          </div>
          
          <div className="text-center">
            <div className="flex items-center justify-center gap-1 mb-1">
              <Flame className="h-4 w-4 text-orange-500" />
              <span className="text-sm text-gray-600">連続日数</span>
            </div>
            <div className="text-xl">{points.streak}日</div>
          </div>
        </div>

        <div className="pt-4 border-t">
          <p className="text-sm text-gray-600 text-center">
            {motivation >= 80
              ? '素晴らしい調子です！この調子で続けましょう！'
              : motivation >= 60
              ? '順調です！もう少しでゴールドランクですよ！'
              : motivation >= 40
              ? 'プレイを続けてモチベーションを上げましょう！'
              : motivation >= 20
              ? 'スケジュールを守ってポイントを獲得しましょう！'
              : '新しいゲームに挑戦してみませんか？'}
          </p>
        </div>
      </CardContent>
    </Card>
  );
}
