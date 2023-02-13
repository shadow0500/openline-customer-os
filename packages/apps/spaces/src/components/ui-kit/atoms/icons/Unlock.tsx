import * as React from 'react';
import { SVGProps } from 'react';
const SvgUnlock = (props: SVGProps<SVGSVGElement>) => (
  <svg
    width={24}
    height={24}
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <path
      d='M17 10.25H8.75V8a3.25 3.25 0 0 1 6.5 0 .75.75 0 1 0 1.5 0 4.75 4.75 0 1 0-9.5 0v2.25H7A2.75 2.75 0 0 0 4.25 13v5A2.75 2.75 0 0 0 7 20.75h10A2.75 2.75 0 0 0 19.75 18v-5A2.75 2.75 0 0 0 17 10.25ZM18.25 18A1.25 1.25 0 0 1 17 19.25H7A1.25 1.25 0 0 1 5.75 18v-5A1.25 1.25 0 0 1 7 11.75h10A1.25 1.25 0 0 1 18.25 13v5Z'
      fill='currentColor'
    />
  </svg>
);
export default SvgUnlock;
