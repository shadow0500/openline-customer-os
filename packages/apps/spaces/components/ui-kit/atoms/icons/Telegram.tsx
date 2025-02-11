import * as React from 'react';
import { SVGProps } from 'react';
const SvgTelegram = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='none'
    {...props}
  >
    <path
      d='M12 4a8 8 0 0 0-7.391 4.939 8 8 0 0 0 1.734 8.718 8 8 0 0 0 8.718 1.734A8 8 0 0 0 20 12a8 8 0 0 0-8-8zm3.93 5.48-1.31 6.19c-.1.44-.36.54-.73.34l-2-1.48-1 .93a.51.51 0 0 1-.4.2l.14-2 3.7-3.35c.17-.14 0-.22-.24-.08l-4.54 2.85-2-.62c-.43-.13-.44-.43.09-.63l7.71-3c.38-.11.7.11.58.65z'
      fill='currentColor'
    />
  </svg>
);
export default SvgTelegram;
