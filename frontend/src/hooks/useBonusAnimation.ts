import { useState } from "react";

export function useBonusAnimation() {
  const [showBonusAnimation, setShowBonusAnimation] = useState(false);
  const [bonusAnimationData, setBonusAnimationData] = useState<{
    type: "bonus" | "penalty";
    points: number;
    message: string;
  } | null>(null);

  const showAnimation = (
    type: "bonus" | "penalty",
    points: number,
    message: string
  ) => {
    setBonusAnimationData({ type, points, message });
    setShowBonusAnimation(true);
  };

  const hideAnimation = () => {
    setShowBonusAnimation(false);
  };

  return {
    showBonusAnimation,
    bonusAnimationData,
    showAnimation,
    hideAnimation,
  };
}
