import { useState } from 'react';
import { Game, PlayStatus, GameGenre } from '../types';
import { Button } from './ui/button';
import { Badge } from './ui/badge';
import { Card, CardContent, CardHeader, CardTitle } from './ui/card';
import { Trash2, Filter } from 'lucide-react';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from './ui/select';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuRadioGroup,
  DropdownMenuRadioItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from './ui/dropdown-menu';

interface GameListAdvancedProps {
  games: Game[];
  onDeleteGame: (id: string) => void;
  onUpdateStatus: (id: string, status: PlayStatus) => void;
  onUpdateGenre: (id: string, genre: GameGenre) => void;
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

const genres: GameGenre[] = ['RPG', 'アクション', 'アドベンチャー', 'シミュレーション', 'パズル', 'スポーツ', 'その他'];

export function GameListAdvanced({
  games,
  onDeleteGame,
  onUpdateStatus,
  onUpdateGenre,
}: GameListAdvancedProps) {
  const [statusFilter, setStatusFilter] = useState<PlayStatus | 'all'>('all');
  const [genreFilter, setGenreFilter] = useState<GameGenre | 'all'>('all');

  const filteredGames = games.filter(game => {
    const statusMatch = statusFilter === 'all' || game.status === statusFilter;
    const genreMatch = genreFilter === 'all' || game.genre === genreFilter;
    return statusMatch && genreMatch;
  });

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
    <div className="space-y-4">
      <Card>
        <CardContent className="pt-6">
          <div className="flex flex-wrap gap-4">
            <div className="flex-1 min-w-[200px]">
              <Label className="text-sm mb-2 block">ステータスでフィルタ</Label>
              <Select
                value={statusFilter}
                onValueChange={(value) => setStatusFilter(value as PlayStatus | 'all')}
              >
                <SelectTrigger>
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="all">すべて</SelectItem>
                  {Object.entries(statusLabels).map(([value, label]) => (
                    <SelectItem key={value} value={value}>
                      {label}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>

            <div className="flex-1 min-w-[200px]">
              <Label className="text-sm mb-2 block">ジャンルでフィルタ</Label>
              <Select
                value={genreFilter}
                onValueChange={(value) => setGenreFilter(value as GameGenre | 'all')}
              >
                <SelectTrigger>
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="all">すべて</SelectItem>
                  {genres.map((genre) => (
                    <SelectItem key={genre} value={genre}>
                      {genre}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          </div>

          <div className="mt-4 text-sm text-gray-600">
            {filteredGames.length}件のゲームを表示中（全{games.length}件）
          </div>
        </CardContent>
      </Card>

      <div className="space-y-3">
        {filteredGames.map(game => (
          <Card key={game.id}>
            <CardContent className="pt-6">
              <div className="flex items-start justify-between gap-4">
                <div className="flex-1 min-w-0">
                  <h3 className="truncate mb-2">{game.title}</h3>
                  {game.genre && (
                    <Badge variant="outline" className="mb-2">
                      {game.genre}
                    </Badge>
                  )}
                </div>
                
                <div className="flex items-start gap-2">
                  <div className="space-y-2">
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

                    <Select
                      value={game.genre || 'none'}
                      onValueChange={(value) => {
                        if (value !== 'none') {
                          onUpdateGenre(game.id, value as GameGenre);
                        }
                      }}
                    >
                      <SelectTrigger className="w-[140px]">
                        <SelectValue placeholder="ジャンル" />
                      </SelectTrigger>
                      <SelectContent>
                        {genres.map((genre) => (
                          <SelectItem key={genre} value={genre}>
                            {genre}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
                  </div>

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
    </div>
  );
}

function Label({ children, className }: { children: React.ReactNode; className?: string }) {
  return <div className={className}>{children}</div>;
}
