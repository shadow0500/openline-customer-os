import * as React from 'react';
import { SVGProps } from 'react';
const SvgMobile = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='currentColor'
    {...props}
  >
    <path d='M16 3.25H8A1.76 1.76 0 0 0 6.25 5v14A1.76 1.76 0 0 0 8 20.75h8A1.76 1.76 0 0 0 17.75 19V5A1.76 1.76 0 0 0 16 3.25zM16.25 19a.25.25 0 0 1-.25.25H8a.25.25 0 0 1-.25-.25V5A.25.25 0 0 1 8 4.75h8a.25.25 0 0 1 .25.25v14zM12 14.5a1.5 1.5 0 0 0-1.061 2.561A1.5 1.5 0 0 0 13.5 16a1.5 1.5 0 0 0-1.5-1.5z' />
  </svg>
);
export default SvgMobile;
