import { useState } from 'react';
import { Button } from './ui/button';
import { Input } from './ui/input';
import { Label } from './ui/label';
import { Card, CardContent, CardHeader, CardTitle } from './ui/card';
import { Plus } from 'lucide-react';

interface GameFormProps {
  onAddGame: (title: string) => void;
}

export function GameForm({ onAddGame }: GameFormProps) {
  const [title, setTitle] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (title.trim()) {
      onAddGame(title.trim());
      setTitle('');
    }
  };

  return (
    <Card>
      <CardHeader>
        <CardTitle>ゲームを追加</CardTitle>
      </CardHeader>
      <CardContent>
        <form onSubmit={handleSubmit} className="flex gap-2">
          <div className="flex-1">
            <Input
              type="text"
              placeholder="ゲームタイトルを入力"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
            />
          </div>
          <Button type="submit" disabled={!title.trim()}>
            <Plus className="h-4 w-4 mr-2" />
            追加
          </Button>
        </form>
      </CardContent>
    </Card>
  );
}
