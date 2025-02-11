import * as React from 'react';
import { SVGProps } from 'react';
const SvgEuro = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='currentColor'
    {...props}
  >
    <path d='M13 20.75h-.15a8.75 8.75 0 0 1-7.946-5.531A8.75 8.75 0 0 1 6.92 5.75a8.54 8.54 0 0 1 6.23-2.46 8.75 8.75 0 0 1 6 2.5.75.75 0 0 1 0 1.06.75.75 0 0 1-1.06 0 7.26 7.26 0 0 0-7.837-1.546 7.26 7.26 0 0 0-4.514 6.59 7.26 7.26 0 0 0 4.273 6.749A7.26 7.26 0 0 0 17.9 17.38l.22-.21a.79.79 0 0 1 1.09 0 .7.7 0 0 1 .235.489.7.7 0 0 1-.185.511l-.05.05-.29.28A8.72 8.72 0 0 1 13 20.75z' />
    <path d='M17 11.25H3a.75.75 0 1 1 0-1.5h14a.75.75 0 1 1 0 1.5zm-1.5 3H3a.75.75 0 1 1 0-1.5h12.5a.75.75 0 1 1 0 1.5z' />
  </svg>
);
export default SvgEuro;
