import * as React from 'react';
import { SVGProps } from 'react';
const SvgSync = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='currentColor'
    {...props}
  >
    <path d='M18.43 4.25a.76.76 0 0 0-.75.75v2.43l-.84-.84a7.24 7.24 0 0 0-12 2.78.74.74 0 0 0 .46 1 .73.73 0 0 0 .25 0 .76.76 0 0 0 .71-.51 5.63 5.63 0 0 1 1.37-2.2 5.76 5.76 0 0 1 8.13 0l.84.84h-2.41a.75.75 0 1 0 0 1.5h4.24a.74.74 0 0 0 .75-.75V5a.75.75 0 0 0-.75-.75zm.25 9.43a.76.76 0 0 0-1 .47 5.63 5.63 0 0 1-1.37 2.2 5.76 5.76 0 0 1-8.13 0l-.84-.84h2.47a.75.75 0 1 0 0-1.5H5.57a.74.74 0 0 0-.75.75V19a.75.75 0 1 0 1.5 0v-2.43l.84.84a7.24 7.24 0 0 0 12-2.78.74.74 0 0 0-.48-.95z' />
  </svg>
);
export default SvgSync;
