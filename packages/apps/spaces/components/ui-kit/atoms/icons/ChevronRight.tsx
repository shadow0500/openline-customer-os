import * as React from 'react';
import { SVGProps } from 'react';
const SvgChevronRight = (props: SVGProps<SVGSVGElement>) => (
  <svg
    viewBox='0 0 24 24'
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <path
      d='M10 17.75a.74.74 0 0 1-.53-.22.75.75 0 0 1 0-1.06L13.94 12 9.47 7.53a.75.75 0 0 1 1.06-1.06l5 5a.75.75 0 0 1 0 1.06l-5 5a.74.74 0 0 1-.53.22z'
      fill='currentColor'
    />
  </svg>
);
export default SvgChevronRight;
