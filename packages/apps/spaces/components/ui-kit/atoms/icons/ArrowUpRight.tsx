import * as React from 'react';
import { SVGProps } from 'react';
const SvgArrowUpRight = (props: SVGProps<SVGSVGElement>) => (
  <svg
    viewBox='0 0 24 24'
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <path
      d='M8.46 6.3a.75.75 0 0 0 0 1.5h6.68l-8.62 8.62a.75.75 0 1 0 1.06 1.06l8.62-8.62v6.68a.75.75 0 1 0 1.5 0V7.05a.75.75 0 0 0-.316-.61A.76.76 0 0 0 17 6.3H8.46z'
      fill='currentColor'
    />
  </svg>
);
export default SvgArrowUpRight;
