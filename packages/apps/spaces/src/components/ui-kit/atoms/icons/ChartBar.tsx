import * as React from 'react';
import { SVGProps } from 'react';
const SvgChartBar = (props: SVGProps<SVGSVGElement>) => (
  <svg
    width={24}
    height={24}
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <g fill='currentColor'>
      <path d='M4.5 20.25a.76.76 0 0 1-.75-.75v-15a.75.75 0 0 1 1.5 0v15a.76.76 0 0 1-.75.75Z' />
      <path d='M19.5 20.25h-15a.75.75 0 1 1 0-1.5h15a.75.75 0 1 1 0 1.5ZM8 16.75a.76.76 0 0 1-.75-.75v-4a.75.75 0 1 1 1.5 0v4a.76.76 0 0 1-.75.75ZM11.5 16.75a.76.76 0 0 1-.75-.75V8a.75.75 0 1 1 1.5 0v8a.76.76 0 0 1-.75.75ZM15 16.75a.76.76 0 0 1-.75-.75v-4a.75.75 0 1 1 1.5 0v4a.76.76 0 0 1-.75.75ZM18.5 16.75a.76.76 0 0 1-.75-.75V8a.75.75 0 1 1 1.5 0v8a.76.76 0 0 1-.75.75Z' />
    </g>
  </svg>
);
export default SvgChartBar;
