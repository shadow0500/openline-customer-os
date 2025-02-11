import * as React from 'react';
import { SVGProps } from 'react';
const SvgExternalLink = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='currentColor'
    {...props}
  >
    <path d='M18 20.75H6A2.75 2.75 0 0 1 3.25 18V6A2.75 2.75 0 0 1 6 3.25h6a.75.75 0 1 1 0 1.5H6A1.25 1.25 0 0 0 4.75 6v12A1.25 1.25 0 0 0 6 19.25h12A1.25 1.25 0 0 0 19.25 18v-6a.75.75 0 1 1 1.5 0v6A2.75 2.75 0 0 1 18 20.75zm2-12a.76.76 0 0 1-.75-.75V4.75H16a.75.75 0 1 1 0-1.5h4a.76.76 0 0 1 .75.75v4a.76.76 0 0 1-.75.75z' />
    <path d='M13.5 11.25A.74.74 0 0 1 13 11a.75.75 0 0 1 0-1l6.5-6.5a.75.75 0 0 1 .535-.239.75.75 0 0 1 .544.22.75.75 0 0 1 .22.544.75.75 0 0 1-.239.535L14 11a.74.74 0 0 1-.5.25z' />
  </svg>
);
export default SvgExternalLink;
