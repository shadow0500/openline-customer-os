import * as React from 'react';
import { SVGProps } from 'react';
const SvgLock = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='none'
    {...props}
  >
    <path
      d='M17 10.25h-.25V8A4.75 4.75 0 0 0 12 3.25 4.75 4.75 0 0 0 7.25 8v2.25H7A2.75 2.75 0 0 0 4.25 13v5A2.75 2.75 0 0 0 7 20.75h10A2.75 2.75 0 0 0 19.75 18v-5A2.75 2.75 0 0 0 17 10.25zM8.75 8A3.25 3.25 0 0 1 12 4.75 3.25 3.25 0 0 1 15.25 8v2.25h-6.5V8zm9.5 10A1.25 1.25 0 0 1 17 19.25H7A1.25 1.25 0 0 1 5.75 18v-5A1.25 1.25 0 0 1 7 11.75h10A1.25 1.25 0 0 1 18.25 13v5z'
      fill='currentColor'
    />
  </svg>
);
export default SvgLock;
