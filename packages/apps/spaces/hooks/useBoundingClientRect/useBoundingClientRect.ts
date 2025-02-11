import { RefObject, useLayoutEffect, useState } from 'react';
import { BoundingClientRect } from './types';

export const useBoundingClientRect = (
  ref: RefObject<HTMLElement>,
): BoundingClientRect => {
  const [domRects, setDOMRects] = useState<BoundingClientRect>({
    height: 0,
    width: 0,
    x: 0,
    y: 0,
    bottom: 0,
    left: 0,
    right: 0,
    top: 0,
  });
  useLayoutEffect(() => {
    if (ref?.current) {
      setDOMRects(ref.current.getBoundingClientRect());
    }
  }, [ref]);

  return domRects;
};
