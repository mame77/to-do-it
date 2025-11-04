import { BonusPenaltyRecord } from '../types';
import { Card, CardContent, CardHeader, CardTitle } from './ui/card';
import { Badge } from './ui/badge';
import { TrendingUp, TrendingDown, Award } from 'lucide-react';

interface BonusPenaltyHistoryProps {
  records: BonusPenaltyRecord[];
}

export function BonusPenaltyHistory({ records }: BonusPenaltyHistoryProps) {
  const formatDate = (dateStr: string) => {
    const date = new Date(dateStr);
    return `${date.getMonth() + 1}/${date.getDate()} ${date.getHours()}:${String(date.getMinutes()).padStart(2, '0')}`;
  };

  const sortedRecords = [...records].sort((a, b) => 
    new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime()
  );

  const totalBonus = records.filter(r => r.type === 'bonus').reduce((sum, r) => sum + r.points, 0);
  const totalPenalty = records.filter(r => r.type === 'penalty').reduce((sum, r) => sum + Math.abs(r.points), 0);

  return (
    <div className="space-y-4">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm">累計ボーナス</CardTitle>
            <TrendingUp className="h-4 w-4 text-green-500" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl text-green-600">+{totalBonus}</div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm">累計ペナルティ</CardTitle>
            <TrendingDown className="h-4 w-4 text-red-500" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl text-red-600">-{totalPenalty}</div>
          </CardContent>
        </Card>
      </div>

      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Award className="h-5 w-5" />
            ボーナス・ペナルティ履歴
          </CardTitle>
        </CardHeader>
        <CardContent>
          {sortedRecords.length === 0 ? (
            <p className="text-center text-gray-500 py-4">履歴がありません</p>
          ) : (
            <div className="space-y-2 max-h-96 overflow-y-auto">
              {sortedRecords.map(record => (
                <div
                  key={record.id}
                  className={`p-3 rounded-lg border ${
                    record.type === 'bonus'
                      ? 'bg-green-50 border-green-200'
                      : 'bg-red-50 border-red-200'
                  }`}
                >
                  <div className="flex items-center justify-between mb-1">
                    <Badge
                      variant={record.type === 'bonus' ? 'default' : 'destructive'}
                      className={record.type === 'bonus' ? 'bg-green-500' : ''}
                    >
                      {record.type === 'bonus' ? 'ボーナス' : 'ペナルティ'}
                    </Badge>
                    <span
                      className={`${
                        record.type === 'bonus' ? 'text-green-600' : 'text-red-600'
                      }`}
                    >
                      {record.points > 0 ? '+' : ''}{record.points}pt
                    </span>
                  </div>
                  <p className="text-sm">{record.reason}</p>
                  {record.gameTitle && (
                    <p className="text-sm text-gray-600 mt-1">ゲーム: {record.gameTitle}</p>
                  )}
                  <p className="text-xs text-gray-400 mt-1">{formatDate(record.createdAt)}</p>
                </div>
              ))}
            </div>
          )}
        </CardContent>
      </Card>
    </div>
  );
}
