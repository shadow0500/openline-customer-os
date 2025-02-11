import * as React from 'react';
import { SVGProps } from 'react';
const SvgCamera = (props: SVGProps<SVGSVGElement>) => (
  <svg
    viewBox='0 0 24 24'
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <g fill='currentColor'>
      <path d='M18 19.25H6a2.75 2.75 0 0 1-2.75-2.75v-6A2.75 2.75 0 0 1 6 7.75h.6L7.78 6a2.76 2.76 0 0 1 2.29-1.22h3.86A2.76 2.76 0 0 1 16.22 6l1.18 1.75h.6a2.75 2.75 0 0 1 2.75 2.75v6A2.75 2.75 0 0 1 18 19.25zm-12-10a1.25 1.25 0 0 0-1.25 1.25v6A1.25 1.25 0 0 0 6 17.75h12a1.25 1.25 0 0 0 1.25-1.25v-6A1.25 1.25 0 0 0 18 9.25h-1a.75.75 0 0 1-.62-.33L15 6.81a1.24 1.24 0 0 0-1-.56h-3.93a1.24 1.24 0 0 0-1 .56L7.62 8.92a.75.75 0 0 1-.62.33H6z' />
      <path d='M12 16.25A3.25 3.25 0 1 1 15.25 13 3.26 3.26 0 0 1 12 16.25zm0-5A1.75 1.75 0 1 0 13.75 13 1.76 1.76 0 0 0 12 11.25z' />
    </g>
  </svg>
);
export default SvgCamera;
