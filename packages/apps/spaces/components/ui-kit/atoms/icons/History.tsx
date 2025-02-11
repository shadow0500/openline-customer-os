import * as React from 'react';
import { SVGProps } from 'react';
const SvgHistory = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='currentColor'
    {...props}
  >
    <path d='M18.05 6.08a8.39 8.39 0 0 0-11.84 0L5 7.29V4.38a.75.75 0 1 0-1.5 0v4.74a.76.76 0 0 0 .75.76H9a.75.75 0 1 0 0-1.5H6l1.27-1.24a6.88 6.88 0 0 1 9.72 0c6.19 6.69-3 15.91-9.72 9.72a.75.75 0 0 0-1.06 0 .74.74 0 0 0 0 1.06A8.37 8.37 0 1 0 18.05 6.08zM12 7.75a.76.76 0 0 0-.75.75V12a.75.75 0 0 0 .22.53L14 15a.74.74 0 0 0 .53.22A.75.75 0 0 0 15 14l-2.28-2.28V8.5a.76.76 0 0 0-.72-.75z' />
  </svg>
);
export default SvgHistory;
