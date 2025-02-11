import * as React from 'react';
import { SVGProps } from 'react';
const SvgQuestion = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='currentColor'
    {...props}
  >
    <path d='M16.07 4.93A5.75 5.75 0 0 0 6.25 9a.75.75 0 1 0 1.5 0 4.26 4.26 0 0 1 7.272-3.022A4.26 4.26 0 0 1 12 13.25a.76.76 0 0 0-.75.75v2a.75.75 0 1 0 1.5 0v-1.3a5.76 5.76 0 0 0 3.32-9.77zM12 20.75a1.25 1.25 0 1 0 0-2.5 1.25 1.25 0 1 0 0 2.5z' />
  </svg>
);
export default SvgQuestion;
