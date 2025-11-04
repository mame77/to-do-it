import { UserPoints } from '../types';
import { Card, CardContent, CardHeader, CardTitle } from './ui/card';
import { Trophy, Flame, Star } from 'lucide-react';

interface PointsDisplayProps {
  points: UserPoints;
}

export function PointsDisplay({ points }: PointsDisplayProps) {
  return (
    <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm">総ポイント</CardTitle>
          <Star className="h-4 w-4 text-yellow-500" />
        </CardHeader>
        <CardContent>
          <div className="text-2xl">{points.total}</div>
          <p className="text-xs text-gray-600 mt-1">
            完了: +10pt / スキップ: -5pt
          </p>
        </CardContent>
      </Card>

      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm">連続プレイ</CardTitle>
          <Flame className="h-4 w-4 text-orange-500" />
        </CardHeader>
        <CardContent>
          <div className="text-2xl">{points.streak}日</div>
          <p className="text-xs text-gray-600 mt-1">
            継続してプレイしよう！
          </p>
        </CardContent>
      </Card>

      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm">ランク</CardTitle>
          <Trophy className="h-4 w-4 text-purple-500" />
        </CardHeader>
        <CardContent>
          <div className="text-2xl">
            {points.total >= 100 ? 'ゴールド' : points.total >= 50 ? 'シルバー' : 'ブロンズ'}
          </div>
          <p className="text-xs text-gray-600 mt-1">
            {points.total >= 100 ? '最高ランク！' : `次まであと${points.total >= 50 ? 100 - points.total : 50 - points.total}pt`}
          </p>
        </CardContent>
      </Card>
    </div>
  );
}
