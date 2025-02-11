import * as React from 'react';
import { SVGProps } from 'react';
const SvgChartPie = (props: SVGProps<SVGSVGElement>) => (
  <svg
    viewBox='0 0 24 24'
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <path
      d='M13.25 4.79V4.5a.76.76 0 0 0-.75-.75 8.8 8.8 0 0 0-7.61 13.13.75.75 0 0 0 1.02.27L6.2 17a7.74 7.74 0 1 0 7.05-12.2v-.01zm-1.5.5v6.78l-5.89 3.38a7.28 7.28 0 0 1 5.89-10.16zm.75 13.46a6.27 6.27 0 0 1-5-2.51l5.37-3.09a.73.73 0 0 0 .38-.65V6.3a6.25 6.25 0 0 1-.75 12.45z'
      fill='currentColor'
    />
  </svg>
);
export default SvgChartPie;
