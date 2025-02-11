import * as React from 'react';
import { SVGProps } from 'react';
const SvgCar = (props: SVGProps<SVGSVGElement>) => (
  <svg
    viewBox='0 0 24 24'
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <g fill='currentColor'>
      <path d='m19.78 9.44-1.84-5a1.75 1.75 0 0 0-1.64-1.19H7.7A1.75 1.75 0 0 0 6.06 4.4l-1.84 5a1.76 1.76 0 0 0-1 1.56v4.5A1.73 1.73 0 0 0 4 16.94V19a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1.75h10V19a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-2.06a1.73 1.73 0 0 0 .76-1.44V11a1.76 1.76 0 0 0-.98-1.56zm-.53 6.06a.25.25 0 0 1-.25.25H5a.25.25 0 0 1-.25-.25V11a.25.25 0 0 1 .25-.25h14a.25.25 0 0 1 .25.25v4.5zM7.47 4.91a.25.25 0 0 1 .23-.16h8.6a.25.25 0 0 1 .23.16l1.4 3.84H6.07l1.4-3.84z' />
      <path d='M8 14.75a1.5 1.5 0 1 0 0-3 1.5 1.5 0 0 0 0 3zm8 0a1.5 1.5 0 1 0 0-3 1.5 1.5 0 0 0 0 3z' />
    </g>
  </svg>
);
export default SvgCar;
