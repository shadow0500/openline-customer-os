import * as React from 'react';
import { SVGProps } from 'react';
const SvgTimes = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='none'
    {...props}
  >
    <path
      d='m13.06 12 4.42-4.42a.75.75 0 0 0 .239-.535.75.75 0 0 0-.22-.544.75.75 0 0 0-.543-.22.75.75 0 0 0-.535.239L12 10.94 7.58 6.52a.75.75 0 0 0-1.042.018.75.75 0 0 0-.018 1.042L10.94 12l-4.42 4.42a.75.75 0 0 0 0 1.06.75.75 0 0 0 1.06 0L12 13.06l4.42 4.42a.75.75 0 0 0 1.06 0 .75.75 0 0 0 0-1.06L13.06 12z'
      fill='currentColor'
    />
  </svg>
);
export default SvgTimes;
