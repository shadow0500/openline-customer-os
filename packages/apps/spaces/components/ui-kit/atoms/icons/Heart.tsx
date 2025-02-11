import * as React from 'react';
import { SVGProps } from 'react';
const SvgHeart = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='none'
    {...props}
  >
    <path
      d='M12 19.75a.75.75 0 0 1-.53-.22L4.7 12.74a5 5 0 0 1 0-7 4.95 4.95 0 0 1 7 0L12 6l.28-.28a4.92 4.92 0 0 1 3.51-1.46 4.92 4.92 0 0 1 3.51 1.45 5 5 0 0 1 0 7l-6.77 6.79a.75.75 0 0 1-.53.25zm-3.79-14a3.44 3.44 0 0 0-2.45 1 3.48 3.48 0 0 0 0 4.91L12 17.94l6.23-6.26a3.47 3.47 0 0 0 1.018-2.455A3.47 3.47 0 0 0 18.23 6.77a3.4 3.4 0 0 0-2.44-1 3.44 3.44 0 0 0-2.45 1l-.81.81a.77.77 0 0 1-1.06 0l-.81-.81a3.44 3.44 0 0 0-2.45-1.02z'
      fill='currentColor'
    />
  </svg>
);
export default SvgHeart;
