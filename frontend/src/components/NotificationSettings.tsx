import { useState, useEffect } from 'react';
import { NotificationSettings as NotificationSettingsType } from '../types';
import { Card, CardContent, CardHeader, CardTitle } from './ui/card';
import { Label } from './ui/label';
import { Switch } from './ui/switch';
import { Button } from './ui/button';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from './ui/select';
import { Bell, Volume2 } from 'lucide-react';
import { toast } from 'sonner@2.0.3';

interface NotificationSettingsProps {
  settings: NotificationSettingsType;
  onUpdateSettings: (settings: NotificationSettingsType) => void;
}

export function NotificationSettings({
  settings,
  onUpdateSettings,
}: NotificationSettingsProps) {
  const [localSettings, setLocalSettings] = useState(settings);
  const [permissionStatus, setPermissionStatus] = useState<'granted' | 'denied' | 'default'>('default');

  useEffect(() => {
    if ('Notification' in window) {
      setPermissionStatus(Notification.permission);
    }
  }, []);

  const requestPermission = async () => {
    if ('Notification' in window) {
      const permission = await Notification.requestPermission();
      setPermissionStatus(permission);
      
      if (permission === 'granted') {
        toast.success('通知が有効になりました');
      } else {
        toast.error('通知の許可が拒否されました');
      }
    }
  };

  const handleSave = () => {
    onUpdateSettings(localSettings);
    toast.success('設定を保存しました');
  };

  const testNotification = () => {
    if ('Notification' in window && Notification.permission === 'granted') {
      new Notification('テスト通知', {
        body: 'これはテスト通知です。実際の通知はこのように表示されます。',
        icon: '/favicon.ico',
      });
      toast.success('テスト通知を送信しました');
    } else {
      toast.error('通知の許可が必要です');
    }
  };

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center gap-2">
          <Bell className="h-5 w-5" />
          通知設定
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-6">
        {permissionStatus !== 'granted' && (
          <div className="p-4 bg-yellow-50 border border-yellow-200 rounded-lg">
            <p className="text-sm text-yellow-800 mb-2">
              ブラウザの通知許可が必要です
            </p>
            <Button onClick={requestPermission} size="sm">
              通知を許可
            </Button>
          </div>
        )}

        <div className="flex items-center justify-between">
          <div className="space-y-0.5">
            <Label>通知を有効にする</Label>
            <p className="text-sm text-gray-500">
              スケジュールの開始前に通知を受け取ります
            </p>
          </div>
          <Switch
            checked={localSettings.enabled}
            onCheckedChange={(checked) =>
              setLocalSettings({ ...localSettings, enabled: checked })
            }
            disabled={permissionStatus !== 'granted'}
          />
        </div>

        <div className="space-y-2">
          <Label htmlFor="minutes-before">通知タイミング</Label>
          <Select
            value={String(localSettings.minutesBefore)}
            onValueChange={(value) =>
              setLocalSettings({ ...localSettings, minutesBefore: Number(value) })
            }
            disabled={!localSettings.enabled}
          >
            <SelectTrigger id="minutes-before">
              <SelectValue />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="5">5分前</SelectItem>
              <SelectItem value="10">10分前</SelectItem>
              <SelectItem value="15">15分前</SelectItem>
              <SelectItem value="30">30分前</SelectItem>
              <SelectItem value="60">1時間前</SelectItem>
            </SelectContent>
          </Select>
        </div>

        <div className="flex items-center justify-between">
          <div className="space-y-0.5">
            <Label className="flex items-center gap-2">
              <Volume2 className="h-4 w-4" />
              サウンド
            </Label>
            <p className="text-sm text-gray-500">
              通知時に音を鳴らします
            </p>
          </div>
          <Switch
            checked={localSettings.soundEnabled}
            onCheckedChange={(checked) =>
              setLocalSettings({ ...localSettings, soundEnabled: checked })
            }
            disabled={!localSettings.enabled}
          />
        </div>

        <div className="flex gap-2 pt-4 border-t">
          <Button onClick={handleSave} className="flex-1">
            設定を保存
          </Button>
          <Button
            onClick={testNotification}
            variant="outline"
            disabled={permissionStatus !== 'granted'}
          >
            テスト通知
          </Button>
        </div>
      </CardContent>
    </Card>
  );
}
