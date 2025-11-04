import { Schedule, Game } from '../types';
import { Card, CardContent, CardHeader, CardTitle } from './ui/card';
import { Badge } from './ui/badge';
import { Button } from './ui/button';
import { Check, X, Clock } from 'lucide-react';
import { useState } from 'react';

interface ScheduleCalendarProps {
  schedules: Schedule[];
  games: Game[];
  onCompleteSchedule: (scheduleId: string) => void;
  onSkipSchedule: (scheduleId: string) => void;
}

export function ScheduleCalendar({
  schedules,
  games,
  onCompleteSchedule,
  onSkipSchedule,
}: ScheduleCalendarProps) {
  const [viewMode, setViewMode] = useState<'week' | 'list'>('list');

  const getGameTitle = (gameId: string) => {
    const game = games.find(g => g.id === gameId);
    return game?.title || '不明なゲーム';
  };

  const formatDate = (dateStr: string) => {
    const date = new Date(dateStr);
    const today = new Date();
    today.setHours(0, 0, 0, 0);
    const scheduleDate = new Date(date);
    scheduleDate.setHours(0, 0, 0, 0);

    const diffTime = scheduleDate.getTime() - today.getTime();
    const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));

    if (diffDays === 0) return '今日';
    if (diffDays === 1) return '明日';
    if (diffDays === -1) return '昨日';

    const weekday = ['日', '月', '火', '水', '木', '金', '土'][date.getDay()];
    return `${date.getMonth() + 1}/${date.getDate()}(${weekday})`;
  };

  const isPast = (dateStr: string, endTime: string) => {
    const scheduleDateTime = new Date(`${dateStr}T${endTime}`);
    return scheduleDateTime < new Date();
  };

  // 日付ごとにグループ化
  const groupedSchedules = schedules.reduce((acc, schedule) => {
    if (!acc[schedule.date]) {
      acc[schedule.date] = [];
    }
    acc[schedule.date].push(schedule);
    return acc;
  }, {} as Record<string, Schedule[]>);

  const sortedDates = Object.keys(groupedSchedules).sort();

  if (schedules.length === 0) {
    return (
      <Card>
        <CardContent className="pt-6">
          <p className="text-center text-gray-500">
            スケジュールが生成されていません。ゲームを追加して「スケジュール生成」ボタンを押してください。
          </p>
        </CardContent>
      </Card>
    );
  }

  return (
    <div className="space-y-4">
      {sortedDates.map(date => (
        <Card key={date}>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Clock className="h-5 w-5" />
              {formatDate(date)}
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-3">
              {groupedSchedules[date].map(schedule => (
                <div
                  key={schedule.id}
                  className={`flex items-center justify-between p-3 rounded-lg border ${
                    schedule.completed
                      ? 'bg-green-50 border-green-200'
                      : schedule.skipped
                      ? 'bg-red-50 border-red-200'
                      : 'bg-white'
                  }`}
                >
                  <div className="flex-1">
                    <div className="flex items-center gap-2 mb-1">
                      <span className="text-sm text-gray-600">
                        {schedule.startTime} - {schedule.endTime}
                      </span>
                      {schedule.completed && (
                        <Badge variant="outline" className="bg-green-100 text-green-700 border-green-300">
                          完了
                        </Badge>
                      )}
                      {schedule.skipped && (
                        <Badge variant="outline" className="bg-red-100 text-red-700 border-red-300">
                          スキップ
                        </Badge>
                      )}
                    </div>
                    <p>{getGameTitle(schedule.gameId)}</p>
                  </div>

                  {!schedule.completed && !schedule.skipped && (
                    <div className="flex gap-2">
                      <Button
                        size="sm"
                        variant="outline"
                        onClick={() => onCompleteSchedule(schedule.id)}
                        className="text-green-600 hover:text-green-700"
                      >
                        <Check className="h-4 w-4 mr-1" />
                        完了
                      </Button>
                      <Button
                        size="sm"
                        variant="outline"
                        onClick={() => onSkipSchedule(schedule.id)}
                        className="text-red-600 hover:text-red-700"
                      >
                        <X className="h-4 w-4 mr-1" />
                        スキップ
                      </Button>
                    </div>
                  )}
                </div>
              ))}
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
}
