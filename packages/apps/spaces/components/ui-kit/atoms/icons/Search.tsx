import * as React from 'react';
import { SVGProps } from 'react';
const SvgSearch = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='currentColor'
    {...props}
  >
    <path d='M10.77 18.3a7.53 7.53 0 0 1-6.957-4.648 7.53 7.53 0 0 1 1.632-8.206 7.53 7.53 0 0 1 8.206-1.632A7.53 7.53 0 0 1 18.3 10.77a7.53 7.53 0 0 1-2.206 5.325A7.53 7.53 0 0 1 10.77 18.3zm0-13.55a6 6 0 0 0-5.543 3.704 6 6 0 0 0 1.301 6.539 6 6 0 0 0 6.539 1.301 6 6 0 0 0 3.704-5.543 6 6 0 0 0-6-6z' />
    <path d='M20 20.75a.74.74 0 0 1-.53-.22l-4.13-4.13a.75.75 0 0 1 .018-1.042.75.75 0 0 1 1.042-.018l4.13 4.13a.75.75 0 0 1 0 1.06.74.74 0 0 1-.53.22z' />
  </svg>
);
export default SvgSearch;
