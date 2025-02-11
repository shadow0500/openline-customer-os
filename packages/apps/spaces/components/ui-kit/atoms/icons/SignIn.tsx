import * as React from 'react';
import { SVGProps } from 'react';
const SvgSignIn = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='currentColor'
    {...props}
  >
    <path d='M18 20.75h-3a.75.75 0 1 1 0-1.5h3a1.16 1.16 0 0 0 1.25-1V5.78a1.16 1.16 0 0 0-1.25-1h-3a.75.75 0 1 1 0-1.5h3a2.64 2.64 0 0 1 2.75 2.53v12.41A2.64 2.64 0 0 1 18 20.75zm-7-4a.74.74 0 0 1-.53-.22.75.75 0 0 1 0-1.06L13.94 12l-3.47-3.47a.75.75 0 0 1 1.06-1.06l4 4a.75.75 0 0 1 0 1.06l-4 4a.74.74 0 0 1-.53.22z' />
    <path d='M15 12.75H4a.75.75 0 0 1-.75-.75.75.75 0 0 1 .75-.75h11a.75.75 0 0 1 .75.75.75.75 0 0 1-.75.75z' />
  </svg>
);
export default SvgSignIn;
