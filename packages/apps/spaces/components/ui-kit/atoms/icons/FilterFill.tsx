import * as React from 'react';
import { SVGProps } from 'react';
const SvgFilterFill = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='none'
    {...props}
  >
    <path
      d='M20.17 3.91a.76.76 0 0 0-.67-.41h-15a.76.76 0 0 0-.67.41.73.73 0 0 0 .07.78L9.25 12v7.75a.76.76 0 0 0 .75.75h4a.76.76 0 0 0 .75-.75V12l5.35-7.31a.73.73 0 0 0 .07-.78z'
      fill='currentColor'
    />
  </svg>
);
export default SvgFilterFill;
