import { useState } from 'react';
import { Schedule, Game } from '../types';
import { Card, CardContent, CardHeader, CardTitle } from './ui/card';
import { Button } from './ui/button';
import { Badge } from './ui/badge';
import { ChevronLeft, ChevronRight, Calendar, Check, X } from 'lucide-react';
import { Tabs, TabsContent, TabsList, TabsTrigger } from './ui/tabs';

interface CalendarViewProps {
  schedules: Schedule[];
  games: Game[];
  onCompleteSchedule: (scheduleId: string) => void;
  onSkipSchedule: (scheduleId: string) => void;
  onMoveSchedule: (scheduleId: string, newDate: string, newStartTime: string) => void;
}

export function CalendarView({
  schedules,
  games,
  onCompleteSchedule,
  onSkipSchedule,
  onMoveSchedule,
}: CalendarViewProps) {
  const [currentDate, setCurrentDate] = useState(new Date());
  const [viewMode, setViewMode] = useState<'week' | 'month'>('week');

  const getGameTitle = (gameId: string) => {
    const game = games.find(g => g.id === gameId);
    return game?.title || '不明なゲーム';
  };

  const formatDate = (date: Date) => {
    return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`;
  };

  const getWeekDates = (date: Date): Date[] => {
    const start = new Date(date);
    start.setDate(date.getDate() - date.getDay()); // 日曜日から開始
    
    const dates: Date[] = [];
    for (let i = 0; i < 7; i++) {
      const d = new Date(start);
      d.setDate(start.getDate() + i);
      dates.push(d);
    }
    return dates;
  };

  const getMonthDates = (date: Date): Date[] => {
    const year = date.getFullYear();
    const month = date.getMonth();
    
    const firstDay = new Date(year, month, 1);
    const lastDay = new Date(year, month + 1, 0);
    
    const startDate = new Date(firstDay);
    startDate.setDate(firstDay.getDate() - firstDay.getDay());
    
    const dates: Date[] = [];
    let current = new Date(startDate);
    
    while (dates.length < 42) { // 6週間分
      dates.push(new Date(current));
      current.setDate(current.getDate() + 1);
    }
    
    return dates;
  };

  const getSchedulesForDate = (date: Date): Schedule[] => {
    const dateStr = date.toISOString().split('T')[0];
    return schedules.filter(s => s.date === dateStr).sort((a, b) => 
      a.startTime.localeCompare(b.startTime)
    );
  };

  const navigateWeek = (direction: 'prev' | 'next') => {
    const newDate = new Date(currentDate);
    newDate.setDate(currentDate.getDate() + (direction === 'next' ? 7 : -7));
    setCurrentDate(newDate);
  };

  const navigateMonth = (direction: 'prev' | 'next') => {
    const newDate = new Date(currentDate);
    newDate.setMonth(currentDate.getMonth() + (direction === 'next' ? 1 : -1));
    setCurrentDate(newDate);
  };

  const isToday = (date: Date): boolean => {
    const today = new Date();
    return date.toDateString() === today.toDateString();
  };

  const weekdays = ['日', '月', '火', '水', '木', '金', '土'];

  if (schedules.length === 0) {
    return (
      <Card>
        <CardContent className="pt-6">
          <p className="text-center text-gray-500">
            スケジュールが生成されていません。
          </p>
        </CardContent>
      </Card>
    );
  }

  return (
    <Card>
      <CardHeader>
        <div className="flex items-center justify-between">
          <CardTitle className="flex items-center gap-2">
            <Calendar className="h-5 w-5" />
            {viewMode === 'week' ? '週間表示' : '月間表示'}
          </CardTitle>
          <Tabs value={viewMode} onValueChange={(v) => setViewMode(v as 'week' | 'month')}>
            <TabsList>
              <TabsTrigger value="week">週</TabsTrigger>
              <TabsTrigger value="month">月</TabsTrigger>
            </TabsList>
          </Tabs>
        </div>
      </CardHeader>
      <CardContent>
        <div className="mb-4 flex items-center justify-between">
          <Button
            variant="outline"
            size="sm"
            onClick={() => viewMode === 'week' ? navigateWeek('prev') : navigateMonth('prev')}
          >
            <ChevronLeft className="h-4 w-4" />
          </Button>
          
          <div className="text-center">
            <p>{currentDate.getFullYear()}年{currentDate.getMonth() + 1}月</p>
          </div>
          
          <Button
            variant="outline"
            size="sm"
            onClick={() => viewMode === 'week' ? navigateWeek('next') : navigateMonth('next')}
          >
            <ChevronRight className="h-4 w-4" />
          </Button>
        </div>

        {viewMode === 'week' ? (
          <WeekView
            dates={getWeekDates(currentDate)}
            getSchedulesForDate={getSchedulesForDate}
            getGameTitle={getGameTitle}
            isToday={isToday}
            onCompleteSchedule={onCompleteSchedule}
            onSkipSchedule={onSkipSchedule}
          />
        ) : (
          <MonthView
            dates={getMonthDates(currentDate)}
            getSchedulesForDate={getSchedulesForDate}
            currentMonth={currentDate.getMonth()}
            isToday={isToday}
          />
        )}
      </CardContent>
    </Card>
  );
}

function WeekView({
  dates,
  getSchedulesForDate,
  getGameTitle,
  isToday,
  onCompleteSchedule,
  onSkipSchedule,
}: any) {
  const weekdays = ['日', '月', '火', '水', '木', '金', '土'];
  
  return (
    <div className="grid grid-cols-7 gap-2">
      {dates.map((date: Date, index: number) => {
        const daySchedules = getSchedulesForDate(date);
        const today = isToday(date);
        
        return (
          <div
            key={index}
            className={`border rounded-lg p-2 min-h-[200px] ${
              today ? 'bg-blue-50 border-blue-300' : 'bg-white'
            }`}
          >
            <div className="text-center mb-2">
              <div className="text-xs text-gray-500">{weekdays[date.getDay()]}</div>
              <div className={`${today ? 'text-blue-600' : ''}`}>
                {date.getDate()}
              </div>
            </div>
            
            <div className="space-y-1">
              {daySchedules.map((schedule: Schedule) => (
                <div
                  key={schedule.id}
                  className={`text-xs p-2 rounded ${
                    schedule.completed
                      ? 'bg-green-100 border border-green-300'
                      : schedule.skipped
                      ? 'bg-red-100 border border-red-300'
                      : 'bg-gray-100 border border-gray-300'
                  }`}
                >
                  <div className="truncate mb-1">{getGameTitle(schedule.gameId)}</div>
                  <div className="text-gray-600">{schedule.startTime}</div>
                  
                  {!schedule.completed && !schedule.skipped && (
                    <div className="flex gap-1 mt-1">
                      <button
                        onClick={() => onCompleteSchedule(schedule.id)}
                        className="flex-1 p-1 bg-green-500 text-white rounded hover:bg-green-600"
                      >
                        <Check className="h-3 w-3 mx-auto" />
                      </button>
                      <button
                        onClick={() => onSkipSchedule(schedule.id)}
                        className="flex-1 p-1 bg-red-500 text-white rounded hover:bg-red-600"
                      >
                        <X className="h-3 w-3 mx-auto" />
                      </button>
                    </div>
                  )}
                </div>
              ))}
            </div>
          </div>
        );
      })}
    </div>
  );
}

function MonthView({
  dates,
  getSchedulesForDate,
  currentMonth,
  isToday,
}: any) {
  const weekdays = ['日', '月', '火', '水', '木', '金', '土'];
  
  return (
    <div>
      <div className="grid grid-cols-7 gap-1 mb-2">
        {weekdays.map((day, index) => (
          <div key={index} className="text-center text-sm text-gray-600 py-2">
            {day}
          </div>
        ))}
      </div>
      
      <div className="grid grid-cols-7 gap-1">
        {dates.map((date: Date, index: number) => {
          const daySchedules = getSchedulesForDate(date);
          const today = isToday(date);
          const isCurrentMonth = date.getMonth() === currentMonth;
          
          return (
            <div
              key={index}
              className={`border rounded p-2 min-h-[80px] ${
                today
                  ? 'bg-blue-50 border-blue-300'
                  : isCurrentMonth
                  ? 'bg-white'
                  : 'bg-gray-50'
              }`}
            >
              <div
                className={`text-sm mb-1 ${
                  today ? 'text-blue-600' : !isCurrentMonth ? 'text-gray-400' : ''
                }`}
              >
                {date.getDate()}
              </div>
              
              <div className="space-y-1">
                {daySchedules.slice(0, 3).map((schedule: Schedule) => (
                  <div
                    key={schedule.id}
                    className={`text-xs p-1 rounded truncate ${
                      schedule.completed
                        ? 'bg-green-100'
                        : schedule.skipped
                        ? 'bg-red-100'
                        : 'bg-blue-100'
                    }`}
                  >
                    {schedule.startTime}
                  </div>
                ))}
                {daySchedules.length > 3 && (
                  <div className="text-xs text-gray-500">
                    +{daySchedules.length - 3}件
                  </div>
                )}
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
}
