import { useState } from 'react';
import { GameGenre } from '../types';
import { Button } from './ui/button';
import { Input } from './ui/input';
import { Label } from './ui/label';
import { Card, CardContent, CardHeader, CardTitle } from './ui/card';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from './ui/select';
import { Plus } from 'lucide-react';

interface GameFormAdvancedProps {
  onAddGame: (title: string, genre?: GameGenre) => void;
}

const genres: GameGenre[] = ['RPG', 'アクション', 'アドベンチャー', 'シミュレーション', 'パズル', 'スポーツ', 'その他'];

export function GameFormAdvanced({ onAddGame }: GameFormAdvancedProps) {
  const [title, setTitle] = useState('');
  const [genre, setGenre] = useState<GameGenre | undefined>(undefined);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (title.trim()) {
      onAddGame(title.trim(), genre);
      setTitle('');
      setGenre(undefined);
    }
  };

  return (
    <Card>
      <CardHeader>
        <CardTitle>ゲームを追加</CardTitle>
      </CardHeader>
      <CardContent>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <Label htmlFor="game-title">ゲームタイトル</Label>
            <Input
              id="game-title"
              type="text"
              placeholder="例: ゼルダの伝説"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
            />
          </div>

          <div>
            <Label htmlFor="game-genre">ジャンル（任意）</Label>
            <Select
              value={genre}
              onValueChange={(value) => setGenre(value as GameGenre)}
            >
              <SelectTrigger id="game-genre">
                <SelectValue placeholder="ジャンルを選択" />
              </SelectTrigger>
              <SelectContent>
                {genres.map((g) => (
                  <SelectItem key={g} value={g}>
                    {g}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          <Button type="submit" disabled={!title.trim()} className="w-full">
            <Plus className="h-4 w-4 mr-2" />
            ゲームを追加
          </Button>
        </form>
      </CardContent>
    </Card>
  );
}
