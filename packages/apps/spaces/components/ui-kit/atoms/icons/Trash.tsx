import * as React from 'react';
import { SVGProps } from 'react';
const SvgTrash = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='currentColor'
    {...props}
  >
    <path d='M20 8.7H4a.75.75 0 0 1-.53-.22.75.75 0 0 1-.22-.53.75.75 0 0 1 .22-.53A.75.75 0 0 1 4 7.2h16a.75.75 0 0 1 .75.75.75.75 0 0 1-.75.75z' />
    <path d='M16.44 20.75H7.56A2.4 2.4 0 0 1 5 18.49V8a.75.75 0 1 1 1.5 0v10.49c0 .41.47.76 1 .76h8.88c.56 0 1-.35 1-.76V8A.81.81 0 0 1 19 8v10.49a2.4 2.4 0 0 1-2.56 2.26zm.12-13a.74.74 0 0 1-.75-.75V5.51c0-.41-.48-.76-1-.76H9.22c-.55 0-1 .35-1 .76V7a.75.75 0 1 1-1.5 0V5.51a2.41 2.41 0 0 1 2.5-2.26h5.56a2.41 2.41 0 0 1 2.53 2.26V7a.75.75 0 0 1-.75.76v-.01zM10.22 17a.76.76 0 0 1-.75-.75v-4.53a.75.75 0 1 1 1.5 0v4.52a.75.75 0 0 1-.75.76zm3.56 0a.75.75 0 0 1-.75-.75v-4.53a.75.75 0 1 1 1.5 0v4.52a.76.76 0 0 1-.75.76z' />
  </svg>
);
export default SvgTrash;
