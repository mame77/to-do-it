import { useEffect, useState } from 'react';
import { motion, AnimatePresence } from 'motion/react';
import { Sparkles, X } from 'lucide-react';

interface BonusAnimationProps {
  show: boolean;
  type: 'bonus' | 'penalty';
  points: number;
  message: string;
  onClose: () => void;
}

export function BonusAnimation({ show, type, points, message, onClose }: BonusAnimationProps) {
  useEffect(() => {
    if (show) {
      const timer = setTimeout(() => {
        onClose();
      }, 3000);
      return () => clearTimeout(timer);
    }
  }, [show, onClose]);

  return (
    <AnimatePresence>
      {show && (
        <motion.div
          initial={{ opacity: 0, scale: 0.5, y: 50 }}
          animate={{ opacity: 1, scale: 1, y: 0 }}
          exit={{ opacity: 0, scale: 0.5, y: -50 }}
          className="fixed bottom-8 right-8 z-50"
        >
          <div
            className={`relative p-6 rounded-lg shadow-2xl border-2 ${
              type === 'bonus'
                ? 'bg-gradient-to-br from-green-400 to-green-600 border-green-300'
                : 'bg-gradient-to-br from-red-400 to-red-600 border-red-300'
            }`}
          >
            <button
              onClick={onClose}
              className="absolute top-2 right-2 text-white hover:bg-white/20 rounded-full p-1"
            >
              <X className="h-4 w-4" />
            </button>

            <div className="flex items-center gap-4">
              {type === 'bonus' && (
                <motion.div
                  animate={{
                    rotate: [0, 10, -10, 10, 0],
                    scale: [1, 1.2, 1, 1.2, 1],
                  }}
                  transition={{
                    duration: 0.5,
                    repeat: Infinity,
                    repeatDelay: 1,
                  }}
                >
                  <Sparkles className="h-12 w-12 text-yellow-300" />
                </motion.div>
              )}

              <div className="text-white">
                <div className="text-3xl mb-1">
                  {type === 'bonus' ? '+' : ''}{points}pt
                </div>
                <p className="text-sm opacity-90">{message}</p>
              </div>
            </div>

            {type === 'bonus' && (
              <>
                {[...Array(5)].map((_, i) => (
                  <motion.div
                    key={i}
                    initial={{ opacity: 1, y: 0, x: 0 }}
                    animate={{
                      opacity: 0,
                      y: -100,
                      x: (Math.random() - 0.5) * 100,
                    }}
                    transition={{
                      duration: 1,
                      delay: i * 0.1,
                      repeat: Infinity,
                      repeatDelay: 2,
                    }}
                    className="absolute bottom-0 left-1/2"
                  >
                    <Sparkles className="h-4 w-4 text-yellow-300" />
                  </motion.div>
                ))}
              </>
            )}
          </div>
        </motion.div>
      )}
    </AnimatePresence>
  );
}
