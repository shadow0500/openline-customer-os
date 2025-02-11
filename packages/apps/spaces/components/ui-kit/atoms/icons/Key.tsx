import * as React from 'react';
import { SVGProps } from 'react';
const SvgKey = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='currentColor'
    {...props}
  >
    <path d='M15 14.75a5.74 5.74 0 0 1-4.07-1.68 5.77 5.77 0 0 1-.733-7.287 5.77 5.77 0 0 1 7.005-2.137 5.77 5.77 0 0 1 3.458 6.456A5.77 5.77 0 0 1 15 14.75zm0-10a4.25 4.25 0 0 0-3.917 2.625A4.25 4.25 0 0 0 12 12a4.25 4.25 0 0 0 3.837 1.169 4.25 4.25 0 0 0 3.101-2.544 4.25 4.25 0 0 0-.397-3.991A4.25 4.25 0 0 0 15 4.75z' />
    <path d='M4.5 20.25A.74.74 0 0 1 4 20a.75.75 0 0 1 0-1l6.46-6.47a.755.755 0 0 1 1.227.25c.037.091.056.19.056.289s-.02.197-.059.288-.094.174-.164.243L5 20a.74.74 0 0 1-.5.25z' />
    <path d='M8 20.75a.74.74 0 0 1-.53-.22l-2-2a.75.75 0 0 1 1.06-1.06l2 2a.75.75 0 0 1 0 1.06.74.74 0 0 1-.53.22zm2-2a.74.74 0 0 1-.53-.22l-2-2a.75.75 0 0 1-.019-1.079.75.75 0 0 1 1.079.019l2 2a.75.75 0 0 1 0 1.06.74.74 0 0 1-.53.22z' />
  </svg>
);
export default SvgKey;
