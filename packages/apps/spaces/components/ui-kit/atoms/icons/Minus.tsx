import * as React from 'react';
import { SVGProps } from 'react';
const SvgMinus = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='none'
    {...props}
  >
    <path
      d='M20 13H4a1 1 0 0 1-1-1 1 1 0 0 1 1-1h16a1 1 0 0 1 1 1 1 1 0 0 1-1 1z'
      fill='currentColor'
    />
  </svg>
);
export default SvgMinus;
