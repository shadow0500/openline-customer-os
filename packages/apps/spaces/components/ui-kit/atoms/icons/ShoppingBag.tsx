import * as React from 'react';
import { SVGProps } from 'react';
const SvgShoppingBag = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='none'
    {...props}
  >
    <path
      d='M19.5 8.25h-3v-.5a4.5 4.5 0 0 0-4.5-4.5 4.5 4.5 0 0 0-4.5 4.5v.5h-3A1.25 1.25 0 0 0 3.25 9.5V18A2.75 2.75 0 0 0 6 20.75h12A2.75 2.75 0 0 0 20.75 18V9.5a1.25 1.25 0 0 0-1.25-1.25zM9 7.75a3 3 0 0 1 3-3 3 3 0 0 1 3 3v.5H9v-.5zM19.25 18A1.25 1.25 0 0 1 18 19.25H6A1.25 1.25 0 0 1 4.75 18V9.75H7.5V12a.75.75 0 0 0 .75.75A.75.75 0 0 0 9 12V9.75h6V12a.75.75 0 0 0 .75.75.75.75 0 0 0 .75-.75V9.75h2.75V18z'
      fill='currentColor'
    />
  </svg>
);
export default SvgShoppingBag;
