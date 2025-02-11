import * as React from 'react';
import { SVGProps } from 'react';
const SvgTags = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='currentColor'
    {...props}
  >
    <path d='m21.07 10.3-6-6a.75.75 0 0 0-.53-.22H3a.76.76 0 0 0-.75.75v7.88a.75.75 0 0 0 .22.53l6 6a2.31 2.31 0 0 0 1.65.68 2.34 2.34 0 0 0 1.64-.68L12 19a1.14 1.14 0 0 0 .15.23 2.33 2.33 0 0 0 3.29 0l5.65-5.66a2.33 2.33 0 0 0-.02-3.27zM10.7 18.17a.81.81 0 0 1-1.17 0L3.75 12.4V5.58h6.82l5.78 5.78a.83.83 0 0 1 0 1.17l-5.65 5.64zm9.3-5.65-5.65 5.65a.82.82 0 0 1-1.17 0A.54.54 0 0 0 13 18l4.44-4.45a2.33 2.33 0 0 0 0-3.28l-4.75-4.69h1.54L20 11.36a.82.82 0 0 1 0 1.16zM7 9.75a1.25 1.25 0 1 0 0-2.5 1.25 1.25 0 1 0 0 2.5z' />
  </svg>
);
export default SvgTags;
