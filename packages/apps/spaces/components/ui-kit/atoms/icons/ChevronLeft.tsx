import * as React from 'react';
import { SVGProps } from 'react';
const SvgChevronLeft = (props: SVGProps<SVGSVGElement>) => (
  <svg
    viewBox='0 0 24 24'
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <path
      d='M14 17.75a.74.74 0 0 1-.53-.22l-5-5a.75.75 0 0 1 0-1.06l5-5a.75.75 0 0 1 1.06 1.06L10.06 12l4.47 4.47a.75.75 0 0 1-.53 1.28z'
      fill='currentColor'
    />
  </svg>
);
export default SvgChevronLeft;
