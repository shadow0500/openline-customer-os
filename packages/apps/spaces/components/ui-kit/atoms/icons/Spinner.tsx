import * as React from 'react';
import { SVGProps } from 'react';
const SvgSpinner = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='none'
    {...props}
  >
    <path
      d='M12 21a9 9 0 0 1-7.441-3.949 9 9 0 0 1-.915-8.375 9 9 0 0 1 6.414-5.462A9 9 0 0 1 18.18 5.45a.75.75 0 0 1 0 1.06.74.74 0 0 1-.53.224.74.74 0 0 1-.53-.224 7.51 7.51 0 0 0-9.456-.654A7.51 7.51 0 0 0 5.1 14.982a7.51 7.51 0 0 0 8.413 4.367A7.51 7.51 0 0 0 19.5 12a.75.75 0 0 1 .75-.75.75.75 0 0 1 .75.75 9 9 0 0 1-9 9z'
      fill='currentColor'
    />
  </svg>
);
export default SvgSpinner;
