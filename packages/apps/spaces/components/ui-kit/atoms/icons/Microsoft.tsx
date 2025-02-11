import * as React from 'react';
import { SVGProps } from 'react';
const SvgMicrosoft = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='none'
    {...props}
  >
    <path
      d='M4 4h7.5v7.5H4V4zm8.5 0H20v7.5h-7.5V4zM4 12.5h7.5V20H4v-7.5zm8.5 0H20V20h-7.5v-7.5z'
      fill='currentColor'
    />
  </svg>
);
export default SvgMicrosoft;
