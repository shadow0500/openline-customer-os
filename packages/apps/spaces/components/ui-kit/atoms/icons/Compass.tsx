import * as React from 'react';
import { SVGProps } from 'react';
const SvgCompass = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='currentColor'
    {...props}
  >
    <path d='m15.94 7.62-4.88 2a2.63 2.63 0 0 0-1.48 1.48l-2 4.88a.34.34 0 0 0 .19.44.36.36 0 0 0 .25 0l4.88-2a2.63 2.63 0 0 0 1.48-1.48l2-4.88a.34.34 0 0 0-.19-.44.36.36 0 0 0-.25 0zM12 13a1 1 0 0 1-.707-1.707A1 1 0 0 1 13 12a1 1 0 0 1-1 1zm0 8A9 9 0 0 1 5.636 5.636 9 9 0 0 1 21 12a9 9 0 0 1-9 9zm0-16.5a7.5 7.5 0 0 0-5.303 12.803A7.5 7.5 0 0 0 19.5 12 7.5 7.5 0 0 0 12 4.5z' />
  </svg>
);
export default SvgCompass;
