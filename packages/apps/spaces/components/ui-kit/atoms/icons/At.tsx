import * as React from 'react';
import { SVGProps } from 'react';
const SvgAt = (props: SVGProps<SVGSVGElement>) => (
  <svg
    viewBox='0 0 24 24'
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <path
      d='M12 3.25A8.75 8.75 0 0 0 3.25 12 8.65 8.65 0 0 0 12 20.75a.75.75 0 1 0 0-1.5A7.171 7.171 0 0 1 4.75 12 7.26 7.26 0 0 1 12 4.75c4.81 0 7.25 2.44 7.25 7.25v1.38a1.46 1.46 0 1 1-2.91 0v-5a.75.75 0 1 0-1.5 0v.34A4.34 4.34 0 1 0 12 16.34a4.298 4.298 0 0 0 3.24-1.49 2.95 2.95 0 0 0 5.51-1.47V12c0-5.64-3.11-8.75-8.75-8.75zm0 11.59A2.84 2.84 0 1 1 14.84 12 2.85 2.85 0 0 1 12 14.84z'
      fill='currentColor'
    />
  </svg>
);
export default SvgAt;
