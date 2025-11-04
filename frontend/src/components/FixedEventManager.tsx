import { useState } from 'react';
import { FixedEvent } from '../types';
import { Button } from './ui/button';
import { Input } from './ui/input';
import { Label } from './ui/label';
import { Card, CardContent, CardHeader, CardTitle } from './ui/card';
import { Switch } from './ui/switch';
import { Checkbox } from './ui/checkbox';
import { Trash2, Plus } from 'lucide-react';

interface FixedEventManagerProps {
  events: FixedEvent[];
  onAddEvent: (event: Omit<FixedEvent, 'id'>) => void;
  onDeleteEvent: (id: string) => void;
}

const weekdays = ['日', '月', '火', '水', '木', '金', '土'];

export function FixedEventManager({
  events,
  onAddEvent,
  onDeleteEvent,
}: FixedEventManagerProps) {
  const [title, setTitle] = useState('');
  const [startTime, setStartTime] = useState('09:00');
  const [endTime, setEndTime] = useState('10:00');
  const [isRecurring, setIsRecurring] = useState(true);
  const [selectedDays, setSelectedDays] = useState<number[]>([1, 2, 3, 4, 5]); // 平日デフォルト
  const [specificDate, setSpecificDate] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!title.trim()) return;
    if (isRecurring && selectedDays.length === 0) return;
    if (!isRecurring && !specificDate) return;

    onAddEvent({
      title: title.trim(),
      startTime,
      endTime,
      isRecurring,
      dayOfWeek: isRecurring ? selectedDays : [],
      specificDate: !isRecurring ? specificDate : undefined,
    });

    setTitle('');
    setSelectedDays([1, 2, 3, 4, 5]);
    setSpecificDate('');
  };

  const toggleDay = (day: number) => {
    setSelectedDays(prev =>
      prev.includes(day) ? prev.filter(d => d !== day) : [...prev, day]
    );
  };

  const formatEventTime = (event: FixedEvent) => {
    if (event.isRecurring) {
      const days = event.dayOfWeek.map(d => weekdays[d]).join(', ');
      return `毎週${days} ${event.startTime}-${event.endTime}`;
    } else {
      const date = new Date(event.specificDate!);
      return `${date.getMonth() + 1}/${date.getDate()} ${event.startTime}-${event.endTime}`;
    }
  };

  return (
    <div className="space-y-4">
      <Card>
        <CardHeader>
          <CardTitle>固定予定を追加</CardTitle>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <Label htmlFor="event-title">予定名</Label>
              <Input
                id="event-title"
                type="text"
                placeholder="例: 授業、仕事"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
              />
            </div>

            <div className="flex items-center space-x-2">
              <Switch
                id="recurring"
                checked={isRecurring}
                onCheckedChange={setIsRecurring}
              />
              <Label htmlFor="recurring">繰り返し予定</Label>
            </div>

            {isRecurring ? (
              <div>
                <Label>曜日を選択</Label>
                <div className="flex gap-2 mt-2">
                  {weekdays.map((day, index) => (
                    <button
                      key={index}
                      type="button"
                      onClick={() => toggleDay(index)}
                      className={`w-10 h-10 rounded-full border-2 transition-colors ${
                        selectedDays.includes(index)
                          ? 'bg-blue-500 text-white border-blue-500'
                          : 'bg-white text-gray-700 border-gray-300'
                      }`}
                    >
                      {day}
                    </button>
                  ))}
                </div>
              </div>
            ) : (
              <div>
                <Label htmlFor="specific-date">日付</Label>
                <Input
                  id="specific-date"
                  type="date"
                  value={specificDate}
                  onChange={(e) => setSpecificDate(e.target.value)}
                />
              </div>
            )}

            <div className="grid grid-cols-2 gap-4">
              <div>
                <Label htmlFor="start-time">開始時刻</Label>
                <Input
                  id="start-time"
                  type="time"
                  value={startTime}
                  onChange={(e) => setStartTime(e.target.value)}
                />
              </div>
              <div>
                <Label htmlFor="end-time">終了時刻</Label>
                <Input
                  id="end-time"
                  type="time"
                  value={endTime}
                  onChange={(e) => setEndTime(e.target.value)}
                />
              </div>
            </div>

            <Button type="submit" className="w-full">
              <Plus className="h-4 w-4 mr-2" />
              予定を追加
            </Button>
          </form>
        </CardContent>
      </Card>

      {events.length > 0 && (
        <Card>
          <CardHeader>
            <CardTitle>登録済み固定予定</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-2">
              {events.map(event => (
                <div
                  key={event.id}
                  className="flex items-center justify-between p-3 border rounded-lg"
                >
                  <div>
                    <p>{event.title}</p>
                    <p className="text-sm text-gray-600">{formatEventTime(event)}</p>
                  </div>
                  <Button
                    variant="ghost"
                    size="icon"
                    onClick={() => onDeleteEvent(event.id)}
                  >
                    <Trash2 className="h-4 w-4" />
                  </Button>
                </div>
              ))}
            </div>
          </CardContent>
        </Card>
      )}
    </div>
  );
}
