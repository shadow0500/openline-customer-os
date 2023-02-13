import * as React from 'react';
import { SVGProps } from 'react';
const SvgUsers = (props: SVGProps<SVGSVGElement>) => (
  <svg
    width={24}
    height={24}
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <g fill='currentColor'>
      <path d='M14 12.25a3.75 3.75 0 1 1 0-7.5 3.75 3.75 0 0 1 0 7.5Zm0-6a2.25 2.25 0 1 0 0 4.5 2.25 2.25 0 0 0 0-4.5ZM21 19.25a.76.76 0 0 1-.75-.75c0-1.95-1.06-3.25-6.25-3.25-5.19 0-6.25 1.3-6.25 3.25a.75.75 0 1 1-1.5 0c0-4.75 5.43-4.75 7.75-4.75 2.32 0 7.75 0 7.75 4.75a.76.76 0 0 1-.75.75ZM8.32 13.06H8a3.014 3.014 0 1 1 .58-6 .75.75 0 1 1-.15 1.49 1.46 1.46 0 0 0-1.09.34 1.47 1.47 0 0 0-.54 1 1.49 1.49 0 0 0 1.35 1.64c.326.028.651-.05.93-.22a.752.752 0 0 1 .79 1.28 3 3 0 0 1-1.55.47ZM3 18.5a.76.76 0 0 1-.75-.75c0-2.7.72-4.5 4.25-4.5a.75.75 0 1 1 0 1.5c-2.35 0-2.75.75-2.75 3a.76.76 0 0 1-.75.75Z' />
    </g>
  </svg>
);
export default SvgUsers;
