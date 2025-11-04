import { Game, PlayStatus } from '../types';
import { Button } from './ui/button';
import { Badge } from './ui/badge';
import { Card, CardContent, CardHeader, CardTitle } from './ui/card';
import { Trash2, Edit } from 'lucide-react';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from './ui/select';

interface GameListProps {
  games: Game[];
  onDeleteGame: (id: string) => void;
  onUpdateStatus: (id: string, status: PlayStatus) => void;
}

const statusLabels: Record<PlayStatus, string> = {
  unstarted: '未開始',
  playing: 'プレイ中',
  completed: 'クリア済み',
};

const statusColors: Record<PlayStatus, string> = {
  unstarted: 'bg-gray-500',
  playing: 'bg-blue-500',
  completed: 'bg-green-500',
};

export function GameList({ games, onDeleteGame, onUpdateStatus }: GameListProps) {
  if (games.length === 0) {
    return (
      <Card>
        <CardContent className="pt-6">
          <p className="text-center text-gray-500">
            ゲームが登録されていません。まずはゲームを追加してください。
          </p>
        </CardContent>
      </Card>
    );
  }

  return (
    <div className="space-y-3">
      {games.map(game => (
        <Card key={game.id}>
          <CardContent className="pt-6">
            <div className="flex items-center justify-between gap-4">
              <div className="flex-1 min-w-0">
                <h3 className="truncate">{game.title}</h3>
              </div>
              
              <div className="flex items-center gap-2">
                <Select
                  value={game.status}
                  onValueChange={(value: PlayStatus) => onUpdateStatus(game.id, value)}
                >
                  <SelectTrigger className="w-[140px]">
                    <SelectValue />
                  </SelectTrigger>
                  <SelectContent>
                    {Object.entries(statusLabels).map(([value, label]) => (
                      <SelectItem key={value} value={value}>
                        <div className="flex items-center gap-2">
                          <div className={`w-2 h-2 rounded-full ${statusColors[value as PlayStatus]}`} />
                          {label}
                        </div>
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>

                <Button
                  variant="ghost"
                  size="icon"
                  onClick={() => onDeleteGame(game.id)}
                >
                  <Trash2 className="h-4 w-4" />
                </Button>
              </div>
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
}
