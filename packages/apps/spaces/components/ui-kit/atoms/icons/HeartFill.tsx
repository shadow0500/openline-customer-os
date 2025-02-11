import * as React from 'react';
import { SVGProps } from 'react';
const SvgHeartFill = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='none'
    {...props}
  >
    <path
      d='M19.3 5.71a4.92 4.92 0 0 0-3.51-1.46 4.92 4.92 0 0 0-3.51 1.46L12 6l-.28-.28a4.95 4.95 0 0 0-7 0 5 5 0 0 0 0 7l6.77 6.79a.75.75 0 0 0 1.06 0l6.77-6.79a5 5 0 0 0-.02-7.01z'
      fill='currentColor'
    />
  </svg>
);
export default SvgHeartFill;
