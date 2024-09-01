import { useEffect, useRef } from "react"

export const useEffectOnce = (callback: () => void) => {
  const isCalledRef = useRef(false);

  useEffect(() => {
    if (!isCalledRef.current) {
      isCalledRef.current = true;
      callback();
    }
  }, [callback]);
}
